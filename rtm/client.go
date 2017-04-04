package rtm

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
	"github.com/lestrrat/go-slack"
	"github.com/pkg/errors"
)

func New(cl *slack.Client) *Client {
	return &Client{
		client:   cl,
		eventsCh: make(chan *Event),
	}
}

func (c *Client) Events() <-chan *Event {
	return c.eventsCh
}

func (c *Client) Run(octx context.Context) error {
	octxwc, cancel := context.WithCancel(octx)
	defer cancel()

	ctx := newRtmCtx(octxwc, c.eventsCh)
	go ctx.run()

	for {
		select {
		case <-ctx.Done():
			return nil
		default:
		}

		ctx.emit(&Event{typ: RTMConnectingEvent})
		res, err := c.client.RTM().Start(ctx)
		if err != nil {
			// TODO: exponential backoff
			log.Printf("%s", err)
			return errors.Wrap(err, `failed to start rtm session`)
		}

		conn, _, err := websocket.DefaultDialer.Dial(res.URL, nil)
		if err != nil {
			// TODO: exponential backoff
			log.Printf("%s", err)
			return errors.Wrap(err, `failed to dial websocket`)
		}
		ctx.handleConn(conn)
	}

	return nil
}

func (ctx *rtmCtx) handleConn(conn *websocket.Conn) {
	defer conn.Close()

	in := make(chan []byte)
	go func(ch chan []byte, conn *websocket.Conn) {
		defer close(ch)

		for {
			typ, data, err := conn.ReadMessage()
			if err != nil {
				return
			}
			// we only understand text messages
			if typ != websocket.TextMessage {
				continue
			}
			ch <- data
		}
	}(in, conn)

	for {
		select {
		case <-ctx.Done():
			return
		case payload, ok := <-in:
			if !ok {
				return
			}
			log.Printf("%s", payload)
			var event Event
			if err := json.Unmarshal(payload, &event); err != nil {
				log.Printf("failed to unmarshal: %s", err)
			}
		}
	}
}

type rtmCtx struct {
	context.Context
	inbuf        chan *Event
	outbuf       chan<- *Event
	writeTimeout time.Duration
}

func newRtmCtx(octx context.Context, outch chan<- *Event) *rtmCtx {
	return &rtmCtx{
		Context:      octx,
		inbuf:        make(chan *Event),
		outbuf:       outch,
		writeTimeout: 500 * time.Millisecond,
	}
}

func (ctx *rtmCtx) run() {
	// callback to write to the outgoing channel
	trywrite := func(ctx *rtmCtx, e *Event) error {
		fmt.Println("Attempting to write to out channel")
		wrtimeout := time.NewTimer(ctx.writeTimeout)
		defer wrtimeout.Stop()
		select {
		case <-ctx.Done():
			return ctx.Err()
		case ctx.outbuf <- e:
			return nil
		case <-wrtimeout.C:
			return errors.New("timeout")
		}
	}

	periodic := time.NewTicker(time.Second)
	defer periodic.Stop()

	var events []*Event
	for {
		fmt.Println("Proxy loop")
		select {
		case <-ctx.Done():
			return
		case e := <-ctx.inbuf:
			events = append(events, e)
		case <-periodic.C:
			// attempt to write periodically. only comes here in case we
			// were unable to write to the channel during the alloted time
		}

		for len(events) > 0 {
			e := events[0]
			// Try writing. if we fail, bail out of this write loop
			if err := trywrite(ctx, e); err != nil {
				break
			}
			// if we were successful, pop the current one and try the next one
			events = events[1:]
		}

		// shink the slice if we're too big
		if l := len(events); l > 16 && cap(events) > 64 {
			events = append([]*Event(nil), events...)
		}
	}
}

// emit sends the event e to a channel. This method doesn't "fail" because
// we expect the the proxy loop in run() to handle things gracefully
func (ctx *rtmCtx) emit(e *Event) {
	fmt.Println("emit")
	defer fmt.Println("done emit")
	select {
	case <-ctx.Done():
		return
	case ctx.inbuf <- e:
		return
	}
}
