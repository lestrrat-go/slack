package objects

import (
	"bytes"
	"encoding/json"

	"github.com/pkg/errors"
)

type messageElementType string

const (
	messageElementTImage      messageElementType = "image"
	messageElementTButton     messageElementType = "button"
	messageElementTOverflow   messageElementType = "overflow"
	messageElementTDatepicker messageElementType = "datepicker"
)

type contextElementType string

const (
	contextElementImage contextElementType = "mixed_image"
	contextElementText  contextElementType = "mixed_text"
)

type Accessory struct {
	ImageElement      *ImageBlockElement
	ButtonElement     *ButtonBlockElement
	OverflowElement   *OverflowBlockElement
	DatePickerElement *DatePickerElement
	SelectElement     SelectBlockElement
}

func (b *ImageBlockElement) Encode() (string, error) {
	block := struct {
		*ImageBlockElement
		Type messageElementType `json:"type"`
	}{
		ImageBlockElement: &ImageBlockElement{
			ImageURL: b.ImageURL,
			AltText:  b.AltText,
		},
		Type: b.messageElementType(),
	}
	buf, err := json.Marshal(&block)
	if err != nil {
		return "", errors.Wrap(err, `failed to encode image block element`)
	}
	return string(buf), nil
}

func (b *ImageBlockElement) Decode(buf string) error {
	if buf == "" {
		return nil
	}
	r := bytes.NewReader([]byte(buf))
	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()

	return dec.Decode(b)
}

type ButtonBlockStyle string

const (
	StyleDefault ButtonBlockStyle = "default"
	StylePrimary ButtonBlockStyle = "primary"
	StyleDanger  ButtonBlockStyle = "danger"
)

func (b *ButtonBlockElement) Encode() (string, error) {
	block := struct {
		*ButtonBlockElement
		Type messageElementType `json:"type"`
	}{
		ButtonBlockElement: &ButtonBlockElement{
			Text:     b.Text,
			ActionID: b.ActionID,
			URL:      b.URL,
			Value:    b.Value,
			Confirm:  b.Confirm,
			Style:    b.Style,
		},
		Type: b.messageElementType(),
	}
	buf, err := json.Marshal(&block)
	if err != nil {
		return "", errors.Wrap(err, `failed to encode button block element`)
	}
	return string(buf), nil
}

func (b *ButtonBlockElement) Decode(buf string) error {
	if buf == "" {
		return nil
	}
	r := bytes.NewReader([]byte(buf))
	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()

	return dec.Decode(b)
}

type selectBlockElementType string

const (
	selectBlockElementTypeStatic        selectBlockElementType = "static_select"
	selectBlockElementTypeExternal      selectBlockElementType = "external_select"
	selectBlockElementTypeUsers         selectBlockElementType = "users_select"
	selectBlockElementTypeConversations selectBlockElementType = "conversations_select"
	selectBlockElementTypeChannels      selectBlockElementType = "channels_select"
)

type multiSelectBlockElementType string

const (
	multiSelectBlockElementTypeStatic        multiSelectBlockElementType = "multi_static_select"
	multiSelectBlockElementTypeExternal      multiSelectBlockElementType = "multi_external_select"
	multiSelectBlockElementTypeUsers         multiSelectBlockElementType = "multi_users_select"
	multiSelectBlockElementTypeConversations multiSelectBlockElementType = "multi_conversations_select"
	multiSelectBlockElementTypeChannels      multiSelectBlockElementType = "multi_channels_select"
)

type selectBlockElementStatic struct {
	Placeholder   TextBlockObject                `json:"placeholder,omitempty"`
	ActionID      string                         `json:"action_id,omitempty"`
	Options       []*OptionBlockObject           `json:"options,omitempty"`
	OptionGroups  []*OptionGroupBlockObject      `json:"option_groups"`
	InitialOption *OptionBlockObject             `json:"initial_option"`
	Confirm       *ConfirmationDialogBlockObject `json:"confirm"`
}

func (b *selectBlockElementStatic) Encode(elementType messageElementType) (string, error) {
	block := struct {
		*selectBlockElementStatic
		Type messageElementType `json:"type"`
	}{
		selectBlockElementStatic: &selectBlockElementStatic{
			Placeholder:   b.Placeholder,
			ActionID:      b.ActionID,
			Options:       b.Options,
			OptionGroups:  b.OptionGroups,
			InitialOption: b.InitialOption,
			Confirm:       b.Confirm,
		},
		Type: elementType,
	}
	buf, err := json.Marshal(&block)
	if err != nil {
		return "", errors.Wrap(err, `failed to encode image block element`)
	}
	return string(buf), nil
}

