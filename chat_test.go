package slack_test

import (
	"context"
	"testing"

	"github.com/lestrrat-go/slack"
	"github.com/lestrrat-go/slack/objects"
	"github.com/stretchr/testify/assert"
)

// Test message create, update, delete
func TestChatMessage(t *testing.T) {
	if !requireSlackToken(t) || !requireDMUser(t) {
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := slack.New(slackToken)
	t.Run("basic usage", func(t *testing.T) {
		res, err := c.Chat().PostMessage(testDmUser).Text("hello").Do(ctx)
		if !assert.NoError(t, err, "Chat.PostMessage failed") {
			return
		}
		t.Logf("%#v", res)

		/*
			res, err = c.Chat().Update(ctx, res.Channel, res.Timestamp, "hello, world!")
			if !assert.NoError(t, err, "Chat.Update failed") {
				return
			}
			t.Logf("%#v", res)

				res, err = c.Chat().Delete(ctx, res.Channel, res.Timestamp)
				if !assert.NoError(t, err, "Chat.Delete failed") {
					return
				}
				t.Logf("%#v", res)
		*/
	})

	t.Run("buttons", func(t *testing.T) {
		var attachment objects.Attachment
		attachment.CallbackID = "wopr_game"
		attachment.Color = "#3AA3E3"
		attachment.Fallback = "You are unable to choose a game"
		attachment.Text = "Choose a game to play"
		attachment.Actions.
			Append(&objects.Action{
				Name:  "game",
				Text:  "Chess",
				Type:  objects.ButtonActionType,
				Value: "chess",
			}).
			Append(&objects.Action{
				Name:  "game",
				Text:  "Falken's Maze",
				Type:  objects.ButtonActionType,
				Value: "maze",
			}).
			Append(&objects.Action{
				Name:  "game",
				Text:  "Thermonuclear War",
				Style: "danger",
				Type:  objects.ButtonActionType,
				Value: "war",
				Confirm: &objects.Confirmation{
					Title:       "Are you sure?",
					Text:        "Wouldn't you prefer a good game of chess?",
					OkText:      "Yes",
					DismissText: "No",
				},
			})

		res, err := c.Chat().PostMessage(testDmUser).
			AsUser(true).
			Attachment(&attachment).
			Text("Would you like to play a game?").
			Do(ctx)
		if !assert.NoError(t, err, "Chat.PostMessage failed") {
			return
		}
		t.Logf("%#v", res)
	})
}

// MeMessages require chat:write:user, so I think I'm going to punt it
/*
func TestChatMeMessage(t *testing.T) {
	if !requireSlackToken(t) || !requireDMUser(t) || !requireRealUser(t) {
		return
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := slack.New(slackToken)

	res, err := c.Chat().MeMessage(ctx, testDmUser, "hello")
	if !assert.NoError(t, err, "Chat.MeMessage failed") {
		return
	}
	t.Logf("%#v", res)
}
*/
