package rtm_test

import (
	"context"
	"testing"

	"github.com/lestrrat/go-slack"
	"github.com/lestrrat/go-slack/rtm"
	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	if !requireSlackToken(t) {
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cl := slack.New(slackToken)
	rtm := rtm.New(cl)

	ch := make(chan error)
	go func(ch chan error) {
		defer close(ch)
		ch <- rtm.Run(ctx)
	}(ch)

	for {
		select {
		case err := <-ch:
			assert.NoError(t, err, "rtm.Run returned an error")
			return
		case e := <-rtm.Events():
			t.Logf("%#v", e)
		}
	}
}
