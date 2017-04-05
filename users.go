package slack

import (
	"context"
	"net/url"

	"github.com/lestrrat/go-slack/objects"
	"github.com/pkg/errors"
)

type UsersGetPresenceCall struct {
	service *UsersService
	user    string // user ID
}

// GetPresence returns the result of users.getPresence API
func (s *UsersService) GetPresence(id string) *UsersGetPresenceCall {
	return &UsersGetPresenceCall{
		service: s,
		user:    id,
	}
}

func (c *UsersGetPresenceCall) Values() url.Values {
	v := url.Values{
		"token": {c.service.token},
		"user":  {c.user},
	}
	return v
}

func (c *UsersGetPresenceCall) Do(ctx context.Context) (*objects.UserPresence, error) {
	const endpoint = "users.getPresence"
	var res struct {
		SlackResponse
		*objects.UserPresence
	}

	if err := c.service.client.postForm(ctx, endpoint, c.Values(), &res); err != nil {
		return nil, errors.Wrapf(err, `failed to post to %s`, endpoint)
	}

	if !res.OK {
		return nil, errors.New(res.Error.String())
	}

	return res.UserPresence, nil
}

type UsersInfoCall struct {
	service *UsersService
	user    string // user ID
}

// Info returns the result of users.info API
func (s *UsersService) Info(id string) *UsersInfoCall {
	return &UsersInfoCall{
		service: s,
		user:    id,
	}
}

func (c *UsersInfoCall) Values() url.Values {
	v := url.Values{
		"token": {c.service.token},
		"user":  {c.user},
	}
	return v
}

func (c *UsersInfoCall) Do(ctx context.Context) (*objects.User, error) {
	const endpoint = "users.info"
	var res struct {
		SlackResponse
		*objects.User `json:"user"`
	}

	if err := c.service.client.postForm(ctx, endpoint, c.Values(), &res); err != nil {
		return nil, errors.Wrapf(err, `failed to post to %s`, endpoint)
	}

	if !res.OK {
		return nil, errors.New(res.Error.String())
	}

	return res.User, nil
}

type UsersListCall struct {
	service      *UsersService
	inclPresence bool
}

// List returns the result of users.list API
func (s *UsersService) List() *UsersListCall {
	return &UsersListCall{
		service: s,
	}
}

func (c *UsersListCall) Values() url.Values {
	v := url.Values{
		"token": {c.service.token},
	}
	if c.inclPresence {
		v.Set("presence", "true")
	}
	return v
}

func (c *UsersListCall) IncludePresence(b bool) *UsersListCall {
	c.inclPresence = b
	return c
}

func (c *UsersListCall) Do(ctx context.Context) (objects.UserList, error) {
	const endpoint = "users.list"
	var res struct {
		SlackResponse
		objects.UserList `json:"members"`
	}

	if err := c.service.client.postForm(ctx, endpoint, c.Values(), &res); err != nil {
		return nil, errors.Wrapf(err, `failed to post to %s`, endpoint)
	}

	if !res.OK {
		return nil, errors.New(res.Error.String())
	}

	return res.UserList, nil
}
