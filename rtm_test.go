package slack_test

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRTMStartUnit(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	s := httptest.NewServer(newDummyServer())
	defer s.Close()

	c := newSlackWithDummy(s)
	_, err := c.RTM().Start().Do(ctx)
	if !assert.NoError(t, err, "RTM.Start should succeed") {
		return
	}
}