package slack_test

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStarsAddUnit(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	s := httptest.NewServer(newDummyServer())
	defer s.Close()

	c := newSlackWithDummy(s)
	err := c.Stars().Add().File("bar").FileComment("baz").Channel("quux").Timestamp("farb").Do(ctx)
	if !assert.NoError(t, err, "Stars.Add should succeed") {
		return
	}
}

func TestStarsListUnit(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	s := httptest.NewServer(newDummyServer())
	defer s.Close()

	c := newSlackWithDummy(s)
	_, err := c.Stars().List().Count(100).Page(5).Do(ctx)
	if !assert.NoError(t, err, "Stars.List should succeed") {
		return
	}
}

func TestStarsRemoveUnit(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	s := httptest.NewServer(newDummyServer())
	defer s.Close()

	c := newSlackWithDummy(s)
	err := c.Stars().Remove().File("bar").FileComment("baz").Channel("quux").Timestamp("farb").Do(ctx)
	if !assert.NoError(t, err, "Stars.Remove should succeed") {
		return
	}
}
