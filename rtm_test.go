package slack_test

import (
	"context"
	"testing"

	"github.com/lestrrat/go-slack"
)

func TestRTM(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	rtm := slack.NewRTM()
	go rtm.Run(ctx)

	e := <-rtm.IncomingEvents()
	t.Logf("%#v", e)
}
