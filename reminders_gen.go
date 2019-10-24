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

// RemindersAddCall is created by RemindersService.Add method call
type RemindersAddCall struct {
	service *RemindersService
	text    string
	time    int
	user    string
}

// RemindersCompleteCall is created by RemindersService.Complete method call
type RemindersCompleteCall struct {
	service  *RemindersService
	reminder string
}

// RemindersDeleteCall is created by RemindersService.Delete method call
type RemindersDeleteCall struct {
	service  *RemindersService
	reminder string
}

// RemindersInfoCall is created by RemindersService.Info method call
type RemindersInfoCall struct {
	service  *RemindersService
	reminder string
}

// RemindersListCall is created by RemindersService.List method call
type RemindersListCall struct {
	service *RemindersService
}

// Add creates a RemindersAddCall object in preparation for accessing the reminders.add endpoint
func (s *RemindersService) Add(text string, time int) *RemindersAddCall {
	var call RemindersAddCall
	call.service = s
	call.text = text
	call.time = time
	return &call
}

// User sets the value for optional user parameter
func (c *RemindersAddCall) User(user string) *RemindersAddCall {
	c.user = user
	return c
}

// ValidateArgs checks that all required fields are set in the RemindersAddCall object
func (c *RemindersAddCall) ValidateArgs() error {
	if len(c.text) <= 0 {
		return errors.New(`required field text not initialized`)
	}
	if c.time == 0 {
		return errors.New(`required field time not initialized`)
	}
	return nil
}

// Values returns the RemindersAddCall object as url.Values
func (c *RemindersAddCall) Values() (url.Values, error) {
	if err := c.ValidateArgs(); err != nil {
		return nil, errors.Wrap(err, `failed validation`)
	}
	v := url.Values{}
	v.Set(`token`, c.service.token)

	v.Set("text", c.text)

	v.Set("time", strconv.Itoa(c.time))

	if len(c.user) > 0 {
		v.Set("user", c.user)
	}
	return v, nil
}

type RemindersAddCallResponse interface {
	OK() bool
	ReplyTo() int
	Error() *objects.ErrorResponse
	Timestamp() string
}

type remindersAddCallResponseProxy struct {
	OK        bool                   `json:"ok"`
	ReplyTo   int                    `json:"reply_to"`
	Error     *objects.ErrorResponse `json:"error"`
	Timestamp string                 `json:"ts"`
	Payload0  json.RawMessage        `json:"-"`
}
type remindersAddCallResponse struct {
	ok      bool
	replyTo int
	error   *objects.ErrorResponse
	ts      string
}
type RemindersAddCallResponseBuilder struct {
	resp *remindersAddCallResponse
}

func BuildRemindersAddCallResponse() *RemindersAddCallResponseBuilder {
	return &RemindersAddCallResponseBuilder{resp: &remindersAddCallResponse{}}
}
func (v *remindersAddCallResponse) OK() bool {
	return v.ok
}
func (v *remindersAddCallResponse) ReplyTo() int {
	return v.replyTo
}
func (v *remindersAddCallResponse) Error() *objects.ErrorResponse {
	return v.error
}
func (v *remindersAddCallResponse) Timestamp() string {
	return v.ts
}
func (b *RemindersAddCallResponseBuilder) OK(v bool) *RemindersAddCallResponseBuilder {
	b.resp.ok = v
	return b
}
func (b *RemindersAddCallResponseBuilder) ReplyTo(v int) *RemindersAddCallResponseBuilder {
	b.resp.replyTo = v
	return b
}
func (b *RemindersAddCallResponseBuilder) Error(v *objects.ErrorResponse) *RemindersAddCallResponseBuilder {
	b.resp.error = v
	return b
}
func (b *RemindersAddCallResponseBuilder) Timestamp(v string) *RemindersAddCallResponseBuilder {
	b.resp.ts = v
	return b
}
func (b *RemindersAddCallResponseBuilder) Build() RemindersAddCallResponse {
	v := b.resp
	b.resp = &remindersAddCallResponse{}
	return v
}
func (r *remindersAddCallResponseProxy) parse(data []byte) error {
	if err := json.Unmarshal(data, r); err != nil {
		return errors.Wrap(err, `failed to unmarshal RemindersAddCallResponse`)
	}
	r.Payload0 = data
	return nil
}
func (r *remindersAddCallResponseProxy) payload() (*objects.Reminder, error) {
	var res0 objects.Reminder
	if err := json.Unmarshal(r.Payload0, &res0); err != nil {
		return nil, errors.Wrap(err, `failed to ummarshal objects.Reminder from response`)
	}
	return &res0, nil
}

