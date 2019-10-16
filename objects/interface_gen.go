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

type Block interface {
	Type() BlockType
}
type BlockList []Block

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

type Action struct {
	confirm         *Confirmation   `json:"confirm,omitempty"`
	dataSource      string          `json:"data_source,omitempty"`
	minQueryLength  int             `json:"min_query_length,omitempty"`
	name            string          `json:"name,omitempty"`
	optionGroups    OptionGroupList `json:"option_groups,omitempty"`
	options         OptionList      `json:"options,omitempty"`
	selectedOptions OptionList      `json:"selected_options,omitempty"`
	style           string          `json:"style,omitempty"`
	text            string          `json:"text,omitempty"`
	typ             string          `json:"typ,omitempty"`
	value           string          `json:"value,omitempty"`
}

type ActionBuilder struct {
	confirm         *Confirmation
	dataSource      string
	minQueryLength  int
	name            string
	optionGroups    OptionGroupList
	options         OptionList
	selectedOptions OptionList
	style           string
	text            string
	typ             string
	value           string
}

type ActionList []*Action

type ActionsBlock struct {
	elements []BlockElement `json:"elements"`
	blockId  string         `json:"block_id,omitempty"`
}

type ActionsBlockBuilder struct {
	elements []BlockElement
	blockId  string
}

type Attachment struct {
	actions        ActionList          `json:"actions,omitempty"`
	attachmentType string              `json:"attachment_type,omitempty"`
	authorName     string              `json:"author_name,omitempty"`
	authorLink     string              `json:"author_link,omitempty"`
	authorIcon     string              `json:"author_icon,omitempty"`
	callbackId     string              `json:"callback_id,omitempty"`
	color          string              `json:"color,omitempty"`
	fallback       string              `json:"fallback,omitempty"`
	fields         AttachmentFieldList `json:"fields,omitempty"`
	footer         string              `json:"footer,omitempty"`
	footerIcon     string              `json:"footer_icon,omitempty"`
	imageUrl       string              `json:"image_url,omitempty"`
	thumbUrl       string              `json:"thumb_url,omitempty"`
	pretext        string              `json:"pretext,omitempty"`
	text           string              `json:"text,omitempty"`
	timestamp      Timestamp           `json:"timestamp,omitempty"`
	title          string              `json:"title,omitempty"`
	titleLink      string              `json:"title_link,omitempty"`
}

type AttachmentBuilder struct {
	actions        ActionList
	attachmentType string
	authorName     string
	authorLink     string
	authorIcon     string
	callbackId     string
	color          string
	fallback       string
	fields         AttachmentFieldList
	footer         string
	footerIcon     string
	imageUrl       string
	thumbUrl       string
	pretext        string
	text           string
	timestamp      Timestamp
	title          string
	titleLink      string
}

type AttachmentList []*Attachment

type AttachmentField struct {
	title string `json:"title,omitempty"`
	value string `json:"value,omitempty"`
	short string `json:"short,omitempty"`
}

type AttachmentFieldBuilder struct {
	title string
	value string
	short string
}

type AttachmentFieldList []*AttachmentField

type AuthTestResponse struct {
	url    string `json:"url,omitempty"`
	team   string `json:"team,omitempty"`
	user   string `json:"user,omitempty"`
	teamId string `json:"team_id,omitempty"`
	userId string `json:"user_id,omitempty"`
}

type AuthTestResponseBuilder struct {
	url    string
	team   string
	user   string
	teamId string
	userId string
}

type Bot struct {
	id      string `json:"id,omitempty"`
	appId   string `json:"app_id,omitempty"`
	deleted bool   `json:"deleted,omitempty"`
	name    string `json:"name,omitempty"`
	icons   Icons  `json:"icons,omitempty"`
}

