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

// DialogOpenCall is created by DialogService.Open method call
type DialogOpenCall struct {
	service    *DialogService
	dialog     *objects.Dialog
	trigger_id string
}

// Open creates a DialogOpenCall object in preparation for accessing the dialog.open endpoint
func (s *DialogService) Open(dialog *objects.Dialog, trigger_id string) *DialogOpenCall {
	var call DialogOpenCall
	call.service = s
	call.dialog = dialog
	call.trigger_id = trigger_id
	return &call
}

// ValidateArgs checks that all required fields are set in the DialogOpenCall object
func (c *DialogOpenCall) ValidateArgs() error {
	if c.dialog == nil {
		return errors.New(`required field dialog not initialized`)
	}
	if len(c.trigger_id) <= 0 {
		return errors.New(`required field trigger_id not initialized`)
	}
	return nil
}

// Values returns the DialogOpenCall object as url.Values
func (c *DialogOpenCall) Values() (url.Values, error) {
	if err := c.ValidateArgs(); err != nil {
		return nil, errors.Wrap(err, `failed validation`)
	}
	v := url.Values{}
	v.Set(`token`, c.service.token)

	dialogEncoded, err := json.Marshal(c.dialog)
	if err != nil {
		return nil, errors.Wrap(err, `failed to encode field`)
	}
	v.Set("dialog", string(dialogEncoded))

	v.Set("trigger_id", c.trigger_id)
	return v, nil
}

type DialogOpenCallResponse interface {
	OK() bool
	ReplyTo() int
	Error() *objects.ErrorResponse
	Timestamp() string
}

type dialogOpenCallResponseProxy struct {
	OK        bool                   `json:"ok"`
	ReplyTo   int                    `json:"reply_to"`
	Error     *objects.ErrorResponse `json:"error"`
	Timestamp string                 `json:"ts"`
	Payload0  json.RawMessage        `json:"-"`
}
type dialogOpenCallResponse struct {
	ok      bool
	replyTo int
	error   *objects.ErrorResponse
	ts      string
}
type DialogOpenCallResponseBuilder struct {
	resp *dialogOpenCallResponse
}

func BuildDialogOpenCallResponse() *DialogOpenCallResponseBuilder {
	return &DialogOpenCallResponseBuilder{resp: &dialogOpenCallResponse{}}
}
func (v *dialogOpenCallResponse) OK() bool {
	return v.ok
}
func (v *dialogOpenCallResponse) ReplyTo() int {
	return v.replyTo
}
func (v *dialogOpenCallResponse) Error() *objects.ErrorResponse {
	return v.error
}
func (v *dialogOpenCallResponse) Timestamp() string {
	return v.ts
}
func (b *DialogOpenCallResponseBuilder) OK(v bool) *DialogOpenCallResponseBuilder {
	b.resp.ok = v
	return b
}
func (b *DialogOpenCallResponseBuilder) ReplyTo(v int) *DialogOpenCallResponseBuilder {
	b.resp.replyTo = v
	return b
}
func (b *DialogOpenCallResponseBuilder) Error(v *objects.ErrorResponse) *DialogOpenCallResponseBuilder {
	b.resp.error = v
	return b
}
func (b *DialogOpenCallResponseBuilder) Timestamp(v string) *DialogOpenCallResponseBuilder {
	b.resp.ts = v
	return b
}
func (b *DialogOpenCallResponseBuilder) Build() DialogOpenCallResponse {
	v := b.resp
	b.resp = &dialogOpenCallResponse{}
	return v
}
func (r *dialogOpenCallResponseProxy) parse(data []byte) error {
	if err := json.Unmarshal(data, r); err != nil {
		return errors.Wrap(err, `failed to unmarshal DialogOpenCallResponse`)
	}
	r.Payload0 = data
	return nil
}
func (r *dialogOpenCallResponseProxy) payload() (*objects.DialogResponse, error) {
	var res0 objects.DialogResponse
	if err := json.Unmarshal(r.Payload0, &res0); err != nil {
		return nil, errors.Wrap(err, `failed to ummarshal objects.DialogResponse from response`)
	}
	return &res0, nil
}
func (r *dialogOpenCallResponse) MarshalJSON() ([]byte, error) {
	var p dialogOpenCallResponseProxy
	p.OK = r.ok
	p.ReplyTo = r.replyTo
	p.Error = r.error
	p.Timestamp = r.ts
	return json.Marshal(p)
}

// Do executes the call to access dialog.open endpoint
func (c *DialogOpenCall) Do(ctx context.Context) (*objects.DialogResponse, error) {
	const endpoint = "dialog.open"
	v, err := c.Values()
	if err != nil {
		return nil, err
	}
	var res dialogOpenCallResponseProxy
	if err := c.service.client.postForm(ctx, endpoint, v, &res); err != nil {
		return nil, errors.Wrap(err, `failed to post to dialog.open`)
	}
	if !res.OK {
		var err error
		if errresp := res.Error; errresp != nil {
			err = errors.New(errresp.String())
		} else {
			err = errors.New(`unknown error while posting to dialog.open`)
		}
		return nil, err
	}

	return res.payload()
}

// FromValues parses the data in v and populates `c`
func (c *DialogOpenCall) FromValues(v url.Values) error {
	var tmp DialogOpenCall
	if raw := strings.TrimSpace(v.Get("dialog")); len(raw) > 0 {
		if err := json.Unmarshal([]byte(raw), &tmp.dialog); err != nil {
			return errors.Wrap(err, `failed to decode value "dialog"`)
		}
	}
	if raw := strings.TrimSpace(v.Get("trigger_id")); len(raw) > 0 {
		tmp.trigger_id = raw
	}
	*c = tmp
	return nil
}
