package slack

import (
	"context"
	"net/url"

	"github.com/pkg/errors"
)

// Start returns the result of users.getPresence API
func (s *RTMService) Start(ctx context.Context) (*RTMResponse, error) {
	v := url.Values{
		"token": {s.token},
	}

	const endpoint = "rtm.start"

	var res struct {
		SlackResponse
		*RTMResponse
	}
	if err := s.client.postForm(ctx, endpoint, v, &res); err != nil {
		return nil, errors.Wrapf(err, `failed to post to %s`, endpoint)
	}

	if !res.OK {
		return nil, errors.New(res.Error)
	}

	return res.RTMResponse, nil
}
