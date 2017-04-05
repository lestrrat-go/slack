package slack

import (
	"context"
	"net/url"

	"github.com/pkg/errors"
)

// Revoke posts to auth.revoke endpoint.
func (s *AuthService) Revoke(ctx context.Context, test bool) error {
	// https://api.slack.com/methods/auth.revoke
	const endpoint = "auth.revoke"
	vars := url.Values{
		"token": {s.token},
	}
	if test {
		vars.Set("test", "true")
	}

	// This method returns something like the following
	// {
	//    "ok": true,
	//    "revoked": true
	// }
	// but the "revoked" field is really useless... so we just
	// check SlackResponse
	var res SlackResponse
	if err := s.client.postForm(ctx, endpoint, vars, &res); err != nil {
		return errors.Wrapf(err, `error while posting to %s`, endpoint)
	}

	if !res.OK {
		return errors.New(res.Error.String())
	}
	return nil
}

// Test posts to auth.test endpoint.
func (s *AuthService) Test(ctx context.Context) (*AuthTestResponse, error) {
	const endpoint = "auth.test"
	vars := url.Values{
		"token": {s.token},
	}
	var res struct {
		SlackResponse
		*AuthTestResponse
	}

	if err := s.client.postForm(ctx, endpoint, vars, &res); err != nil {
		return nil, errors.Wrapf(err, `error while posting to %s`, endpoint)
	}

	if !res.OK {
		return nil, errors.New(res.Error.String())
	}

	return res.AuthTestResponse, nil
}
