package slack_test

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUsersProfileGetUnit(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	s := httptest.NewServer(newDummyServer())
	defer s.Close()

	c := newSlackWithDummy(s)
	_, err := c.UsersProfile().Get().User("foo").IncludeLabels(true).Do(ctx)
	if !assert.NoError(t, err, "UsersProfile.Get should succeed") {
		return
	}
}

func TestUsersProfileSetUnit(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	s := httptest.NewServer(newDummyServer())
	defer s.Close()

	c := newSlackWithDummy(s)
	_, err := c.UsersProfile().Set().User("foo").Profile(nil).Name("bar").Value("baz").Do(ctx)
	if !assert.NoError(t, err, "UsersProfile.Get should succeed") {
		return
	}
}
