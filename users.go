package slack

import (
	"context"
	"net/url"

	"github.com/pkg/errors"
)

func (s *UsersService) GetPresence(ctx context.Context, id string) (*UserPresence, error) {
	v := url.Values{
		"token": {s.token},
		"user":  {id},
	}
	const endpoint = "users.getPresence"

	var res struct {
		SlackResponse
		*UserPresence
	}

	if err := s.client.postForm(ctx, endpoint, v, &res); err != nil {
		return nil, errors.Wrapf(err, `failed to post to %s`, endpoint)
	}

	if !res.OK {
		return nil, errors.New(res.Error)
	}

	return res.UserPresence, nil
}

func (s *UsersService) Info(ctx context.Context, id string) (*User, error) {
	v := url.Values{
		"token": {s.token},
		"user":  {id},
	}
	const endpoint = "users.info"

	var res struct {
		SlackResponse
		*User `json:"user"`
	}

	if err := s.client.postForm(ctx, endpoint, v, &res); err != nil {
		return nil, errors.Wrapf(err, `failed to post to %s`, endpoint)
	}

	if !res.OK {
		return nil, errors.New(res.Error)
	}

	return res.User, nil
}

func (s *UsersService) List(ctx context.Context, inclPresence bool) (UserList, error) {
	v := url.Values{
		"token": {s.token},
	}
	if inclPresence {
		v.Set("presence", "true")
	}
	const endpoint = "users.list"

	var res struct {
		SlackResponse
		UserList `json:"members"`
	}

	if err := s.client.postForm(ctx, endpoint, v, &res); err != nil {
		return nil, errors.Wrapf(err, `failed to post to %s`, endpoint)
	}

	if !res.OK {
		return nil, errors.New(res.Error)
	}

	return res.UserList, nil
}
