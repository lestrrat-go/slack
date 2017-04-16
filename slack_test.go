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

func checkArgs(w http.ResponseWriter, r *http.Request, requiredList []*expectedArg, optionalList []*expectedArg) error {
	r.ParseForm()
	f := r.Form
	for _, arg := range requiredList {
		v, ok := f[arg.name]
		if !ok {
			return errors.Errorf(`argument %s is required`, arg.name)
		}

		if err := arg.check(v); err != nil {
			return errors.Wrapf(err, `validation for required argument %s failed`, arg.name)
		}
		delete(f, arg.name)
	}
	for _, arg := range optionalList {
		if v, ok := f[arg.name]; ok {
			if err := arg.check(v); err != nil {
				return errors.Wrapf(err, `validation for optional argument %s failed`, arg.name)
			}
		}
		delete(f, arg.name)
	}

	for fk := range f {
		return errors.Errorf(`extra argument %s found`, fk)
	}
	return nil
}

type argcheck struct {
	mux  *http.ServeMux
	path string
	req  []*expectedArg
	opt  []*expectedArg
}

func newArgCheck(mux *http.ServeMux, path string) *argcheck {
	return &argcheck{
		mux:  mux,
		path: path,
	}
}

func (c *argcheck) required(args ...*expectedArg) *argcheck {
	c.req = append(c.req, args...)
	return c
}

func (c *argcheck) optional(args ...*expectedArg) *argcheck {
	c.opt = append(c.opt, args...)
	return c
}

func (c *argcheck) do() {
	required := c.req
	optional := c.opt
	c.mux.HandleFunc(c.path, func(w http.ResponseWriter, r *http.Request) {
		if err := checkArgs(w, r, required, optional); err != nil {
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

func makeExactlyOneCheck(name string, allowEmpty bool) *expectedArg {
	return newArg(name, func(in []string) error {
		if len(in) != 1 {
			return errors.Errorf("expected 1 %s (got %d)", name, len(in))
		}
		if !allowEmpty && len(in[0]) <= 0 {
			return errors.Errorf("empty %s not allowed", name)
		}
		return nil
	})
}

func makeIntCheck(name string) *expectedArg {
	return newArg(name, func(in []string) error {
		if len(in) != 1 {
			return errors.Errorf("expected 1 %s (got %d)", name, len(in))
		}
		_, err := strconv.ParseInt(in[0], 10, 64)
		return err
	})
}

func newDummyServer() *dummyServer {
	var s dummyServer
	tokenArg := makeExactlyOneCheck("token", false)
	channelIDArg := makeExactlyOneCheck("channel", false)
	nameArg := makeExactlyOneCheck("name", false)
	userIDArg := makeExactlyOneCheck("user", false)
	mux := http.NewServeMux()
	newArgCheck(mux, "/api/channels.archive").required(tokenArg, channelIDArg).do()
	newArgCheck(mux, "/api/channels.create").required(tokenArg, nameArg).optional(makeExactlyOneCheck("validate", true)).do()
	newArgCheck(mux, "/api/channels.kick").required(tokenArg, channelIDArg, userIDArg).do()
	newArgCheck(mux, "/api/channels.history").
		required(tokenArg, channelIDArg).
		optional(makeIntCheck("count"), makeExactlyOneCheck("inclusive", true), makeExactlyOneCheck("latest", true), makeExactlyOneCheck("oldest", true), makeExactlyOneCheck("ts", true), makeExactlyOneCheck("unreads", true)).
		do()

//	_, err := c.Channels().History("foo").Count(100).Inclusive(true).Latest("dummy").Oldest("dummy").Timestamp("dummy").Unreads(true).Do(ctx)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
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
