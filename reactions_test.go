package slack_test

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReactionsAddUnit(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	s := httptest.NewServer(newDummyServer())
	defer s.Close()

	c := newSlackWithDummy(s)
	err := c.Reactions().Add("foo").File("bar").FileComment("baz").Channel("quux").Timestamp("farb").Do(ctx)
	if !assert.NoError(t, err, "Reactions.Add should succeed") {
		return
	}
}

func TestReactionsGetUnit(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	s := httptest.NewServer(newDummyServer())
	defer s.Close()

	c := newSlackWithDummy(s)
	_, err := c.Reactions().Get().File("bar").FileComment("baz").Channel("quux").Timestamp("farb").Full(true).Do(ctx)
	if !assert.NoError(t, err, "Reactions.Get should succeed") {
		return
	}
}

func TestReactionsListUnit(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	s := httptest.NewServer(newDummyServer())
	defer s.Close()

	c := newSlackWithDummy(s)
	_, err := c.Reactions().List().User("bar").Full(true).Count(100).Page(5).Do(ctx)
	if !assert.NoError(t, err, "Reactions.List should succeed") {
		return
	}
}

func TestReactionsRemoveUnit(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	s := httptest.NewServer(newDummyServer())
	defer s.Close()

	c := newSlackWithDummy(s)
	err := c.Reactions().Remove("foo").File("bar").FileComment("baz").Channel("quux").Timestamp("farb").Do(ctx)
	if !assert.NoError(t, err, "Reactions.Remove should succeed") {
		return
	}
}
