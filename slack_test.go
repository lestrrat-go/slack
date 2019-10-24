package slack_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"regexp"
	"strconv"
	"testing"
	"time"

	"github.com/lestrrat-go/slack"
	"github.com/lestrrat-go/slack/objects"
	"github.com/lestrrat-go/slack/server"
	"github.com/lestrrat-go/slack/server/mockserver"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

const token = "AbCdEfG"

var reLooksLikeChannelID = regexp.MustCompile(`^C[A-Z0-9]+`)

func looksLikeChannelID(s string) bool {
	return reLooksLikeChannelID.MatchString(s)
}

func checkChannel(t *testing.T, channel *objects.Channel) bool {
	if !assert.NotNil(t, channel, "channel should be non-nil") {
		return false
	}
	if !assert.NotEmpty(t, channel.Name, "channel.Name should be populated") {
		return false
	}
	if !assert.True(t, looksLikeChannelID(channel.ID()), "channel.ID looks like an ID") {
		return false
	}
	return true
}

func checkReminder(t *testing.T, reminder *objects.Reminder) bool {
	if !assert.NotNil(t, reminder, "reminder should be non-nil") {
		return false
	}

	if !assert.NotEmpty(t, reminder.ID, "reminder.ID should be populated") {
		return false
	}
	if !assert.NotEmpty(t, reminder.Text, "reminder.Text should be populated") {
		return false
	}
	if !assert.NotEmpty(t, reminder.Time, "reminder.Time should be populated") {
		return false
	}
	if !assert.NotEmpty(t, reminder.User, "reminder.User should be populated") {
		return false
	}
	return true
}

