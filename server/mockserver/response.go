package mockserver

import (
	"strconv"
	"time"

	"github.com/lestrrat-go/slack/objects"
)

// These stock responses do not necessarily represent the actual format used
// by the real Slack API responses. PRs welcome to fix them to something more
// appropriate

// https://github.com/golang/go/commit/6983b9a57955fa12ecd81ab8394ee09e64ef21b9
const aLongLongTimeAgo = objects.EpochTime(233431200)

var ReminderMeetMaceWindu = objects.BuildReminder().
	ID("Rm72119BBY").
	Creator(UserObiwanKenobi.ID()).
	User(UserObiwanKenobi.ID()).
	Text("Meet Mace Windu over lunch").
	Recurring(false).
	Time(aLongLongTimeAgo.Add(86400 * 25 * 365)).
	CompleteTimestamp(objects.EpochTime(0)).
	MustBuild()

var FileComputer = objects.BuildFile().
	Channels(ChannelJedis.ID()).
	CommentsCount(1).
	Created(aLongLongTimeAgo.Add(86400).Int()).
	ID("F00000001").
	Name("computer.gif").
	Timestamp(aLongLongTimeAgo.Add(86400).Int()).
	Title("computer.gif").
	User(UserLukeSkywalker.ID()).
	MustBuild()

var ChannelJedis = objects.BuildChannel().
	ID("C1J3D1ZRUL3").
	Created(aLongLongTimeAgo).
	Creator(UserYoda.ID()).
	Members(
		UserLukeSkywalker.ID(),
		UserObiwanKenobi.ID(),
		UserYoda.ID(),
	).
	Name("jedis").
	NameNormalized("jedis").
	NumMembers(3).
	Purpose(objects.BuildPurpose().
		Creator("yoda").
		LastSet(aLongLongTimeAgo).
		Value("There is no emotion, there is peace.\nThere is no ignorance, there is knowledge.\nThere is no passion, there is serenity.\nThere is no chaos, there is harmony.\nThere is no death, there is the Force.").
		MustBuild(),
	).
	Topic(objects.BuildTopic().
		Creator("yoda").
		Value("Jedi meetup and drinks next Tuesday").
		LastSet(aLongLongTimeAgo).
		MustBuild(),
	).
	IsChannel(true).
	IsMember(true).
	MustBuild()

var TeamJedi = objects.BuildTeam().
	ID("T0123456").
	Name("Jedis").
	Domain("jedi.mock-slack-library.com").
	MustBuild()

var UserLukeSkywalker = objects.BuildUser().
	ID("U0123456").
	Name("luke.skywalker").
	MustBuild()

var UserObiwanKenobi = objects.BuildUser().
	ID("U0012345").
	Name("obiwan.kenobi").
	MustBuild()

var UserYoda = objects.BuildUser().
	ID("U0000001").
	Name("yoda").
	MustBuild()

func stockObjectsAuthTestResponse() interface{} {
	var r = struct {
		objects.GenericResponse
		*objects.AuthTestResponse
	}{
		GenericResponse: StockResponse("dummy").(objects.GenericResponse),
		AuthTestResponse: objects.BuildAuthTestResponse().
			URL("https://jedi.dummy.mock-slack.com").
			Team(TeamJedi.Name()).
			TeamID(TeamJedi.ID()).
			User(UserLukeSkywalker.Name()).
			UserID(UserLukeSkywalker.ID()).
			MustBuild(),
	}
	return r
}

func stockObjectsChannel() interface{} {
	var r = struct {
		objects.GenericResponse
		*objects.Channel `json:"channel"`
	}{
		GenericResponse: StockResponse("dummy").(objects.GenericResponse),
		Channel:         ChannelJedis,
	}
	return r
}

