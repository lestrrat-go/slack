package slack

import (
	"context"
	"net/url"
	"strconv"

	"github.com/lestrrat/go-slack/objects"
	"github.com/pkg/errors"
)

type ChannelsHistoryCall struct {
	service   *ChannelsService
	channel   string // channel ID
	count     int    // 1-1000
	inclusive bool
	latest    string // range of time (end)
	oldest    string // range of time (start)
	timestamp string // used only when retrieving a single message
	unreads   bool   // Include unread_count_display in the output
}

func (s *ChannelsService) History(id string) *ChannelsHistoryCall {
	return &ChannelsHistoryCall{
		service: s,
		channel: id,
	}
}

func (c *ChannelsHistoryCall) Values() url.Values {
	v := url.Values{
		"token":   {c.service.token},
		"channel": {c.channel},
	}

	if c.count > 0 {
		v.Set("count", strconv.Itoa(c.count))
	}

	if c.inclusive {
		v.Set("inclusive", "1")
	}

	if len(c.latest) > 0 {
		v.Set("latest", c.latest)
	}

	if len(c.oldest) > 0 {
		v.Set("oldest", c.oldest)
	}

	if len(c.timestamp) > 0 {
		v.Set("ts", c.timestamp)
	}

	if c.unreads {
		v.Set("unreads", "1")
	}

	return v
}

func (c *ChannelsHistoryCall) Latest(s string) *ChannelsHistoryCall {
	c.latest = s
	return c
}

func (c *ChannelsHistoryCall) Oldest(s string) *ChannelsHistoryCall {
	c.oldest = s
	return c
}

func (c *ChannelsHistoryCall) Inclusive(b bool) *ChannelsHistoryCall {
	c.inclusive = b
	return c
}

func (c *ChannelsHistoryCall) Count(i int) *ChannelsHistoryCall {
	c.count = i
	return c
}

func (c *ChannelsHistoryCall) Unreads(b bool) *ChannelsHistoryCall {
	c.unreads = b
	return c
}

func (c *ChannelsHistoryCall) Do(ctx context.Context) (*ChannelsHistoryResponse, error) {
	const endpoint = "channels.history"

	var res struct {
		SlackResponse
		*ChannelsHistoryResponse
	}

	if err := c.service.client.postForm(ctx, endpoint, c.Values(), &res); err != nil {
		return nil, errors.Wrapf(err, `failed to post to %s`, endpoint)
	}

	if !res.OK {
		return nil, errors.New(res.Error.String())
	}

	return res.ChannelsHistoryResponse, nil
}

// Info returns the result of channels.info API
func (s *ChannelsService) Info(ctx context.Context, id string) (*objects.Channel, error) {
	v := url.Values{
		"token":   {s.token},
		"channel": {id},
	}
	const endpoint = "channels.info"

	var res struct {
		SlackResponse
		*objects.Channel `json:"channel"`
	}

	if err := s.client.postForm(ctx, endpoint, v, &res); err != nil {
		return nil, errors.Wrapf(err, `failed to post to %s`, endpoint)
	}

	if !res.OK {
		return nil, errors.New(res.Error.String())
	}

	return res.Channel, nil
}

// List returns the result of channels.list API
func (s *ChannelsService) List(ctx context.Context, exclArchived bool) (objects.ChannelList, error) {
	v := url.Values{
		"token": {s.token},
	}
	if exclArchived {
		v.Set("exclude_archived", "true")
	}
	const endpoint = "channels.list"

	var res struct {
		SlackResponse
		objects.ChannelList `json:"channels"`
	}

	if err := s.client.postForm(ctx, endpoint, v, &res); err != nil {
		return nil, errors.Wrapf(err, `failed to post to %s`, endpoint)
	}

	if !res.OK {
		return nil, errors.New(res.Error.String())
	}

	return res.ChannelList, nil
}
