package slack

// Auto-generated by internal/cmd/genmethods/genmethods.go. DO NOT EDIT!

import (
	"context"
	"net/url"
	"strconv"
	"strings"

	"github.com/lestrrat/go-slack/objects"
	"github.com/pkg/errors"
)

var _ = strconv.Itoa
var _ = strings.Index
var _ = objects.EpochTime(0)

// BotsInfoCall is created by BotsService.Info method call
type BotsInfoCall struct {
	service *BotsService
	bot     string
}

// Info creates a BotsInfoCall object in preparation for accessing the bots.info endpoint
func (s *BotsService) Info(bot string) *BotsInfoCall {
	var call BotsInfoCall
	call.service = s
	call.bot = bot
	return &call
}

// ValidateArgs checks that all required fields are set in the BotsInfoCall object
func (c *BotsInfoCall) ValidateArgs() error {
	if len(c.bot) <= 0 {
		return errors.New(`required field bot not initialized`)
	}
	return nil
}

// Values returns the BotsInfoCall object as url.Values
func (c *BotsInfoCall) Values() (url.Values, error) {
	if err := c.ValidateArgs(); err != nil {
		return nil, errors.Wrap(err, `failed validation`)
	}
	v := url.Values{}
	v.Set(`token`, c.service.token)

	v.Set("bot", c.bot)
	return v, nil
}

// Do executes the call to access bots.info endpoint
func (c *BotsInfoCall) Do(ctx context.Context) (*objects.Bot, error) {
	const endpoint = "bots.info"
	v, err := c.Values()
	if err != nil {
		return nil, err
	}
	var res struct {
		objects.GenericResponse
		*objects.Bot
	}
	if err := c.service.client.postForm(ctx, endpoint, v, &res); err != nil {
		return nil, errors.Wrap(err, `failed to post to bots.info`)
	}
	if !res.OK {
		return nil, errors.New(res.Error.String())
	}

	return res.Bot, nil
}

// FromValues parses the data in v and populates `c`
func (c *BotsInfoCall) FromValues(v url.Values) error {
	var tmp BotsInfoCall
	if raw := strings.TrimSpace(v.Get("bot")); len(raw) > 0 {
		tmp.bot = raw
	}
	*c = tmp
	return nil
}
