package slack

import (
	"context"
	"net/url"

	"github.com/pkg/errors"
)

type OAuthAccessCall struct {
	service      *OAuthService
	clientID     string
	clientSecret string
	code         string
	redirectURI  string
}

func (s *OAuthService) Access(clientID, clientSecret, code string) *OAuthAccessCall {
	return &OAuthAccessCall{
		service:      s,
		clientID:     clientID,
		clientSecret: clientSecret,
		code:         code,
	}
}

func (c *OAuthAccessCall) Values() url.Values {
	var v url.Values
	v.Set("client_id", c.clientID)
	v.Set("client_secret", c.clientSecret)
	v.Set("code", c.code)
	if len(c.redirectURI) > 0 {
		v.Set("redirect_uri", c.redirectURI)
	}
	return v
}

func (c *OAuthAccessCall) RedirectURI(s string) *OAuthAccessCall {
	c.redirectURI = s
	return c
}

func (c *OAuthAccessCall) Do(ctx context.Context) (*OAuthStartResponse, error) {
	var res struct {
		SlackResponse
		*OAuthStartResponse
	}

	const endpoint = "oauth.start"
	if err := c.service.client.postForm(ctx, endpoint, c.Values(), &res); err != nil {
		return nil, errors.Wrapf(err, `failed to post to %s`, endpoint)
	}

	if !res.OK {
		return nil, errors.New(res.Error.String())
	}

	return res.OAuthStartResponse, nil
}
