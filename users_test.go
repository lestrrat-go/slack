package slack_test

import (
	"context"
	"testing"
	"time"

	"github.com/lestrrat/go-slack"
	"github.com/stretchr/testify/assert"
)

func TestUsersList_Info_Presence(t *testing.T) {
	if !requireSlackToken(t) {
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := slack.New(slackToken)

	list, err := c.Users().List().Do(ctx)
	if !assert.NoError(t, err, "Users.List failed") {
		return
	}

	if !assert.True(t, len(list) > 1, "There should be more than 1 users (slackbot and more)") {
		return
	}

	timeout := time.NewTimer(5 * time.Second)
	defer timeout.Stop()

	for i, user := range list {
		select {
		case <-timeout.C:
			assert.True(t, i > 0, "processed at last 1 user")
			return
		default:
		}

		fromInfo, err := c.Users().Info(user.ID).Do(ctx)
		if !assert.NoError(t, err, "Users.Info failed") {
			return
		}

		if !assert.Equal(t, user.ID, fromInfo.ID, "User.Info should produce identical users") {
			return
		}

		presence, err := c.Users().GetPresence(user.ID).Do(ctx)
		if !assert.NoError(t, err, "Users.GetPresence failed") {
			return
		}
		t.Logf("%#v", presence)
	}
}
