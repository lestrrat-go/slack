package slack_test

import (
	"context"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/lestrrat-go/slack"
	"github.com/stretchr/testify/assert"
)

func TestChannelsList_Info(t *testing.T) {
	if !requireSlackToken(t) {
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := slack.New(slackToken)

	list, err := c.Channels().List().Do(ctx)
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

		fromInfo, err := c.Channels().Info(channel.ID).Do(ctx)
		if !assert.NoError(t, err, "Channels.Info failed") {
			return
		}
		t.Logf("%#v", fromInfo)

		if !assert.Equal(t, channel.ID, fromInfo.ID, "Channel.Info should produce identical channels") {
			return
		}

		history, err := c.Channels().History(channel.ID).Do(ctx)
		if !assert.NoError(t, err, "Channels.History failed") {
			return
		}
		t.Logf("%#v", history)
	}
}

func TestChannelsListUnit(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	s := httptest.NewServer(newDummyServer())
	defer s.Close()

	c := newSlackWithDummy(s)
	_, err := c.Channels().List().Do(ctx)
	if !assert.NoError(t, err, "Channels.List should succeed") {
		return
	}

	_, err = c.Channels().List().ExcludeArchive(true).ExcludeMembers(true).Do(ctx)
	assert.NoError(t, err, "Channels.List should succeed")
}

func TestChannelsMarkUnit(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	s := httptest.NewServer(newDummyServer())
	defer s.Close()

	c := newSlackWithDummy(s)
	err := c.Channels().Mark("general").Do(ctx)
	assert.NoError(t, err, "Channels.Mark should succeed")

	err = c.Channels().Mark("general").Timestamp("194290").Do(ctx)
	assert.NoError(t, err, "Channels.Mark should succeed")
}

func TestChannelsRenameUnit(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	s := httptest.NewServer(newDummyServer())
	defer s.Close()

	c := newSlackWithDummy(s)
	_, err := c.Channels().Rename("old", "new").Do(ctx)
	assert.NoError(t, err, "Channels.Rename should succeed")

	_, err = c.Channels().Rename("old", "new").Validate(true).Do(ctx)
	assert.NoError(t, err, "Channels.Rename should succeed")
}

func TestChannelsRepliesUnit(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	s := httptest.NewServer(newDummyServer())
	defer s.Close()

	c := newSlackWithDummy(s)
	_, err := c.Channels().Replies("general", "78347289").Do(ctx)
	assert.NoError(t, err, "Channels.Replies should succeed")
}

func TestChannelsArchiveUnit(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	s := httptest.NewServer(newDummyServer())
	defer s.Close()

	c := newSlackWithDummy(s)
	if !assert.NoError(t, c.Channels().Archive("foo").Do(ctx), "Archive should succeed") {
		return
	}
}

func TestChannelsCreateUnit(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	s := httptest.NewServer(newDummyServer())
	defer s.Close()

	c := newSlackWithDummy(s)
	if !assert.NoError(t, c.Channels().Create("foo").Validate(true).Do(ctx), "Create should succeed") {
		return
	}

	if !assert.NoError(t, c.Channels().Create("foo").Do(ctx), "Create should succeed") {
		return
	}
}

func TestChannelsHistoryUnit(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	s := httptest.NewServer(newDummyServer())
	defer s.Close()

	c := newSlackWithDummy(s)
	_, err := c.Channels().History("foo").Count(100).Inclusive(true).Latest("dummy").Oldest("dummy").Timestamp("dummy").Unreads(true).Do(ctx)
	if !assert.NoError(t, err, "History should succeed") {
		return
	}
	_, err = c.Channels().History("foo").Do(ctx)
	if !assert.NoError(t, err, "History should succeed") {
		return
	}
}

func TestChannelsInviteUnit(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	s := httptest.NewServer(newDummyServer())
	defer s.Close()

	c := newSlackWithDummy(s)
	_, err := c.Channels().Invite("foo", "bar").Do(ctx)
	if !assert.NoError(t, err, "Invite should succeed") {
		return
	}
}

func TestChannelsKickUnit(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	s := httptest.NewServer(newDummyServer())
	defer s.Close()

	c := newSlackWithDummy(s)
	if !assert.NoError(t, c.Channels().Kick("foo", "bar").Do(ctx), "Kick should succeed") {
		return
	}
}

func TestChannelsLeaveUnit(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	s := httptest.NewServer(newDummyServer())
	defer s.Close()

	c := newSlackWithDummy(s)
	if !assert.NoError(t, c.Channels().Leave("foo").Do(ctx), "Leave should succeed") {
		return
	}
}

func TestChannelsSetTopicUnit(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	s := httptest.NewServer(newDummyServer())
	defer s.Close()

	c := newSlackWithDummy(s)
	_, err := c.Channels().SetTopic("general", "new topic").Do(ctx)
	if !assert.NoError(t, err, "channels.SetTopic should succeed") {
		return
	}
}

func TestChannelsUnarchiveUnit(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	s := httptest.NewServer(newDummyServer())
	defer s.Close()

	c := newSlackWithDummy(s)
	err := c.Channels().Unarchive("general").Do(ctx)
	if !assert.NoError(t, err, "channels.Unarchive should succeed") {
		return
	}
}
