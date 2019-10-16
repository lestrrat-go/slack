package objects

type TextType string

const (
	MarkdownTextType = "mrkdwn"
	PlainTextType    = "plain_text"
)

type BlockType string

const (
	ActionsBlockType BlockType = "actions"
	ContextBlockType BlockType = "context"
	DividerBlockType BlockType = "divider"
	FileBlockType    BlockType = "file"
	ImageBlockType   BlockType = "image"
	InputBlockType   BlockType = "input"
	SectionBlockType BlockType = "section"
)

type ElementType string

const (
	ImageElementType           ElementType = "image"
	ButtonElementType          ElementType = "button"
	SelectMenuElementType      ElementType = "select_menu"
	MultiSelectMenuElementType ElementType = "multi_select_menu"
	OverflowMenuElementType    ElementType = "overflow_menu"
	DatePickerElementType      ElementType = "date_picker"
	InputElementType           ElementType = "input"
)

type BlockElement interface {
	Type() ElementType
}

type ActionsBlock struct {
	elements []BlockElement `json:"elements"`
	blockId  string         `json:"block_id,omitempty"`
}

type ActionsBlockBuilder struct {
	elements []BlockElement
	blockId  string
}

type ButtonElement struct {
	text     *Text         `json:"text"`
	actionId string        `json:"action_id"`
	url      string        `json:"url,omitempty"`
	value    string        `json:"value,omitempty"`
	style    string        `json:"style,omitempty"`
	confirm  ConfirmObject `json:"confirm,omitempty"`
}

type ButtonElementBuilder struct {
	text     *Text
	actionId string
	url      string
	value    string
	style    string
	confirm  ConfirmObject
}

type ConfirmObject struct {
	title   string `json:"title"`
	text    string `json:"text"`
	confirm string `json:"confirm"`
	deny    string `json:"deny"`
}

type ConfirmObjectBuilder struct {
	title   string
	text    string
	confirm string
	deny    string
}

type ContextBlock struct {
	elements []interface{} `json:"elements"`
	blockId  string        `json:"block_id,omitempty"`
}

type ContextBlockBuilder struct {
	elements []interface{}
	blockId  string
}

type DividerBlock struct {
	blockId string `json:"block_id,omitempty"`
}

type DividerBlockBuilder struct {
	blockId string
}

type FileBlock struct {
	externalId string `json:"external_id"`
	source     string `json:"source,omitempty"`
	blockId    string `json:"block_id,omitempty"`
}

type FileBlockBuilder struct {
	externalId string
	source     string
	blockId    string
}

type ImageBlock struct {
	imageUrl string `json:"image_url"`
	altText  string `json:"alt_text"`
	title    string `json:"title,omitempty"`
	blockId  string `json:"block_id,omitempty"`
}

type ImageBlockBuilder struct {
	imageUrl string
	altText  string
	title    string
	blockId  string
}

type ImageElement struct {
	imageUrl string `json:"image_url"`
	altText  string `json:"alt_text"`
}

type ImageElementBuilder struct {
	imageUrl string
	altText  string
}

type InputBlock struct {
	label    string      `json:"label"`
	element  interface{} `json:"element,omitempty"`
	hint     *Text       `json:"hint,omitempty"`
	optional bool        `json:"optional,omitempty"`
}

type InputBlockBuilder struct {
	label    string
	element  interface{}
	hint     *Text
	optional bool
}

type OptionGroupObject struct {
	label   string         `json:"label"`
	options []OptionObject `json:"options"`
}

type OptionGroupObjectBuilder struct {
	label   string
	options []OptionObject
}

type OptionObject struct {
	text  string `json:"text"`
	value string `json:"value"`
	url   string `json:"url,omitempty"`
}

type OptionObjectBuilder struct {
	text  string
	value string
	url   string
}

type SectionBlock struct {
	text      *Text        `json:"text"`
	fields    []*Text      `json:"fields,omitempty"`
	blockId   string       `json:"block_id,omitempty"`
	accessory BlockElement `json:"accessory,omitempty"`
}

type SectionBlockBuilder struct {
	text      *Text
	fields    []*Text
	blockId   string
	accessory BlockElement
}

type SelectElement struct {
	placeholder   *Text               `json:"placeholder"`
	actionId      string              `json:"action_id"`
	options       []OptionObject      `json:"options"`
	optionGroups  []OptionGroupObject `json:"option_groups,omitempty"`
	initialOption interface{}         `json:"initial_option,omitempty"`
	confirm       ConfirmObject       `json:"confirm,omitempty"`
}

type SelectElementBuilder struct {
	placeholder   *Text
	actionId      string
	options       []OptionObject
	optionGroups  []OptionGroupObject
	initialOption interface{}
	confirm       ConfirmObject
}

type Text struct {
	typ      string `json:"typ"`
	text     string `json:"text"`
	emoji    bool   `json:"emoji,omitempty"`
	verbatim bool   `json:"verbatim,omitempty"`
}

type TextBuilder struct {
	typ      string
	text     string
	emoji    bool
	verbatim bool
}

type UserProfile struct {
	alwaysActive       bool   `json:"always_active,omitempty"`
	avatarHash         string `json:"avatar_hash,omitempty"`
	email              string `json:"email,omitempty"`
	firstName          string `json:"first_name,omitempty"`
	image24            string `json:"image_24,omitempty"`
	image32            string `json:"image_32,omitempty"`
	image48            string `json:"image_48,omitempty"`
	image72            string `json:"image_72,omitempty"`
	image192           string `json:"image_192,omitempty"`
	image512           string `json:"image_512,omitempty"`
	lastName           string `json:"last_name,omitempty"`
	realName           string `json:"real_name,omitempty"`
	realNameNormalized string `json:"real_name_normalized,omitempty"`
	statusText         string `json:"status_text,omitempty"`
	statusEmoji        string `json:"status_emoji,omitempty"`
}

type UserProfileBuilder struct {
	alwaysActive       bool
	avatarHash         string
	email              string
	firstName          string
	image24            string
	image32            string
	image48            string
	image72            string
	image192           string
	image512           string
	lastName           string
	realName           string
	realNameNormalized string
	statusText         string
	statusEmoji        string
}
