package slack

// Auto-generated by internal/cmd/genmethods/genmethods.go (generateServiceDetailsFile). DO NOT EDIT!

import (
	"context"
	"encoding/json"
	"net/url"
	"strconv"
	"strings"

	"github.com/lestrrat-go/slack/objects"
	"github.com/pkg/errors"
)

var _ = strconv.Itoa
var _ = strings.Index
var _ = json.Marshal
var _ = objects.EpochTime(0)

// UsergroupsUsersListCall is created by UsergroupsUsersService.List method call
type UsergroupsUsersListCall struct {
	service         *UsergroupsUsersService
	includeDisabled bool
	usergroup       string
}

// UsergroupsUsersUpdateCall is created by UsergroupsUsersService.Update method call
type UsergroupsUsersUpdateCall struct {
	service      *UsergroupsUsersService
	includeCount bool
	usergroup    string
	users        string // Comma-separated list of user IDs
}

// List creates a UsergroupsUsersListCall object in preparation for accessing the usergroups.users.list endpoint
func (s *UsergroupsUsersService) List(usergroup string) *UsergroupsUsersListCall {
	var call UsergroupsUsersListCall
	call.service = s
	call.usergroup = usergroup
	return &call
}

// IncludeDisabled sets the value for optional includeDisabled parameter
func (c *UsergroupsUsersListCall) IncludeDisabled(includeDisabled bool) *UsergroupsUsersListCall {
	c.includeDisabled = includeDisabled
	return c
}

// ValidateArgs checks that all required fields are set in the UsergroupsUsersListCall object
func (c *UsergroupsUsersListCall) ValidateArgs() error {
	if len(c.usergroup) <= 0 {
		return errors.New(`required field usergroup not initialized`)
	}
	return nil
}

// Values returns the UsergroupsUsersListCall object as url.Values
func (c *UsergroupsUsersListCall) Values() (url.Values, error) {
	if err := c.ValidateArgs(); err != nil {
		return nil, errors.Wrap(err, `failed validation`)
	}
	v := url.Values{}
	v.Set(`token`, c.service.token)

	if c.includeDisabled {
		v.Set("include_disabled", "true")
	}

	v.Set("usergroup", c.usergroup)
	return v, nil
}

type UsergroupsUsersListCallResponse interface {
	OK() bool
	ReplyTo() int
	Error() *objects.ErrorResponse
	Timestamp() string
	Users() *objects.UsergroupUsersList
}

type usergroupsUsersListCallResponseProxy struct {
	OK        bool                   `json:"ok"`
	ReplyTo   int                    `json:"reply_to"`
	Error     *objects.ErrorResponse `json:"error"`
	Timestamp string                 `json:"ts"`
	Payload0  json.RawMessage        `json:"-"`
	Payload1  json.RawMessage        `json:"users"`
}
type usergroupsUsersListCallResponse struct {
	ok      bool
	replyTo int
	error   *objects.ErrorResponse
	ts      string
	users   *objects.UsergroupUsersList
}
type UsergroupsUsersListCallResponseBuilder struct {
	resp *usergroupsUsersListCallResponse
}

func BuildUsergroupsUsersListCallResponse() *UsergroupsUsersListCallResponseBuilder {
	return &UsergroupsUsersListCallResponseBuilder{resp: &usergroupsUsersListCallResponse{}}
}
func (v *usergroupsUsersListCallResponse) OK() bool {
	return v.ok
}
func (v *usergroupsUsersListCallResponse) ReplyTo() int {
	return v.replyTo
}
func (v *usergroupsUsersListCallResponse) Error() *objects.ErrorResponse {
	return v.error
}
func (v *usergroupsUsersListCallResponse) Timestamp() string {
	return v.ts
}
func (v *usergroupsUsersListCallResponse) Users() *objects.UsergroupUsersList {
	return v.users
}
func (b *UsergroupsUsersListCallResponseBuilder) OK(v bool) *UsergroupsUsersListCallResponseBuilder {
	b.resp.ok = v
	return b
}
func (b *UsergroupsUsersListCallResponseBuilder) ReplyTo(v int) *UsergroupsUsersListCallResponseBuilder {
	b.resp.replyTo = v
	return b
}
func (b *UsergroupsUsersListCallResponseBuilder) Error(v *objects.ErrorResponse) *UsergroupsUsersListCallResponseBuilder {
	b.resp.error = v
	return b
}
func (b *UsergroupsUsersListCallResponseBuilder) Timestamp(v string) *UsergroupsUsersListCallResponseBuilder {
	b.resp.ts = v
	return b
}
func (b *UsergroupsUsersListCallResponseBuilder) Users(v *objects.UsergroupUsersList) *UsergroupsUsersListCallResponseBuilder {
	b.resp.users = v
	return b
}
func (b *UsergroupsUsersListCallResponseBuilder) Build() UsergroupsUsersListCallResponse {
	v := b.resp
	b.resp = &usergroupsUsersListCallResponse{}
	return v
}
func (r *usergroupsUsersListCallResponseProxy) parse(data []byte) error {
	if err := json.Unmarshal(data, r); err != nil {
		return errors.Wrap(err, `failed to unmarshal UsergroupsUsersListCallResponse`)
	}
	r.Payload0 = data
	return nil
}
func (r *usergroupsUsersListCallResponseProxy) payload() (objects.UsergroupUsersList, error) {
	var res1 objects.UsergroupUsersList
	if err := json.Unmarshal(r.Payload1, &res1); err != nil {
		return nil, errors.Wrap(err, `failed to ummarshal objects.UsergroupUsersList from response`)
	}
	return res1, nil
}
func (r *usergroupsUsersListCallResponse) MarshalJSON() ([]byte, error) {
	var p usergroupsUsersListCallResponseProxy
	p.OK = r.ok
	p.ReplyTo = r.replyTo
	p.Error = r.error
	p.Timestamp = r.ts
	payload1, err := json.Marshal(r.users)
	if err != nil {
		return nil, errors.Wrap(err, `failed to marshal 'users' field`)
	}
	p.Payload1 = payload1
	return json.Marshal(p)
}