func (b *selectBlockElementStatic) Decode(buf string) error {
	if buf == "" {
		return nil
	}
	r := bytes.NewReader([]byte(buf))
	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()

	return dec.Decode(b)
}

func (b *SelectBlockElementStatic) Encode() (string, error) {
	return b.selectBlockElementStatic.Encode(b.selectBlockElementType())
}

func (b *SelectBlockElementStatic) Decode(buf string) error {
	return b.selectBlockElementStatic.Decode(buf)
}

func (b *MultiSelectBlockElementStatic) Encode() (string, error) {
	return b.selectBlockElementStatic.Encode(b.selectBlockElementType())
}

func (b *MultiSelectBlockElementStatic) Decode(buf string) error {
	return b.selectBlockElementStatic.Decode(buf)
}

type selectBlockElementExternal struct {
	Placeholder    TextBlockObject                `json:"placeholder"`
	ActionID       string                         `json:"action_id"`
	InitialOption  *OptionBlockObject             `json:"initial_option,omitempty"`
	MinQueryLength int                            `json:"min_query_length,omitempty"`
	Confirm        *ConfirmationDialogBlockObject `json:"confirm,omitempty"`
}

func (b *selectBlockElementExternal) Encode(elementType messageElementType) (string, error) {
	block := struct {
		*selectBlockElementExternal
		Type messageElementType `json:"type"`
	}{
		selectBlockElementExternal: &selectBlockElementExternal{
			Placeholder:    b.Placeholder,
			ActionID:       b.ActionID,
			InitialOption:  b.InitialOption,
			MinQueryLength: b.MinQueryLength,
			Confirm:        b.Confirm,
		},
		Type: elementType,
	}
	buf, err := json.Marshal(&block)
	if err != nil {
		return "", errors.Wrap(err, `failed to encode image block element`)
	}
	return string(buf), nil
}

func (b *selectBlockElementExternal) Decode(buf string) error {
	if buf == "" {
		return nil
	}
	r := bytes.NewReader([]byte(buf))
	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()

	return dec.Decode(b)
}

func (b *SelectBlockElementExternal) Encode() (string, error) {
	return b.selectBlockElementExternal.Encode(b.selectBlockElementType())
}

func (b *SelectBlockElementExternal) Decode(buf string) error {
	return b.selectBlockElementExternal.Decode(buf)
}

func (b *MultiSelectBlockElementExternal) Encode() (string, error) {
	return b.selectBlockElementExternal.Encode(b.selectBlockElementType())
}

func (b *MultiSelectBlockElementExternal) Decode(buf string) error {
	return b.selectBlockElementExternal.Decode(buf)
}

type selectBlockElementUsers struct {
	Placeholder TextBlockObject                `json:"placeholder"`
	ActionID    string                         `json:"action_id"`
	InitialUser string                         `json:"initial_user,omitempty"`
	Confirm     *ConfirmationDialogBlockObject `json:"confirm,omitempty"`
}

func (b *selectBlockElementUsers) Encode(elementType messageElementType) (string, error) {
	block := struct {
		*selectBlockElementUsers
		Type messageElementType `json:"type"`
	}{
		selectBlockElementUsers: &selectBlockElementUsers{
			Placeholder: b.Placeholder,
			ActionID:    b.ActionID,
			InitialUser: b.InitialUser,
			Confirm:     b.Confirm,
		},
		Type: elementType,
	}
	buf, err := json.Marshal(&block)
	if err != nil {
		return "", errors.Wrap(err, `failed to encode image block element`)
	}
	return string(buf), nil
}

func (b *selectBlockElementUsers) Decode(buf string) error {
	if buf == "" {
		return nil
	}
	r := bytes.NewReader([]byte(buf))
	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()

	return dec.Decode(b)
}

func (b *SelectBlockElementUsers) Encode() (string, error) {
	return b.selectBlockElementUsers.Encode(b.selectBlockElementType())
}

func (b *SelectBlockElementUsers) Decode(buf string) error {
	return b.selectBlockElementUsers.Decode(buf)
}

func (b *MultiSelectBlockElementUsers) Encode() (string, error) {
	return b.selectBlockElementUsers.Encode(b.selectBlockElementType())
}

func (b *MultiSelectBlockElementUsers) Decode(buf string) error {
	return b.selectBlockElementUsers.Decode(buf)
}

type selectBlockElementConversations struct {
	Placeholder         TextBlockObject                `json:"placeholder"`
	ActionID            string                         `json:"action_id"`
	InitialConversation string                         `json:"initial_conversation,omitempty"`
	Confirm             *ConfirmationDialogBlockObject `json:"confirm,omitempty"`
}

