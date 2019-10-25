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

// UsersDeletePhotoCall is created by UsersService.DeletePhoto method call
type UsersDeletePhotoCall struct {
	service *UsersService
}

// UsersGetPresenceCall is created by UsersService.GetPresence method call
type UsersGetPresenceCall struct {
	service *UsersService
	user    string
}

// UsersIdentityCall is created by UsersService.Identity method call
type UsersIdentityCall struct {
	service *UsersService
}

// UsersInfoCall is created by UsersService.Info method call
type UsersInfoCall struct {
	service       *UsersService
	includeLocale bool
	user          string
}

// UsersListCall is created by UsersService.List method call
type UsersListCall struct {
	service       *UsersService
	includeLocale bool
	limit         int
	presence      bool
}

// UsersLookupByEmailCall is created by UsersService.LookupByEmail method call
type UsersLookupByEmailCall struct {
	service *UsersService
	email   string
}

// UsersSetActiveCall is created by UsersService.SetActive method call
type UsersSetActiveCall struct {
	service *UsersService
}

// UsersSetPresenceCall is created by UsersService.SetPresence method call
type UsersSetPresenceCall struct {
	service  *UsersService
	presence string
}

// DeletePhoto creates a UsersDeletePhotoCall object in preparation for accessing the users.deletePhoto endpoint
func (s *UsersService) DeletePhoto() *UsersDeletePhotoCall {
	var call UsersDeletePhotoCall
	call.service = s
	return &call
}

// ValidateArgs checks that all required fields are set in the UsersDeletePhotoCall object
func (c *UsersDeletePhotoCall) ValidateArgs() error {
	return nil
}

// Values returns the UsersDeletePhotoCall object as url.Values
func (c *UsersDeletePhotoCall) Values() (url.Values, error) {
	if err := c.ValidateArgs(); err != nil {
		return nil, errors.Wrap(err, `failed validation`)
	}
	v := url.Values{}
	v.Set(`token`, c.service.token)
	return v, nil
}

type UsersDeletePhotoCallResponse interface {
	OK() bool
	ReplyTo() int
	Error() *objects.ErrorResponse
	Timestamp() string
}

type usersDeletePhotoCallResponseProxy struct {
	OK        bool                   `json:"ok"`
	ReplyTo   int                    `json:"reply_to"`
	Error     *objects.ErrorResponse `json:"error"`
	Timestamp string                 `json:"ts"`
	Payload0  json.RawMessage        `json:"-"`
}
type usersDeletePhotoCallResponse struct {
	ok      bool
	replyTo int
	error   *objects.ErrorResponse
	ts      string
}
type UsersDeletePhotoCallResponseBuilder struct {
	resp *usersDeletePhotoCallResponse
}

func BuildUsersDeletePhotoCallResponse() *UsersDeletePhotoCallResponseBuilder {
	return &UsersDeletePhotoCallResponseBuilder{resp: &usersDeletePhotoCallResponse{}}
}
func (v *usersDeletePhotoCallResponse) OK() bool {
	return v.ok
}
func (v *usersDeletePhotoCallResponse) ReplyTo() int {
	return v.replyTo
}
func (v *usersDeletePhotoCallResponse) Error() *objects.ErrorResponse {
	return v.error
}
func (v *usersDeletePhotoCallResponse) Timestamp() string {
	return v.ts
}
func (b *UsersDeletePhotoCallResponseBuilder) OK(v bool) *UsersDeletePhotoCallResponseBuilder {
	b.resp.ok = v
	return b
}
func (b *UsersDeletePhotoCallResponseBuilder) ReplyTo(v int) *UsersDeletePhotoCallResponseBuilder {
	b.resp.replyTo = v
	return b
}
func (b *UsersDeletePhotoCallResponseBuilder) Error(v *objects.ErrorResponse) *UsersDeletePhotoCallResponseBuilder {
	b.resp.error = v
	return b
}
func (b *UsersDeletePhotoCallResponseBuilder) Timestamp(v string) *UsersDeletePhotoCallResponseBuilder {
	b.resp.ts = v
	return b
}
func (b *UsersDeletePhotoCallResponseBuilder) Build() UsersDeletePhotoCallResponse {
	v := b.resp
	b.resp = &usersDeletePhotoCallResponse{}
	return v
}
func (r *usersDeletePhotoCallResponseProxy) parse(data []byte) error {
	if err := json.Unmarshal(data, r); err != nil {
		return errors.Wrap(err, `failed to unmarshal UsersDeletePhotoCallResponse`)
	}
	r.Payload0 = data
	return nil
}
func (r *usersDeletePhotoCallResponse) MarshalJSON() ([]byte, error) {
	var p usersDeletePhotoCallResponseProxy
	p.OK = r.ok
	p.ReplyTo = r.replyTo
	p.Error = r.error
	p.Timestamp = r.ts
	return json.Marshal(p)
}

// Do executes the call to access users.deletePhoto endpoint
func (c *UsersDeletePhotoCall) Do(ctx context.Context) error {
	const endpoint = "users.deletePhoto"
	v, err := c.Values()
	if err != nil {
		return err
	}
	var res usersDeletePhotoCallResponseProxy
	if err := c.service.client.postForm(ctx, endpoint, v, &res); err != nil {
		return errors.Wrap(err, `failed to post to users.deletePhoto`)
	}
	if !res.OK {
		var err error
		if errresp := res.Error; errresp != nil {
			err = errors.New(errresp.String())
		} else {
			err = errors.New(`unknown error while posting to users.deletePhoto`)
		}
		return err
	}

	return nil
}

