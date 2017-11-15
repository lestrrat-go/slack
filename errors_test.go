package slack_test

import (
	"encoding/json"
	"testing"

	slack "github.com/lestrrat/go-slack"
	"github.com/stretchr/testify/assert"
)

func TestStructuredError(t *testing.T) {
	b := []byte(`{"code": 1049, "msg": "hello, world"}`)
	var e slack.ErrorResponse
	if !assert.NoError(t, json.Unmarshal(b, &e), `json unmarshal should succeed`) {
		return
	}

	if !assert.Equal(t, 1049, e.Code, "code should match") {
		return
	}
	if !assert.Equal(t, "hello, world", e.Message, "message should match") {
		return
	}
}
