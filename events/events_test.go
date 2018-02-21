package events_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/lestrrat-go/slack/events"
	"github.com/stretchr/testify/assert"
)

func TestEventPayload(t *testing.T) {
	//
	// https://api.slack.com/events/message#stars__pins__and_reactions
	const src = `{
		"type": "message.channels",
		"user": "U061F1EUR",
		"item": {
        "type": "message",
        "channel": "C2147483705",
        "user": "U2147483697",
        "text": "Hello world",
        "ts": "1355517523.000005",
        "is_starred": true,
        "pinned_to": ["C024BE7LT"],
        "reactions": [
            {
                "name": "astonished",
                "count": 3,
                "users": [ "U1", "U2", "U3" ]
            },
            {
                "name": "facepalm",
                "count": 1034,
                "users": [ "U1", "U2", "U3", "U4", "U5" ]
            }
        ]
    }
}`

	var p events.Event
	err := json.NewDecoder(strings.NewReader(src)).Decode(&p)
	assert.NoError(t, err, "decode should suceed")
}