// FromValues parses the data in v and populates `c`
func (c *UsersDeletePhotoCall) FromValues(v url.Values) error {
	var tmp UsersDeletePhotoCall
	*c = tmp
	return nil
}

// GetPresence creates a UsersGetPresenceCall object in preparation for accessing the users.getPresence endpoint
func (s *UsersService) GetPresence(user string) *UsersGetPresenceCall {
	var call UsersGetPresenceCall
	call.service = s
	call.user = user
	return &call
}

// ValidateArgs checks that all required fields are set in the UsersGetPresenceCall object
func (c *UsersGetPresenceCall) ValidateArgs() error {
	if len(c.user) <= 0 {
		return errors.New(`required field user not initialized`)
	}
	return nil
}

// Values returns the UsersGetPresenceCall object as url.Values
func (c *UsersGetPresenceCall) Values() (url.Values, error) {
	if err := c.ValidateArgs(); err != nil {
		return nil, errors.Wrap(err, `failed validation`)
	}
	v := url.Values{}
	v.Set(`token`, c.service.token)

	v.Set("user", c.user)
	return v, nil
}

type UsersGetPresenceCallResponse interface {
	OK() bool
	ReplyTo() int
	Error() *objects.ErrorResponse
	Timestamp() string
}

type usersGetPresenceCallResponseProxy struct {
	OK        bool                   `json:"ok"`
	ReplyTo   int                    `json:"reply_to"`
	Error     *objects.ErrorResponse `json:"error"`
	Timestamp string                 `json:"ts"`
	Payload0  json.RawMessage        `json:"-"`
}
type usersGetPresenceCallResponse struct {
	ok      bool
	replyTo int
	error   *objects.ErrorResponse
	ts      string
}
type UsersGetPresenceCallResponseBuilder struct {
	resp *usersGetPresenceCallResponse
}

func BuildUsersGetPresenceCallResponse() *UsersGetPresenceCallResponseBuilder {
	return &UsersGetPresenceCallResponseBuilder{resp: &usersGetPresenceCallResponse{}}
}
func (v *usersGetPresenceCallResponse) OK() bool {
	return v.ok
}
func (v *usersGetPresenceCallResponse) ReplyTo() int {
	return v.replyTo
}
func (v *usersGetPresenceCallResponse) Error() *objects.ErrorResponse {
	return v.error
}
func (v *usersGetPresenceCallResponse) Timestamp() string {
	return v.ts
}
func (b *UsersGetPresenceCallResponseBuilder) OK(v bool) *UsersGetPresenceCallResponseBuilder {
	b.resp.ok = v
	return b
}
func (b *UsersGetPresenceCallResponseBuilder) ReplyTo(v int) *UsersGetPresenceCallResponseBuilder {
	b.resp.replyTo = v
	return b
}
func (b *UsersGetPresenceCallResponseBuilder) Error(v *objects.ErrorResponse) *UsersGetPresenceCallResponseBuilder {
	b.resp.error = v
	return b
}
func (b *UsersGetPresenceCallResponseBuilder) Timestamp(v string) *UsersGetPresenceCallResponseBuilder {
	b.resp.ts = v
	return b
}
func (b *UsersGetPresenceCallResponseBuilder) Build() UsersGetPresenceCallResponse {
	v := b.resp
	b.resp = &usersGetPresenceCallResponse{}
	return v
}
func (r *usersGetPresenceCallResponseProxy) parse(data []byte) error {
	if err := json.Unmarshal(data, r); err != nil {
		return errors.Wrap(err, `failed to unmarshal UsersGetPresenceCallResponse`)
	}
	r.Payload0 = data
	return nil
}
func (r *usersGetPresenceCallResponseProxy) payload() (*objects.UserPresence, error) {
	var res0 objects.UserPresence
	if err := json.Unmarshal(r.Payload0, &res0); err != nil {
		return nil, errors.Wrap(err, `failed to ummarshal objects.UserPresence from response`)
	}
	return &res0, nil
}
func (r *usersGetPresenceCallResponse) MarshalJSON() ([]byte, error) {
	var p usersGetPresenceCallResponseProxy
	p.OK = r.ok
	p.ReplyTo = r.replyTo
	p.Error = r.error
	p.Timestamp = r.ts
	return json.Marshal(p)
}

// Do executes the call to access users.getPresence endpoint
func (c *UsersGetPresenceCall) Do(ctx context.Context) (*objects.UserPresence, error) {
	const endpoint = "users.getPresence"
	v, err := c.Values()
	if err != nil {
		return nil, err
	}
	var res usersGetPresenceCallResponseProxy
	if err := c.service.client.postForm(ctx, endpoint, v, &res); err != nil {
		return nil, errors.Wrap(err, `failed to post to users.getPresence`)
	}
	if !res.OK {
		var err error
		if errresp := res.Error; errresp != nil {
			err = errors.New(errresp.String())
		} else {
			err = errors.New(`unknown error while posting to users.getPresence`)
		}
		return nil, err
	}

	return res.payload()
}

// FromValues parses the data in v and populates `c`
func (c *UsersGetPresenceCall) FromValues(v url.Values) error {
	var tmp UsersGetPresenceCall
	if raw := strings.TrimSpace(v.Get("user")); len(raw) > 0 {
		tmp.user = raw
	}
	*c = tmp
	return nil
}

