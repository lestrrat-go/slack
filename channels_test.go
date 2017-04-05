package slack_test

import (
	"context"
	"testing"
	"time"

	"github.com/lestrrat/go-slack"
	"github.com/stretchr/testify/assert"
)

func TestChannelsList_Info(t *testing.T) {
	if !requireSlackToken(t) {
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := slack.New(slackToken)

	list, err := c.Channels().List(ctx, false)
	if !assert.NoError(t, err, "Channels.List failed") {
		return
	}

	if !assert.True(t, len(list) > 1, "There should be more than 1 channels (slackbot and more)") {
		return
	}

	timeout := time.NewTimer(5 * time.Second)
	defer timeout.Stop()

	for i, channel := range list {
		select {
		case <-timeout.C:
			assert.True(t, i > 0, "processed at last 1 channel")
			return
		default:
		}

		fromInfo, err := c.Channels().Info(ctx, channel.ID)
		if !assert.NoError(t, err, "Channels.Info failed") {
			return
		}
		t.Logf("%#v", fromInfo)

		if !assert.Equal(t, channel.ID, fromInfo.ID, "Channel.Info should produce identical channels") {
			return
		}
	}
}
