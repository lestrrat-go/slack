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
	mux *http.ServeMux
}

type expectedArg struct {
	name  string
	check func([]string) error
}

func nilcheck(_ []string) error { return nil }
func newArg(name string, check func([]string) error) *expectedArg {
	if check == nil {
		check = nilcheck
	}
	return &expectedArg{
		name:  name,
		check: check,
	}
}

func (s *dummyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func requireArguments(w http.ResponseWriter, r *http.Request, args ...*expectedArg) error {
	r.ParseForm()
	for _, arg := range args {
		if err := arg.check(r.Form[arg.name]); err != nil {
			return errors.Wrapf(err, `validation for argument %s failed`, arg.name)
		}
	}
	return nil
}

func makeArgCheckHandler(args ...*expectedArg) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := requireArguments(w, r, args...); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, `{"error": %s}`, strconv.Quote(err.Error()))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, `{"ok":true}`)
	})
}

func makeExactlyOneCheck(name string, allowEmpty bool) func([]string)error {
	return func(in []string) error {
		if len(in) != 1 {
			return errors.Errorf("expected 1 %s (got %d)", name, len(in))
		}
		if !allowEmpty && len(in[0]) <= 0 {
			return errors.Errorf("empty %s not allowed", name)
		}
		return nil
	}
}

func newDummyServer() *dummyServer {
	var s dummyServer
	tokenArg := newArg("token", makeExactlyOneCheck("token", false))
	channelIDArg := newArg("channel", makeExactlyOneCheck("channel", false))
	userIDArg := newArg("user", makeExactlyOneCheck("user", false))
	mux := http.NewServeMux()
	mux.HandleFunc(
		"/api/channels.kick",
		makeArgCheckHandler(tokenArg, channelIDArg, userIDArg),
	)
	s.mux = mux
	return &s
}

func newSlackWithDummy(token string, s *httptest.Server) *slack.Client {
	return slack.New(token, slack.WithAPIEndpoint(s.URL+"/api/"))
}
