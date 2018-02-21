package rtm

import (
	"context"
	"encoding/json"
	"log"
	"math"
	"time"

	"github.com/cenkalti/backoff"
	"github.com/gorilla/websocket"
	pdebug "github.com/lestrrat-go/pdebug"
	"github.com/lestrrat-go/slack"
	"github.com/pkg/errors"
)

// Option defines an interface of optional parameters to the
// `rtm.New` constructor.
type Option interface {
	Name() string
	Value() interface{}
}

type option struct {
	name  string
	value interface{}
}

func (o *option) Name() string {
	return o.name
}
func (o *option) Value() interface{} {
	return o.value
}

const (
	backoffStrategyKey = "backoff_strategy"
	pingIntervalKey    = "ping_interval"
)

func New(cl *slack.Client, options ...Option) *Client {
	pingInterval := 5 * time.Minute
	var strategy backoff.BackOff
	for _, o := range options {
		switch o.Name() {
		case pingIntervalKey:
			pingInterval = o.Value().(time.Duration)
		case backoffStrategyKey:
			strategy = o.Value().(backoff.BackOff)
		}
	}

	if strategy == nil {
		expback := backoff.NewExponentialBackOff()
		expback.InitialInterval = 100 * time.Millisecond
		expback.MaxInterval = 5 * time.Second
		expback.MaxElapsedTime = 0
		strategy = expback
	}

	return &Client{
		backoffStrategy: strategy,
		client:          cl,
		eventsCh:        make(chan *Event),
		pingInterval:    pingInterval,
	}
}

func (c *Client) Events() <-chan *Event {
	return c.eventsCh
}

// Run starts the RTM run loop.
func (c *Client) Run(octx context.Context) error {
	if pdebug.Enabled {
		pdebug.Printf("rtm client: Run()")
		defer pdebug.Printf("rtm client: end Run()")
	}
	octxwc, cancel := context.WithCancel(octx)
	defer cancel()

	ctx := newRtmCtx(octxwc, c.eventsCh)
	ctx.backoffStrategy = c.backoffStrategy
	ctx.pingInterval = c.pingInterval
	go ctx.run()

	// start a message ID generator
	go func(ctx context.Context, ch chan int) {
		defer close(ch)

		msgid := 1
		for {
			select {
			case <-ctx.Done():
				return
			case ch <- int(msgid):
				// max is defined as int32, just to be small enough to not
				// overflow the server side (which, we can't know about)
				if msgid == math.MaxInt32 {
					msgid = 0
				} else {
					msgid++
				}
			}
		}
	}(ctx, ctx.msgidCh)

	for {
		select {
		case <-ctx.Done():
			return nil
		default:
		}

		if err := emitTimeout(ctx, &Event{typ: ClientConnectingEventType}, 5*time.Second); err != nil {
			return errors.Wrap(err, `failed to emit connecting event`)
		}

		var conn *websocket.Conn

		strategy := ctx.backoffStrategy
		err := backoff.Retry(func() error {
			res, err := c.client.RTM().Start().Do(ctx)
			if err != nil {
				log.Printf("failed to start RTM sesson: %s", err)
				return err
			}
			conn, _, err = websocket.DefaultDialer.Dial(res.URL, nil)
			if err != nil {
				log.Printf("failed to dial to websocket: %s", err)
				return err
			}
			return nil
		}, backoff.WithContext(strategy, ctx))

		if err != nil {
			return errors.Wrap(err, `failed to connect to RTM endpoint`)
		}

		ctx.handleConn(conn)
	}

	return nil
}