// Do executes the call to access reminders.add endpoint
func (c *RemindersAddCall) Do(ctx context.Context) (*objects.Reminder, error) {
	const endpoint = "reminders.add"
	v, err := c.Values()
	if err != nil {
		return nil, err
	}
	var res remindersAddCallResponseProxy
	if err := c.service.client.postForm(ctx, endpoint, v, &res); err != nil {
		return nil, errors.Wrap(err, `failed to post to reminders.add`)
	}
	if !res.OK {
		var err error
		if errresp := res.Error; errresp != nil {
			err = errors.New(errresp.String())
		} else {
			err = errors.New(`unknown error while posting to reminders.add`)
		}
		return nil, err
	}

	return res.payload()
}

// FromValues parses the data in v and populates `c`
func (c *RemindersAddCall) FromValues(v url.Values) error {
	var tmp RemindersAddCall
	if raw := strings.TrimSpace(v.Get("text")); len(raw) > 0 {
		tmp.text = raw
	}
	if raw := strings.TrimSpace(v.Get("time")); len(raw) > 0 {
		parsed, err := strconv.ParseInt(raw, 10, 64)
		if err != nil {
			return errors.Wrap(err, `failed to parse integer value "time"`)
		}
		tmp.time = int(parsed)
	}
	if raw := strings.TrimSpace(v.Get("user")); len(raw) > 0 {
		tmp.user = raw
	}
	*c = tmp
	return nil
}

// Complete creates a RemindersCompleteCall object in preparation for accessing the reminders.complete endpoint
func (s *RemindersService) Complete(reminder string) *RemindersCompleteCall {
	var call RemindersCompleteCall
	call.service = s
	call.reminder = reminder
	return &call
}

// ValidateArgs checks that all required fields are set in the RemindersCompleteCall object
func (c *RemindersCompleteCall) ValidateArgs() error {
	if len(c.reminder) <= 0 {
		return errors.New(`required field reminder not initialized`)
	}
	return nil
}

// Values returns the RemindersCompleteCall object as url.Values
func (c *RemindersCompleteCall) Values() (url.Values, error) {
	if err := c.ValidateArgs(); err != nil {
		return nil, errors.Wrap(err, `failed validation`)
	}
	v := url.Values{}
	v.Set(`token`, c.service.token)

	v.Set("reminder", c.reminder)
	return v, nil
}

type RemindersCompleteCallResponse interface {
	OK() bool
	ReplyTo() int
	Error() *objects.ErrorResponse
	Timestamp() string
}

type remindersCompleteCallResponseProxy struct {
	OK        bool                   `json:"ok"`
	ReplyTo   int                    `json:"reply_to"`
	Error     *objects.ErrorResponse `json:"error"`
	Timestamp string                 `json:"ts"`
	Payload0  json.RawMessage        `json:"-"`
}
type remindersCompleteCallResponse struct {
	ok      bool
	replyTo int
	error   *objects.ErrorResponse
	ts      string
}
type RemindersCompleteCallResponseBuilder struct {
	resp *remindersCompleteCallResponse
}

