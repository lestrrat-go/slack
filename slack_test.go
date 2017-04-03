package slack_test

import (
	"context"
	"crypto/rand"
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/lestrrat/go-slack"
	"golang.org/x/oauth2"
)

var dmUser string
var isBot bool
var slackToken string

func init() {
	slackToken = os.Getenv("SLACK_TOKEN")
	dmUser = os.Getenv("TEST_DM_USER") // don't forget to include an "@"

	if len(slackToken) > 0 {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		cl := slack.New(slackToken)
		res, err := cl.Auth().Test(ctx)
		if err == nil {
			user, err := cl.Users().Info(ctx, res.UserID)
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
	if dmUser == "" {
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

func ExampleClient() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	token := os.Getenv("SLACK_TOKEN")
	cl := slack.New(token)

	// check if we are connected
	authres, err := cl.Auth().Test(ctx)
	if err != nil {
		fmt.Printf("failed to test authentication: %s\n", err)
		return
	}
	fmt.Printf("%#v\n", authres)

	// simplest possible message
	chatres, err := cl.Chat().PostMessage(ctx, "@username", "Hello, World!", nil)
	if err != nil {
		fmt.Printf("failed to post messsage: %s\n", err)
		return
	}
	fmt.Printf("%#v\n", chatres)
}

func ExampleOAuth2() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// When you create a Slack App, you need to authorize your app through OAuth2
	//
	// If you installed your app via the Web UI, you should be able to see the
	// tokens generated when you did so at https://api.slack.com/apps/XXXXXX/oauth
	// where XXXXXX is a random ID generated for your app.
	//
	// You could used these tokens, or you can do a manual OAuth2 flow, which is
	// shown in pseudo-working form below. (note: most it just straight oauth2
	// taken from https://godoc.org/golang.org/x/oauth2#example-Config)
	// However, Slack does not allow offline flow, so you will need to actually
	// run this in a webserver, unlike the example in the above URL.

	conf := oauth2.Config{
		ClientID:     os.Getenv("SLACK_APP_CLIENT_ID"),
		ClientSecret: os.Getenv("SLACK_APP_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("SLACK_APP_REDIRECT_URL"),
		Scopes: []string{
			slack.ChatWriteBotScope,
		},
		Endpoint: slack.OAuth2Endpoint,
	}

	http.HandleFunc("/oauth/start", func(w http.ResponseWriter, r *http.Request) {
		// Poor man's UUID
		b := make([]byte, 16)
		rand.Reader.Read(b)
		b[6] = (b[6] & 0x0F) | 0x40
		b[8] = (b[8] &^ 0x40) | 0x80
		state := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
		// TODO: Use session or whatever to save "state", so the user
		// can be verified

		// Redirect user to consent page to ask for permission
		// for the scopes specified above.
		url := conf.AuthCodeURL(state)
		w.Header().Set("Location", url)
		w.WriteHeader(http.StatusFound)
	})

	http.HandleFunc("/oauth/callback", func(w http.ResponseWriter, r *http.Request) {
		code := r.FormValue("code")

		// TODO: Use session or whatever to restore "state", so the user
		// can be verified
		tok, err := conf.Exchange(ctx, code)
		if err != nil {
			http.Error(w, "failed to exchange tokens", http.StatusInternalServerError)
			return
		}

		// You could store tok.AccessToken for later use, or you can immediately
		// start a client like this
		cl := slack.New(tok.AccessToken)
		if _, err := cl.Auth().Test(ctx); err != nil {
			http.Error(w, "failed to test auth", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Contenxt-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Successfully connected to Slack"))
	})

	http.ListenAndServe(":8080", nil)
}
