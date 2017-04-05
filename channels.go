package slack

import (
	"context"
	"net/url"

	"github.com/lestrrat/go-slack/objects"
	"github.com/pkg/errors"
)

// Info returns the result of users.info API
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

