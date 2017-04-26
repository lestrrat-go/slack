package slack_test

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmojiListUnit(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	s := httptest.NewServer(newDummyServer())
	defer s.Close()

	c := newSlackWithDummy(s)
	_, err := c.Emoji().List().Do(ctx)
	if !assert.NoError(t, err, "Emoji.List should succeed") {
		return
	}
}