package slack

import (
	"context"
	"net/url"

	"github.com/pkg/errors"
)

type RTMStartCall struct {
	service *RTMService
}

// Start returns the result of users.getPresence API
func (s *RTMService) Start() *RTMStartCall {
	return &RTMStartCall{
		service: s,
	}
}

func (c *RTMStartCall) Values() url.Values {
	v := url.Values{
		"token": {c.service.token},
	}
	return v
}

func (c *RTMStartCall) Do(ctx context.Context) (*RTMResponse, error) {
	const endpoint = "rtm.start"
	var res struct {
		SlackResponse
		*RTMResponse
	}
	if err := c.service.client.postForm(ctx, endpoint, c.Values(), &res); err != nil {
		return nil, errors.Wrapf(err, `failed to post to %s`, endpoint)
	}

	if !res.OK {
		return nil, errors.New(res.Error.String())
	}

	return res.RTMResponse, nil
}
