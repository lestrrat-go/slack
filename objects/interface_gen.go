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
	icons   *Icons `json:"icons,omitempty"`
}

type BotBuilder struct {
	id      string
	appId   string
	deleted bool
	name    string
	icons   *Icons
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
	purpose            *Purpose  `json:"purpose,omitempty"`
	topic              *Topic    `json:"topic,omitempty"`
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
	purpose            *Purpose
	topic              *Topic
	isChannel          bool
	isGeneral          bool
	isMember           bool
	isOrgShared        bool
	isShared           bool
}

type ChannelList []*Channel

type ChannelsHistoryResponse struct {
	hasMore  bool        `json:"has_more,omitempty"`
	latest   string      `json:"latest,omitempty"`
	messages MessageList `json:"messages,omitempty"`
}

type ChannelsHistoryResponseBuilder struct {
	hasMore  bool
	latest   string
	messages MessageList
}

type ChatResponse struct {
	channel string      `json:"channel,omitempty"`
	ts      string      `json:"ts,omitempty"`
	message interface{} `json:"message,omitempty"`
}

type ChatResponseBuilder struct {
	channel string
	ts      string
	message interface{}
}

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

type Dialog struct {
	callbackId  string            `json:"callback_id,omitempty"`
	title       string            `json:"title,omitempty"`
	submitLabel string            `json:"submit_label,omitempty"`
	elements    DialogElementList `json:"elements,omitempty"`
}

type DialogBuilder struct {
	callbackId  string
	title       string
	submitLabel string
	elements    DialogElementList
}

type DialogElement struct {
	label       string `json:"label,omitempty"`
	typ         string `json:"typ,omitempty"`
	name        string `json:"name,omitempty"`
	hint        string `json:"hint,omitempty"`
	maxLength   string `json:"max_length,omitempty"`
	minLength   string `json:"min_length,omitempty"`
	optional    bool   `json:"optional,omitempty"`
	placeholder string `json:"placeholder,omitempty"`
	subtype     string `json:"subtype,omitempty"`
	value       string `json:"value,omitempty"`
}

type DialogElementBuilder struct {
	label       string
	typ         string
	name        string
	hint        string
	maxLength   string
	minLength   string
	optional    bool
	placeholder string
	subtype     string
	value       string
}

type DialogElementList []*DialogElement

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

type EphemeralResponse struct {
	messageTs string `json:"message_ts,omitempty"`
}

type EphemeralResponseBuilder struct {
	messageTs string
}

type File struct {
	id                 string       `json:"id,omitempty"`
	name               string       `json:"name,omitempty"`
	user               string       `json:"user,omitempty"`
	created            int          `json:"created,omitempty"`
	timestamp          int          `json:"timestamp,omitempty"`
	updated            int          `json:"updated,omitempty"`
	mimetype           string       `json:"mimetype,omitempty"`
	filetype           string       `json:"filetype,omitempty"`
	prettyType         string       `json:"pretty_type,omitempty"`
	mode               string       `json:"mode,omitempty"`
	editable           bool         `json:"editable,omitempty"`
	isExternal         bool         `json:"is_external,omitempty"`
	externalType       string       `json:"external_type,omitempty"`
	size               int          `json:"size,omitempty"`
	url                string       `json:"url,omitempty"`
	urlDownload        string       `json:"url_download,omitempty"`
	urlPrivate         string       `json:"url_private,omitempty"`
	urlPrivateDownload string       `json:"url_private_download,omitempty"`
	imageExifRotation  string       `json:"image_exif_rotation,omitempty"`
	originalW          int          `json:"original_w,omitempty"`
	originalH          int          `json:"original_h,omitempty"`
	thumb64            string       `json:"thumb_64,omitempty"`
	thumb80            string       `json:"thumb_80,omitempty"`
	thumb160           string       `json:"thumb_160,omitempty"`
	thumb360           string       `json:"thumb_360,omitempty"`
	thumb360Gif        string       `json:"thumb_360_gif,omitempty"`
	thumb360W          int          `json:"thumb_360_w,omitempty"`
	thumb360H          int          `json:"thumb_360_h,omitempty"`
	thumb480           string       `json:"thumb_480,omitempty"`
	thumb480W          int          `json:"thumb_480_w,omitempty"`
	thumb480H          int          `json:"thumb_480_h,omitempty"`
	thumb720           string       `json:"thumb_720,omitempty"`
	thumb720W          int          `json:"thumb_720_w,omitempty"`
	thumb720H          int          `json:"thumb_720_h,omitempty"`
	thumb960           string       `json:"thumb_960,omitempty"`
	thumb960W          int          `json:"thumb_960_w,omitempty"`
	thumb960H          int          `json:"thumb_960_h,omitempty"`
	thumb1024          string       `json:"thumb_1024,omitempty"`
	thumb1024W         int          `json:"thumb_1024_w,omitempty"`
	thumb1024H         int          `json:"thumb_1024_h,omitempty"`
	permalink          string       `json:"permalink,omitempty"`
	permalinkPublic    string       `json:"permalink_public,omitempty"`
	editLink           string       `json:"edit_link,omitempty"`
	preview            string       `json:"preview,omitempty"`
	previewHighlight   string       `json:"preview_highlight,omitempty"`
	lines              int          `json:"lines,omitempty"`
	linesMore          int          `json:"lines_more,omitempty"`
	isPublic           bool         `json:"is_public,omitempty"`
	publicUrlShared    bool         `json:"public_url_shared,omitempty"`
	channels           []string     `json:"channels,omitempty"`
	groups             []string     `json:"groups,omitempty"`
	ims                []string     `json:"ims,omitempty"`
	initialComment     Comment      `json:"initial_comment,omitempty"`
	commentsCount      int          `json:"comments_count,omitempty"`
	numStars           int          `json:"num_stars,omitempty"`
	isStarred          bool         `json:"is_starred,omitempty"`
	title              string       `json:"title,omitempty"`
	reactions          ReactionList `json:"reactions,omitempty"`
}