func (ctx *rtmCtx) handleConn(conn *websocket.Conn) error {
	// we get here if we manually canceled the context
	// of if the websocket ReadMessage returned an error
	defer emitTimeout(ctx, &Event{typ: ClientDisconnectedEventType}, time.Second)
	defer conn.Close()

	in := make(chan []byte)

	// This goroutine is responsible for reading from the
	// websocket connection. It's separated because the
	// ReadMessage() operation is blocking.
	go func(ch chan []byte, conn *websocket.Conn) {
		defer close(ch)

		for {
			typ, data, err := conn.ReadMessage()
			if err != nil {
				// There was an error. we need to bail out
				if pdebug.Enabled {
					pdebug.Printf("error while reading message from websocket: %s", err)
				}
				return
			}

			// we only understand text messages
			if typ != websocket.TextMessage {
				if pdebug.Enabled {
					pdebug.Printf("received websocket message, but it is not a text payload. refusing to process")
				}
				continue
			}
			if pdebug.Enabled {
				pdebug.Printf("forwarding new websocket message")
			}
			ch <- data
		}
	}(in, conn)

	pingTick := time.NewTicker(ctx.pingInterval)
	var pingMessage struct {
		ID    int    `json:"id"`
		Type  string `json:"type"`
		acked bool
	}
	pingMessage.Type = "ping"

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-pingTick.C:
			// it's time to ping.
			if pingMessage.ID != 0 {
				if !pingMessage.acked {
					if pdebug.Enabled {
						pdebug.Printf("websocket proxy: did not get pong in time")
					}
					return errors.New("websocket proxy: did not get pong in time") // can't continue
				}
			}
			// Send the ping request, and wait for the pong
			pingMessage.ID = <-ctx.msgidCh
			pingMessage.acked = false
			if pdebug.Enabled {
				pdebug.Printf("websocket proxy: sending ping %d", pingMessage.ID)
			}
			if err := conn.WriteJSON(pingMessage); err != nil {
				if pdebug.Enabled {
					pdebug.Printf("websocket proxy: failed to write ping JSON: %s", err)
				}
				return errors.Wrap(err, `websocket proxy: failed to write ping JSON`) // can't continue
			}
			continue
		case payload, ok := <-in:
			if !ok {
				if pdebug.Enabled {
					pdebug.Printf("websocket proxy: detected incoming channel close.")
				}
				// if the channel is closed, we probably had some
				// problems in the ReadMessage proxy. bail out
				return errors.New(`websocket proxy: detected incoming channel close`)
			}

			if pdebug.Enabled {
				pdebug.Printf("websocket proxy: received raw payload: %s", payload)
			}

			var event Event
			if err := json.Unmarshal(payload, &event); err != nil {
				if pdebug.Enabled {
					pdebug.Printf("websocket proxy: failed to unmarshal payload: %s", err)
				}
				continue
			}

			// intercept pong replies
			if pong, ok := event.Data().(*PongEvent); ok {
				if pong.ReplyTo == pingMessage.ID {
					if pdebug.Enabled {
						pdebug.Printf("websocket proxy: got pong for message ID %d", pingMessage.ID)
					}
					pingMessage.acked = true
				}
				continue
			}

			emit(ctx, &event)
		}
	}
	return nil
}

type rtmCtx struct {
	context.Context
	backoffStrategy backoff.BackOff
	inbuf           chan *Event
	msgidCh         chan int
	outbuf          chan<- *Event
	pingInterval    time.Duration
	writeTimeout    time.Duration
}

func newRtmCtx(octx context.Context, outch chan<- *Event) *rtmCtx {
	return &rtmCtx{
		Context:      octx,
		inbuf:        make(chan *Event),
		msgidCh:      make(chan int),
		outbuf:       outch,
		writeTimeout: 500 * time.Millisecond,
	}
}

// Attempt to write to the outgoing channel, within the
// alloted time frame.
func (ctx *rtmCtx) trywrite(e *Event) error {
	tctx, cancel := context.WithTimeout(ctx, ctx.writeTimeout)
	defer cancel()

	select {
	case <-tctx.Done():
		switch err := tctx.Err(); err {
		case context.DeadlineExceeded:
			return errors.New("write timeout")
		default:
			return err
		}
	case ctx.outbuf <- e:
		return nil
	}

	return errors.New("unreachable")
}

// The point of this loop is to ensure the writer (the loop receiving
// events from the websocket connection) can safely write the events
// to a channel without worrying about blocking.
//
// Inside this loop, we read from the channel receiving the events,
// and we either write to the consumer channel, or buffer in our
// in memory queue (list) for later consumption
func (ctx *rtmCtx) run() {
	defer close(ctx.outbuf) // make sure the reader of Events() gets notified

	periodic := time.NewTicker(time.Second)
	defer periodic.Stop()

	var events []*Event
	for {
		select {
		case <-ctx.Done():
			return
		case e := <-ctx.inbuf:
			events = append(events, e)
		case <-periodic.C:
			// attempt to flush the buffer periodically.
		}

		// events should only contain more than one item if we
		// failed to write to the outgoing channel within the
		// allotted time
		for len(events) > 0 {
			e := events[0]
			// Try writing. if we fail, bail out of this write loop
			if err := ctx.trywrite(e); err != nil {
				break
			}
			// if we were successful, pop the current one and try the next one
			events = events[1:]
		}

		// shrink the slice if we're too big
		if l := len(events); l > 16 && cap(events) > 2*l {
			events = append([]*Event(nil), events...)
		}
	}
}

func (ctx *rtmCtx) WithTimeout(t time.Duration) (*rtmCtx, func()) {
	octx, cancel := context.WithTimeout(ctx, t)

	var newCtx rtmCtx
	newCtx.Context = octx
	newCtx.backoffStrategy = ctx.backoffStrategy
	newCtx.inbuf = ctx.inbuf
	newCtx.msgidCh = ctx.msgidCh
	newCtx.outbuf = ctx.outbuf
	newCtx.pingInterval = ctx.pingInterval
	newCtx.writeTimeout = ctx.writeTimeout

	return &newCtx, cancel
}

func emitTimeout(ctx *rtmCtx, e *Event, t time.Duration) error {
	newCtx, cancel := ctx.WithTimeout(t)
	defer cancel()

	return emit(newCtx, e)
}

// emit sends the event e to a channel. This method doesn't "fail" to
// write because we expect the the proxy loop in run() to read these
// requests as quickly as possible under normal circumstances
func emit(ctx *rtmCtx, e *Event) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case ctx.inbuf <- e:
		return nil
	}
}
