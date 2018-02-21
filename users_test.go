package slack_test

import (
	"context"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/lestrrat-go/slack"
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

func TestUsersListUnit(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	s := httptest.NewServer(newDummyServer())
	defer s.Close()

	c := newSlackWithDummy(s)
	_, err := c.Users().List().Do(ctx)
	assert.NoError(t, err, "Users.List should succeed")

	_, err = c.Users().List().Limit(100).IncludeLocale(true).Do(ctx)
	assert.NoError(t, err, "Users.List should succeed")
}

func TestUsersDeletePhotoUnit(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	s := httptest.NewServer(newDummyServer())
	defer s.Close()

	c := newSlackWithDummy(s)
	err := c.Users().DeletePhoto().Do(ctx)
	assert.NoError(t, err, "Users.DeletePhoto should succeed")
}

func TestUsersGetPresenceUnit(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	s := httptest.NewServer(newDummyServer())
	defer s.Close()

	c := newSlackWithDummy(s)
	_, err := c.Users().GetPresence("foo").Do(ctx)
	if !assert.NoError(t, err, "Users.GetPresence should succeed") {
		return
	}
}

func TestUsersIdentityUnit(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	s := httptest.NewServer(newDummyServer())
	defer s.Close()

	c := newSlackWithDummy(s)
	_, _, err := c.Users().Identity().Do(ctx)
	assert.NoError(t, err, "Users.Identity should succeed")
}

func TestUsersInfoUnit(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	s := httptest.NewServer(newDummyServer())
	defer s.Close()

	c := newSlackWithDummy(s)
	_, err := c.Users().Info("foo").Do(ctx)
	assert.NoError(t, err, "Users.Info should succeed")

	_, err = c.Users().Info("foo").IncludeLocale(true).Do(ctx)
	assert.NoError(t, err, "Users.Info should succeed")
}
