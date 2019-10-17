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

// Do executes the call to access reminders.add endpoint
func (c *RemindersAddCall) Do(ctx context.Context) (*objects.Reminder, error) {
	const endpoint = "reminders.add"
	v, err := c.Values()
	if err != nil {
		return nil, err
	}
	var res struct {
		objects.GenericResponse
		*objects.Reminder
	}
	if err := c.service.client.postForm(ctx, endpoint, v, &res); err != nil {
		return nil, errors.Wrap(err, `failed to post to reminders.add`)
	}
	if !res.OK() {
		return nil, errors.New(res.Error().String())
	}

	return res.Reminder, nil
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

// Do executes the call to access reminders.complete endpoint
func (c *RemindersCompleteCall) Do(ctx context.Context) error {
	const endpoint = "reminders.complete"
	v, err := c.Values()
	if err != nil {
		return err
	}
	var res struct {
		objects.GenericResponse
	}
	if err := c.service.client.postForm(ctx, endpoint, v, &res); err != nil {
		return errors.Wrap(err, `failed to post to reminders.complete`)
	}
	if !res.OK() {
		return errors.New(res.Error().String())
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

// Do executes the call to access reminders.delete endpoint
func (c *RemindersDeleteCall) Do(ctx context.Context) error {
	const endpoint = "reminders.delete"
	v, err := c.Values()
	if err != nil {
		return err
	}
	var res struct {
		objects.GenericResponse
	}
	if err := c.service.client.postForm(ctx, endpoint, v, &res); err != nil {
		return errors.Wrap(err, `failed to post to reminders.delete`)
	}
	if !res.OK() {
		return errors.New(res.Error().String())
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

// Do executes the call to access reminders.info endpoint
func (c *RemindersInfoCall) Do(ctx context.Context) (*objects.Reminder, error) {
	const endpoint = "reminders.info"
	v, err := c.Values()
	if err != nil {
		return nil, err
	}
	var res struct {
		objects.GenericResponse
		*objects.Reminder
	}
	if err := c.service.client.postForm(ctx, endpoint, v, &res); err != nil {
		return nil, errors.Wrap(err, `failed to post to reminders.info`)
	}
	if !res.OK() {
		return nil, errors.New(res.Error().String())
	}

	return res.Reminder, nil
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

// Do executes the call to access reminders.list endpoint
func (c *RemindersListCall) Do(ctx context.Context) (objects.ReminderList, error) {
	const endpoint = "reminders.list"
	v, err := c.Values()
	if err != nil {
		return nil, err
	}
	var res struct {
		objects.GenericResponse
		objects.ReminderList `json:"reminders"`
	}
	if err := c.service.client.postForm(ctx, endpoint, v, &res); err != nil {
		return nil, errors.Wrap(err, `failed to post to reminders.list`)
	}
	if !res.OK() {
		return nil, errors.New(res.Error().String())
	}

	return res.ReminderList, nil
}

// FromValues parses the data in v and populates `c`
func (c *RemindersListCall) FromValues(v url.Values) error {
	var tmp RemindersListCall
	*c = tmp
	return nil
}