func (b *selectBlockElementConversations) Encode(elementType messageElementType) (string, error) {
	block := struct {
		*selectBlockElementConversations
		Type messageElementType `json:"type"`
	}{
		selectBlockElementConversations: &selectBlockElementConversations{
			Placeholder:         b.Placeholder,
			ActionID:            b.ActionID,
			InitialConversation: b.InitialConversation,
			Confirm:             b.Confirm,
		},
		Type: elementType,
	}
	buf, err := json.Marshal(&block)
	if err != nil {
		return "", errors.Wrap(err, `failed to encode image block element`)
	}
	return string(buf), nil
}

func (b *selectBlockElementConversations) Decode(buf string) error {
	if buf == "" {
		return nil
	}
	r := bytes.NewReader([]byte(buf))
	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()

	return dec.Decode(b)
}

func (b *SelectBlockElementConversations) Encode() (string, error) {
	return b.selectBlockElementConversations.Encode(b.selectBlockElementType())
}

func (b *SelectBlockElementConversations) Decode(buf string) error {
	return b.selectBlockElementConversations.Decode(buf)
}

func (b *MultiSelectBlockElementConversations) Encode() (string, error) {
	return b.selectBlockElementConversations.Encode(b.selectBlockElementType())
}

func (b *MultiSelectBlockElementConversations) Decode(buf string) error {
	return b.selectBlockElementConversations.Decode(buf)
}

type selectBlockElementChannels struct {
	Placeholder    TextBlockObject                `json:"placeholder"`
	ActionID       string                         `json:"action_id"`
	InitialChannel string                         `json:"initial_channel,omitempty"`
	Confirm        *ConfirmationDialogBlockObject `json:"confirm,omitempty"`
}

func (b *selectBlockElementChannels) Encode(elementType messageElementType) (string, error) {
	block := struct {
		*selectBlockElementChannels
		Type messageElementType `json:"type"`
	}{
		selectBlockElementChannels: &selectBlockElementChannels{
			Placeholder:    b.Placeholder,
			ActionID:       b.ActionID,
			InitialChannel: b.InitialChannel,
			Confirm:        b.Confirm,
		},
		Type: elementType,
	}
	buf, err := json.Marshal(&block)
	if err != nil {
		return "", errors.Wrap(err, `failed to encode image block element`)
	}
	return string(buf), nil
}

func (b *selectBlockElementChannels) Decode(buf string) error {
	if buf == "" {
		return nil
	}
	r := bytes.NewReader([]byte(buf))
	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()

	return dec.Decode(b)
}

func (b *SelectBlockElementChannels) Encode() (string, error) {
	return b.selectBlockElementChannels.Encode(b.selectBlockElementType())
}

func (b *SelectBlockElementChannels) Decode(buf string) error {
	return b.selectBlockElementChannels.Decode(buf)
}

func (b *MultiSelectBlockElementChannels) Encode() (string, error) {
	return b.selectBlockElementChannels.Encode(b.selectBlockElementType())
}

func (b *MultiSelectBlockElementChannels) Decode(buf string) error {
	return b.selectBlockElementChannels.Decode(buf)
}

func (b *OverflowBlockElement) Encode() (string, error) {
	block := struct {
		*OverflowBlockElement
		Type messageElementType `json:"type"`
	}{
		OverflowBlockElement: &OverflowBlockElement{
			ActionID: b.ActionID,
			Options:  b.Options,
			Confirm:  b.Confirm,
		},
		Type: b.messageElementType(),
	}
	buf, err := json.Marshal(&block)
	if err != nil {
		return "", errors.Wrap(err, `failed to encode overflow menu element`)
	}
	return string(buf), nil
}

func (b *OverflowBlockElement) Decode(buf string) error {
	if buf == "" {
		return nil
	}
	r := bytes.NewReader([]byte(buf))
	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()

	return dec.Decode(b)
}

func (b *DatePickerElement) Encode() (string, error) {
	block := struct {
		*DatePickerElement
		Type messageElementType `json:"type"`
	}{
		DatePickerElement: &DatePickerElement{
			ActionID:    b.ActionID,
			Placeholder: b.Placeholder,
			InitialDate: b.InitialDate,
			Confirm:     b.Confirm,
		},
		Type: b.messageElementType(),
	}
	buf, err := json.Marshal(&block)
	if err != nil {
		return "", errors.Wrap(err, `failed to encode DatePicker element`)
	}
	return string(buf), nil
}

func (b *DatePickerElement) Decode(buf string) error {
	if buf == "" {
		return nil
	}
	r := bytes.NewReader([]byte(buf))
	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()

	return dec.Decode(b)
}
