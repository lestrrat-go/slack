package slack

import (
	"context"
	"net/url"
	"strconv"

	"github.com/lestrrat/go-slack/objects"
	"github.com/pkg/errors"
)

// ChannelsArchiveCall is created via Channels.Archive() method
type ChannelsArchiveCall struct {
	service *ChannelsService
	channel string // channel ID
}

func (s *ChannelsService) Archive(id string) *ChannelsArchiveCall {
	return &ChannelsArchiveCall{
		service: s,
		channel: id,
	}
}

func (c *ChannelsArchiveCall) Values() url.Values {
	v := url.Values{
		"token":   {c.service.token},
		"channel": {c.channel},
	}
	return v
}

func (c *ChannelsArchiveCall) Do(ctx context.Context) error {
	const endpoint = "channels.archive"

	var res SlackResponse
	if err := genericPost(ctx, c.service.client, endpoint, c.Values(), &res); err != nil {
		return err
	}

	return nil
}

// ChannelsCreateCall is created via Channels.Create() method
type ChannelsCreateCall struct {
	service  *ChannelsService
	name     string // channel name
	validate bool
}

func (s *ChannelsService) Create(name string) *ChannelsCreateCall {
	return &ChannelsCreateCall{
		service: s,
		name:    name,
	}
}

func (c *ChannelsCreateCall) Values() url.Values {
	v := url.Values{
		"token": {c.service.token},
		"name":  {c.name},
	}
	if c.validate {
		v.Set("validate", "true")
	}
	return v
}

func (c *ChannelsCreateCall) Validate(b bool) *ChannelsCreateCall {
	c.validate = b
	return c
}

func (c *ChannelsCreateCall) Do(ctx context.Context) error {
	const endpoint = "channels.create"
	var res SlackResponse
	if err := genericPost(ctx, c.service.client, endpoint, c.Values(), &res); err != nil {
		return err
	}
	return nil
}

// ChannelsHistoryCall is created via Channels.History() method
type ChannelsHistoryCall struct {
	service   *ChannelsService
	channel   string // channel ID
	count     int    // 1-1000
	inclusive bool
	latest    string // range of time (end)
	oldest    string // range of time (start)
	timestamp string // used only when retrieving a single message
	unreads   bool   // Include unread_count_display in the output
}

func (s *ChannelsService) History(id string) *ChannelsHistoryCall {
	return &ChannelsHistoryCall{
		service: s,
		channel: id,
	}
}

func (c *ChannelsHistoryCall) Values() url.Values {
	v := url.Values{
		"token":   {c.service.token},
		"channel": {c.channel},
	}

	if c.count > 0 {
		v.Set("count", strconv.Itoa(c.count))
	}

	if c.inclusive {
		v.Set("inclusive", "1")
	}

	if len(c.latest) > 0 {
		v.Set("latest", c.latest)
	}

	if len(c.oldest) > 0 {
		v.Set("oldest", c.oldest)
	}

	if len(c.timestamp) > 0 {
		v.Set("ts", c.timestamp)
	}

	if c.unreads {
		v.Set("unreads", "1")
	}

	return v
}

func (c *ChannelsHistoryCall) Latest(s string) *ChannelsHistoryCall {
	c.latest = s
	return c
}

func (c *ChannelsHistoryCall) Oldest(s string) *ChannelsHistoryCall {
	c.oldest = s
	return c
}

func (c *ChannelsHistoryCall) Inclusive(b bool) *ChannelsHistoryCall {
	c.inclusive = b
	return c
}

func (c *ChannelsHistoryCall) Count(i int) *ChannelsHistoryCall {
	c.count = i
	return c
}

func (c *ChannelsHistoryCall) Timestamp(s string) *ChannelsHistoryCall {
	c.timestamp = s
	return c
}

func (c *ChannelsHistoryCall) Unreads(b bool) *ChannelsHistoryCall {
	c.unreads = b
	return c
}

func (c *ChannelsHistoryCall) Do(ctx context.Context) (*ChannelsHistoryResponse, error) {
	const endpoint = "channels.history"

	var res struct {
		SlackResponse
		*ChannelsHistoryResponse
	}

	if err := genericPost(ctx, c.service.client, endpoint, c.Values(), &res); err != nil {
		return nil, err
	}

	return res.ChannelsHistoryResponse, nil
}

// ChannelsInfoCall is created via Channels.Info() method
type ChannelsInfoCall struct {
	service *ChannelsService
	channel string // channel ID
}

// Info returns the result of channels.info API
func (s *ChannelsService) Info(id string) *ChannelsInfoCall {
	return &ChannelsInfoCall{
		service: s,
		channel: id,
	}
}

func (c *ChannelsInfoCall) Values() url.Values {
	return url.Values{
		"token":   {c.service.token},
		"channel": {c.channel},
	}
}

