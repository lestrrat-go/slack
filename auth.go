package slack

import (
	"context"
	"net/url"

	"github.com/pkg/errors"
)

type AuthRevokeCall struct {
	service *AuthService
	test    bool
}

// Revoke posts to auth.revoke endpoint.
func (s *AuthService) Revoke() *AuthRevokeCall {
	return &AuthRevokeCall{
		service: s,
	}
}

func (c *AuthRevokeCall) Values() url.Values {
	v := url.Values{
		"token": {c.service.token},
	}
	if c.test {
		v.Set("test", "true")
	}
	return v
}

func (c *AuthRevokeCall) Test(b bool) *AuthRevokeCall {
	c.test = b
	return c
}

func (c *AuthRevokeCall) Do(ctx context.Context) error {
	// https://api.slack.com/methods/auth.revoke
	const endpoint = "auth.revoke"

	// This method returns something like the following
	// {
	//    "ok": true,
	//    "revoked": true
	// }
	// but the "revoked" field is really useless... so we just
	// check SlackResponse
	var res SlackResponse
	if err := c.service.client.postForm(ctx, endpoint, c.Values(), &res); err != nil {
		return errors.Wrapf(err, `error while posting to %s`, endpoint)
	}

	if !res.OK {
		return errors.New(res.Error.String())
	}
	return nil
}

type AuthTestCall struct {
	service *AuthService
}

// Test posts to auth.test endpoint.
func (s *AuthService) Test() *AuthTestCall {
	return &AuthTestCall{
		service: s,
	}
}

func (c *AuthTestCall) Values() url.Values {
	v := url.Values{
		"token": {c.service.token},
	}
	return v
}

func (c *AuthTestCall) Do(ctx context.Context) (*AuthTestResponse, error) {
	const endpoint = "auth.test"
	var res struct {
		SlackResponse
		*AuthTestResponse
	}

	if err := c.service.client.postForm(ctx, endpoint, c.Values(), &res); err != nil {
		return nil, errors.Wrapf(err, `error while posting to %s`, endpoint)
	}

	if !res.OK {
		return nil, errors.New(res.Error.String())
	}

	return res.AuthTestResponse, nil
}
