package slack

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/lestrrat/go-slack/objects"
	"github.com/pkg/errors"
)

func (s *UsersProfileService) Get() *UsersProfileGetCall {
	return &UsersProfileGetCall{
		service: s,
	}
}

type UsersProfileGetCall struct {
	service       *UsersProfileService
	user          string
	includeLabels bool
}

func (c *UsersProfileGetCall) Values() url.Values {
	v := url.Values{
		"token": {c.service.token},
	}

	if val := c.user; len(val) > 0 {
		v.Set("user", val)
	}

	if c.includeLabels {
		v.Set("name", "true")
	}

	return v
}

func (c *UsersProfileGetCall) User(s string) *UsersProfileGetCall {
	c.user = s
	return c
}

func (c *UsersProfileGetCall) IncludeLabels(b bool) *UsersProfileGetCall {
	c.includeLabels = b
	return c
}

func (c *UsersProfileGetCall) Do(ctx context.Context) (*objects.UserProfile, error) {
	const endpoint = "users.profile.get"
	var res struct {
		SlackResponse
		*objects.UserProfile
	}

	if err := c.service.client.postForm(ctx, endpoint, c.Values(), &res); err != nil {
		return nil, errors.Wrapf(err, `failed to post to %s`, endpoint)
	}

	if !res.OK {
		return nil, errors.New(res.Error.String())
	}

	return res.UserProfile, nil
}

func (s *UsersProfileService) Set() *UsersProfileSetCall {
	return &UsersProfileSetCall{
		service: s,
	}
}

type UsersProfileSetCall struct {
	service *UsersProfileService
	user    string
	profile *objects.UserProfile
	name    string
	value   string
}

func (c *UsersProfileSetCall) Values() (url.Values, error) {
	v := url.Values{
		"token": {c.service.token},
	}

	if val := c.user; len(val) > 0 {
		v.Set("user", val)
	}

	if val := c.profile; val != nil {
		buf, err := json.Marshal(val)
		if err != nil {
			return nil, errors.Wrap(err, `failed to unmarshal profile`)
		}
		v.Set("profile", string(buf))

		if len(c.name) > 0 || len(c.value) > 0 {
			return nil, errors.New(`"name"/"value" can only be set if "profile" is not set`)
		}
	}

	if val := c.name; len(val) > 0 {
		v.Set("name", val)
	}

	if val := c.value; len(val) > 0 {
		v.Set("value", val)
	}

	return v, nil
}

func (c *UsersProfileSetCall) User(s string) *UsersProfileSetCall {
	c.user = s
	return c
}

func (c *UsersProfileSetCall) Profile(p *objects.UserProfile) *UsersProfileSetCall {
	c.profile = p
	return c
}

func (c *UsersProfileSetCall) Name(s string) *UsersProfileSetCall {
	c.name = s
	return c
}

func (c *UsersProfileSetCall) Value(s string) *UsersProfileSetCall {
	c.value = s
	return c
}

func (c *UsersProfileSetCall) Do(ctx context.Context) (*objects.UserProfile, error) {
	const endpoint = "users.profile.set"
	var res struct {
		SlackResponse
		*objects.UserProfile
	}

	v, err := c.Values()
	if err != nil {
		return nil, errors.Wrap(err, `failed to prepare post parameters`)
	}

	if err := c.service.client.postForm(ctx, endpoint, v, &res); err != nil {
		return nil, errors.Wrapf(err, `failed to post to %s`, endpoint)
	}

	if !res.OK {
		return nil, errors.New(res.Error.String())
	}

	return res.UserProfile, nil
}
