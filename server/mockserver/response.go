package mockserver

import (
	"strconv"
	"time"

	"github.com/lestrrat/go-slack"
	"github.com/lestrrat/go-slack/objects"
)

// These stock responses do not necessarily represent the actual format used
// by the real Slack API responses. PRs welcome to fix them to something more
// appropriate

// https://github.com/golang/go/commit/6983b9a57955fa12ecd81ab8394ee09e64ef21b9
const aLongLongTimeAgo = objects.EpochTime(233431200)
const channelID = "C0123456"

var TeamJedi = objects.Team{
	ID:     "T0123456",
	Name:   "Jedis",
	Domain: "jedi.mock-slack-library.com",
}
var UserLukeSkywalker = objects.User{
	ID:   "U0123456",
	Name: "luke.skywalker",
}

func stockAuthTestResponse() interface{} {
	var r = struct {
		slack.SlackResponse
		slack.AuthTestResponse
	}{
		SlackResponse: StockResponse("dummy").(slack.SlackResponse),
		AuthTestResponse: slack.AuthTestResponse{
			URL:    "https://jedi.dummy.mock-slack.com",
			Team:   TeamJedi.Name,
			TeamID: TeamJedi.ID,
			User:   UserLukeSkywalker.Name,
			UserID: UserLukeSkywalker.ID,
		},
	}
	return r
}

func stockObjectsChannel() interface{} {
	var r = struct {
		slack.SlackResponse
		objects.Channel
	}{
		SlackResponse: StockResponse("dummy").(slack.SlackResponse),
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
	return r
}

func stockReactionsGetResponse() interface{}          { return StockResponse("dummy") }
func stockObjectsUserProfileObjectsTeam() interface{} { return StockResponse("dummy") }
func stockObjectsUserList() interface{}               { return StockResponse("dummy") }
func stockObjectsBot() interface{} {
	var r = struct {
		slack.SlackResponse
		objects.Bot
	}{
		SlackResponse: StockResponse("dummy").(slack.SlackResponse),
		Bot: objects.Bot{
			ID:      "B0123456",
			AppID:   "A0123456",
			Deleted: false,
			Name:    "jabbathehutt-bot",
			Icons: objects.Icons{
				Image36: "https://upload.wikimedia.org/wikipedia/commons/f/f0/Jabba_the_Hutt_%288175228157%29.jpg",
				Image48: "https://upload.wikimedia.org/wikipedia/commons/f/f0/Jabba_the_Hutt_%288175228157%29.jpg",
				Image72: "https://upload.wikimedia.org/wikipedia/commons/f/f0/Jabba_the_Hutt_%288175228157%29.jpg",
			},
		},
	}
	return r
}
func stockString() interface{}                              { return StockResponse("dummy") }
func stockChatResponse() interface{}                        { return StockResponse("dummy") }
func stockObjectsGroupBool() interface{}                    { return StockResponse("dummy") }
func stockOAuthAccessResponse() interface{}                 { return StockResponse("dummy") }
func stockObjectsUsergroupUsersList() interface{}           { return StockResponse("dummy") }
func stockObjectsUser() interface{}                         { return StockResponse("dummy") }
func stockEmojiListResponse() interface{}                   { return StockResponse("dummy") }
func stockObjectsMessageListObjectsThreadInfo() interface{} { return StockResponse("dummy") }
func stockRTMResponse() interface{}                         { return StockResponse("dummy") }
func stockObjectsMessageList() interface{}                  { return StockResponse("dummy") }
func stockStringObjectsMessageList() interface{}            { return StockResponse("dummy") }
func stockObjectsUsergroup() interface{}                    { return StockResponse("dummy") }
func stockObjectsUserProfile() interface{}                  { return StockResponse("dummy") }
func stockObjectsUsergroupList() interface{}                { return StockResponse("dummy") }
func stockObjectsUserPresence() interface{}                 { return StockResponse("dummy") }
func stockChannelsHistoryResponse() interface{} {
	var r = struct {
		slack.SlackResponse
		slack.ChannelsHistoryResponse
	}{
		SlackResponse: StockResponse("dummy").(slack.SlackResponse),
		ChannelsHistoryResponse: slack.ChannelsHistoryResponse{
			HasMore: true,
			Latest:  "dummy",
			Messages: objects.MessageList{
				&objects.Message{
					Channel:   channelID,
					Timestamp: strconv.FormatInt(time.Now().Unix()-7*86400, 10),
				},
			},
		},
	}
	return r
}
func stockObjectsGroup() interface{}          { return StockResponse("dummy") }
func stockObjectsGroupList() interface{}      { return StockResponse("dummy") }
func stockReactionsListResponse() interface{} { return StockResponse("dummy") }
func stockObjectsChannelList() interface{}    { return StockResponse("dummy") }
