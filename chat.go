package slack

import (
	"bytes"
	"context"
	"encoding/json"
	"net/url"
	"strings"

	"github.com/lestrrat/go-slack/objects"
	"github.com/pkg/errors"
)

type ChatUpdateCall struct {
	service   *ChatService
	channel   string
	timestamp string
	text      string
}

// Update returns the result of chat.update API
func (s *ChatService) Update(channel, text, ts string) *ChatUpdateCall {
	return &ChatUpdateCall{
		service:   s,
		channel:   channel,
		text:      text,
		timestamp: ts,
	}
}

func (c *ChatUpdateCall) Values() url.Values {
	v := url.Values{
		"token":   {c.service.token},
		"text":    {c.text},
		"channel": {c.channel},
		"ts":      {c.timestamp},
	}
	return v
}

func (c *ChatUpdateCall) Do(ctx context.Context) (*ChatResponse, error) {
	const endpoint = "chat.update"
	var res fullChatResponse
	if err := c.service.client.postForm(ctx, endpoint, c.Values(), &res); err != nil {
		return nil, errors.Wrap(err, `failed to post to chat.delete`)
	}

	if !res.OK {
		return nil, errors.New(res.Error.String())
	}

	return res.ChatResponse, nil
}

type ChatDeleteCall struct {
	service   *ChatService
	channel   string
	timestamp string
	asUser    bool
}

// Delete returns the result of chat.delete API
func (s *ChatService) Delete(channel, ts string) *ChatDeleteCall {
	return &ChatDeleteCall{
		service:   s,
		channel:   channel,
		timestamp: ts,
	}
}

func (c *ChatDeleteCall) Values() url.Values {
	v := url.Values{
		"token":   {c.service.token},
		"channel": {c.channel},
		"ts":      {c.timestamp},
	}

	if c.asUser {
		v.Set("as_user", "true")
	}
	return v
}

func (c *ChatDeleteCall) AsUser(b bool) *ChatDeleteCall {
	c.asUser = b
	return c
}

func (c *ChatDeleteCall) Do(ctx context.Context) (*ChatResponse, error) {
	const endpoint = "chat.delete"
	var res fullChatResponse
	if err := c.service.client.postForm(ctx, endpoint, c.Values(), &res); err != nil {
		return nil, errors.Wrap(err, `failed to post to chat.delete`)
	}

	if !res.OK {
		return nil, errors.New(res.Error.String())
	}

	return res.ChatResponse, nil
}

type ChatMeMessageCall struct {
	service *ChatService
	channel string
	text    string
}

// MeMessage returns the result of users.meMessage API
func (s *ChatService) MeMessage(channel, text string) *ChatMeMessageCall {
	return &ChatMeMessageCall{
		service: s,
		channel: channel,
		text:    text,
	}
}

func (c *ChatMeMessageCall) Values() url.Values {
	v := url.Values{
		"token":   {c.service.token},
		"channel": {c.channel},
		"text":    {c.text},
	}
	return v
}