type FileBuilder struct {
	id                 string
	name               string
	user               string
	created            int
	timestamp          int
	updated            int
	mimetype           string
	filetype           string
	prettyType         string
	mode               string
	editable           bool
	isExternal         bool
	externalType       string
	size               int
	url                string
	urlDownload        string
	urlPrivate         string
	urlPrivateDownload string
	imageExifRotation  string
	originalW          int
	originalH          int
	thumb64            string
	thumb80            string
	thumb160           string
	thumb360           string
	thumb360Gif        string
	thumb360W          int
	thumb360H          int
	thumb480           string
	thumb480W          int
	thumb480H          int
	thumb720           string
	thumb720W          int
	thumb720H          int
	thumb960           string
	thumb960W          int
	thumb960H          int
	thumb1024          string
	thumb1024W         int
	thumb1024H         int
	permalink          string
	permalinkPublic    string
	editLink           string
	preview            string
	previewHighlight   string
	lines              int
	linesMore          int
	isPublic           bool
	publicUrlShared    bool
	channels           []string
	groups             []string
	ims                []string
	initialComment     Comment
	commentsCount      int
	numStars           int
	isStarred          bool
	title              string
	reactions          ReactionList
}

type FileList []*File

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

type GenericResponse struct {
	ok      bool           `json:"ok,omitempty"`
	replyTo int            `json:"reply_to,omitempty"`
	error   *ErrorResponse `json:"error,omitempty"`
	ts      string         `json:"ts,omitempty"`
}

