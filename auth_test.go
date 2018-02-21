package slack_test

import (
	"context"
	"testing"

	"github.com/lestrrat-go/slack"
	"github.com/stretchr/testify/assert"
)

func TestAuthTest(t *testing.T) {
	if !requireSlackToken(t) {
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := debuggingClient(slackToken)
	res, err := c.Auth().Test().Do(ctx)
	if !assert.NoError(t, err, "Auth.Test failed") {
		return
	}
	t.Logf("%#v", res)
}

func TestAuthRevoke(t *testing.T) {
	t.Skip()

	if !requireSlackToken(t) {
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := slack.New(slackToken)
	err := c.Auth().Revoke().Do(ctx)
	if !assert.NoError(t, err, "Auth.Revoke failed") {
		return
	}
}