// Identity creates a UsersIdentityCall object in preparation for accessing the users.identity endpoint
func (s *UsersService) Identity() *UsersIdentityCall {
	var call UsersIdentityCall
	call.service = s
	return &call
}

// ValidateArgs checks that all required fields are set in the UsersIdentityCall object
func (c *UsersIdentityCall) ValidateArgs() error {
	return nil
}

// Values returns the UsersIdentityCall object as url.Values
func (c *UsersIdentityCall) Values() (url.Values, error) {
	if err := c.ValidateArgs(); err != nil {
		return nil, errors.Wrap(err, `failed validation`)
	}
	v := url.Values{}
	v.Set(`token`, c.service.token)
	return v, nil
}

type UsersIdentityCallResponse interface {
	OK() bool
	ReplyTo() int
	Error() *objects.ErrorResponse
	Timestamp() string
	User() *objects.UserProfile
	Team() *objects.Team
}

type usersIdentityCallResponseProxy struct {
	OK        bool                   `json:"ok"`
	ReplyTo   int                    `json:"reply_to"`
	Error     *objects.ErrorResponse `json:"error"`
	Timestamp string                 `json:"ts"`
	Payload0  json.RawMessage        `json:"-"`
	Payload1  json.RawMessage        `json:"user"`
	Payload2  json.RawMessage        `json:"team"`
}
type usersIdentityCallResponse struct {
	ok      bool
	replyTo int
	error   *objects.ErrorResponse
	ts      string
	user    *objects.UserProfile
	team    *objects.Team
}
type UsersIdentityCallResponseBuilder struct {
	resp *usersIdentityCallResponse
}

func BuildUsersIdentityCallResponse() *UsersIdentityCallResponseBuilder {
	return &UsersIdentityCallResponseBuilder{resp: &usersIdentityCallResponse{}}
}
func (v *usersIdentityCallResponse) OK() bool {
	return v.ok
}
func (v *usersIdentityCallResponse) ReplyTo() int {
	return v.replyTo
}
func (v *usersIdentityCallResponse) Error() *objects.ErrorResponse {
	return v.error
}
func (v *usersIdentityCallResponse) Timestamp() string {
	return v.ts
}
func (v *usersIdentityCallResponse) User() *objects.UserProfile {
	return v.user
}
func (v *usersIdentityCallResponse) Team() *objects.Team {
	return v.team
}
func (b *UsersIdentityCallResponseBuilder) OK(v bool) *UsersIdentityCallResponseBuilder {
	b.resp.ok = v
	return b
}
func (b *UsersIdentityCallResponseBuilder) ReplyTo(v int) *UsersIdentityCallResponseBuilder {
	b.resp.replyTo = v
	return b
}
func (b *UsersIdentityCallResponseBuilder) Error(v *objects.ErrorResponse) *UsersIdentityCallResponseBuilder {
	b.resp.error = v
	return b
}
func (b *UsersIdentityCallResponseBuilder) Timestamp(v string) *UsersIdentityCallResponseBuilder {
	b.resp.ts = v
	return b
}
func (b *UsersIdentityCallResponseBuilder) User(v *objects.UserProfile) *UsersIdentityCallResponseBuilder {
	b.resp.user = v
	return b
}
func (b *UsersIdentityCallResponseBuilder) Team(v *objects.Team) *UsersIdentityCallResponseBuilder {
	b.resp.team = v
	return b
}
func (b *UsersIdentityCallResponseBuilder) Build() UsersIdentityCallResponse {
	v := b.resp
	b.resp = &usersIdentityCallResponse{}
	return v
}
func (r *usersIdentityCallResponseProxy) parse(data []byte) error {
	if err := json.Unmarshal(data, r); err != nil {
		return errors.Wrap(err, `failed to unmarshal UsersIdentityCallResponse`)
	}
	r.Payload0 = data
	return nil
}
func (r *usersIdentityCallResponseProxy) payload() (*objects.UserProfile, *objects.Team, error) {
	var res1 objects.UserProfile
	if err := json.Unmarshal(r.Payload1, &res1); err != nil {
		return nil, nil, errors.Wrap(err, `failed to ummarshal objects.UserProfile from response`)
	}
	var res2 objects.Team
	if err := json.Unmarshal(r.Payload2, &res2); err != nil {
		return nil, nil, errors.Wrap(err, `failed to ummarshal objects.Team from response`)
	}
	return &res1, &res2, nil
}
func (r *usersIdentityCallResponse) MarshalJSON() ([]byte, error) {
	var p usersIdentityCallResponseProxy
	p.OK = r.ok
	p.ReplyTo = r.replyTo
	p.Error = r.error
	p.Timestamp = r.ts
	payload1, err := json.Marshal(r.user)
	if err != nil {
		return nil, errors.Wrap(err, `failed to marshal 'user' field`)
	}
	p.Payload1 = payload1
	payload2, err := json.Marshal(r.team)
	if err != nil {
		return nil, errors.Wrap(err, `failed to marshal 'team' field`)
	}
	p.Payload2 = payload2
	return json.Marshal(p)
}

// Do executes the call to access users.identity endpoint
func (c *UsersIdentityCall) Do(ctx context.Context) (*objects.UserProfile, *objects.Team, error) {
	const endpoint = "users.identity"
	v, err := c.Values()
	if err != nil {
		return nil, nil, err
	}
	var res usersIdentityCallResponseProxy
	if err := c.service.client.postForm(ctx, endpoint, v, &res); err != nil {
		return nil, nil, errors.Wrap(err, `failed to post to users.identity`)
	}
	if !res.OK {
		var err error
		if errresp := res.Error; errresp != nil {
			err = errors.New(errresp.String())
		} else {
			err = errors.New(`unknown error while posting to users.identity`)
		}
		return nil, nil, err
	}

	return res.payload()
}

