package slack

import (
	"bytes"
	"context"
	"encoding/json"
	"net/url"
	"strings"

	"github.com/pkg/errors"
)

func (s *ChatService) Update(ctx context.Context, channel, ts, text string) (*ChatResponse, error) {
	v := url.Values{
		"token":   {s.token},
		"text":    {text},
		"channel": {channel},
		"ts":      {ts},
	}
	const endpoint = "chat.update"

	var res fullChatResponse
	if err := s.client.postForm(ctx, endpoint, v, &res); err != nil {
		return nil, errors.Wrap(err, `failed to post to chat.delete`)
	}

	if !res.OK {
		return nil, errors.New(res.Error)
	}

	return res.ChatResponse, nil
}

func (s *ChatService) Delete(ctx context.Context, channel, ts string) (*ChatResponse, error) {
	v := url.Values{
		"token":   {s.token},
		"channel": {channel},
		"ts":      {ts},
		"as_user": {"true"},
	}
	const endpoint = "chat.delete"
	var res fullChatResponse
	if err := s.client.postForm(ctx, endpoint, v, &res); err != nil {
		return nil, errors.Wrap(err, `failed to post to chat.delete`)
	}

	if !res.OK {
		return nil, errors.New(res.Error)
	}

	return res.ChatResponse, nil
}

func NewMessageParams() *MessageParams {
	return &MessageParams{
		// everything else should be initialzed to the zero value
		EscapeText:  true,
		Markdown:    true,
		UnfurlMedia: true,
	}
}

func (p *MessageParams) toValues(v url.Values) error {
	if p.AsUser {
		v.Set("as_user", "true")
	}
	if len(p.Attachments) > 0 {
		var buf bytes.Buffer
		if err := json.NewEncoder(&buf).Encode(p.Attachments); err != nil {
			return errors.Wrap(err, `failed to serialize attachments`)
		}
		v.Set("attachments", buf.String())
	}
	if len(p.IconEmoji) > 0 {
		v.Set("icon_emoji", p.IconEmoji)
	}
	if len(p.IconURL) > 0 {
		v.Set("icon_url", p.IconURL)
	}
	if p.LinkNames {
		v.Set("link_names", "true")
	}

	if !p.Markdown {
		v.Set("mrkdwn", "false")
	}

	if len(p.Parse) > 0 {
		v.Set("parse", p.Parse)
	}

	// taken from github.com/nlopes/slack:
	//    I want to send a message with explicit `as_user` `true` and
	//    `unfurl_links` `false` in request. Because setting `as_user` to
	//    `true` will change the default value for `unfurl_links` to `true`
	//    on Slack API side.
	if p.AsUser && !p.UnfurlLinks {
		v.Set("unfurl_link", "false")
	} else if p.UnfurlLinks {
		v.Set("unfurl_link", "true")
	}

	if p.UnfurlMedia {
		v.Set("unfurl_media", "true")
	}
	if len(p.Username) > 0 {
		v.Set("username", p.Username)
	}
	return nil
}

var replacer = strings.NewReplacer("&", "&amp;", "<", "&lt;", ">", "&gt;")

func escapeMessage(s string) string {
	return replacer.Replace(s)
}

type fullChatResponse struct {
	SlackResponse
	*ChatResponse
}

func (s *ChatService) PostMessage(ctx context.Context, channel, txt string, p *MessageParams) (*ChatResponse, error) {
	v := url.Values{
		"token":   {s.token},
		"channel": {channel},
	}

	if p != nil {
		if p.EscapeText {
			txt = escapeMessage(txt)
		}

		if err := p.toValues(v); err != nil {
			return nil, errors.Wrap(err, `failed to serialize message`)
		}
	}

	v.Set("text", txt)

	const endpoint = "chat.postMessage"
	var res fullChatResponse
	if err := s.client.postForm(ctx, endpoint, v, &res); err != nil {
		return nil, errors.Wrap(err, `failed to post to chat.postMessage`)
	}

	if !res.OK {
		return nil, errors.New(res.Error)
	}

	return res.ChatResponse, nil
}
