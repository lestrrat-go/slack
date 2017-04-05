package slack

import (
	"context"
	"net/url"
	"strings"

	"github.com/lestrrat/go-slack/objects"
	"github.com/pkg/errors"
)

// Update returns the result of chat.update API
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

// Delete returns the result of chat.delete API
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

// MeMessage returns the result of users.meMessage API
func (s *ChatService) MeMessage(ctx context.Context, channel, text string) (*ChatResponse, error) {
	v := url.Values{
		"token":   {s.token},
		"channel": {channel},
		"text":    {text},
	}
	const endpoint = "chat.meMessage"
	var res fullChatResponse
	if err := s.client.postForm(ctx, endpoint, v, &res); err != nil {
		return nil, errors.Wrap(err, `failed to post to chat.meMessage`)
	}

	if !res.OK {
		return nil, errors.New(res.Error)
	}

	return res.ChatResponse, nil
}

var replacer = strings.NewReplacer("&", "&amp;", "<", "&lt;", ">", "&gt;")

func escapeMessage(s string) string {
	return replacer.Replace(s)
}

type fullChatResponse struct {
	SlackResponse
	*ChatResponse
}

// PostMessage returns the result of chat.postMessage API
func (s *ChatService) PostMessage(ctx context.Context, p *objects.MessageParams) (*ChatResponse, error) {
	v := url.Values{
		"token":   {s.token},
		"channel": {p.Channel},
	}

	var txt = p.Text
	if p != nil {
		if p.EscapeText {
			txt = escapeMessage(txt)
		}

		if err := p.Values(v); err != nil {
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