// FromValues parses the data in v and populates `c`
func (c *UsersIdentityCall) FromValues(v url.Values) error {
	var tmp UsersIdentityCall
	*c = tmp
	return nil
}

// Info creates a UsersInfoCall object in preparation for accessing the users.info endpoint
func (s *UsersService) Info(user string) *UsersInfoCall {
	var call UsersInfoCall
	call.service = s
	call.user = user
	return &call
}

// IncludeLocale sets the value for optional includeLocale parameter
func (c *UsersInfoCall) IncludeLocale(includeLocale bool) *UsersInfoCall {
	c.includeLocale = includeLocale
	return c
}

// ValidateArgs checks that all required fields are set in the UsersInfoCall object
func (c *UsersInfoCall) ValidateArgs() error {
	if len(c.user) <= 0 {
		return errors.New(`required field user not initialized`)
	}
	return nil
}

// Values returns the UsersInfoCall object as url.Values
func (c *UsersInfoCall) Values() (url.Values, error) {
	if err := c.ValidateArgs(); err != nil {
		return nil, errors.Wrap(err, `failed validation`)
	}
	v := url.Values{}
	v.Set(`token`, c.service.token)

	if c.includeLocale {
		v.Set("include_locale", "true")
	}

	v.Set("user", c.user)
	return v, nil
}

type UsersInfoCallResponse interface {
	OK() bool
	ReplyTo() int
	Error() *objects.ErrorResponse
	Timestamp() string
	User() *objects.User
}

type usersInfoCallResponseProxy struct {
	OK        bool                   `json:"ok"`
	ReplyTo   int                    `json:"reply_to"`
	Error     *objects.ErrorResponse `json:"error"`
	Timestamp string                 `json:"ts"`
	Payload0  json.RawMessage        `json:"-"`
	Payload1  json.RawMessage        `json:"user"`
}
type usersInfoCallResponse struct {
	ok      bool
	replyTo int
	error   *objects.ErrorResponse
	ts      string
	user    *objects.User
}
type UsersInfoCallResponseBuilder struct {
	resp *usersInfoCallResponse
}

func BuildUsersInfoCallResponse() *UsersInfoCallResponseBuilder {
	return &UsersInfoCallResponseBuilder{resp: &usersInfoCallResponse{}}
}
func (v *usersInfoCallResponse) OK() bool {
	return v.ok
}
func (v *usersInfoCallResponse) ReplyTo() int {
	return v.replyTo
}
func (v *usersInfoCallResponse) Error() *objects.ErrorResponse {
	return v.error
}
func (v *usersInfoCallResponse) Timestamp() string {
	return v.ts
}
func (v *usersInfoCallResponse) User() *objects.User {
	return v.user
}
func (b *UsersInfoCallResponseBuilder) OK(v bool) *UsersInfoCallResponseBuilder {
	b.resp.ok = v
	return b
}
func (b *UsersInfoCallResponseBuilder) ReplyTo(v int) *UsersInfoCallResponseBuilder {
	b.resp.replyTo = v
	return b
}
func (b *UsersInfoCallResponseBuilder) Error(v *objects.ErrorResponse) *UsersInfoCallResponseBuilder {
	b.resp.error = v
	return b
}
func (b *UsersInfoCallResponseBuilder) Timestamp(v string) *UsersInfoCallResponseBuilder {
	b.resp.ts = v
	return b
}
func (b *UsersInfoCallResponseBuilder) User(v *objects.User) *UsersInfoCallResponseBuilder {
	b.resp.user = v
	return b
}
func (b *UsersInfoCallResponseBuilder) Build() UsersInfoCallResponse {
	v := b.resp
	b.resp = &usersInfoCallResponse{}
	return v
}
func (r *usersInfoCallResponseProxy) parse(data []byte) error {
	if err := json.Unmarshal(data, r); err != nil {
		return errors.Wrap(err, `failed to unmarshal UsersInfoCallResponse`)
	}
	r.Payload0 = data
	return nil
}
func (r *usersInfoCallResponseProxy) payload() (*objects.User, error) {
	var res1 objects.User
	if err := json.Unmarshal(r.Payload1, &res1); err != nil {
		return nil, errors.Wrap(err, `failed to ummarshal objects.User from response`)
	}
	return &res1, nil
}
func (r *usersInfoCallResponse) MarshalJSON() ([]byte, error) {
	var p usersInfoCallResponseProxy
	p.OK = r.ok
	p.ReplyTo = r.replyTo
	p.Error = r.error
	p.Timestamp = r.ts
	payload1, err := json.Marshal(r.user)
	if err != nil {
		return nil, errors.Wrap(err, `failed to marshal 'user' field`)
	}
	p.Payload1 = payload1
	return json.Marshal(p)
}

// Do executes the call to access users.info endpoint
func (c *UsersInfoCall) Do(ctx context.Context) (*objects.User, error) {
	const endpoint = "users.info"
	v, err := c.Values()
	if err != nil {
		return nil, err
	}
	var res usersInfoCallResponseProxy
	if err := c.service.client.postForm(ctx, endpoint, v, &res); err != nil {
		return nil, errors.Wrap(err, `failed to post to users.info`)
	}
	if !res.OK {
		var err error
		if errresp := res.Error; errresp != nil {
			err = errors.New(errresp.String())
		} else {
			err = errors.New(`unknown error while posting to users.info`)
		}
		return nil, err
	}

	return res.payload()
}

