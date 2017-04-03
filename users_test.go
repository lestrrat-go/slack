package slack_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/lestrrat/go-slack"
	"github.com/stretchr/testify/assert"
)

func TestUsersList(t *testing.T) {
	if !requireSlackToken(t) {
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := slack.New(slackToken)

	res, err := c.Users().List(ctx, false)
	if !assert.NoError(t, err, "Users.List failed") {
		return
	}

	buf, _ := json.MarshalIndent(res, "", "  ")
	t.Logf("%s", buf)
}