func (c *ChannelsInfoCall) Do(ctx context.Context) (*objects.Channel, error) {
	const endpoint = "channels.info"
	var res struct {
		SlackResponse
		*objects.Channel `json:"channel"`
	}

	if err := genericPost(ctx, c.service.client, endpoint, c.Values(), &res); err != nil {
		return nil, errors.Wrapf(err, `failed to post to %s`, endpoint)
	}

	if !res.OK {
		return nil, errors.New(res.Error.String())
	}

	return res.Channel, nil
}

// ChannelsInviteCall is created via Channels.Invite() method
type ChannelsInviteCall struct {
	service *ChannelsService
	channel string
	user    string
}

func (s *ChannelsService) Invite(channelID, userID string) *ChannelsInviteCall {
	return &ChannelsInviteCall{
		service: s,
		channel: channelID,
		user:    userID,
	}
}

func (c *ChannelsInviteCall) Values() url.Values {
	v := url.Values{
		"token":   {c.service.token},
		"channel": {c.channel},
		"user":    {c.user},
	}
	return v
}

func (c *ChannelsInviteCall) Do(ctx context.Context) (*objects.Channel, error) {
	const endpoint = "channels.invite"
	var res struct {
		SlackResponse
		*objects.Channel `json:"channel"`
	}

	if err := genericPost(ctx, c.service.client, endpoint, c.Values(), &res); err != nil {
		return nil, err
	}
	return res.Channel, nil
}

// ChannelsJoinCall is created via Channels.Join() method
type ChannelsJoinCall struct {
	service  *ChannelsService
	name     string // channel name
	validate bool
}

func (s *ChannelsService) Join(name string) *ChannelsJoinCall {
	return &ChannelsJoinCall{
		service: s,
		name:    name,
	}
}

func (c *ChannelsJoinCall) Values() url.Values {
	v := url.Values{
		"token": {c.service.token},
		"name":  {c.name},
	}
	if c.validate {
		v.Set("validate", "true")
	}
	return v
}

func (c *ChannelsJoinCall) Validate(b bool) *ChannelsJoinCall {
	c.validate = b
	return c
}

func (c *ChannelsJoinCall) Do(ctx context.Context) (*objects.Channel, error) {
	const endpoint = "channels.join"
	var res struct {
		SlackResponse
		*objects.Channel `json:"channel"`
	}

	if err := genericPost(ctx, c.service.client, endpoint, c.Values(), &res); err != nil {
		return nil, err
	}
	return res.Channel, nil
}

// ChannelsKickCall is created via Channels.Kick() method
type ChannelsKickCall struct {
	service *ChannelsService
	channel string // channel ID
	user    string
}

func (s *ChannelsService) Kick(id, user string) *ChannelsKickCall {
	return &ChannelsKickCall{
		service: s,
		channel: id,
		user:    user,
	}
}

func (c *ChannelsKickCall) Values() url.Values {
	v := url.Values{
		"token":   {c.service.token},
		"channel": {c.channel},
		"user":    {c.user},
	}
	return v
}

func (c *ChannelsKickCall) Do(ctx context.Context) error {
	const endpoint = "channels.kick"

	var res SlackResponse
	if err := genericPost(ctx, c.service.client, endpoint, c.Values(), &res); err != nil {
		return err
	}

	return nil
}

// ChannelsLeaveCall is created via Channels.Leave() method
type ChannelsLeaveCall struct {
	service *ChannelsService
	channel string // channel ID
}

func (s *ChannelsService) Leave(id string) *ChannelsLeaveCall {
	return &ChannelsLeaveCall{
		service: s,
		channel: id,
	}
}

func (c *ChannelsLeaveCall) Values() url.Values {
	v := url.Values{
		"token":   {c.service.token},
		"channel": {c.channel},
	}
	return v
}

func (c *ChannelsLeaveCall) Do(ctx context.Context) error {
	const endpoint = "channels.leave"

	var res SlackResponse
	if err := genericPost(ctx, c.service.client, endpoint, c.Values(), &res); err != nil {
		return err
	}

	return nil
}

// ChannelsListCall is created via Channels.List() method
type ChannelsListCall struct {
	service      *ChannelsService
	exclArchived bool
}

// List returns the result of channels.list API
func (s *ChannelsService) List() *ChannelsListCall {
	return &ChannelsListCall{
		service: s,
	}
}

func (c *ChannelsListCall) Values() url.Values {
	v := url.Values{
		"token": {c.service.token},
	}
	if c.exclArchived {
		v.Set("exclude_archived", "true")
	}
	return v
}

func (c *ChannelsListCall) ExcludeArchive(b bool) *ChannelsListCall {
	c.exclArchived = b
	return c
}

func (c *ChannelsListCall) Do(ctx context.Context) (objects.ChannelList, error) {
	const endpoint = "channels.list"
	var res struct {
		SlackResponse
		objects.ChannelList `json:"channels"`
	}

	if err := genericPost(ctx, c.service.client, endpoint, c.Values(), &res); err != nil {
		return nil, err
	}
	return res.ChannelList, nil
}