// FromValues parses the data in v and populates `c`
func (c *UsersInfoCall) FromValues(v url.Values) error {
	var tmp UsersInfoCall
	if raw := strings.TrimSpace(v.Get("include_locale")); len(raw) > 0 {
		parsed, err := strconv.ParseBool(raw)
		if err != nil {
			return errors.Wrap(err, `failed to parse boolean value "include_locale"`)
		}
		tmp.includeLocale = parsed
	}
	if raw := strings.TrimSpace(v.Get("user")); len(raw) > 0 {
		tmp.user = raw
	}
	*c = tmp
	return nil
}

// List creates a UsersListCall object in preparation for accessing the users.list endpoint
func (s *UsersService) List() *UsersListCall {
	var call UsersListCall
	call.service = s
	return &call
}

// IncludeLocale sets the value for optional includeLocale parameter
func (c *UsersListCall) IncludeLocale(includeLocale bool) *UsersListCall {
	c.includeLocale = includeLocale
	return c
}

// Limit sets the value for optional limit parameter
func (c *UsersListCall) Limit(limit int) *UsersListCall {
	c.limit = limit
	return c
}

// Presence sets the value for optional presence parameter
func (c *UsersListCall) Presence(presence bool) *UsersListCall {
	c.presence = presence
	return c
}

// ValidateArgs checks that all required fields are set in the UsersListCall object
func (c *UsersListCall) ValidateArgs() error {
	return nil
}

// Values returns the UsersListCall object as url.Values
func (c *UsersListCall) Values() (url.Values, error) {
	if err := c.ValidateArgs(); err != nil {
		return nil, errors.Wrap(err, `failed validation`)
	}
	v := url.Values{}
	v.Set(`token`, c.service.token)

	if c.includeLocale {
		v.Set("include_locale", "true")
	}

	if c.limit > 0 {
		v.Set("limit", strconv.Itoa(c.limit))
	}

	if c.presence {
		v.Set("presence", "true")
	}
	return v, nil
}

type UsersListCallResponse interface {
	OK() bool
	ReplyTo() int
	Error() *objects.ErrorResponse
	Timestamp() string
	Members() *objects.UserList
}

type usersListCallResponseProxy struct {
	OK        bool                   `json:"ok"`
	ReplyTo   int                    `json:"reply_to"`
	Error     *objects.ErrorResponse `json:"error"`
	Timestamp string                 `json:"ts"`
	Payload0  json.RawMessage        `json:"-"`
	Payload1  json.RawMessage        `json:"members"`
}
type usersListCallResponse struct {
	ok      bool
	replyTo int
	error   *objects.ErrorResponse
	ts      string
	members *objects.UserList
}
type UsersListCallResponseBuilder struct {
	resp *usersListCallResponse
}

func BuildUsersListCallResponse() *UsersListCallResponseBuilder {
	return &UsersListCallResponseBuilder{resp: &usersListCallResponse{}}
}
func (v *usersListCallResponse) OK() bool {
	return v.ok
}
func (v *usersListCallResponse) ReplyTo() int {
	return v.replyTo
}
func (v *usersListCallResponse) Error() *objects.ErrorResponse {
	return v.error
}
func (v *usersListCallResponse) Timestamp() string {
	return v.ts
}
func (v *usersListCallResponse) Members() *objects.UserList {
	return v.members
}
func (b *UsersListCallResponseBuilder) OK(v bool) *UsersListCallResponseBuilder {
	b.resp.ok = v
	return b
}
func (b *UsersListCallResponseBuilder) ReplyTo(v int) *UsersListCallResponseBuilder {
	b.resp.replyTo = v
	return b
}
func (b *UsersListCallResponseBuilder) Error(v *objects.ErrorResponse) *UsersListCallResponseBuilder {
	b.resp.error = v
	return b
}
func (b *UsersListCallResponseBuilder) Timestamp(v string) *UsersListCallResponseBuilder {
	b.resp.ts = v
	return b
}
func (b *UsersListCallResponseBuilder) Members(v *objects.UserList) *UsersListCallResponseBuilder {
	b.resp.members = v
	return b
}
func (b *UsersListCallResponseBuilder) Build() UsersListCallResponse {
	v := b.resp
	b.resp = &usersListCallResponse{}
	return v
}
func (r *usersListCallResponseProxy) parse(data []byte) error {
	if err := json.Unmarshal(data, r); err != nil {
		return errors.Wrap(err, `failed to unmarshal UsersListCallResponse`)
	}
	r.Payload0 = data
	return nil
}
func (r *usersListCallResponseProxy) payload() (objects.UserList, error) {
	var res1 objects.UserList
	if err := json.Unmarshal(r.Payload1, &res1); err != nil {
		return nil, errors.Wrap(err, `failed to ummarshal objects.UserList from response`)
	}
	return res1, nil
}
func (r *usersListCallResponse) MarshalJSON() ([]byte, error) {
	var p usersListCallResponseProxy
	p.OK = r.ok
	p.ReplyTo = r.replyTo
	p.Error = r.error
	p.Timestamp = r.ts
	payload1, err := json.Marshal(r.members)
	if err != nil {
		return nil, errors.Wrap(err, `failed to marshal 'members' field`)
	}
	p.Payload1 = payload1
	return json.Marshal(p)
}

