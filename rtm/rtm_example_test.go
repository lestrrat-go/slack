package rtm_test

import (
	"context"
	"fmt"

	"github.com/lestrrat-go/slack/rtm"
	"github.com/lestrrat-go/slack"
)

func processMessageEvent(e *rtm.MessageEvent) {
	// Dummy
}

func ExampleClient() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Create a new Slack REST client. This is required,
	// as the RTM client needs to hit the rtm.start API
	// endpoint in order to initiate a websocket session
	restcl := slack.New(slackToken)

	// Create a new Slack RTM client. Use `Run()` to start
	// listening to incoming events. If you need to stop the
	// RTM client, use the cancel function for the context.Context
	// object that you passed
	rtmcl := rtm.New(restcl)
	go rtmcl.Run(ctx)

	// Now, consume your events. They come from a channel.
	// If you care about properly bailing out of a processing
	// loop like below, you should be using select {} on the
	// channel returned by `rtmcl.Events()`
	for e := range rtmcl.Events() {
		// Event types are listed in event_types.go file, which
		// is auto-generated from Slack's web page describing
		// the events.
		switch typ := e.Type(); typ {
		case rtm.MessageType:
			// It's a message. Event objects carry a unique payload
			// depending on the event type, so you will have to do
			// a type assertion.
			// TODO: come up with a mapping of Type -> Data's underlying
			// type, so it's easier for the users to see
			processMessageEvent(e.Data().(*rtm.MessageEvent))
		default:
			fmt.Printf("Unhandled event: %s", typ)
		}
	}
}