// These tests just excercise the "regular" code path
func TestWithMockServer(t *testing.T) {
	h := mockserver.New(mockserver.WithToken(token))
	s := server.New()
	h.InstallHandlers(s)

	ts := httptest.NewServer(s)
	defer ts.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cl := slack.New(token, slack.WithAPIEndpoint(ts.URL))
	t.Run("Auth", func(t *testing.T) {
		t.Run("Revoke", func(t *testing.T) {
			err := cl.Auth().Revoke().Test(true).Do(ctx)
			if !assert.NoError(t, err, "auth.revoke should succeed") {
				return
			}
		})
		t.Run("Test", func(t *testing.T) {
			res, err := cl.Auth().Test().Do(ctx)
			if !assert.NoError(t, err, "auth.test should succeed") {
				return
			}
			if !assert.NotEmpty(t, res.URL, "res.URL should be populated") {
				return
			}
			if !assert.NotEmpty(t, res.Team, "res.Team should be populated") {
				return
			}
			if !assert.NotEmpty(t, res.User, "res.User should be populated") {
				return
			}
			if !assert.NotEmpty(t, res.TeamID, "res.TeamID should be populated") {
				return
			}
			if !assert.NotEmpty(t, res.UserID, "res.UserID should be populated") {
				return
			}
		})
	})
	t.Run("Bots", func(t *testing.T) {
		t.Run("Info", func(t *testing.T) {
			res, err := cl.Bots().Info("B0123456").Do(ctx)
			if !assert.NoError(t, err, "bots.info should succeed") {
				return
			}

			if !assert.NotEmpty(t, res.ID, "res.ID should be populated") {
				return
			}
			if !assert.NotEmpty(t, res.AppID, "res.AppID should be populated") {
				return
			}
			if !assert.NotEmpty(t, res.Name, "res.Name should be populated") {
				return
			}
			if !assert.False(t, res.Deleted(), "res.Delete should be false") {
				return
			}
			if !assert.NotEmpty(t, res.Icons().Image36(), "res.Icons.Image36 should be populated") {
				return
			}
			if !assert.NotEmpty(t, res.Icons().Image48(), "res.Icons.Image48 should be populated") {
				return
			}
			if !assert.NotEmpty(t, res.Icons().Image72(), "res.Icons.Image72 should be populated") {
				return
			}
		})
	})
	t.Run("Channels", func(t *testing.T) {
		t.Run("Archive", func(t *testing.T) {
			err := cl.Channels().Archive("C0123456").Do(ctx)
			if !assert.NoError(t, err, "channels.archive should succeed") {
				return
			}
		})
		t.Run("Create", func(t *testing.T) {
			err := cl.Channels().Create("siths").Validate(true).Do(ctx)
			if !assert.NoError(t, err, "channels.create should succeed") {
				return
			}
		})
		t.Run("History", func(t *testing.T) {
			res, err := cl.Channels().History(mockserver.ChannelJedis.ID()).
				Count(1000).
				Inclusive(true).
				Latest("dummy").
				Oldest("dummy").
				Unreads(true).
				Timestamp("dummy").
				Do(ctx)
			if !assert.NoError(t, err, "channels.history should succeed") {
				return
			}
			if !assert.True(t, res.HasMore(), "res.HasMore should be true") {
				return
			}
			if !assert.NotEmpty(t, res.Latest, "res.Latest should be populated") {
				return
			}
			if !assert.NotEmpty(t, res.Messages, "res.Messages should be populated") {
				return
			}
		})
		t.Run("Info", func(t *testing.T) {
			res, err := cl.Channels().Info(mockserver.ChannelJedis.ID()).
				IncludeLocale(true).
				Do(ctx)
			if !assert.NoError(t, err, "channels.info should succeed") {
				return
			}
			if !checkChannel(t, res) {
				return
			}
		})
		t.Run("Invite", func(t *testing.T) {
			res, err := cl.Channels().Invite("C0123456", "U0123456").
				Do(ctx)
			if !assert.NoError(t, err, "channels.invite should succeed") {
				return
			}
			if !checkChannel(t, res) {
				return
			}
		})
	})
	t.Run("Reactions", func(t *testing.T) {
		t.Run("Get", func(t *testing.T) {
			res, err := cl.Reactions().Get().
				Channel(mockserver.ChannelJedis.ID()).
				// File, FileComment
				Full(true).
				// Timestamp
				Do(ctx)
			if !assert.NoError(t, err, "reactions.get should succeed") {
				return
			}
			if !assert.NotNil(t, res, "reaction should be non-nil") {
				return
			}
		})
	})
	t.Run("Reminder", func(t *testing.T) {
		t.Run("Add", func(t *testing.T) {
			res, err := cl.Reminders().Add("Meet Mace Windu over lunch", mockserver.ReminderMeetMaceWindu.Time().Int()).Do(ctx)
			if !assert.NoError(t, err, "reminders.add should succeed") {
				return
			}
			if !assert.NotNil(t, res, "reminder should be non-nil") {
				return
			}
		})
	})
	t.Run("Users", func(t *testing.T) {
		t.Run("List", func(t *testing.T) {
			res, err := cl.Users().List().Do(ctx)
			if !assert.NoError(t, err, "users.list should succeed") {
				return
			}
			if !assert.NotNil(t, res, "users should be non-nil") {
				return
			}

			for _, u := range res {
				u2, err := cl.Users().Info(u.ID()).Do(ctx)
				if !assert.NoError(t, err, `users.info should succeed`) {
					return
				}

				// XXX the mock server can only return the same response regardless
				// of the input, so we punt testing for equality and make sure it's
				// non-nill
				if !assert.NotNil(t, u2, `user should be non-nil`) {
					return
				}
			}
		})
	})
}

var testDmUser string
var testChannel string
var isBot bool
var slackToken string

func init() {
	slackToken = os.Getenv("SLACK_TOKEN")
	testDmUser = os.Getenv("TEST_DM_USER") // don't forget to include an "@"

	if len(slackToken) > 0 {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		cl := slack.New(slackToken)
		res, err := cl.Auth().Test().Do(ctx)
		if err == nil {
			user, err := cl.Users().Info(res.UserID()).Do(ctx)
			if err == nil {
				isBot = user.IsBot()
			}
		}
	}
}

func requireSlackToken(t *testing.T) bool {
	if slackToken == "" {
		t.Skip("SLACK_TOKEN not available")
		return false
	}
	return true
}

func requireDMUser(t *testing.T) bool {
	if testDmUser == "" {
		t.Skip("TEST_DM_USER not available")
		return false
	}
	return true
}

func requireRealUser(t *testing.T) bool {
	if !requireSlackToken(t) {
		return false
	}

	if isBot {
		t.Skip("User authenticated by the token is a bot.")
		return false
	}
	return true
}