type GenericResponseBuilder struct {
	ok      bool
	replyTo int
	error   *ErrorResponse
	ts      string
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

type GroupList []*Group

type IM struct {
	id                 string    `json:"id,omitempty"`
	created            EpochTime `json:"created,omitempty"`
	isOpen             bool      `json:"is_open,omitempty"`
	lastRead           string    `json:"last_read,omitempty"`
	latest             *Message  `json:"latest,omitempty"`
	unreadCount        int       `json:"unread_count,omitempty"`
	unreadCountDisplay int       `json:"unread_count_display,omitempty"`
	isIm               bool      `json:"is_im,omitempty"`
	user               string    `json:"user,omitempty"`
	isUserDeleted      bool      `json:"is_user_deleted,omitempty"`
}

type IMBuilder struct {
	id                 string
	created            EpochTime
	isOpen             bool
	lastRead           string
	latest             *Message
	unreadCount        int
	unreadCountDisplay int
	isIm               bool
	user               string
	isUserDeleted      bool
}

type IMList []*IM

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

type OAuthAccessResponse struct {
	accessToken string `json:"access_token,omitempty"`
	scope       string `json:"scope,omitempty"`
}

type OAuthAccessResponseBuilder struct {
	accessToken string
	scope       string
}

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

type Paging struct {
	count int `json:"count,omitempty"`
	total int `json:"total,omitempty"`
	page  int `json:"page,omitempty"`
	pages int `json:"pages,omitempty"`
}

type PagingBuilder struct {
	count int
	total int
	page  int
	pages int
}

type PermalinkResponse struct {
	channel   string `json:"channel,omitempty"`
	permalink string `json:"permalink,omitempty"`
}

type PermalinkResponseBuilder struct {
	channel   string
	permalink string
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

type RTMResponse struct {
	url      string       `json:"url,omitempty"`
	self     *UserDetails `json:"self,omitempty"`
	team     *Team        `json:"team,omitempty"`
	users    []*User      `json:"users,omitempty"`
	channels []*Channel   `json:"channels,omitempty"`
	groups   []*Group     `json:"groups,omitempty"`
	bots     []*Bot       `json:"bots,omitempty"`
	ims      []*IM        `json:"ims,omitempty"`
}

type RTMResponseBuilder struct {
	url      string
	self     *UserDetails
	team     *Team
	users    []*User
	channels []*Channel
	groups   []*Group
	bots     []*Bot
	ims      []*IM
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

type ReactionsGetResponse struct {
	channel string   `json:"channel,omitempty"`
	comment string   `json:"comment,omitempty"`
	file    *File    `json:"file,omitempty"`
	message *Message `json:"message,omitempty"`
	typ     string   `json:"typ,omitempty"`
}

type ReactionsGetResponseBuilder struct {
	channel string
	comment string
	file    *File
	message *Message
	typ     string
}

type ReactionsGetResponseList []*ReactionsGetResponse

type ReactionsListResponse struct {
	items  ReactionsGetResponseList `json:"items,omitempty"`
	paging *Paging                  `json:"paging,omitempty"`
}

type ReactionsListResponseBuilder struct {
	items  ReactionsGetResponseList
	paging *Paging
}

type Reminder struct {
	id                string    `json:"id,omitempty"`
	creator           string    `json:"creator,omitempty"`
	user              string    `json:"user,omitempty"`
	text              string    `json:"text,omitempty"`
	recurring         bool      `json:"recurring,omitempty"`
	time              EpochTime `json:"time,omitempty"`
	completeTimestamp EpochTime `json:"complete_timestamp,omitempty"`
}

type ReminderBuilder struct {
	id                string
	creator           string
	user              string
	text              string
	recurring         bool
	time              EpochTime
	completeTimestamp EpochTime
}

type ReminderList []*Reminder

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

type Team struct {
	id                string                 `json:"id,omitempty"`
	name              string                 `json:"name,omitempty"`
	domain            string                 `json:"domain,omitempty"`
	emailDomain       string                 `json:"email_domain,omitempty"`
	enterpriseId      string                 `json:"enterprise_id,omitempty"`
	enterpriseName    string                 `json:"enterprise_name,omitempty"`
	icon              map[string]interface{} `json:"icon,omitempty"`
	msgEditWindowMins int                    `json:"msg_edit_window_mins,omitempty"`
	overStorageLimit  bool                   `json:"over_storage_limit,omitempty"`
	prefs             interface{}            `json:"prefs,omitempty"`
	plan              string                 `json:"plan,omitempty"`
}

type TeamBuilder struct {
	id                string
	name              string
	domain            string
	emailDomain       string
	enterpriseId      string
	enterpriseName    string
	icon              map[string]interface{}
	msgEditWindowMins int
	overStorageLimit  bool
	prefs             interface{}
	plan              string
}

type TeamList []*Team

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

type ThreadInfo struct {
	complete bool `json:"complete,omitempty"`
	count    int  `json:"count,omitempty"`
}

type ThreadInfoBuilder struct {
	complete bool
	count    int
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

type User struct {
	color             string       `json:"color,omitempty"`
	deleted           bool         `json:"deleted,omitempty"`
	id                string       `json:"id,omitempty"`
	isAdmin           bool         `json:"is_admin,omitempty"`
	isBot             bool         `json:"is_bot,omitempty"`
	isOwner           bool         `json:"is_owner,omitempty"`
	isRestricted      bool         `json:"is_restricted,omitempty"`
	isUltraRestricted bool         `json:"is_ultra_restricted,omitempty"`
	name              string       `json:"name,omitempty"`
	profile           *UserProfile `json:"profile,omitempty"`
	realName          string       `json:"real_name,omitempty"`
	status            string       `json:"status,omitempty"`
	teamId            string       `json:"team_id,omitempty"`
	tz                string       `json:"tz,omitempty"`
	tzLabel           string       `json:"tz_label,omitempty"`
	tzOffset          int          `json:"tz_offset,omitempty"`
	updated           int          `json:"updated,omitempty"`
}

type UserBuilder struct {
	color             string
	deleted           bool
	id                string
	isAdmin           bool
	isBot             bool
	isOwner           bool
	isRestricted      bool
	isUltraRestricted bool
	name              string
	profile           *UserProfile
	realName          string
	status            string
	teamId            string
	tz                string
	tzLabel           string
	tzOffset          int
	updated           int
}

type UserList []*User

type UserDetails struct {
	id             string     `json:"id,omitempty"`
	name           string     `json:"name,omitempty"`
	created        EpochTime  `json:"created,omitempty"`
	manualPresence string     `json:"manual_presence,omitempty"`
	prefs          *UserPrefs `json:"prefs,omitempty"`
}

type UserDetailsBuilder struct {
	id             string
	name           string
	created        EpochTime
	manualPresence string
	prefs          *UserPrefs
}

type UserDetailsList []*UserDetails

type UserPrefs struct {
}

type UserPrefsBuilder struct {
}

type UserPresence struct {
	autoAway        bool      `json:"auto_away,omitempty"`
	connectionCount int       `json:"connection_count,omitempty"`
	lastActivity    int       `json:"last_activity,omitempty"`
	manualAway      bool      `json:"manual_away,omitempty"`
	online          bool      `json:"online,omitempty"`
	presence        *Presence `json:"presence,omitempty"`
}

type UserPresenceBuilder struct {
	autoAway        bool
	connectionCount int
	lastActivity    int
	manualAway      bool
	online          bool
	presence        *Presence
}

type UserPresenceList []*UserPresence

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

type Usergroup struct {
	autoProvision       bool            `json:"auto_provision,omitempty"`
	autoType            string          `json:"auto_type,omitempty"`
	createdBy           string          `json:"created_by,omitempty"`
	dateCreate          EpochTime       `json:"date_create,omitempty"`
	dateDelete          EpochTime       `json:"date_delete,omitempty"`
	dateUpdate          EpochTime       `json:"date_update,omitempty"`
	deletedBy           string          `json:"deleted_by,omitempty"`
	description         string          `json:"description,omitempty"`
	enterpriseSubteamId string          `json:"enterprise_subteam_id,omitempty"`
	handle              string          `json:"handle,omitempty"`
	id                  string          `json:"id,omitempty"`
	isExternal          bool            `json:"is_external,omitempty"`
	isSubteam           bool            `json:"is_subteam,omitempty"`
	isUsergroup         bool            `json:"is_usergroup,omitempty"`
	name                string          `json:"name,omitempty"`
	prefs               *UsergroupPrefs `json:"prefs,omitempty"`
	teamId              string          `json:"team_id,omitempty"`
	updatedBy           string          `json:"updated_by,omitempty"`
	users               []string        `json:"users,omitempty"`
	userCount           int             `json:"user_count,omitempty"`
}

type UsergroupBuilder struct {
	autoProvision       bool
	autoType            string
	createdBy           string
	dateCreate          EpochTime
	dateDelete          EpochTime
	dateUpdate          EpochTime
	deletedBy           string
	description         string
	enterpriseSubteamId string
	handle              string
	id                  string
	isExternal          bool
	isSubteam           bool
	isUsergroup         bool
	name                string
	prefs               *UsergroupPrefs
	teamId              string
	updatedBy           string
	users               []string
	userCount           int
}

type UsergroupList []*Usergroup

type UsergroupPrefs struct {
	channels []string `json:"channels,omitempty"`
	groups   []string `json:"groups,omitempty"`
}

type UsergroupPrefsBuilder struct {
	channels []string
	groups   []string
}
