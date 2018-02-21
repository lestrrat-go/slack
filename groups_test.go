package slack_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/lestrrat-go/slack"
)

func TestGroups_Integration(t *testing.T) {
	if v := os.Getenv("SLACK_TOKEN"); v == "" {
		t.Skip("Set SLACK_TOKEN to run integration tests")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client := slack.New(slackToken)

	list, err := client.Groups().List().Do(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if len(list) < 1 {
		t.Fatalf("expected more than one group")
	}

	group := list[0]

	t.Run("history", func(t *testing.T) {
		t.Parallel()

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		_, hist, err := client.Groups().History(group.ID).Do(ctx)
		if err != nil {
			t.Fatal(err)
		}
		if len(hist) < 1 {
			t.Errorf("expected history")
		}
	})

	t.Run("info", func(t *testing.T) {
		t.Parallel()

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		info, err := client.Groups().Info(group.ID).Do(ctx)
		if err != nil {
			t.Fatal(err)
		}
		if !info.IsGroup {
			t.Errorf("expected to be group")
		}
	})
}

func TestGroups_Archive(t *testing.T) {
	ctx, client, closer := testClient(t)
	defer closer()

	if err := client.Groups().Archive("group").Do(ctx); err != nil {
		t.Fatal(err)
	}
}

func TestGroups_Create(t *testing.T) {
	ctx, client, closer := testClient(t)
	defer closer()

	if _, err := client.Groups().Create("group").Do(ctx); err != nil {
		t.Fatal(err)
	}

	if _, err := client.Groups().Create("group").
		Validate(true).
		Do(ctx); err != nil {
		t.Fatal(err)
	}
}

func TestGroups_CreateChild(t *testing.T) {
	ctx, client, closer := testClient(t)
	defer closer()

	if _, err := client.Groups().CreateChild("child").Do(ctx); err != nil {
		t.Fatal(err)
	}
}

func TestGroups_History(t *testing.T) {
	ctx, client, closer := testClient(t)
	defer closer()

	if _, _, err := client.Groups().History("group").Do(ctx); err != nil {
		t.Fatal(err)
	}

	if _, _, err := client.Groups().History("group").
		Count(100).
		Inclusive(true).
		Latest("12345").
		Oldest("89102").
		Unreads(true).
		Do(ctx); err != nil {
		t.Fatal(err)
	}
}

func TestGroups_Info(t *testing.T) {
	ctx, client, closer := testClient(t)
	defer closer()

	if _, err := client.Groups().Info("group").Do(ctx); err != nil {
		t.Fatal(err)
	}

	if _, err := client.Groups().Info("group").
		IncludeLocale(true).
		Do(ctx); err != nil {
		t.Fatal(err)
	}
}

func TestGroups_Invite(t *testing.T) {
	ctx, client, closer := testClient(t)
	defer closer()

	if _, _, err := client.Groups().Invite("group", "user").Do(ctx); err != nil {
		t.Fatal(err)
	}
}

func TestGroups_Kick(t *testing.T) {
	ctx, client, closer := testClient(t)
	defer closer()

	if err := client.Groups().Kick("group", "user").Do(ctx); err != nil {
		t.Fatal(err)
	}
}

func TestGroups_Leave(t *testing.T) {
	ctx, client, closer := testClient(t)
	defer closer()

	if err := client.Groups().Leave("group").Do(ctx); err != nil {
		t.Fatal(err)
	}
}

func TestGroups_List(t *testing.T) {
	ctx, client, closer := testClient(t)
	defer closer()

	if _, err := client.Groups().List().Do(ctx); err != nil {
		t.Fatal(err)
	}

	if _, err := client.Groups().List().
		ExcludeArchived(true).
		ExcludeMembers(true).
		Do(ctx); err != nil {
		t.Fatal(err)
	}
}

func TestGroups_Mark(t *testing.T) {
	ctx, client, closer := testClient(t)
	defer closer()

	if err := client.Groups().Mark("group", "12345").Do(ctx); err != nil {
		t.Fatal(err)
	}
}

func TestGroups_Open(t *testing.T) {
	ctx, client, closer := testClient(t)
	defer closer()

	if err := client.Groups().Open("group").Do(ctx); err != nil {
		t.Fatal(err)
	}
}

func TestGroups_Rename(t *testing.T) {
	ctx, client, closer := testClient(t)
	defer closer()

	if _, err := client.Groups().Rename("group", "new-group").Do(ctx); err != nil {
		t.Fatal(err)
	}

	if _, err := client.Groups().Rename("group", "new-group").
		Validate(true).
		Do(ctx); err != nil {
		t.Fatal(err)
	}
}

func TestGroups_Replies(t *testing.T) {
	ctx, client, closer := testClient(t)
	defer closer()

	if _, _, err := client.Groups().Replies("group", "12345").Do(ctx); err != nil {
		t.Fatal(err)
	}
}

func TestGroups_SetPurpose(t *testing.T) {
	ctx, client, closer := testClient(t)
	defer closer()

	if _, err := client.Groups().SetPurpose("group", "purpose").Do(ctx); err != nil {
		t.Fatal(err)
	}
}

func TestGroups_SetTopic(t *testing.T) {
	ctx, client, closer := testClient(t)
	defer closer()

	if _, err := client.Groups().SetTopic("group", "topic").Do(ctx); err != nil {
		t.Fatal(err)
	}
}

func TestGroups_Unarchive(t *testing.T) {
	ctx, client, closer := testClient(t)
	defer closer()

	if err := client.Groups().Unarchive("group").Do(ctx); err != nil {
		t.Fatal(err)
	}
}
