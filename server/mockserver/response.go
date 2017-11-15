package mockserver

import (
	"github.com/lestrrat/go-slack"
	"github.com/lestrrat/go-slack/objects"
)

// https://github.com/golang/go/commit/6983b9a57955fa12ecd81ab8394ee09e64ef21b9
var aLongLongTimeAgo = objects.EpochTime(233431200)

func stockAuthTestResponse() interface{} {
	var r = struct {
		slack.SlackResponse
		slack.AuthTestResponse
	}{
		SlackResponse: StockResponse("dummy").(slack.SlackResponse),
		AuthTestResponse: slack.AuthTestResponse{
			URL:    "https://jedi.dummy.mock-slack.com",
			Team:   "jedi",
			User:   "lukeskywalker",
			TeamID: "T01234567",
			UserID: "U01234567",
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

func stockReactionsGetResponse() interface{}                { return StockResponse("dummy") }
func stockObjectsUserProfileObjectsTeam() interface{}       { return StockResponse("dummy") }
func stockObjectsUserList() interface{}                     { return StockResponse("dummy") }
func stockObjectsBot() interface{}                          { return StockResponse("dummy") }
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
func stockChannelsHistoryResponse() interface{}             { return StockResponse("dummy") }
func stockObjectsGroup() interface{}                        { return StockResponse("dummy") }
func stockObjectsGroupList() interface{}                    { return StockResponse("dummy") }
func stockReactionsListResponse() interface{}               { return StockResponse("dummy") }
func stockObjectsChannelList() interface{}                  { return StockResponse("dummy") }
