package rtm_test

import (
	"context"
	"testing"

	"github.com/lestrrat/go-slack"
	"github.com/lestrrat/go-slack/rtm"
)

func TestRTM(t *testing.T) {
	if !requireSlackToken(t) {
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cl := slack.New(slackToken)
	rtm := rtm.New(cl)
	go rtm.Run(ctx)

	for e := range rtm.Events() {
		t.Logf("%#v", e)
	}
}
