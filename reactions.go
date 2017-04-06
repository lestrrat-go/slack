package slack

import (
	"context"
	"net/url"
	"strconv"

	"github.com/pkg/errors"
)

type ReactionsAddCall struct {
	service     *ReactionsService
	name        string
	file        string
	fileComment string
	channel     string
	timestamp   string
}

func (s *ReactionsService) Add(name string) *ReactionsAddCall {
	return &ReactionsAddCall{
		service: s,
		name:    name,
	}
}

func (c *ReactionsAddCall) Values() url.Values {
	v := url.Values{
		"token": {c.service.token},
		"name":  {c.name},
	}
	if len(c.file) > 0 {
		v.Set("file", c.file)
	}

	if len(c.fileComment) > 0 {
		v.Set("file_comment", c.fileComment)
	}

	if len(c.channel) > 0 {
		v.Set("channel", c.channel)
	}

	if len(c.timestamp) > 0 {
		v.Set("timestamp", c.timestamp)
	}
	return v
}

func (c *ReactionsAddCall) File(s string) *ReactionsAddCall {
	c.file = s
	return c
}

func (c *ReactionsAddCall) FileComment(s string) *ReactionsAddCall {
	c.fileComment = s
	return c
}

func (c *ReactionsAddCall) Channel(s string) *ReactionsAddCall {
	c.channel = s
	return c
}

func (c *ReactionsAddCall) Timestamp(s string) *ReactionsAddCall {
	c.timestamp = s
	return c
}

func (c *ReactionsAddCall) Do(ctx context.Context) error {
	var res SlackResponse

	const endpoint = "reactions.add"
	if err := c.service.client.postForm(ctx, endpoint, c.Values(), &res); err != nil {
		return errors.Wrapf(err, `failed to post to %s`, endpoint)
	}

	if !res.OK {
		return errors.New(res.Error.String())
	}

	return nil
}

type ReactionsGetCall struct {
	service     *ReactionsService
	file        string
	fileComment string
	channel     string
	timestamp   string
	full        bool
}

func (s *ReactionsService) Get() *ReactionsGetCall {
	return &ReactionsGetCall{
		service: s,
	}
}

func (c *ReactionsGetCall) Values() url.Values {
	v := url.Values{
		"token": {c.service.token},
	}
	if len(c.file) > 0 {
		v.Set("file", c.file)
	}

	if len(c.fileComment) > 0 {
		v.Set("file_comment", c.fileComment)
	}

	if len(c.channel) > 0 {
		v.Set("channel", c.channel)
	}

	if len(c.timestamp) > 0 {
		v.Set("timestamp", c.timestamp)
	}

	if c.full {
		v.Set("full", "true")
	}
	return v
}

func (c *ReactionsGetCall) File(s string) *ReactionsGetCall {
	c.file = s
	return c
}

func (c *ReactionsGetCall) FileComment(s string) *ReactionsGetCall {
	c.fileComment = s
	return c
}

func (c *ReactionsGetCall) Channel(s string) *ReactionsGetCall {
	c.channel = s
	return c
}

func (c *ReactionsGetCall) Timestamp(s string) *ReactionsGetCall {
	c.timestamp = s
	return c
}

func (c *ReactionsGetCall) Full(b bool) *ReactionsGetCall {
	c.full = b
	return c
}

func (c *ReactionsGetCall) Do(ctx context.Context) (*ReactionsGetResponse, error) {
	var res struct {
		SlackResponse
		*ReactionsGetResponse
	}

	const endpoint = "reactions.get"
	if err := c.service.client.postForm(ctx, endpoint, c.Values(), &res); err != nil {
		return nil, errors.Wrapf(err, `failed to post to %s`, endpoint)
	}

	if !res.OK {
		return nil, errors.New(res.Error.String())
	}

	return res.ReactionsGetResponse, nil
}

type ReactionsListCall struct {
	service *ReactionsService
	user    string
	full    bool
	count   int
	page    int
}

func (s *ReactionsService) List() *ReactionsListCall {
	return &ReactionsListCall{
		service: s,
	}
}

func (c *ReactionsListCall) Values() url.Values {
	v := url.Values{
		"token": {c.service.token},
	}
	if len(c.user) > 0 {
		v.Set("user", c.user)
	}

	if c.full {
		v.Set("full", "true")
	}

	if c.count > 0 {
		v.Set("count", strconv.Itoa(c.count))
	}

	if c.page > 0 {
		v.Set("page", strconv.Itoa(c.page))
	}
	return v
}

func (c *ReactionsListCall) User(s string) *ReactionsListCall {
	c.user = s
	return c
}

func (c *ReactionsListCall) Full(b bool) *ReactionsListCall {
	c.full = b
	return c
}

func (c *ReactionsListCall) Count(i int) *ReactionsListCall {
	c.count = i
	return c
}

func (c *ReactionsListCall) Page(i int) *ReactionsListCall {
	c.page = i
	return c
}

func (c *ReactionsListCall) Do(ctx context.Context) (*ReactionsListResponse, error) {
	var res struct {
		SlackResponse
		*ReactionsListResponse
	}

	const endpoint = "reactions.list"
	if err := c.service.client.postForm(ctx, endpoint, c.Values(), &res); err != nil {
		return nil, errors.Wrapf(err, `failed to post to %s`, endpoint)
	}

	if !res.OK {
		return nil, errors.New(res.Error.String())
	}

	return res.ReactionsListResponse, nil
}

type ReactionsRemoveCall struct {
	service     *ReactionsService
	name        string
	file        string
	fileComment string
	channel     string
	timestamp   string
}

func (s *ReactionsService) Remove(name string) *ReactionsRemoveCall {
	return &ReactionsRemoveCall{
		service: s,
		name:    name,
	}
}

func (c *ReactionsRemoveCall) Values() url.Values {
	v := url.Values{
		"token": {c.service.token},
		"name":  {c.name},
	}
	if len(c.file) > 0 {
		v.Set("file", c.file)
	}

	if len(c.fileComment) > 0 {
		v.Set("file_comment", c.fileComment)
	}

	if len(c.channel) > 0 {
		v.Set("channel", c.channel)
	}

	if len(c.timestamp) > 0 {
		v.Set("timestamp", c.timestamp)
	}
	return v
}

func (c *ReactionsRemoveCall) File(s string) *ReactionsRemoveCall {
	c.file = s
	return c
}

func (c *ReactionsRemoveCall) FileComment(s string) *ReactionsRemoveCall {
	c.fileComment = s
	return c
}

func (c *ReactionsRemoveCall) Channel(s string) *ReactionsRemoveCall {
	c.channel = s
	return c
}

func (c *ReactionsRemoveCall) Timestamp(s string) *ReactionsRemoveCall {
	c.timestamp = s
	return c
}

func (c *ReactionsRemoveCall) Do(ctx context.Context) error {
	var res SlackResponse

	const endpoint = "reactions.removet"
	if err := c.service.client.postForm(ctx, endpoint, c.Values(), &res); err != nil {
		return errors.Wrapf(err, `failed to post to %s`, endpoint)
	}

	if !res.OK {
		return errors.New(res.Error.String())
	}

	return nil
}