// Do executes the call to access users.list endpoint
func (c *UsersListCall) Do(ctx context.Context) (objects.UserList, error) {
	const endpoint = "users.list"
	v, err := c.Values()
	if err != nil {
		return nil, err
	}
	var res usersListCallResponseProxy
	if err := c.service.client.postForm(ctx, endpoint, v, &res); err != nil {
		return nil, errors.Wrap(err, `failed to post to users.list`)
	}
	if !res.OK {
		var err error
		if errresp := res.Error; errresp != nil {
			err = errors.New(errresp.String())
		} else {
			err = errors.New(`unknown error while posting to users.list`)
		}
		return nil, err
	}

	return res.payload()
}

// FromValues parses the data in v and populates `c`
func (c *UsersListCall) FromValues(v url.Values) error {
	var tmp UsersListCall
	if raw := strings.TrimSpace(v.Get("include_locale")); len(raw) > 0 {
		parsed, err := strconv.ParseBool(raw)
		if err != nil {
			return errors.Wrap(err, `failed to parse boolean value "include_locale"`)
		}
		tmp.includeLocale = parsed
	}
	if raw := strings.TrimSpace(v.Get("limit")); len(raw) > 0 {
		parsed, err := strconv.ParseInt(raw, 10, 64)
		if err != nil {
			return errors.Wrap(err, `failed to parse integer value "limit"`)
		}
		tmp.limit = int(parsed)
	}
	if raw := strings.TrimSpace(v.Get("presence")); len(raw) > 0 {
		parsed, err := strconv.ParseBool(raw)
		if err != nil {
			return errors.Wrap(err, `failed to parse boolean value "presence"`)
		}
		tmp.presence = parsed
	}
	*c = tmp
	return nil
}

// LookupByEmail creates a UsersLookupByEmailCall object in preparation for accessing the users.lookupByEmail endpoint
func (s *UsersService) LookupByEmail(email string) *UsersLookupByEmailCall {
	var call UsersLookupByEmailCall
	call.service = s
	call.email = email
	return &call
}

// ValidateArgs checks that all required fields are set in the UsersLookupByEmailCall object
func (c *UsersLookupByEmailCall) ValidateArgs() error {
	if len(c.email) <= 0 {
		return errors.New(`required field email not initialized`)
	}
	return nil
}

// Values returns the UsersLookupByEmailCall object as url.Values
func (c *UsersLookupByEmailCall) Values() (url.Values, error) {
	if err := c.ValidateArgs(); err != nil {
		return nil, errors.Wrap(err, `failed validation`)
	}
	v := url.Values{}
	v.Set(`token`, c.service.token)

	v.Set("email", c.email)
	return v, nil
}

type UsersLookupByEmailCallResponse interface {
	OK() bool
	ReplyTo() int
	Error() *objects.ErrorResponse
	Timestamp() string
	User() *objects.User
}

type usersLookupByEmailCallResponseProxy struct {
	OK        bool                   `json:"ok"`
	ReplyTo   int                    `json:"reply_to"`
	Error     *objects.ErrorResponse `json:"error"`
	Timestamp string                 `json:"ts"`
	Payload0  json.RawMessage        `json:"-"`
	Payload1  json.RawMessage        `json:"user"`
}
type usersLookupByEmailCallResponse struct {
	ok      bool
	replyTo int
	error   *objects.ErrorResponse
	ts      string
	user    *objects.User
}
type UsersLookupByEmailCallResponseBuilder struct {
	resp *usersLookupByEmailCallResponse
}

func BuildUsersLookupByEmailCallResponse() *UsersLookupByEmailCallResponseBuilder {
	return &UsersLookupByEmailCallResponseBuilder{resp: &usersLookupByEmailCallResponse{}}
}
func (v *usersLookupByEmailCallResponse) OK() bool {
	return v.ok
}
func (v *usersLookupByEmailCallResponse) ReplyTo() int {
	return v.replyTo
}
func (v *usersLookupByEmailCallResponse) Error() *objects.ErrorResponse {
	return v.error
}
func (v *usersLookupByEmailCallResponse) Timestamp() string {
	return v.ts
}
func (v *usersLookupByEmailCallResponse) User() *objects.User {
	return v.user
}
func (b *UsersLookupByEmailCallResponseBuilder) OK(v bool) *UsersLookupByEmailCallResponseBuilder {
	b.resp.ok = v
	return b
}
func (b *UsersLookupByEmailCallResponseBuilder) ReplyTo(v int) *UsersLookupByEmailCallResponseBuilder {
	b.resp.replyTo = v
	return b
}
func (b *UsersLookupByEmailCallResponseBuilder) Error(v *objects.ErrorResponse) *UsersLookupByEmailCallResponseBuilder {
	b.resp.error = v
	return b
}
func (b *UsersLookupByEmailCallResponseBuilder) Timestamp(v string) *UsersLookupByEmailCallResponseBuilder {
	b.resp.ts = v
	return b
}
func (b *UsersLookupByEmailCallResponseBuilder) User(v *objects.User) *UsersLookupByEmailCallResponseBuilder {
	b.resp.user = v
	return b
}
func (b *UsersLookupByEmailCallResponseBuilder) Build() UsersLookupByEmailCallResponse {
	v := b.resp
	b.resp = &usersLookupByEmailCallResponse{}
	return v
}
func (r *usersLookupByEmailCallResponseProxy) parse(data []byte) error {
	if err := json.Unmarshal(data, r); err != nil {
		return errors.Wrap(err, `failed to unmarshal UsersLookupByEmailCallResponse`)
	}
	r.Payload0 = data
	return nil
}
func (r *usersLookupByEmailCallResponseProxy) payload() (*objects.User, error) {
	var res1 objects.User
	if err := json.Unmarshal(r.Payload1, &res1); err != nil {
		return nil, errors.Wrap(err, `failed to ummarshal objects.User from response`)
	}
	return &res1, nil
}
func (r *usersLookupByEmailCallResponse) MarshalJSON() ([]byte, error) {
	var p usersLookupByEmailCallResponseProxy
	p.OK = r.ok
	p.ReplyTo = r.replyTo
	p.Error = r.error
	p.Timestamp = r.ts
	payload1, err := json.Marshal(r.user)
	if err != nil {
		return nil, errors.Wrap(err, `failed to marshal 'user' field`)
	}
	p.Payload1 = payload1
	return json.Marshal(p)
}