func BuildRemindersCompleteCallResponse() *RemindersCompleteCallResponseBuilder {
	return &RemindersCompleteCallResponseBuilder{resp: &remindersCompleteCallResponse{}}
}
func (v *remindersCompleteCallResponse) OK() bool {
	return v.ok
}
func (v *remindersCompleteCallResponse) ReplyTo() int {
	return v.replyTo
}
func (v *remindersCompleteCallResponse) Error() *objects.ErrorResponse {
	return v.error
}
func (v *remindersCompleteCallResponse) Timestamp() string {
	return v.ts
}
func (b *RemindersCompleteCallResponseBuilder) OK(v bool) *RemindersCompleteCallResponseBuilder {
	b.resp.ok = v
	return b
}
func (b *RemindersCompleteCallResponseBuilder) ReplyTo(v int) *RemindersCompleteCallResponseBuilder {
	b.resp.replyTo = v
	return b
}
func (b *RemindersCompleteCallResponseBuilder) Error(v *objects.ErrorResponse) *RemindersCompleteCallResponseBuilder {
	b.resp.error = v
	return b
}
func (b *RemindersCompleteCallResponseBuilder) Timestamp(v string) *RemindersCompleteCallResponseBuilder {
	b.resp.ts = v
	return b
}
func (b *RemindersCompleteCallResponseBuilder) Build() RemindersCompleteCallResponse {
	v := b.resp
	b.resp = &remindersCompleteCallResponse{}
	return v
}
func (r *remindersCompleteCallResponseProxy) parse(data []byte) error {
	if err := json.Unmarshal(data, r); err != nil {
		return errors.Wrap(err, `failed to unmarshal RemindersCompleteCallResponse`)
	}
	r.Payload0 = data
	return nil
}

// Do executes the call to access reminders.complete endpoint
func (c *RemindersCompleteCall) Do(ctx context.Context) error {
	const endpoint = "reminders.complete"
	v, err := c.Values()
	if err != nil {
		return err
	}
	var res remindersCompleteCallResponseProxy
	if err := c.service.client.postForm(ctx, endpoint, v, &res); err != nil {
		return errors.Wrap(err, `failed to post to reminders.complete`)
	}
	if !res.OK {
		var err error
		if errresp := res.Error; errresp != nil {
			err = errors.New(errresp.String())
		} else {
			err = errors.New(`unknown error while posting to reminders.complete`)
		}
		return err
	}

	return nil
}

// FromValues parses the data in v and populates `c`
func (c *RemindersCompleteCall) FromValues(v url.Values) error {
	var tmp RemindersCompleteCall
	if raw := strings.TrimSpace(v.Get("reminder")); len(raw) > 0 {
		tmp.reminder = raw
	}
	*c = tmp
	return nil
}

// Delete creates a RemindersDeleteCall object in preparation for accessing the reminders.delete endpoint
func (s *RemindersService) Delete(reminder string) *RemindersDeleteCall {
	var call RemindersDeleteCall
	call.service = s
	call.reminder = reminder
	return &call
}

// ValidateArgs checks that all required fields are set in the RemindersDeleteCall object
func (c *RemindersDeleteCall) ValidateArgs() error {
	if len(c.reminder) <= 0 {
		return errors.New(`required field reminder not initialized`)
	}
	return nil
}

// Values returns the RemindersDeleteCall object as url.Values
func (c *RemindersDeleteCall) Values() (url.Values, error) {
	if err := c.ValidateArgs(); err != nil {
		return nil, errors.Wrap(err, `failed validation`)
	}
	v := url.Values{}
	v.Set(`token`, c.service.token)

	v.Set("reminder", c.reminder)
	return v, nil
}

type RemindersDeleteCallResponse interface {
	OK() bool
	ReplyTo() int
	Error() *objects.ErrorResponse
	Timestamp() string
}

type remindersDeleteCallResponseProxy struct {
	OK        bool                   `json:"ok"`
	ReplyTo   int                    `json:"reply_to"`
	Error     *objects.ErrorResponse `json:"error"`
	Timestamp string                 `json:"ts"`
	Payload0  json.RawMessage        `json:"-"`
}
type remindersDeleteCallResponse struct {
	ok      bool
	replyTo int
	error   *objects.ErrorResponse
	ts      string
}
type RemindersDeleteCallResponseBuilder struct {
	resp *remindersDeleteCallResponse
}