func stockObjectsReactionsGetResponse() interface{} {
	reaction := objects.BuildReaction().
		Count(1).
		Name("stuck_out_tongue_winking_eye").
		Users(UserLukeSkywalker.ID()).
		MustBuild()
	f := objects.BuildFile().
		Channels(FileComputer.Channels()...).
		CommentsCount(FileComputer.CommentsCount()).
		Created(FileComputer.Created()).
		ID(FileComputer.ID()).
		Name(FileComputer.Name()).
		Reactions(reaction).
		Timestamp(FileComputer.Timestamp()).
		Title(FileComputer.Title()).
		User(FileComputer.User()).
		MustBuild()

	var r = struct {
		objects.GenericResponse
		*objects.ReactionsGetResponse
	}{
		GenericResponse: StockResponse("dummy").(objects.GenericResponse),
		ReactionsGetResponse: objects.BuildReactionsGetResponse().
			File(f).
			Type("file").
			MustBuild(),
	}
	return r
}
func stockObjectsUserProfileObjectsTeam() interface{} { return StockResponse("dummy") }
func stockObjectsUserList() interface{} {
	var r = struct {
		objects.GenericResponse
		*objects.UserList `json:"members,omitempty"`
	}{
		GenericResponse: StockResponse("dummy").(objects.GenericResponse),
		UserList: &objects.UserList{
			UserLukeSkywalker,
			UserObiwanKenobi,
			UserYoda,
		},
	}
	return r
}
func stockObjectsBot() interface{} {
	var r = struct {
		objects.GenericResponse
		*objects.Bot
	}{
		GenericResponse: StockResponse("dummy").(objects.GenericResponse),
		Bot: objects.BuildBot().
			ID("B0123456").
			AppID("A0123456").
			Deleted(false).
			Name("jabbathehutt-bot").
			Icons(objects.BuildIcons().
				Image36("https://upload.wikimedia.org/wikipedia/commons/f/f0/Jabba_the_Hutt_%288175228157%29.jpg").
				Image48("https://upload.wikimedia.org/wikipedia/commons/f/f0/Jabba_the_Hutt_%288175228157%29.jpg").
				Image72("https://upload.wikimedia.org/wikipedia/commons/f/f0/Jabba_the_Hutt_%288175228157%29.jpg").
				MustBuild(),
			).
			MustBuild(),
	}
	return r
}
func stockString() interface{} { return StockResponse("dummy") }
func stockObjectsChatResponse() interface{} {
	var r = struct {
		objects.GenericResponse
		*objects.Message
	}{
		GenericResponse: StockResponse("dummy").(objects.GenericResponse),
		Message: objects.BuildMessage().
			Channel(ChannelJedis.ID()).
			MustBuild(),
	}
	return r
}
func stockObjectsGroupBool() interface{}           { return StockResponse("dummy") }
func stockObjectsOAuthAccessResponse() interface{} { return StockResponse("dummy") }
func stockObjectsUsergroupUsersList() interface{}  { return StockResponse("dummy") }
func stockObjectsUser() interface{} {
	var r = struct {
		objects.GenericResponse
		*objects.User `json:"user"`
	}{
		GenericResponse: StockResponse("dummy").(objects.GenericResponse),
		User:            UserLukeSkywalker,
	}
	return r
}
func stockObjectsEmojiListResponse() interface{}            { return StockResponse("dummy") }
func stockObjectsMessageListObjectsThreadInfo() interface{} { return StockResponse("dummy") }
func stockObjectsRTMResponse() interface{}                  { return StockResponse("dummy") }
func stockObjectsMessageList() interface{}                  { return StockResponse("dummy") }
func stockStringObjectsMessageList() interface{}            { return StockResponse("dummy") }
func stockObjectsUsergroup() interface{}                    { return StockResponse("dummy") }
func stockObjectsUserProfile() interface{}                  { return StockResponse("dummy") }
func stockObjectsUsergroupList() interface{}                { return StockResponse("dummy") }
func stockObjectsUserPresence() interface{}                 { return StockResponse("dummy") }
func stockObjectsChannelsHistoryResponse() interface{} {
	var r = struct {
		objects.GenericResponse
		*objects.ChannelsHistoryResponse
	}{
		GenericResponse: StockResponse("dummy").(objects.GenericResponse),
		ChannelsHistoryResponse: objects.BuildChannelsHistoryResponse().
			HasMore(true).
			Latest("dummy").
			Messages(objects.BuildMessage().
				Channel(ChannelJedis.ID()).
				Timestamp(strconv.FormatInt(time.Now().Unix()-7*86400, 10)).
				MustBuild(),
			).
			MustBuild(),
	}
	return r
}
func stockObjectsGroup() interface{}                 { return StockResponse("dummy") }
func stockObjectsGroupList() interface{}             { return StockResponse("dummy") }
func stockObjectsReactionsListResponse() interface{} { return StockResponse("dummy") }
func stockObjectsReminder() interface{} {
	var r = struct {
		objects.GenericResponse
		*objects.Reminder
	}{
		GenericResponse: StockResponse("dummy").(objects.GenericResponse),
		Reminder:        ReminderMeetMaceWindu,
	}
	return r
}
func stockObjectsReminderList() interface{} { return StockResponse("dummy") }
func stockObjectsChannelList() interface{}  { return StockResponse("dummy") }

func stockObjectsDialogResponse() interface{} {
	return struct {
		objects.GenericResponse
		objects.DialogResponse
	}{
		GenericResponse: StockResponse("dummy").(objects.GenericResponse),
		DialogResponse: objects.DialogResponse{
			ResponseMetadata: struct {
				Messages []string `json:"messages"`
			}{[]string{""}},
		},
	}
}

func stockObjectsEphemeralResponse() interface{} { return StockResponse("dummy") }

func stockObjectsPermalinkResponse() interface{} { return StockResponse("dummy") }