// Do executes the call to access users.lookupByEmail endpoint
func (c *UsersLookupByEmailCall) Do(ctx context.Context) (*objects.User, error) {
	const endpoint = "users.lookupByEmail"
	v, err := c.Values()
	if err != nil {
		return nil, err
	}
	var res usersLookupByEmailCallResponseProxy
	if err := c.service.client.postForm(ctx, endpoint, v, &res); err != nil {
		return nil, errors.Wrap(err, `failed to post to users.lookupByEmail`)
	}
	if !res.OK {
		var err error
		if errresp := res.Error; errresp != nil {
			err = errors.New(errresp.String())
		} else {
			err = errors.New(`unknown error while posting to users.lookupByEmail`)
		}
		return nil, err
	}

	return res.payload()
}

// FromValues parses the data in v and populates `c`
func (c *UsersLookupByEmailCall) FromValues(v url.Values) error {
	var tmp UsersLookupByEmailCall
	if raw := strings.TrimSpace(v.Get("email")); len(raw) > 0 {
		tmp.email = raw
	}
	*c = tmp
	return nil
}

// SetActive creates a UsersSetActiveCall object in preparation for accessing the users.setActive endpoint
func (s *UsersService) SetActive() *UsersSetActiveCall {
	var call UsersSetActiveCall
	call.service = s
	return &call
}

// ValidateArgs checks that all required fields are set in the UsersSetActiveCall object
func (c *UsersSetActiveCall) ValidateArgs() error {
	return nil
}

// Values returns the UsersSetActiveCall object as url.Values
func (c *UsersSetActiveCall) Values() (url.Values, error) {
	if err := c.ValidateArgs(); err != nil {
		return nil, errors.Wrap(err, `failed validation`)
	}
	v := url.Values{}
	v.Set(`token`, c.service.token)
	return v, nil
}

type UsersSetActiveCallResponse interface {
	OK() bool
	ReplyTo() int
	Error() *objects.ErrorResponse
	Timestamp() string
}

type usersSetActiveCallResponseProxy struct {
	OK        bool                   `json:"ok"`
	ReplyTo   int                    `json:"reply_to"`
	Error     *objects.ErrorResponse `json:"error"`
	Timestamp string                 `json:"ts"`
	Payload0  json.RawMessage        `json:"-"`
}
type usersSetActiveCallResponse struct {
	ok      bool
	replyTo int
	error   *objects.ErrorResponse
	ts      string
}
type UsersSetActiveCallResponseBuilder struct {
	resp *usersSetActiveCallResponse
}

func BuildUsersSetActiveCallResponse() *UsersSetActiveCallResponseBuilder {
	return &UsersSetActiveCallResponseBuilder{resp: &usersSetActiveCallResponse{}}
}
func (v *usersSetActiveCallResponse) OK() bool {
	return v.ok
}
func (v *usersSetActiveCallResponse) ReplyTo() int {
	return v.replyTo
}
func (v *usersSetActiveCallResponse) Error() *objects.ErrorResponse {
	return v.error
}
func (v *usersSetActiveCallResponse) Timestamp() string {
	return v.ts
}
func (b *UsersSetActiveCallResponseBuilder) OK(v bool) *UsersSetActiveCallResponseBuilder {
	b.resp.ok = v
	return b
}
func (b *UsersSetActiveCallResponseBuilder) ReplyTo(v int) *UsersSetActiveCallResponseBuilder {
	b.resp.replyTo = v
	return b
}
func (b *UsersSetActiveCallResponseBuilder) Error(v *objects.ErrorResponse) *UsersSetActiveCallResponseBuilder {
	b.resp.error = v
	return b
}
func (b *UsersSetActiveCallResponseBuilder) Timestamp(v string) *UsersSetActiveCallResponseBuilder {
	b.resp.ts = v
	return b
}
func (b *UsersSetActiveCallResponseBuilder) Build() UsersSetActiveCallResponse {
	v := b.resp
	b.resp = &usersSetActiveCallResponse{}
	return v
}
func (r *usersSetActiveCallResponseProxy) parse(data []byte) error {
	if err := json.Unmarshal(data, r); err != nil {
		return errors.Wrap(err, `failed to unmarshal UsersSetActiveCallResponse`)
	}
	r.Payload0 = data
	return nil
}
func (r *usersSetActiveCallResponse) MarshalJSON() ([]byte, error) {
	var p usersSetActiveCallResponseProxy
	p.OK = r.ok
	p.ReplyTo = r.replyTo
	p.Error = r.error
	p.Timestamp = r.ts
	return json.Marshal(p)
}

