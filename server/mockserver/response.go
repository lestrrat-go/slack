package mockserver

import (
	"github.com/lestrrat/go-slack"
	"github.com/lestrrat/go-slack/objects"
)

// https://github.com/golang/go/commit/6983b9a57955fa12ecd81ab8394ee09e64ef21b9
var aLongLongTimeAgo = objects.EpochTime(233431200)

var (
	StockChannelResponse = struct {
		slack.SlackResponse
		objects.Channel
	}{
		SlackResponse: slack.SlackResponse{
			OK: true,
		},
		Channel: objects.Channel{
			Group: objects.Group{
				Conversation: objects.Conversation{
					ID:      "123456789ABCDEFG",
					Created: aLongLongTimeAgo,
				},
				Creator: "yoda",
				Members: []string{
					"obiwan",
					"lukeskywalker",
				},
				Name:           "jedis",
				NameNormalized: "jedis",
				NumMembers:     2,
				Purpose: objects.Purpose{
					Creator: "yoda",
					LastSet: aLongLongTimeAgo,
					Value:   "There is no emotion, there is peace.\nThere is no ignorance, there is knowledge.\nThere is no passion, there is serenity.\nThere is no chaos, there is harmony.\nThere is no death, there is the Force.",
				},
				Topic: objects.Topic{
					Creator: "yoda",
					Value:   "Jedi meetup and drinks next Tuesday",
					LastSet: aLongLongTimeAgo,
				},
			},
		},
	}
)
