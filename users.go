package slack

import (
	"context"
	"net/url"

	"github.com/lestrrat/go-slack/objects"
	"github.com/pkg/errors"
)

// GetPresence returns the result of users.getPresence API
func (s *UsersService) GetPresence(ctx context.Context, id string) (*objects.UserPresence, error) {
	v := url.Values{
		"token": {s.token},
		"user":  {id},
	}
	const endpoint = "users.getPresence"

	var res struct {
		SlackResponse
		*objects.UserPresence
	}

	if err := s.client.postForm(ctx, endpoint, v, &res); err != nil {
		return nil, errors.Wrapf(err, `failed to post to %s`, endpoint)
	}

	if !res.OK {
		return nil, errors.New(res.Error.String())
	}

	return res.UserPresence, nil
}

// Info returns the result of users.info API
func (s *UsersService) Info(ctx context.Context, id string) (*objects.User, error) {
	v := url.Values{
		"token": {s.token},
		"user":  {id},
	}
	const endpoint = "users.info"

	var res struct {
		SlackResponse
		*objects.User `json:"user"`
	}

	if err := s.client.postForm(ctx, endpoint, v, &res); err != nil {
		return nil, errors.Wrapf(err, `failed to post to %s`, endpoint)
	}

	if !res.OK {
		return nil, errors.New(res.Error.String())
	}

	return res.User, nil
}

// List returns the result of users.list API
func (s *UsersService) List(ctx context.Context, inclPresence bool) (objects.UserList, error) {
	v := url.Values{
		"token": {s.token},
	}
	if inclPresence {
		v.Set("presence", "true")
	}
	const endpoint = "users.list"

	var res struct {
		SlackResponse
		objects.UserList `json:"members"`
	}

	if err := s.client.postForm(ctx, endpoint, v, &res); err != nil {
		return nil, errors.Wrapf(err, `failed to post to %s`, endpoint)
	}

	if !res.OK {
		return nil, errors.New(res.Error.String())
	}

	return res.UserList, nil
}