// Do executes the call to access usergroups.users.list endpoint
func (c *UsergroupsUsersListCall) Do(ctx context.Context) (objects.UsergroupUsersList, error) {
	const endpoint = "usergroups.users.list"
	v, err := c.Values()
	if err != nil {
		return nil, err
	}
	var res usergroupsUsersListCallResponseProxy
	if err := c.service.client.postForm(ctx, endpoint, v, &res); err != nil {
		return nil, errors.Wrap(err, `failed to post to usergroups.users.list`)
	}
	if !res.OK {
		var err error
		if errresp := res.Error; errresp != nil {
			err = errors.New(errresp.String())
		} else {
			err = errors.New(`unknown error while posting to usergroups.users.list`)
		}
		return nil, err
	}

	return res.payload()
}

// FromValues parses the data in v and populates `c`
func (c *UsergroupsUsersListCall) FromValues(v url.Values) error {
	var tmp UsergroupsUsersListCall
	if raw := strings.TrimSpace(v.Get("include_disabled")); len(raw) > 0 {
		parsed, err := strconv.ParseBool(raw)
		if err != nil {
			return errors.Wrap(err, `failed to parse boolean value "include_disabled"`)
		}
		tmp.includeDisabled = parsed
	}
	if raw := strings.TrimSpace(v.Get("usergroup")); len(raw) > 0 {
		tmp.usergroup = raw
	}
	*c = tmp
	return nil
}

// Update creates a UsergroupsUsersUpdateCall object in preparation for accessing the usergroups.users.update endpoint
func (s *UsergroupsUsersService) Update(usergroup string, users string) *UsergroupsUsersUpdateCall {
	var call UsergroupsUsersUpdateCall
	call.service = s
	call.usergroup = usergroup
	call.users = users
	return &call
}

// IncludeCount sets the value for optional includeCount parameter
func (c *UsergroupsUsersUpdateCall) IncludeCount(includeCount bool) *UsergroupsUsersUpdateCall {
	c.includeCount = includeCount
	return c
}

// ValidateArgs checks that all required fields are set in the UsergroupsUsersUpdateCall object
func (c *UsergroupsUsersUpdateCall) ValidateArgs() error {
	if len(c.usergroup) <= 0 {
		return errors.New(`required field usergroup not initialized`)
	}
	if len(c.users) <= 0 {
		return errors.New(`required field users not initialized`)
	}
	return nil
}

// Values returns the UsergroupsUsersUpdateCall object as url.Values
func (c *UsergroupsUsersUpdateCall) Values() (url.Values, error) {
	if err := c.ValidateArgs(); err != nil {
		return nil, errors.Wrap(err, `failed validation`)
	}
	v := url.Values{}
	v.Set(`token`, c.service.token)

	if c.includeCount {
		v.Set("include_count", "true")
	}

	v.Set("usergroup", c.usergroup)

	v.Set("users", c.users)
	return v, nil
}

type UsergroupsUsersUpdateCallResponse interface {
	OK() bool
	ReplyTo() int
	Error() *objects.ErrorResponse
	Timestamp() string
	Usergroup() *objects.Usergroup
}

type usergroupsUsersUpdateCallResponseProxy struct {
	OK        bool                   `json:"ok"`
	ReplyTo   int                    `json:"reply_to"`
	Error     *objects.ErrorResponse `json:"error"`
	Timestamp string                 `json:"ts"`
	Payload0  json.RawMessage        `json:"-"`
	Payload1  json.RawMessage        `json:"usergroup"`
}
type usergroupsUsersUpdateCallResponse struct {
	ok        bool
	replyTo   int
	error     *objects.ErrorResponse
	ts        string
	usergroup *objects.Usergroup
}
type UsergroupsUsersUpdateCallResponseBuilder struct {
	resp *usergroupsUsersUpdateCallResponse
}