// Do executes the call to access users.setActive endpoint
func (c *UsersSetActiveCall) Do(ctx context.Context) error {
	const endpoint = "users.setActive"
	v, err := c.Values()
	if err != nil {
		return err
	}
	var res usersSetActiveCallResponseProxy
	if err := c.service.client.postForm(ctx, endpoint, v, &res); err != nil {
		return errors.Wrap(err, `failed to post to users.setActive`)
	}
	if !res.OK {
		var err error
		if errresp := res.Error; errresp != nil {
			err = errors.New(errresp.String())
		} else {
			err = errors.New(`unknown error while posting to users.setActive`)
		}
		return err
	}

	return nil
}

// FromValues parses the data in v and populates `c`
func (c *UsersSetActiveCall) FromValues(v url.Values) error {
	var tmp UsersSetActiveCall
	*c = tmp
	return nil
}

// SetPresence creates a UsersSetPresenceCall object in preparation for accessing the users.setPresence endpoint
func (s *UsersService) SetPresence(presence string) *UsersSetPresenceCall {
	var call UsersSetPresenceCall
	call.service = s
	call.presence = presence
	return &call
}

// ValidateArgs checks that all required fields are set in the UsersSetPresenceCall object
func (c *UsersSetPresenceCall) ValidateArgs() error {
	if len(c.presence) <= 0 {
		return errors.New(`required field presence not initialized`)
	}
	return nil
}

// Values returns the UsersSetPresenceCall object as url.Values
func (c *UsersSetPresenceCall) Values() (url.Values, error) {
	if err := c.ValidateArgs(); err != nil {
		return nil, errors.Wrap(err, `failed validation`)
	}
	v := url.Values{}
	v.Set(`token`, c.service.token)

	v.Set("presence", c.presence)
	return v, nil
}

type UsersSetPresenceCallResponse interface {
	OK() bool
	ReplyTo() int
	Error() *objects.ErrorResponse
	Timestamp() string
}

type usersSetPresenceCallResponseProxy struct {
	OK        bool                   `json:"ok"`
	ReplyTo   int                    `json:"reply_to"`
	Error     *objects.ErrorResponse `json:"error"`
	Timestamp string                 `json:"ts"`
	Payload0  json.RawMessage        `json:"-"`
}
type usersSetPresenceCallResponse struct {
	ok      bool
	replyTo int
	error   *objects.ErrorResponse
	ts      string
}
type UsersSetPresenceCallResponseBuilder struct {
	resp *usersSetPresenceCallResponse
}

func BuildUsersSetPresenceCallResponse() *UsersSetPresenceCallResponseBuilder {
	return &UsersSetPresenceCallResponseBuilder{resp: &usersSetPresenceCallResponse{}}
}
func (v *usersSetPresenceCallResponse) OK() bool {
	return v.ok
}
func (v *usersSetPresenceCallResponse) ReplyTo() int {
	return v.replyTo
}
func (v *usersSetPresenceCallResponse) Error() *objects.ErrorResponse {
	return v.error
}
func (v *usersSetPresenceCallResponse) Timestamp() string {
	return v.ts
}
func (b *UsersSetPresenceCallResponseBuilder) OK(v bool) *UsersSetPresenceCallResponseBuilder {
	b.resp.ok = v
	return b
}
func (b *UsersSetPresenceCallResponseBuilder) ReplyTo(v int) *UsersSetPresenceCallResponseBuilder {
	b.resp.replyTo = v
	return b
}
func (b *UsersSetPresenceCallResponseBuilder) Error(v *objects.ErrorResponse) *UsersSetPresenceCallResponseBuilder {
	b.resp.error = v
	return b
}
func (b *UsersSetPresenceCallResponseBuilder) Timestamp(v string) *UsersSetPresenceCallResponseBuilder {
	b.resp.ts = v
	return b
}
func (b *UsersSetPresenceCallResponseBuilder) Build() UsersSetPresenceCallResponse {
	v := b.resp
	b.resp = &usersSetPresenceCallResponse{}
	return v
}
func (r *usersSetPresenceCallResponseProxy) parse(data []byte) error {
	if err := json.Unmarshal(data, r); err != nil {
		return errors.Wrap(err, `failed to unmarshal UsersSetPresenceCallResponse`)
	}
	r.Payload0 = data
	return nil
}
func (r *usersSetPresenceCallResponse) MarshalJSON() ([]byte, error) {
	var p usersSetPresenceCallResponseProxy
	p.OK = r.ok
	p.ReplyTo = r.replyTo
	p.Error = r.error
	p.Timestamp = r.ts
	return json.Marshal(p)
}

// Do executes the call to access users.setPresence endpoint
func (c *UsersSetPresenceCall) Do(ctx context.Context) error {
	const endpoint = "users.setPresence"
	v, err := c.Values()
	if err != nil {
		return err
	}
	var res usersSetPresenceCallResponseProxy
	if err := c.service.client.postForm(ctx, endpoint, v, &res); err != nil {
		return errors.Wrap(err, `failed to post to users.setPresence`)
	}
	if !res.OK {
		var err error
		if errresp := res.Error; errresp != nil {
			err = errors.New(errresp.String())
		} else {
			err = errors.New(`unknown error while posting to users.setPresence`)
		}
		return err
	}

	return nil
}

// FromValues parses the data in v and populates `c`
func (c *UsersSetPresenceCall) FromValues(v url.Values) error {
	var tmp UsersSetPresenceCall
	if raw := strings.TrimSpace(v.Get("presence")); len(raw) > 0 {
		tmp.presence = raw
	}
	*c = tmp
	return nil
}