func debuggingClient(tok string, options ...slack.Option) *slack.Client {
	options = append(options, slack.WithDebug(true))
	return slack.New(tok, options...)
}

type dummyServer struct {
	mux http.Handler
}

type expectedArg struct {
	name     string
	required bool
	check    func([]string) error
}

func nilcheck(_ []string) error { return nil }
func newArg(name string, check func([]string) error) *expectedArg {
	return &expectedArg{
		name:  name,
		check: check,
	}
}
func intArg(name string) *expectedArg {
	return newArg(name, func(l []string) error {
		_, err := strconv.ParseInt(l[0], 10, 64)
		return err
	})
}

func (s *dummyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

type mux struct {
	*http.ServeMux
}

func newMux() *mux {
	return &mux{
		ServeMux: http.NewServeMux(),
	}
}

type argCheck []*expectedArg

func ArgCheck(args ...*expectedArg) argCheck {
	return append(argCheck(nil), args...)
}

func (ac argCheck) Check(r *http.Request) error {
	for _, arg := range ac {
		v, ok := r.Form[arg.name]
		if !ok || len(v) == 0 {
			if arg.required {
				return errors.Errorf("required argument %s was not present", arg.name)
			}
			return nil
		}

		if check := arg.check; check != nil {
			if err := check(v); err != nil {
				return errors.Wrapf(err, "check for %s failed", arg.name)
			}
		}
		delete(r.Form, arg.name)
	}

	for name := range r.Form {
		return errors.Errorf("extra argument %s found", name)
	}

	return nil
}

func (m *mux) HandleFunc(path string, ac argCheck, resp interface{}) {
	if resp == nil {
		resp = objects.BuildGenericResponse().
			OK(true).
			MustBuild()
	}

	f := func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		if err := ac.Check(r); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, `{"error": %s}`, strconv.Quote(err.Error()))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(resp)
	}
	m.ServeMux.HandleFunc(path, f)
}

func required(arg *expectedArg) *expectedArg {
	var newArg expectedArg
	newArg = *arg
	newArg.required = true
	return &newArg
}

