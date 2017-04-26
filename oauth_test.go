package slack_test

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOAuthAccessUnit(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	s := httptest.NewServer(newDummyServer())
	defer s.Close()

	c := newSlackWithDummy(s)
	_, err := c.OAuth().Access("foo", "bar", "baz").RedirectURI("quux").Do(ctx)
	if !assert.NoError(t, err, "OAuth.Access should succeed") {
		return
	}
}
