package slack_test

import (
	"context"
	"os"
	"testing"
	"time"

	slack "github.com/lestrrat-go/slack"
)

func TestUsergroups_Integration(t *testing.T) {
	if v := os.Getenv("SLACK_TOKEN"); v == "" {
		t.Skip("Set SLACK_TOKEN to run integration tests")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client := slack.New(slackToken)

	list, err := client.Usergroups().List().Do(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if len(list) < 1 {
		t.Fatalf("expected more than one group")
	}
}

func TestUsergroups_Create(t *testing.T) {
	ctx, client, closer := testClient(t)
	defer closer()

	if _, err := client.Usergroups().Create("usergroup").Do(ctx); err != nil {
		t.Fatal(err)
	}

	if _, err := client.Usergroups().Create("usergroup").
		Channels("C1234567890,C2345678901,C3456789012").
		Description("desc").
		Handle("my-usergroup").
		IncludeCount(true).
		Do(ctx); err != nil {
		t.Fatal(err)
	}
}

func TestUsergroups_Disable(t *testing.T) {
	ctx, client, closer := testClient(t)
	defer closer()

	if _, err := client.Usergroups().Disable("usergroup").Do(ctx); err != nil {
		t.Fatal(err)
	}

	if _, err := client.Usergroups().Disable("usergroup").
		IncludeCount(true).
		Do(ctx); err != nil {
		t.Fatal(err)
	}
}

func TestUsergroups_Enable(t *testing.T) {
	ctx, client, closer := testClient(t)
	defer closer()

	if _, err := client.Usergroups().Enable("usergroup").Do(ctx); err != nil {
		t.Fatal(err)
	}

	if _, err := client.Usergroups().Enable("usergroup").
		IncludeCount(true).
		Do(ctx); err != nil {
		t.Fatal(err)
	}
}

func TestUsergroups_List(t *testing.T) {
	ctx, client, closer := testClient(t)
	defer closer()

	if _, err := client.Usergroups().List().Do(ctx); err != nil {
		t.Fatal(err)
	}

	if _, err := client.Usergroups().List().
		IncludeCount(true).
		IncludeDisabled(true).
		IncludeUsers(true).
		Do(ctx); err != nil {
		t.Fatal(err)
	}
}

func TestUsergroups_Update(t *testing.T) {
	ctx, client, closer := testClient(t)
	defer closer()

	if _, err := client.Usergroups().Update("usergroup").Do(ctx); err != nil {
		t.Fatal(err)
	}

	if _, err := client.Usergroups().Update("usergroup").
		Channels("C1234567890,C2345678901,C3456789012").
		Description("desc").
		Handle("my-usergroup").
		IncludeCount(true).
		Name("new-usergroup").
		Do(ctx); err != nil {
		t.Fatal(err)
	}
}