func BuildRemindersDeleteCallResponse() *RemindersDeleteCallResponseBuilder {
	return &RemindersDeleteCallResponseBuilder{resp: &remindersDeleteCallResponse{}}
}
func (v *remindersDeleteCallResponse) OK() bool {
	return v.ok
}
func (v *remindersDeleteCallResponse) ReplyTo() int {
	return v.replyTo
}
func (v *remindersDeleteCallResponse) Error() *objects.ErrorResponse {
	return v.error
}
func (v *remindersDeleteCallResponse) Timestamp() string {
	return v.ts
}
func (b *RemindersDeleteCallResponseBuilder) OK(v bool) *RemindersDeleteCallResponseBuilder {
	b.resp.ok = v
	return b
}
func (b *RemindersDeleteCallResponseBuilder) ReplyTo(v int) *RemindersDeleteCallResponseBuilder {
	b.resp.replyTo = v
	return b
}
func (b *RemindersDeleteCallResponseBuilder) Error(v *objects.ErrorResponse) *RemindersDeleteCallResponseBuilder {
	b.resp.error = v
	return b
}
func (b *RemindersDeleteCallResponseBuilder) Timestamp(v string) *RemindersDeleteCallResponseBuilder {
	b.resp.ts = v
	return b
}
func (b *RemindersDeleteCallResponseBuilder) Build() RemindersDeleteCallResponse {
	v := b.resp
	b.resp = &remindersDeleteCallResponse{}
	return v
}
func (r *remindersDeleteCallResponseProxy) parse(data []byte) error {
	if err := json.Unmarshal(data, r); err != nil {
		return errors.Wrap(err, `failed to unmarshal RemindersDeleteCallResponse`)
	}
	r.Payload0 = data
	return nil
}

// Do executes the call to access reminders.delete endpoint
func (c *RemindersDeleteCall) Do(ctx context.Context) error {
	const endpoint = "reminders.delete"
	v, err := c.Values()
	if err != nil {
		return err
	}
	var res remindersDeleteCallResponseProxy
	if err := c.service.client.postForm(ctx, endpoint, v, &res); err != nil {
		return errors.Wrap(err, `failed to post to reminders.delete`)
	}
	if !res.OK {
		var err error
		if errresp := res.Error; errresp != nil {
			err = errors.New(errresp.String())
		} else {
			err = errors.New(`unknown error while posting to reminders.delete`)
		}
		return err
	}

	return nil
}

// FromValues parses the data in v and populates `c`
func (c *RemindersDeleteCall) FromValues(v url.Values) error {
	var tmp RemindersDeleteCall
	if raw := strings.TrimSpace(v.Get("reminder")); len(raw) > 0 {
		tmp.reminder = raw
	}
	*c = tmp
	return nil
}

// Info creates a RemindersInfoCall object in preparation for accessing the reminders.info endpoint
func (s *RemindersService) Info(reminder string) *RemindersInfoCall {
	var call RemindersInfoCall
	call.service = s
	call.reminder = reminder
	return &call
}

// ValidateArgs checks that all required fields are set in the RemindersInfoCall object
func (c *RemindersInfoCall) ValidateArgs() error {
	if len(c.reminder) <= 0 {
		return errors.New(`required field reminder not initialized`)
	}
	return nil
}

// Values returns the RemindersInfoCall object as url.Values
func (c *RemindersInfoCall) Values() (url.Values, error) {
	if err := c.ValidateArgs(); err != nil {
		return nil, errors.Wrap(err, `failed validation`)
	}
	v := url.Values{}
	v.Set(`token`, c.service.token)

	v.Set("reminder", c.reminder)
	return v, nil
}

type RemindersInfoCallResponse interface {
	OK() bool
	ReplyTo() int
	Error() *objects.ErrorResponse
	Timestamp() string
}

