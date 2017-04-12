package slack

import (
	"context"
	"net/url"

	"github.com/lestrrat/go-slack/objects"
	"github.com/pkg/errors"
)

type BotsInfoCall struct {
	service *BotsService
	bot     string
}

func (s *BotsService) Info(bot string) *BotsInfoCall {
	return &BotsInfoCall{
		service: s,
		bot:     bot,
	}
}

func (c *BotsInfoCall) Values() url.Values {
	v := url.Values{
		"token": {c.service.token},
		"bot":   {c.bot},
	}
	return v
}

func (c *BotsInfoCall) Do(ctx context.Context) (*objects.Bot, error) {
	const endpoint = "bots.info"
	var res struct {
		SlackResponse
		*objects.Bot
	}

	if err := c.service.client.postForm(ctx, endpoint, c.Values(), &res); err != nil {
		return nil, errors.Wrapf(err, `failed to post to %s`, endpoint)
	}

	if !res.OK {
		return nil, errors.New(res.Error.String())
	}

	return res.Bot, nil
}
