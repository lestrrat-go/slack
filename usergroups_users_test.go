package slack_test

import (
	"context"
	"os"
	"testing"
	"time"

	slack "github.com/lestrrat-go/slack"
)

func TestUsergroupsUsers_Integration(t *testing.T) {
	if v := os.Getenv("SLACK_TOKEN"); v == "" {
		t.Skip("Set SLACK_TOKEN to run integration tests")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client := slack.New(slackToken)

	usergroups, err := client.Usergroups().List().Do(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if len(usergroups) < 1 {
		t.Fatalf("expected more than one group")
	}

	usergroup := usergroups[0]

	list, err := client.UsergroupsUsers().List(usergroup.ID).Do(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if len(list) < 1 {
		t.Fatalf("expected more than one user")
	}
}

func TestUsergroupsUsers_List(t *testing.T) {
	ctx, client, closer := testClient(t)
	defer closer()

	if _, err := client.UsergroupsUsers().List("group").Do(ctx); err != nil {
		t.Fatal(err)
	}

	if _, err := client.UsergroupsUsers().List("group").
		IncludeDisabled(true).
		Do(ctx); err != nil {
		t.Fatal(err)
	}
}

func TestUsergroupsUsers_Update(t *testing.T) {
	ctx, client, closer := testClient(t)
	defer closer()

	if _, err := client.UsergroupsUsers().Update("group", "U060R4BJ4,U060RNRCZ").Do(ctx); err != nil {
		t.Fatal(err)
	}

	if _, err := client.UsergroupsUsers().Update("group", "U060R4BJ4,U060RNRCZ").
		IncludeCount(true).
		Do(ctx); err != nil {
		t.Fatal(err)
	}
}
