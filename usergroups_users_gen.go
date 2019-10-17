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

// Do executes the call to access usergroups.users.list endpoint
func (c *UsergroupsUsersListCall) Do(ctx context.Context) (objects.UsergroupUsersList, error) {
	const endpoint = "usergroups.users.list"
	v, err := c.Values()
	if err != nil {
		return nil, err
	}
	var res struct {
		objects.GenericResponse
		objects.UsergroupUsersList `json:"users"`
	}
	if err := c.service.client.postForm(ctx, endpoint, v, &res); err != nil {
		return nil, errors.Wrap(err, `failed to post to usergroups.users.list`)
	}
	if !res.OK() {
		return nil, errors.New(res.Error().String())
	}

	return res.UsergroupUsersList, nil
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

// Do executes the call to access usergroups.users.update endpoint
func (c *UsergroupsUsersUpdateCall) Do(ctx context.Context) (*objects.Usergroup, error) {
	const endpoint = "usergroups.users.update"
	v, err := c.Values()
	if err != nil {
		return nil, err
	}
	var res struct {
		objects.GenericResponse
		*objects.Usergroup `json:"usergroup"`
	}
	if err := c.service.client.postForm(ctx, endpoint, v, &res); err != nil {
		return nil, errors.Wrap(err, `failed to post to usergroups.users.update`)
	}
	if !res.OK() {
		return nil, errors.New(res.Error().String())
	}

	return res.Usergroup, nil
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