func BuildUsergroupsUsersUpdateCallResponse() *UsergroupsUsersUpdateCallResponseBuilder {
	return &UsergroupsUsersUpdateCallResponseBuilder{resp: &usergroupsUsersUpdateCallResponse{}}
}
func (v *usergroupsUsersUpdateCallResponse) OK() bool {
	return v.ok
}
func (v *usergroupsUsersUpdateCallResponse) ReplyTo() int {
	return v.replyTo
}
func (v *usergroupsUsersUpdateCallResponse) Error() *objects.ErrorResponse {
	return v.error
}
func (v *usergroupsUsersUpdateCallResponse) Timestamp() string {
	return v.ts
}
func (v *usergroupsUsersUpdateCallResponse) Usergroup() *objects.Usergroup {
	return v.usergroup
}
func (b *UsergroupsUsersUpdateCallResponseBuilder) OK(v bool) *UsergroupsUsersUpdateCallResponseBuilder {
	b.resp.ok = v
	return b
}
func (b *UsergroupsUsersUpdateCallResponseBuilder) ReplyTo(v int) *UsergroupsUsersUpdateCallResponseBuilder {
	b.resp.replyTo = v
	return b
}
func (b *UsergroupsUsersUpdateCallResponseBuilder) Error(v *objects.ErrorResponse) *UsergroupsUsersUpdateCallResponseBuilder {
	b.resp.error = v
	return b
}
func (b *UsergroupsUsersUpdateCallResponseBuilder) Timestamp(v string) *UsergroupsUsersUpdateCallResponseBuilder {
	b.resp.ts = v
	return b
}
func (b *UsergroupsUsersUpdateCallResponseBuilder) Usergroup(v *objects.Usergroup) *UsergroupsUsersUpdateCallResponseBuilder {
	b.resp.usergroup = v
	return b
}
func (b *UsergroupsUsersUpdateCallResponseBuilder) Build() UsergroupsUsersUpdateCallResponse {
	v := b.resp
	b.resp = &usergroupsUsersUpdateCallResponse{}
	return v
}
func (r *usergroupsUsersUpdateCallResponseProxy) parse(data []byte) error {
	if err := json.Unmarshal(data, r); err != nil {
		return errors.Wrap(err, `failed to unmarshal UsergroupsUsersUpdateCallResponse`)
	}
	r.Payload0 = data
	return nil
}
func (r *usergroupsUsersUpdateCallResponseProxy) payload() (*objects.Usergroup, error) {
	var res1 objects.Usergroup
	if err := json.Unmarshal(r.Payload1, &res1); err != nil {
		return nil, errors.Wrap(err, `failed to ummarshal objects.Usergroup from response`)
	}
	return &res1, nil
}
func (r *usergroupsUsersUpdateCallResponse) MarshalJSON() ([]byte, error) {
	var p usergroupsUsersUpdateCallResponseProxy
	p.OK = r.ok
	p.ReplyTo = r.replyTo
	p.Error = r.error
	p.Timestamp = r.ts
	payload1, err := json.Marshal(r.usergroup)
	if err != nil {
		return nil, errors.Wrap(err, `failed to marshal 'usergroup' field`)
	}
	p.Payload1 = payload1
	return json.Marshal(p)
}

// Do executes the call to access usergroups.users.update endpoint
func (c *UsergroupsUsersUpdateCall) Do(ctx context.Context) (*objects.Usergroup, error) {
	const endpoint = "usergroups.users.update"
	v, err := c.Values()
	if err != nil {
		return nil, err
	}
	var res usergroupsUsersUpdateCallResponseProxy
	if err := c.service.client.postForm(ctx, endpoint, v, &res); err != nil {
		return nil, errors.Wrap(err, `failed to post to usergroups.users.update`)
	}
	if !res.OK {
		var err error
		if errresp := res.Error; errresp != nil {
			err = errors.New(errresp.String())
		} else {
			err = errors.New(`unknown error while posting to usergroups.users.update`)
		}
		return nil, err
	}

	return res.payload()
}

// FromValues parses the data in v and populates `c`
func (c *UsergroupsUsersUpdateCall) FromValues(v url.Values) error {
	var tmp UsergroupsUsersUpdateCall
	if raw := strings.TrimSpace(v.Get("include_count")); len(raw) > 0 {
		parsed, err := strconv.ParseBool(raw)
		if err != nil {
			return errors.Wrap(err, `failed to parse boolean value "include_count"`)
		}
		tmp.includeCount = parsed
	}
	if raw := strings.TrimSpace(v.Get("usergroup")); len(raw) > 0 {
		tmp.usergroup = raw
	}
	if raw := strings.TrimSpace(v.Get("users")); len(raw) > 0 {
		tmp.users = raw
	}
	*c = tmp
	return nil
}