func (c *ChatMeMessageCall) Do(ctx context.Context) (*ChatResponse, error) {
	const endpoint = "chat.meMessage"
	var res fullChatResponse
	if err := c.service.client.postForm(ctx, endpoint, c.Values(), &res); err != nil {
		return nil, errors.Wrap(err, `failed to post to chat.meMessage`)
	}

	if !res.OK {
		return nil, errors.New(res.Error.String())
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

type ChatPostMessageCall struct {
	service     *ChatService
	asUser      bool
	attachments objects.AttachmentList
	channel     string
	escapeText  bool
	iconEmoji   string
	iconURL     string
	linkNames   bool
	markdown    bool
	parse       string
	text        string
	unfurlLinks bool
	unfurlMedia bool
	username    string
}

// PostMessage returns the result of chat.postMessage API
func (s *ChatService) PostMessage(channel string) *ChatPostMessageCall {
	return &ChatPostMessageCall{
		service:     s,
		channel:     channel,
		escapeText:  true,
		markdown:    true,
		unfurlMedia: true,
	}
}

func (c *ChatPostMessageCall) Values() (url.Values, error) {
	v := url.Values{
		"token":   {c.service.token},
		"channel": {c.channel},
	}

	if c.asUser {
		v.Set("as_user", "true")
	}

	if len(c.attachments) > 0 {
		var buf bytes.Buffer
		if err := json.NewEncoder(&buf).Encode(c.attachments); err != nil {
			return nil, errors.Wrap(err, `failed to serialize attachments`)
		}
		v.Set("attachments", buf.String())
	}

	if len(c.iconEmoji) > 0 {
		v.Set("icon_emoji", c.iconEmoji)
	}
	if len(c.iconURL) > 0 {
		v.Set("icon_url", c.iconURL)
	}
	if c.linkNames {
		v.Set("link_names", "true")
	}

	if !c.markdown {
		v.Set("mrkdwn", "false")
	}

	if len(c.parse) > 0 {
		v.Set("parse", c.parse)
	}

	// taken from github.com/nlopes/slack:
	//    I want to send a message with explicit `as_user` `true` and
	//    `unfurl_links` `false` in request. Because setting `as_user` to
	//    `true` will change the default value for `unfurl_links` to `true`
	//    on Slack API side.
	if c.asUser && !c.unfurlLinks {
		v.Set("unfurl_link", "false")
	} else if c.unfurlLinks {
		v.Set("unfurl_link", "true")
	}

	if c.unfurlMedia {
		v.Set("unfurl_media", "true")
	}
	if len(c.username) > 0 {
		v.Set("username", c.username)
	}

	var txt = c.text
	if c.escapeText {
		txt = escapeMessage(txt)
	}
	v.Set("text", txt)

	return v, nil
}

func (c *ChatPostMessageCall) AsUser(b bool) *ChatPostMessageCall {
	c.asUser = b
	return c
}

// SetAttachments replaces the attachment list
func (c *ChatPostMessageCall) SetAttachments(l objects.AttachmentList) *ChatPostMessageCall {
	c.attachments = l
	return c
}

// Attachment appends to the attachments
func (c *ChatPostMessageCall) Attachment(a *objects.Attachment) *ChatPostMessageCall {
	c.attachments.Append(a)
	return c
}

func (c *ChatPostMessageCall) EscapeText(b bool) *ChatPostMessageCall {
	c.escapeText = b
	return c
}

func (c *ChatPostMessageCall) IconEmoji(s string) *ChatPostMessageCall {
	c.iconEmoji = s
	return c
}

func (c *ChatPostMessageCall) IconURL(s string) *ChatPostMessageCall {
	c.iconURL = s
	return c
}

func (c *ChatPostMessageCall) LinkNames(b bool) *ChatPostMessageCall {
	c.linkNames = b
	return c
}

func (c *ChatPostMessageCall) Markdown(b bool) *ChatPostMessageCall {
	c.markdown = b
	return c
}

func (c *ChatPostMessageCall) Parse(s string) *ChatPostMessageCall {
	c.parse = s
	return c
}

func (c *ChatPostMessageCall) Text(s string) *ChatPostMessageCall {
	c.text = s
	return c
}

func (c *ChatPostMessageCall) UnfurlLinks(b bool) *ChatPostMessageCall {
	c.unfurlLinks = b
	return c
}

func (c *ChatPostMessageCall) UnfurlMedia(b bool) *ChatPostMessageCall {
	c.unfurlMedia = b
	return c
}

func (c *ChatPostMessageCall) Username(s string) *ChatPostMessageCall {
	c.username = s
	return c
}

func (c *ChatPostMessageCall) Do(ctx context.Context) (*ChatResponse, error) {
	if len(c.channel) <= 0 {
		return nil, errors.New("channel not specified")
	}
	const endpoint = "chat.postMessage"

	v, err := c.Values()
	if err != nil {
		return nil, err
	}

	var res fullChatResponse
	if err := c.service.client.postForm(ctx, endpoint, v, &res); err != nil {
		return nil, errors.Wrap(err, `failed to post to chat.postMessage`)
	}

	if !res.OK {
		return nil, errors.New(res.Error.String())
	}

	return res.ChatResponse, nil
}
