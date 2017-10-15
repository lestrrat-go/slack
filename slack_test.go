package slack_test

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/lestrrat/go-slack"
	"github.com/pkg/errors"
)

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
			user, err := cl.Users().Info(res.UserID).Do(ctx)
			if err == nil {
				isBot = user.IsBot
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

func (m *mux) HandleFunc(path string, args ...*expectedArg) {
	checker := func(r *http.Request) error {
		for _, arg := range args {
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
	f := func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		if err := checker(r); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, `{"error": %s}`, strconv.Quote(err.Error()))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, `{"ok":true}`)
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
	mux.HandleFunc("/api/auth.revoke", required(tokenArg), newArg("test", nil))
	mux.HandleFunc("/api/auth.test", required(tokenArg))
	mux.HandleFunc("/api/channels.archive", required(tokenArg), required(channelArg))
	mux.HandleFunc("/api/channels.create", required(tokenArg), required(nameArg), newArg("validate", nil))
	mux.HandleFunc("/api/channels.history", required(tokenArg), required(channelArg), intArg("count"), newArg("includesive", nil), newArg("latest", nil), newArg("oldest", nil), newArg("ts", nil), newArg("unreads", nil))
	mux.HandleFunc("/api/channels.info", required(tokenArg), required(channelArg), newArg("include_locale", nil))
	mux.HandleFunc("/api/channels.invite", required(tokenArg), required(channelArg), required(userArg))
	mux.HandleFunc("/api/channels.kick", required(tokenArg), required(channelArg), required(userArg))
	mux.HandleFunc("/api/channels.leave", required(tokenArg), required(channelArg))
	mux.HandleFunc("/api/channels.list", required(tokenArg), newArg("exclude_archived", nil), newArg("exclude_members", nil), newArg("limit", nil))
	mux.HandleFunc("/api/channels.mark", required(tokenArg), required(channelArg), newArg("ts", nil))
	mux.HandleFunc("/api/channels.rename", required(tokenArg), required(channelArg), required(newArg("name", nil)), newArg("validate", nil))
	mux.HandleFunc("/api/channels.replies", required(tokenArg), required(channelArg), required(newArg("thread_ts", nil)))
	mux.HandleFunc("/api/channels.setTopic", required(tokenArg), required(channelArg), required(newArg("topic", nil)))
	mux.HandleFunc("/api/channels.unarchive", required(tokenArg), required(channelArg))
	mux.HandleFunc("/api/emoji.list", required(tokenArg))
	mux.HandleFunc("/api/oauth.access",
		required(newArg("client_id", nil)),
		required(newArg("client_secret", nil)),
		required(newArg("code", nil)),
		newArg("redirect_uri", nil),
	)
	mux.HandleFunc("/api/reactions.add",
		required(tokenArg),
		required(nameArg),
		channelArg,
		newArg("file", nil),
		newArg("fileComment", nil),
		newArg("timestamp", nil),
	)
	mux.HandleFunc("/api/reactions.get",
		required(tokenArg),
		channelArg,
		newArg("file", nil),
		newArg("fileComment", nil),
		newArg("timestamp", nil),
		newArg("full", nil),
	)
	mux.HandleFunc("/api/reactions.list",
		required(tokenArg),
		channelArg,
		newArg("user", nil),
		newArg("full", nil),
		newArg("count", nil),
		newArg("page", nil),
	)
	mux.HandleFunc("/api/reactions.remove",
		required(tokenArg),
		required(nameArg),
		channelArg,
		newArg("file", nil),
		newArg("fileComment", nil),
		newArg("timestamp", nil),
	)
	mux.HandleFunc("/api/rtm.start", required(tokenArg))
	mux.HandleFunc("/api/users.deletePhoto", required(tokenArg))
	mux.HandleFunc("/api/users.getPresence",
		required(tokenArg),
		required(userArg),
	)
	mux.HandleFunc("/api/users.identity", required(tokenArg))
	mux.HandleFunc("/api/users.info",
		required(tokenArg),
		required(userArg),
		newArg("include_locale", nil),
	)
	mux.HandleFunc("/api/users.list",
		required(tokenArg),
		newArg("presence", nil),
	)
	mux.HandleFunc("/api/users.profile.get",
		required(tokenArg),
		userArg,
		newArg("include_labels", nil),
	)
	mux.HandleFunc("/api/users.profile.set",
		required(tokenArg),
		userArg,
		newArg("profile", nil),
		newArg("name", nil),
		newArg("value", nil),
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