func newDummyServer() *dummyServer {
	var s dummyServer

	tokenArg := newArg("token", nil)
	channelArg := newArg("channel", nil)
	nameArg := newArg("name", nil)
	userArg := newArg("user", nil)

	mux := newMux()
	mux.HandleFunc(
		"/api/auth.revoke",
		ArgCheck(required(tokenArg), newArg("test", nil)),
		nil,
	)
	mux.HandleFunc(
		"/api/auth.test",
		ArgCheck(required(tokenArg)),
		nil,
	)
	mux.HandleFunc(
		"/api/channels.archive",
		ArgCheck(required(tokenArg), required(channelArg)),
		nil,
	)
	mux.HandleFunc(
		"/api/channels.create",
		ArgCheck(required(tokenArg), required(nameArg), newArg("validate", nil)),
		nil,
	)
	mux.HandleFunc(
		"/api/channels.history",
		ArgCheck(required(tokenArg), required(channelArg), intArg("count"), newArg("includesive", nil), newArg("latest", nil), newArg("oldest", nil), newArg("ts", nil), newArg("unreads", nil)),
		nil,
	)
	mux.HandleFunc(
		"/api/channels.info",
		ArgCheck(required(tokenArg), required(channelArg), newArg("include_locale", nil)),
		nil,
	)
	mux.HandleFunc(
		"/api/channels.invite",
		ArgCheck(required(tokenArg), required(channelArg), required(userArg)),
		nil,
	)
	mux.HandleFunc(
		"/api/channels.kick",
		ArgCheck(required(tokenArg), required(channelArg), required(userArg)),
		nil,
	)
	mux.HandleFunc(
		"/api/channels.leave",
		ArgCheck(required(tokenArg), required(channelArg)),
		nil,
	)
	mux.HandleFunc(
		"/api/channels.list",
		ArgCheck(required(tokenArg), newArg("exclude_archived", nil), newArg("exclude_members", nil), newArg("limit", nil)),
		nil,
	)
	mux.HandleFunc(
		"/api/channels.mark",
		ArgCheck(required(tokenArg), required(channelArg), newArg("ts", nil)),
		nil,
	)
	mux.HandleFunc(
		"/api/channels.rename",
		ArgCheck(required(tokenArg), required(channelArg), required(newArg("name", nil)), newArg("validate", nil)),
		nil,
	)
	mux.HandleFunc(
		"/api/channels.replies",
		ArgCheck(required(tokenArg), required(channelArg), required(newArg("thread_ts", nil))),
		nil,
	)
	mux.HandleFunc(
		"/api/channels.setTopic",
		ArgCheck(required(tokenArg), required(channelArg), required(newArg("topic", nil))),
		nil,
	)
	mux.HandleFunc(
		"/api/channels.unarchive",
		ArgCheck(required(tokenArg), required(channelArg)),
		nil,
	)
	mux.HandleFunc(
		"/api/emoji.list",
		ArgCheck(required(tokenArg)),
		nil,
	)

	// groups
	mux.HandleFunc(
		"/api/groups.archive",
		ArgCheck(required(tokenArg), required(channelArg)),
		nil,
	)
	mux.HandleFunc(
		"/api/groups.create",
		ArgCheck(required(tokenArg), required(newArg("name", nil)), newArg("validate", nil)),
		nil,
	)
	mux.HandleFunc(
		"/api/groups.createChild",
		ArgCheck(required(tokenArg), required(channelArg)),
		nil,
	)
	mux.HandleFunc(
		"/api/groups.history",
		ArgCheck(required(tokenArg), required(channelArg), newArg("count", nil), newArg("inclusive", nil), newArg("latest", nil), newArg("oldest", nil), newArg("unreads", nil)),
		nil,
	)
	mux.HandleFunc(
		"/api/groups.info",
		ArgCheck( required(tokenArg), required(channelArg), newArg("include_locale", nil) ),
		nil,
	)
	mux.HandleFunc(
		"/api/groups.invite",
		ArgCheck( required(tokenArg), required(channelArg), required(newArg("user", nil)) ),
		nil,
	)
	mux.HandleFunc(
		"/api/groups.kick",
		ArgCheck( required(tokenArg), required(channelArg), required(newArg("user", nil))),
		nil,
	)
	mux.HandleFunc(
		"/api/groups.leave",
		ArgCheck( required(tokenArg), required(channelArg)),
		nil,
	)
	mux.HandleFunc(
		"/api/groups.list",
		ArgCheck( required(tokenArg), newArg("exclude_archived", nil), newArg("exclude_members", nil)),
		nil,
	)
	mux.HandleFunc(
		"/api/groups.mark",
		ArgCheck( required(tokenArg), required(channelArg), required(newArg("ts", nil)) ),
		nil,
	)
	mux.HandleFunc(
		"/api/groups.open",
		ArgCheck( required(tokenArg), required(channelArg) ),
		nil,
	)
	mux.HandleFunc(
		"/api/groups.rename",
		ArgCheck( required(tokenArg), required(channelArg), required(newArg("name", nil)), newArg("validate", nil)),
		nil,
	)
	mux.HandleFunc(
		"/api/groups.replies",
		ArgCheck( required(tokenArg), required(channelArg), newArg("thread_ts", nil) ),
		nil,
	)
	mux.HandleFunc(
		"/api/groups.setPurpose",
		ArgCheck( required(tokenArg), required(channelArg), newArg("purpose", nil) ),
		nil,
	)
	mux.HandleFunc(
		"/api/groups.setTopic",
		ArgCheck( required(tokenArg), required(channelArg), newArg("topic", nil) ),
		nil,
	)
	mux.HandleFunc(
		"/api/groups.unarchive",
		ArgCheck( required(tokenArg), required(channelArg) ),
		nil,
	)

	mux.HandleFunc(
		"/api/oauth.access",
		ArgCheck(required(newArg("client_id", nil)), required(newArg("client_secret", nil)), required(newArg("code", nil)), newArg("redirect_uri", nil)),
		nil,
	)
	mux.HandleFunc(
		"/api/reactions.add",
		ArgCheck( required(tokenArg), required(nameArg), channelArg, newArg("file", nil), newArg("fileComment", nil), newArg("timestamp", nil) ),
		nil,
	)
	mux.HandleFunc(
		"/api/reactions.get",
		ArgCheck( required(tokenArg), channelArg, newArg("file", nil), newArg("fileComment", nil), newArg("timestamp", nil), newArg("full", nil)),
		nil,
	)
	mux.HandleFunc(
		"/api/reactions.list",
		ArgCheck( required(tokenArg), channelArg, newArg("user", nil), newArg("full", nil), newArg("count", nil), newArg("page", nil)),
		nil,
	)
	mux.HandleFunc(
		"/api/reactions.remove",
		ArgCheck( required(tokenArg), required(nameArg), channelArg, newArg("file", nil), newArg("fileComment", nil), newArg("timestamp", nil) ),
		nil,
	)
	mux.HandleFunc(
		"/api/rtm.start",
		ArgCheck(required(tokenArg)),
		nil,
	)

	// usergroups
	mux.HandleFunc(
		"/api/usergroups.create",
		ArgCheck( required(tokenArg), required(newArg("name", nil)), newArg("channels", nil), newArg("description", nil), newArg("handle", nil), newArg("include_count", nil) ),
		nil,
	)
	mux.HandleFunc(
		"/api/usergroups.disable",
		ArgCheck( required(tokenArg), required(newArg("usergroup", nil)), newArg("include_count", nil) ),
		nil,
	)
	mux.HandleFunc(
		"/api/usergroups.enable",
		ArgCheck( required(tokenArg), required(newArg("usergroup", nil)), newArg("include_count", nil) ),
		nil,
	)
	mux.HandleFunc(
		"/api/usergroups.list",
		ArgCheck( required(tokenArg), newArg("include_count", nil), newArg("include_disabled", nil), newArg("include_users", nil) ),
		nil,
	)
	mux.HandleFunc(
		"/api/usergroups.update",
		ArgCheck( required(tokenArg), required(newArg("usergroup", nil)), newArg("channels", nil), newArg("description", nil), newArg("handle", nil), newArg("include_count", nil), newArg("name", nil) ),
		nil,
	)

	// usergroups.users
	mux.HandleFunc(
		"/api/usergroups.users.list",
		ArgCheck( required(tokenArg), required(newArg("usergroup", nil)), newArg("include_disabled", nil)),
		nil,
	)
	mux.HandleFunc(
		"/api/usergroups.users.update",
		ArgCheck( required(tokenArg), required(newArg("usergroup", nil)), required(newArg("users", nil)), newArg("include_count", nil)),
		nil,
	)

	mux.HandleFunc(
		"/api/users.deletePhoto",
		ArgCheck(required(tokenArg)),
		nil,
	)
	mux.HandleFunc(
		"/api/users.getPresence",
		ArgCheck(required(tokenArg), required(userArg)),
		nil,
	)
	mux.HandleFunc(
		"/api/users.identity",
		ArgCheck(required(tokenArg)),
		nil,
	)
	mux.HandleFunc(
		"/api/users.info",
		ArgCheck(required(tokenArg), required(userArg), newArg("include_locale", nil)),
		slack.BuildUsersInfoCallResponse().OK(true).User(mockserver.UserLukeSkywalker).Build(),
	)
	mux.HandleFunc(
		"/api/users.list",
		ArgCheck( required(tokenArg), newArg("presence", nil) ),
		nil,
	)
	mux.HandleFunc(
		"/api/users.profile.get",
		ArgCheck( required(tokenArg), userArg, newArg("include_labels", nil) ),
		nil,
	)
	mux.HandleFunc(
		"/api/users.profile.set",
		ArgCheck( required(tokenArg), userArg, newArg("profile", nil), newArg("name", nil), newArg("value", nil) ),
		nil,
	)

	mux.ServeMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"error": "handler does not implement %s"}`, r.URL.Path)
	})
	s.mux = mux
	return &s
}

func newSlackWithDummy(s *httptest.Server) *slack.Client {
	return slack.New("random-token", slack.WithAPIEndpoint(s.URL+"/api/"))
}

// testClient is a test helper for creating a new test server, context, client,
// and closing function.
func testClient(tb testing.TB) (context.Context, *slack.Client, func()) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	s := httptest.NewServer(newDummyServer())
	client := newSlackWithDummy(s)

	return ctx, client, func() {
		cancel()
		s.Close()
	}
}
