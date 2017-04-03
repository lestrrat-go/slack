package slack

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func (e *RTMEvent) Type() EventType {
	return e.typ
}

func NewRTM() *RTM {
	return &RTM{
		outch: make(chan Event),
	}
}

func (r *RTM) IncomingEvents() <-chan Event {
	return r.outch
}

func (r *RTM) Run(octx context.Context) error {
	ctx := newRtmCtx(octx, r.outch)
	go ctx.run()

	for {
		fmt.Println("RTM.Run loop")
		select {
		case <-ctx.Done():
			return nil
		default:
		}

		ctx.emit(&RTMEvent{typ: RTMConnectingEvent})
		time.Sleep(time.Minute)
	}
}

type rtmCtx struct {
	context.Context
	inbuf        chan Event
	outbuf       chan<- Event
	writeTimeout time.Duration
}

func newRtmCtx(octx context.Context, outch chan<- Event) *rtmCtx {
	return &rtmCtx{
		Context:      octx,
		inbuf:        make(chan Event),
		outbuf:       outch,
		writeTimeout: 500 * time.Millisecond,
	}
}

func (ctx *rtmCtx) run() {
	// callback to write to the outgoing channel
	trywrite := func(ctx *rtmCtx, e Event) error {
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

	var events []Event
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
			events = append([]Event(nil), events...)
		}
	}
}

// emit sends the event e to a channel. This method doesn't "fail" because
// we expect the the proxy loop in run() to handle things gracefully
func (ctx *rtmCtx) emit(e Event) {
	fmt.Println("emit")
	defer fmt.Println("done emit")
	select {
	case <-ctx.Done():
		return
	case ctx.inbuf <- e:
		return
	}
}