type remindersInfoCallResponseProxy struct {
	OK        bool                   `json:"ok"`
	ReplyTo   int                    `json:"reply_to"`
	Error     *objects.ErrorResponse `json:"error"`
	Timestamp string                 `json:"ts"`
	Payload0  json.RawMessage        `json:"-"`
}
type remindersInfoCallResponse struct {
	ok      bool
	replyTo int
	error   *objects.ErrorResponse
	ts      string
}
type RemindersInfoCallResponseBuilder struct {
	resp *remindersInfoCallResponse
}

func BuildRemindersInfoCallResponse() *RemindersInfoCallResponseBuilder {
	return &RemindersInfoCallResponseBuilder{resp: &remindersInfoCallResponse{}}
}
func (v *remindersInfoCallResponse) OK() bool {
	return v.ok
}
func (v *remindersInfoCallResponse) ReplyTo() int {
	return v.replyTo
}
func (v *remindersInfoCallResponse) Error() *objects.ErrorResponse {
	return v.error
}
func (v *remindersInfoCallResponse) Timestamp() string {
	return v.ts
}
func (b *RemindersInfoCallResponseBuilder) OK(v bool) *RemindersInfoCallResponseBuilder {
	b.resp.ok = v
	return b
}
func (b *RemindersInfoCallResponseBuilder) ReplyTo(v int) *RemindersInfoCallResponseBuilder {
	b.resp.replyTo = v
	return b
}
func (b *RemindersInfoCallResponseBuilder) Error(v *objects.ErrorResponse) *RemindersInfoCallResponseBuilder {
	b.resp.error = v
	return b
}
func (b *RemindersInfoCallResponseBuilder) Timestamp(v string) *RemindersInfoCallResponseBuilder {
	b.resp.ts = v
	return b
}
func (b *RemindersInfoCallResponseBuilder) Build() RemindersInfoCallResponse {
	v := b.resp
	b.resp = &remindersInfoCallResponse{}
	return v
}
func (r *remindersInfoCallResponseProxy) parse(data []byte) error {
	if err := json.Unmarshal(data, r); err != nil {
		return errors.Wrap(err, `failed to unmarshal RemindersInfoCallResponse`)
	}
	r.Payload0 = data
	return nil
}
func (r *remindersInfoCallResponseProxy) payload() (*objects.Reminder, error) {
	var res0 objects.Reminder
	if err := json.Unmarshal(r.Payload0, &res0); err != nil {
		return nil, errors.Wrap(err, `failed to ummarshal objects.Reminder from response`)
	}
	return &res0, nil
}

// Do executes the call to access reminders.info endpoint
func (c *RemindersInfoCall) Do(ctx context.Context) (*objects.Reminder, error) {
	const endpoint = "reminders.info"
	v, err := c.Values()
	if err != nil {
		return nil, err
	}
	var res remindersInfoCallResponseProxy
	if err := c.service.client.postForm(ctx, endpoint, v, &res); err != nil {
		return nil, errors.Wrap(err, `failed to post to reminders.info`)
	}
	if !res.OK {
		var err error
		if errresp := res.Error; errresp != nil {
			err = errors.New(errresp.String())
		} else {
			err = errors.New(`unknown error while posting to reminders.info`)
		}
		return nil, err
	}

	return res.payload()
}

// FromValues parses the data in v and populates `c`
func (c *RemindersInfoCall) FromValues(v url.Values) error {
	var tmp RemindersInfoCall
	if raw := strings.TrimSpace(v.Get("reminder")); len(raw) > 0 {
		tmp.reminder = raw
	}
	*c = tmp
	return nil
}

// List creates a RemindersListCall object in preparation for accessing the reminders.list endpoint
func (s *RemindersService) List() *RemindersListCall {
	var call RemindersListCall
	call.service = s
	return &call
}

// ValidateArgs checks that all required fields are set in the RemindersListCall object
func (c *RemindersListCall) ValidateArgs() error {
	return nil
}

// Values returns the RemindersListCall object as url.Values
func (c *RemindersListCall) Values() (url.Values, error) {
	if err := c.ValidateArgs(); err != nil {
		return nil, errors.Wrap(err, `failed validation`)
	}
	v := url.Values{}
	v.Set(`token`, c.service.token)
	return v, nil
}

