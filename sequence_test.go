package slack_test

import (
	"testing"

	slack "github.com/lestrrat-go/slack"
	"github.com/stretchr/testify/assert"
)

func TestParseControlSequence(t *testing.T) {
	data := []struct {
		Text     string
		Error    bool
		Expected slack.ControlSequence
	}{
		{
			Text:     `<@U12345678|bob>`,
			Expected: &slack.UserLink{ID: `@U12345678`, Username: `bob`},
		},
		{
			Text:  `<foo|bar`,
			Error: true,
		},
	}

	for _, testcase := range data {
		t.Run(testcase.Text, func(t *testing.T) {
			l, err := slack.ParseControlSequence(testcase.Text)
			if testcase.Error {
				if !assert.Error(t, err, `should fail to parse sequence`) {
					return
				}
				return
			}

			if !assert.NoError(t, err, `failed to parse sequences`) {
				return
			}

			if !assert.Equal(t, testcase.Expected, l, `sequences should match`) {
				return
			}
		})
	}
}

func TestExtractControlSequences(t *testing.T) {
	data := []struct {
		Text     string
		Error    bool
		Expected []slack.ControlSequence
	}{
		{
			Text: `Why not join <#C024BE7LR|general>?`,
			Expected: []slack.ControlSequence{
				&slack.ChannelLink{ID: `#C024BE7LR`, Channel: `general`},
			},
		},
		{
			Text: `Hey <@U024BE7LH|bob>, did you see my file?`,
			Expected: []slack.ControlSequence{
				&slack.UserLink{ID: `@U024BE7LH`, Username: `bob`},
			},
		},
		{
			Text: `This message contains a URL <http://foo.com/>`,
			Expected: []slack.ControlSequence{
				&slack.ExternalLink{URL: `http://foo.com/`, Text: `http://foo.com/`},
			},
		},
		{
			Text: `So does this one: <http://www.foo.com|www.foo.com>`,
			Expected: []slack.ControlSequence{
				&slack.ExternalLink{URL: `http://www.foo.com`, Text: `www.foo.com`},
			},
		},
		{
			Text: `<mailto:bob@example.com|Bob>`,
			Expected: []slack.ControlSequence{
				&slack.ExternalLink{URL: `mailto:bob@example.com`, Text: `Bob`},
			},
		},
	}

	for _, testcase := range data {
		t.Run(testcase.Text, func(t *testing.T) {
			l, err := slack.ExtractControlSequences(testcase.Text)
			if testcase.Error {
				if !assert.Error(t, err, `should fail to extract sequences`) {
					return
				}
				return
			}

			if !assert.NoError(t, err, `failed to extract sequences`) {
				return
			}

			if !assert.Equal(t, testcase.Expected, l, `sequences should match`) {
				return
			}
		})
	}
}
