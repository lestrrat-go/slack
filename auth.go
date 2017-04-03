package slack

import (
	"context"
	"net/url"

	"github.com/pkg/errors"
)

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

	return res.AuthTestResponse, nil
}