type BotBuilder struct {
	id      string
	appId   string
	deleted bool
	name    string
	icons   Icons
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

type Channel struct {
	id                 string    `json:"id,omitempty"`
	created            EpochTime `json:"created,omitempty"`
	isOpen             bool      `json:"is_open,omitempty"`
	lastRead           string    `json:"last_read,omitempty"`
	latest             *Message  `json:"latest,omitempty"`
	unreadCount        int       `json:"unread_count,omitempty"`
	unreadCountDisplay int       `json:"unread_count_display,omitempty"`
	creator            string    `json:"creator,omitempty"`
	isArchived         bool      `json:"is_archived,omitempty"`
	isGroup            bool      `json:"is_group,omitempty"`
	isMpim             bool      `json:"is_mpim,omitempty"`
	members            []string  `json:"members,omitempty"`
	name               string    `json:"name,omitempty"`
	nameNormalized     string    `json:"name_normalized,omitempty"`
	numMembers         int       `json:"num_members,omitempty"`
	previousNames      []string  `json:"previous_names,omitempty"`
	purpose            Purpose   `json:"purpose,omitempty"`
	topic              Topic     `json:"topic,omitempty"`
	isChannel          bool      `json:"is_channel,omitempty"`
	isGeneral          bool      `json:"is_general,omitempty"`
	isMember           bool      `json:"is_member,omitempty"`
	isOrgShared        bool      `json:"is_org_shared,omitempty"`
	isShared           bool      `json:"is_shared,omitempty"`
}

type ChannelBuilder struct {
	id                 string
	created            EpochTime
	isOpen             bool
	lastRead           string
	latest             *Message
	unreadCount        int
	unreadCountDisplay int
	creator            string
	isArchived         bool
	isGroup            bool
	isMpim             bool
	members            []string
	name               string
	nameNormalized     string
	numMembers         int
	previousNames      []string
	purpose            Purpose
	topic              Topic
	isChannel          bool
	isGeneral          bool
	isMember           bool
	isOrgShared        bool
	isShared           bool
}

type ChannelList []*Channel

type Comment struct {
	id         string    `json:"id,omitempty"`
	created    EpochTime `json:"created,omitempty"`
	timestamkp EpochTime `json:"timestamkp,omitempty"`
	user       string    `json:"user,omitempty"`
	comment    string    `json:"comment,omitempty"`
}

type CommentBuilder struct {
	id         string
	created    EpochTime
	timestamkp EpochTime
	user       string
	comment    string
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

type Confirmation struct {
	title       string `json:"title,omitempty"`
	text        string `json:"text,omitempty"`
	okText      string `json:"ok_text,omitempty"`
	dismissText string `json:"dismiss_text,omitempty"`
}

type ConfirmationBuilder struct {
	title       string
	text        string
	okText      string
	dismissText string
}

type ConfirmationList []*Confirmation

type ContextBlock struct {
	elements []interface{} `json:"elements"`
	blockId  string        `json:"block_id,omitempty"`
}

type ContextBlockBuilder struct {
	elements []interface{}
	blockId  string
}

type Conversation struct {
	id                 string    `json:"id,omitempty"`
	created            EpochTime `json:"created,omitempty"`
	isOpen             bool      `json:"is_open,omitempty"`
	lastRead           string    `json:"last_read,omitempty"`
	latest             *Message  `json:"latest,omitempty"`
	unreadCount        int       `json:"unread_count,omitempty"`
	unreadCountDisplay int       `json:"unread_count_display,omitempty"`
}

type ConversationBuilder struct {
	id                 string
	created            EpochTime
	isOpen             bool
	lastRead           string
	latest             *Message
	unreadCount        int
	unreadCountDisplay int
}

type ConversationList []*Conversation

type DividerBlock struct {
	blockId string `json:"block_id,omitempty"`
}

type DividerBlockBuilder struct {
	blockId string
}

type Edited struct {
	ts   string `json:"ts,omitempty"`
	user string `json:"user,omitempty"`
}

type EditedBuilder struct {
	ts   string
	user string
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

type Group struct {
	id                 string    `json:"id,omitempty"`
	created            EpochTime `json:"created,omitempty"`
	isOpen             bool      `json:"is_open,omitempty"`
	lastRead           string    `json:"last_read,omitempty"`
	latest             *Message  `json:"latest,omitempty"`
	unreadCount        int       `json:"unread_count,omitempty"`
	unreadCountDisplay int       `json:"unread_count_display,omitempty"`
	creator            string    `json:"creator,omitempty"`
	isArchived         bool      `json:"is_archived,omitempty"`
	isGroup            bool      `json:"is_group,omitempty"`
	isMpim             bool      `json:"is_mpim,omitempty"`
	members            []string  `json:"members,omitempty"`
	name               string    `json:"name,omitempty"`
	nameNormalized     string    `json:"name_normalized,omitempty"`
	numMembers         int       `json:"num_members,omitempty"`
	previousNames      []string  `json:"previous_names,omitempty"`
	purpose            Purpose   `json:"purpose,omitempty"`
	topic              Topic     `json:"topic,omitempty"`
}

type GroupBuilder struct {
	id                 string
	created            EpochTime
	isOpen             bool
	lastRead           string
	latest             *Message
	unreadCount        int
	unreadCountDisplay int
	creator            string
	isArchived         bool
	isGroup            bool
	isMpim             bool
	members            []string
	name               string
	nameNormalized     string
	numMembers         int
	previousNames      []string
	purpose            Purpose
	topic              Topic
}

type Icons struct {
	image36 string `json:"image_36,omitempty"`
	image48 string `json:"image_48,omitempty"`
	image72 string `json:"image_72,omitempty"`
}

type IconsBuilder struct {
	image36 string
	image48 string
	image72 string
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

type Message struct {
	attachments AttachmentList `json:"attachments,omitempty"`
	channel     string         `json:"channel,omitempty"`
	edited      *Edited        `json:"edited,omitempty"`
	isStarred   bool           `json:"is_starred,omitempty"`
	pinnedTo    []string       `json:"pinned_to,omitempty"`
	text        string         `json:"text,omitempty"`
	ts          string         `json:"ts,omitempty"`
	typ         string         `json:"typ,omitempty"`
	user        string         `json:"user,omitempty"`
	subtype     string         `json:"subtype,omitempty"`
	hidden      bool           `json:"hidden,omitempty"`
	deletedTs   string         `json:"deleted_ts,omitempty"`
	eventTs     string         `json:"event_ts,omitempty"`
	botId       string         `json:"bot_id,omitempty"`
	username    string         `json:"username,omitempty"`
	icons       *Icons         `json:"icons,omitempty"`
	inviter     string         `json:"inviter,omitempty"`
	topic       string         `json:"topic,omitempty"`
	purpose     string         `json:"purpose,omitempty"`
	name        string         `json:"name,omitempty"`
	oldName     string         `json:"old_name,omitempty"`
	members     []string       `json:"members,omitempty"`
	upload      bool           `json:"upload,omitempty"`
	comment     *Comment       `json:"comment,omitempty"`
	itemType    string         `json:"item_type,omitempty"`
	replyTo     int            `json:"reply_to,omitempty"`
	team        string         `json:"team,omitempty"`
	reactions   ReactionList   `json:"reactions,omitempty"`
	blocks      BlockList      `json:"blocks,omitempty"`
}

type MessageBuilder struct {
	attachments AttachmentList
	channel     string
	edited      *Edited
	isStarred   bool
	pinnedTo    []string
	text        string
	ts          string
	typ         string
	user        string
	subtype     string
	hidden      bool
	deletedTs   string
	eventTs     string
	botId       string
	username    string
	icons       *Icons
	inviter     string
	topic       string
	purpose     string
	name        string
	oldName     string
	members     []string
	upload      bool
	comment     *Comment
	itemType    string
	replyTo     int
	team        string
	reactions   ReactionList
	blocks      BlockList
}

type MessageList []*Message

type Option struct {
	text        string `json:"text,omitempty"`
	value       string `json:"value,omitempty"`
	description string `json:"description,omitempty"`
}

type OptionBuilder struct {
	text        string
	value       string
	description string
}

type OptionList []*Option

type OptionGroup struct {
	text    string     `json:"text,omitempty"`
	options OptionList `json:"options,omitempty"`
}

type OptionGroupBuilder struct {
	text    string
	options OptionList
}

type OptionGroupList []*OptionGroup

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

type Purpose struct {
	value   string    `json:"value,omitempty"`
	creator string    `json:"creator,omitempty"`
	lastSet EpochTime `json:"last_set,omitempty"`
}

type PurposeBuilder struct {
	value   string
	creator string
	lastSet EpochTime
}

type Reaction struct {
	count int      `json:"count,omitempty"`
	name  string   `json:"name,omitempty"`
	users []string `json:"users,omitempty"`
}

type ReactionBuilder struct {
	count int
	name  string
	users []string
}

type ReactionList []*Reaction

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

type Topic struct {
	value   string    `json:"value,omitempty"`
	creator string    `json:"creator,omitempty"`
	lastSet EpochTime `json:"last_set,omitempty"`
}

type TopicBuilder struct {
	value   string
	creator string
	lastSet EpochTime
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

type UserProfileList []*UserProfile