type RemindersListCallResponse interface {
	OK() bool
	ReplyTo() int
	Error() *objects.ErrorResponse
	Timestamp() string
	Reminders() *objects.ReminderList
}

type remindersListCallResponseProxy struct {
	OK        bool                   `json:"ok"`
	ReplyTo   int                    `json:"reply_to"`
	Error     *objects.ErrorResponse `json:"error"`
	Timestamp string                 `json:"ts"`
	Payload0  json.RawMessage        `json:"-"`
	Payload1  json.RawMessage        `json:"reminders"`
}
type remindersListCallResponse struct {
	ok        bool
	replyTo   int
	error     *objects.ErrorResponse
	ts        string
	reminders *objects.ReminderList
}
type RemindersListCallResponseBuilder struct {
	resp *remindersListCallResponse
}

func BuildRemindersListCallResponse() *RemindersListCallResponseBuilder {
	return &RemindersListCallResponseBuilder{resp: &remindersListCallResponse{}}
}
func (v *remindersListCallResponse) OK() bool {
	return v.ok
}
func (v *remindersListCallResponse) ReplyTo() int {
	return v.replyTo
}
func (v *remindersListCallResponse) Error() *objects.ErrorResponse {
	return v.error
}
func (v *remindersListCallResponse) Timestamp() string {
	return v.ts
}
func (v *remindersListCallResponse) Reminders() *objects.ReminderList {
	return v.reminders
}
func (b *RemindersListCallResponseBuilder) OK(v bool) *RemindersListCallResponseBuilder {
	b.resp.ok = v
	return b
}
func (b *RemindersListCallResponseBuilder) ReplyTo(v int) *RemindersListCallResponseBuilder {
	b.resp.replyTo = v
	return b
}
func (b *RemindersListCallResponseBuilder) Error(v *objects.ErrorResponse) *RemindersListCallResponseBuilder {
	b.resp.error = v
	return b
}
func (b *RemindersListCallResponseBuilder) Timestamp(v string) *RemindersListCallResponseBuilder {
	b.resp.ts = v
	return b
}
func (b *RemindersListCallResponseBuilder) Reminders(v *objects.ReminderList) *RemindersListCallResponseBuilder {
	b.resp.reminders = v
	return b
}
func (b *RemindersListCallResponseBuilder) Build() RemindersListCallResponse {
	v := b.resp
	b.resp = &remindersListCallResponse{}
	return v
}
func (r *remindersListCallResponseProxy) parse(data []byte) error {
	if err := json.Unmarshal(data, r); err != nil {
		return errors.Wrap(err, `failed to unmarshal RemindersListCallResponse`)
	}
	r.Payload0 = data
	return nil
}
func (r *remindersListCallResponseProxy) payload() (objects.ReminderList, error) {
	var res1 objects.ReminderList
	if err := json.Unmarshal(r.Payload1, &res1); err != nil {
		return nil, errors.Wrap(err, `failed to ummarshal objects.ReminderList from response`)
	}
	return res1, nil
}

// Do executes the call to access reminders.list endpoint
func (c *RemindersListCall) Do(ctx context.Context) (objects.ReminderList, error) {
	const endpoint = "reminders.list"
	v, err := c.Values()
	if err != nil {
		return nil, err
	}
	var res remindersListCallResponseProxy
	if err := c.service.client.postForm(ctx, endpoint, v, &res); err != nil {
		return nil, errors.Wrap(err, `failed to post to reminders.list`)
	}
	if !res.OK {
		var err error
		if errresp := res.Error; errresp != nil {
			err = errors.New(errresp.String())
		} else {
			err = errors.New(`unknown error while posting to reminders.list`)
		}
		return nil, err
	}

	return res.payload()
}

// FromValues parses the data in v and populates `c`
func (c *RemindersListCall) FromValues(v url.Values) error {
	var tmp RemindersListCall
	*c = tmp
	return nil
}
