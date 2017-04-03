package slack_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/lestrrat/go-slack"
)

var slackToken string
var dmUser string

func init() {
	slackToken = os.Getenv("SLACK_TOKEN")
	dmUser = os.Getenv("TEST_DM_USER") // don't forget to include an "@"
}

func hasTestSlackToken(t *testing.T) bool {
	if slackToken == "" {
		t.Skip("SLACK_TOKEN not available")
		return false
	}
	return true
}

func hasTestDMUser(t *testing.T) bool {
	if dmUser == "" {
		t.Skip("TEST_DM_USER not available")
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
