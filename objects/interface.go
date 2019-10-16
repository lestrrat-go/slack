package objects

type EpochTime int64
type Timestamp float64

const (
	ButtonActionType = "button"
)

// Action is used in conjunction with message buttons
type Action struct {
	Confirm         *Confirmation   `json:"confirm,omitempty"`
	DataSource      string          `json:"data_source,omitempty"`
	MinQueryLength  int             `json:"min_query_length,omitempty"`
	Name            string          `json:"name"`
	OptionGroups    OptionGroupList `json:"option_groups,omitempty"`
	Options         OptionList      `json:"options,omitempty"`
	SelectedOptions OptionList      `json:"selected_options,omitempty"`
	Style           string          `json:"style,omitempty"`
	Text            string          `json:"text"`
	Type            string          `json:"type"`
	Value           string          `json:"value"`
}

type ActionList []*Action

type Attachment struct {
	Actions        ActionList          `json:"actions,omitempty"` // for buttons
	AttachmentType string              `json:"attachment_type"`
	AuthorName     string              `json:"author_name"`
	AuthorLink     string              `json:"author_link"`
	AuthorIcon     string              `json:"author_icon"`
	CallbackID     string              `json:"callback_id,omitempty"` // for buttons
	Color          string              `json:"color,omitempty"`
	Fallback       string              `json:"fallback"`
	Fields         AttachmentFieldList `json:"fields"`
	Footer         string              `json:"footer"`
	FooterIcon     string              `json:"footer_icon"`
	ImageURL       string              `json:"image_url"`
	ThumbURL       string              `json:"thumb_url"`
	Pretext        string              `json:"pretext,omitempty"`
	Text           string              `json:"text"`
	Timestamp      Timestamp           `json:"ts"`
	Title          string              `json:"title"`
	TitleLink      string              `json:"title_link"`
}

type AttachmentList []*Attachment

// AuthTestResponse is the data structure response from auth.test
type AuthTestResponse struct {
	URL    string `json:"url"`
	Team   string `json:"team"`
	User   string `json:"user"`
	TeamID string `json:"team_id"`
	UserID string `json:"user_id"`
}

type Block interface {
	blockType() blockType
}

type SectionBlock struct {
	Text      TextBlockObject   `json:"text,omitempty"`
	BlockID   string            `json:"block_id,omitempty"`
	Fields    []TextBlockObject `json:"fields,omitempty"`
	Accessory *Accessory        `json:"accessory,omitempty"`
}

type ActionsBlock struct {
	BlockID  string         `json:"block_id,omitempty"`
	Elements *BlockElements `json:"elements"`
}

type DividerBlock struct {
	BlockID string `json:"block_id,omitempty"`
}

type FileBlock struct {
	BlockID    string `json:"block_id,omitempty"`
	ExternalID string `json:"external_id,omitempty"`
	Source     string `json:"source,omitempty"`
}

type ImageBlock struct {
	ImageURL string          `json:"image_url"`
	AltText  string          `json:"alt_text"`
	BlockID  string          `json:"block_id,omitempty"`
	Title    TextBlockObject `json:"title"`
}

type ContextBlock struct {
	BlockID  string           `json:"block_id,omitempty"`
	Elements *ContextElements `json:"elements,omitempty"`
}

type BlockList []Block

type BlockObjects struct {
	TextBlockObject          []TextBlockObject
	ConfirmationDialogObject []*ConfirmationDialogBlockObject
	OptionObjects            []*OptionBlockObject
	OptionGroupObjects       []*OptionGroupBlockObject
}

type TextBlockObject interface {
	textBlockObjectType() textBlockObjectType
	contextElementType() contextElementType
}

type PlainTextBlock struct {
	Text  string `json:"text"`
	Emoji bool   `json:"emoji,omitempty"`
}

func (PlainTextBlock) textBlockObjectType() textBlockObjectType { return textBlockPlainText }
func (PlainTextBlock) contextElementType() contextElementType   { return contextElementText }

type MarkdownTextBlock struct {
	Text     string `json:"text"`
	Verbatim bool   `json:"verbatim,omitempty"`
}

func (MarkdownTextBlock) textBlockObjectType() textBlockObjectType { return textBlockMarkdown }
func (MarkdownTextBlock) contextElementType() contextElementType   { return contextElementText }

type BlockObject interface {
	blockObjectType() blockObjectType
}

type ConfirmationDialogBlockObject struct {
	Title   *PlainTextBlock `json:"title"`
	Text    TextBlockObject `json:"text"`
	Confirm *PlainTextBlock `json:"confirm"`
	Deny    *PlainTextBlock `json:"deny"`
}

func (ConfirmationDialogBlockObject) blockObjectType() blockObjectType { return objectsConfirmation }

type OptionBlockObject struct {
	Text  *PlainTextBlock `json:"text"`
	Value string          `json:"value"`
	URL   string          `json:"url,omitempty"`
}

func (OptionBlockObject) blockObjectType() blockObjectType { return objectsOptionObjects }

type OptionGroupBlockObject struct {
	Label   *PlainTextBlock      `json:"label"`
	Options []*OptionBlockObject `json:"options"`
}

func (OptionGroupBlockObject) blockObjectType() blockObjectType { return objectsOptionGroupObjects }

type BlockElement interface {
	messageElementType() messageElementType
}

type BlockElements struct {
	ElementSet []BlockElement `json:"elements,omitempty"`
}

type ContextElement interface {
	contextElementType() contextElementType
}

type ContextElements struct {
	Elements []ContextElement
}

type ImageBlockElement struct {
	ImageURL string `json:"image_url"`
	AltText  string `json:"alt_text"`
}

func (ImageBlockElement) messageElementType() messageElementType { return messageElementTImage }
func (ImageBlockElement) contextElementType() contextElementType { return contextElementImage }

type ButtonBlockElement struct {
	Text     TextBlockObject                `json:"text"`
	ActionID string                         `json:"action_id,omitempty"`
	URL      string                         `json:"url,omitempty"`
	Value    string                         `json:"value,omitempty"`
	Confirm  *ConfirmationDialogBlockObject `json:"confirm,omitempty"`
	Style    ButtonBlockStyle               `json:"style,omitempty"`
}

func (ButtonBlockElement) messageElementType() messageElementType { return messageElementTButton }

type SelectBlockElement interface {
	selectBlockElementType() messageElementType
}

type SelectBlockElementStatic struct {
	*selectBlockElementStatic
}

func (SelectBlockElementStatic) selectBlockElementType() messageElementType {
	return messageElementType(selectBlockElementTypeExternal)
}

type MultiSelectBlockElementStatic struct {
	*selectBlockElementStatic
}

func (MultiSelectBlockElementStatic) selectBlockElementType() messageElementType {
	return messageElementType(multiSelectBlockElementTypeStatic)
}

type SelectBlockElementExternal struct {
	*selectBlockElementExternal
}

func (SelectBlockElementExternal) selectBlockElementType() messageElementType {
	return messageElementType(selectBlockElementTypeExternal)
}

type MultiSelectBlockElementExternal struct {
	*selectBlockElementExternal
}

func (MultiSelectBlockElementExternal) selectBlockElementType() messageElementType {
	return messageElementType(multiSelectBlockElementTypeExternal)
}

type SelectBlockElementUsers struct {
	*selectBlockElementUsers
}

func (SelectBlockElementUsers) selectBlockElementType() messageElementType {
	return messageElementType(selectBlockElementTypeUsers)
}

type MultiSelectBlockElementUsers struct {
	*selectBlockElementUsers
}

func (MultiSelectBlockElementUsers) selectBlockElementType() messageElementType {
	return messageElementType(multiSelectBlockElementTypeUsers)
}

type SelectBlockElementConversations struct {
	*selectBlockElementConversations
}

func (SelectBlockElementConversations) selectBlockElementType() messageElementType {
	return messageElementType(selectBlockElementTypeConversations)
}

type MultiSelectBlockElementConversations struct {
	*selectBlockElementConversations
}

func (MultiSelectBlockElementConversations) selectBlockElementType() messageElementType {
	return messageElementType(multiSelectBlockElementTypeConversations)
}

type SelectBlockElementChannels struct {
	*selectBlockElementChannels
}

func (SelectBlockElementChannels) selectBlockElementType() messageElementType {
	return messageElementType(selectBlockElementTypeChannels)
}

type MultiSelectBlockElementChannels struct {
	*selectBlockElementChannels
}

func (MultiSelectBlockElementChannels) selectBlockElementType() messageElementType {
	return messageElementType(multiSelectBlockElementTypeChannels)
}

type OverflowBlockElement struct {
	ActionID string                         `json:"action_id"`
	Options  []*OptionBlockObject           `json:"options"`
	Confirm  *ConfirmationDialogBlockObject `json:"confirm,omitempty"`
}

func (OverflowBlockElement) messageElementType() messageElementType { return messageElementTOverflow }

type DatePickerElement struct {
	ActionID    string                         `json:"action_id"`
	Placeholder TextBlockObject                `json:"placeholder,omitempty"`
	InitialDate string                         `json:"initial_date,omitempty"`
	Confirm     *ConfirmationDialogBlockObject `json:"confirm,omitempty"`
}

func (DatePickerElement) messageElementType() messageElementType { return messageElementTDatepicker }

type Channel struct {
	Group
	IsChannel   bool `json:"is_channel"`
	IsGeneral   bool `json:"is_general"`
	IsMember    bool `json:"is_member"`
	IsOrgShared bool `json:"is_org_shared"`
	IsShared    bool `json:"is_shared"`
}

type ChannelList []*Channel

type ChannelsHistoryResponse struct {
	HasMore  bool        `json:"has_more"`
	Latest   string      `json:"latest"`
	Messages MessageList `json:"messages"`
}

type ChatResponse struct {
	Channel   string      `json:"channel"`
	Timestamp string      `json:"ts"`
	Message   interface{} `json:"message"` // TODO
}

type Comment struct {
	ID        string    `json:"id,omitempty"`
	Created   EpochTime `json:"created,omitempty"`
	Timestamp EpochTime `json:"timestamp,omitempty"`
	User      string    `json:"user,omitempty"`
	Comment   string    `json:"comment,omitempty"`
}

// Confirmation is used in conjunction with message buttons
type Confirmation struct {
	Title       string `json:"title"`
	Text        string `json:"text"`
	OkText      string `json:"ok_text"`
	DismissText string `json:"dismiss_text"`
}

// Conversation is a structure that is never used by itself:
// it's re-used to describe a basic conversation profile
// by being embedded in other objects
type Conversation struct {
	ID                 string    `json:"id"`
	Created            EpochTime `json:"created"`
	IsOpen             bool      `json:"is_open"`
	LastRead           string    `json:"last_read,omitempty"`
	Latest             *Message  `json:"latest,omitempty"`
	UnreadCount        int       `json:"unread_count,omitempty"`
	UnreadCountDisplay int       `json:"unread_count_display,omitempty"`
}

// DialogElements represents elements in dialog, including text, textarea,
// and select
type DialogElement struct {
	Label       string `json:"label"`
	Type        string `json:"type"`
	Name        string `json:"name"`
	Hint        string `json:"hint,omitempty"`
	MaxLength   int    `json:"max_length,omitempty"`
	MinLength   int    `json:"min_length,omitempty"`
	Optional    bool   `json:"optional,omitempty"`
	Placeholder string `json:"placeholder,omitempty"`
	Subtype     string `json:"subtype,omitempty"`
	Value       string `json:"value,omitempty"`
}

type DialogElements []*DialogElement

type Dialog struct {
	CallbackID  string         `json:"callback_id"`
	Title       string         `json:"title"`
	SubmitLabel string         `json:"submit_label"`
	Elements    DialogElements `json:"elements"`
}

type DialogResponse struct {
	ResponseMetadata struct {
		Messages []string `json:"messages"`
	} `json:"response_metadata"`
}

type EmojiListResponse map[string]string

type EphemeralResponse struct {
	MessageTimestamp string `json:"message_ts"`
}

// ErrorResponse wraps errors returned by Slack. It's usually a string,
// but it could be a structure.
// https://api.slack.com/rtm#handling_responses
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

// GenericResponse is the generic response part given by all
// slack API response.
type GenericResponse struct {
	OK        bool          `json:"ok"`
	ReplyTo   int           `json:"reply_to,omitempty"`
	Error     ErrorResponse `json:"error,omitempty"`
	Timestamp string        `json:"ts"`
}

// Group contains information about a private channel. Private channels were
// once known as "private groups."
type Group struct {
	Conversation
	Creator        string   `json:"creator"`
	IsArchived     bool     `json:"is_archived"`
	IsGroup        bool     `json:"is_group"`
	IsMPIM         bool     `json:"is_mpim"`
	Members        []string `json:"members"`
	Name           string   `json:"name"`
	NameNormalized string   `json:"name_normalized"`
	NumMembers     int      `json:"num_members,omitempty"`
	PreviousNames  []string `json:"previous_names,omitempty"`
	Purpose        Purpose  `json:"purpose"`
	Topic          Topic    `json:"topic"`
}

// GroupList is a list of groups.
type GroupList []*Group

// InteractiveButtonRequest is a request that is sent when a user
// hits a Slack button. Note: this is experimental
type InteractiveButtonRequest struct {
	ActionTimestamp string     `json:"action_ts"`
	Actions         ActionList `json:"actions"`
	AttachmentID    int        `json:"attachment_id,string"`
	CallbackID      string     `json:"callback_id"`
	Channel         struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"channel"`
	IsAppUnfurl      bool            `json:"is_app_unfurl"`
	MessageTimestamp string          `json:"message_ts"`
	OriginalMessage  *Message        `json:"original_message"`
	Options          OptionList      `json:"options"`
	OptionGroups     OptionGroupList `json:"option_groups"`
	ResponseURL      string          `json:"response_url"`
	Team             struct {
		Domain string `json:"domain"`
		ID     string `json:"id"`
	} `json:"team"`
	Token     string `json:"token"`
	TriggerID string `json:"trigger_id"`
	User      struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"user"`
}

type DialogSubmission struct {
	ActionTimestamp string            `json:"action_ts"`
	Submission      map[string]string `json:"submission"`
	Token           string            `json:"token"`
	TriggerID       string            `json:"trigger_id"`
	Team            struct {
		Domain string `json:"domain"`
		ID     string `json:"id"`
	} `json:"team"`
	Channel struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"channel"`
	User struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"user"`
}

type OAuthAccessResponse struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
}

type Purpose struct {
	Value   string    `json:"value"`
	Creator string    `json:"creator"`
	LastSet EpochTime `json:"last_set"`
}

type Topic struct {
	Value   string    `json:"value"`
	Creator string    `json:"creator"`
	LastSet EpochTime `json:"last_set"`
}

type Edited struct {
	Timestamp string `json:"ts"`
	User      string `json:"user"`
}

type AttachmentField struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"short"`
}
type AttachmentFieldList []*AttachmentField

// Message is a representation of a message, as obtained
// by the RTM or Events API. This is NOT what you use when
// you are posting a message. See ChatService#PostMessage
// and MessageParams for that.
type Message struct {
	Attachments AttachmentList `json:"attachments"`
	Channel     string         `json:"channel"`
	Edited      *Edited        `json:"edited"`
	IsStarred   bool           `json:"is_starred"`
	PinnedTo    []string       `json:"pinned_to"`
	Text        string         `json:"text"`
	Timestamp   string         `json:"ts"`
	Type        string         `json:"type"`
	User        string         `json:"user"`

	// Message Subtypes
	Subtype string `json:"subtype"`

	// Hidden Subtypes
	Hidden           bool   `json:"hidden,omitempty"`     // message_changed, message_deleted, unpinned_item
	DeletedTimestamp string `json:"deleted_ts,omitempty"` // message_deleted
	EventTimestamp   string `json:"event_ts,omitempty"`

	// bot_message (https://api.slack.com/events/message/bot_message)
	BotID    string `json:"bot_id,omitempty"`
	Username string `json:"username,omitempty"`
	Icons    *Icon  `json:"icons,omitempty"`

	// channel_join, group_join
	Inviter string `json:"inviter,omitempty"`

	// channel_topic, group_topic
	Topic string `json:"topic,omitempty"`

	// channel_purpose, group_purpose
	Purpose string `json:"purpose,omitempty"`

	// channel_name, group_name
	Name    string `json:"name,omitempty"`
	OldName string `json:"old_name,omitempty"`

	// channel_archive, group_archive
	Members []string `json:"members,omitempty"`

	// file_share, file_comment, file_mention
	//	File *File `json:"file,omitempty"`

	// file_share
	Upload bool `json:"upload,omitempty"`

	// file_comment
	Comment *Comment `json:"comment,omitempty"`

	// pinned_item
	ItemType string `json:"item_type,omitempty"`

	// https://api.slack.com/rtm
	ReplyTo int    `json:"reply_to,omitempty"`
	Team    string `json:"team,omitempty"`

	// reactions
	Reactions ReactionList `json:"reactions,omitempty"`

	Blocks BlockList `json:"blocks,omitempty"`
}

type MessageList []*Message

type IM struct {
	Conversation
	IsIM          bool   `json:"is_im"`
	User          string `json:"user"`
	IsUserDeleted bool   `json:"is_user_deleted"`
}

type Icon struct {
	IconURL   string `json:"icon_url,omitempty"`
	IconEmoji string `json:"icon_emoji,omitempty"`
}

type UserProfile struct {
	AlwaysActive       bool   `json:"always_active,omitempty"`
	AvatarHash         string `json:"avatar_hash,omitempty"`
	Email              string `json:"email,omitempty"`
	FirstName          string `json:"first_name,omitempty"`
	Image24            string `json:"image_24,omitempty"`
	Image32            string `json:"image_32,omitempty"`
	Image48            string `json:"image_48,omitempty"`
	Image72            string `json:"image_72,omitempty"`
	Image192           string `json:"image_192,omitempty"`
	Image512           string `json:"image_512,omitempty"`
	LastName           string `json:"last_name,omitempty"`
	RealName           string `json:"real_name,omitempty"`
	RealNameNormalized string `json:"real_name_normalized,omitempty"`
	StatusText         string `json:"status_text,omitempty"`
	StatusEmoji        string `json:"status_emoji,omitempty"`
}

type User struct {
	Color             string      `json:"color"`
	Deleted           bool        `json:"deleted"`
	ID                string      `json:"id"`
	IsAdmin           bool        `json:"is_admin"`
	IsBot             bool        `json:"is_bot"`
	IsOwner           bool        `json:"is_owner"`
	IsPrimaryOwner    bool        `json:"is_primary_owner"`
	IsRestricted      bool        `json:"is_restricted"`
	IsUltraRestricted bool        `json:"is_ultra_restricted"`
	Name              string      `json:"name"`
	Profile           UserProfile `json:"profile"`
	RealName          string      `json:"real_name"`
	Status            string      `json:"status,omitempty"`
	TeamID            string      `json:"team_id"`
	TZ                string      `json:"tz,omitempty"`
	TZLabel           string      `json:"tz_label"`
	TZOffset          int         `json:"tz_offset"`
	Updated           int         `json:"updated"`
}

// UserDetails is only provided by rtm.start response
type UserDetails struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	Created        EpochTime `json:"created"`
	ManualPresence string    `json:"manual_presence"`
	Prefs          UserPrefs `json:"prefs"`
}

type UserPrefs struct{} // TODO

type UserList []*User

type Presence string

const (
	Presencective Presence = "away"
	PresenceAway  Presence = "away"
)

type UserPresence struct {
	AutoAway        bool     `json:"auto_away,omitempty"`
	ConnectionCount int      `json:"connection_count,omitempty"`
	LastActivity    int      `json:"last_activity,omitempty"`
	ManualAway      bool     `json:"manual_away,omitempty"`
	Online          bool     `json:"online"`
	Presence        Presence `json:"presence"`
}

type Team struct {
	ID                string                 `json:"id"`
	Name              string                 `json:"name"`
	Domain            string                 `json:"domain"`
	EmailDomain       string                 `json:"email_domain"`
	EnterpriseID      string                 `json:"enterprise_id,omitempty"`
	EnterpriseName    string                 `json:"enterprise_name,omitempty"`
	Icon              map[string]interface{} `json:"icon"`
	MsgEditWindowMins int                    `json:"msg_edit_window_mins"`
	OverStorageLimit  bool                   `json:"over_storage_limit"`
	Prefs             interface{}            `json:"prefs"`
	Plan              string                 `json:"plan"`
}

type Bot struct {
	ID      string `json:"id"`
	AppID   string `json:"app_id"`
	Deleted bool   `json:"deleted"`
	Name    string `json:"name"`
	Icons   Icons  `json:"icons"`
}

type Icons struct {
	Image36 string `json:"image_36"`
	Image48 string `json:"image_48"`
	Image72 string `json:"image_72"`
}

// File represents a file object (https://api.slack.com/types/file)
type File struct {
	ID           string
	Name         string
	User         string
	Created      int    `json:"created"`   // The created property is a unix timestamp representing when the file was created.
	Timestamp    int    `json:"timestamp"` // Deprecated
	Updated      int    `json:"updated"`   // The updated property (for Post filetypes only) is a unix timestamp of when the Post was last edited.
	MimeType     string `json:"mimetype"`  // The mimetype and filetype props do not have a 1-to-1 mapping, as multiple different files types ('html', 'js', etc.) share the same mime type.
	FileType     string `json:"filetype"`
	PrettyType   string `json:"pretty_type"` // The pretty_type property contains a human-readable version of the type.
	Mode         string `json:"mode"`        // The mode property contains one of hosted, external, snippet or post.
	Editable     bool   `json:"editable"`    // The editable property indicates that files are stored in editable mode.
	IsExternal   bool   `json:"is_external"` // The is_external property indicates whether the master copy of a file is stored within the system or not. If the file is_external, then the url property will point to the externally-hosted master file. Further, the external_type property will indicate what kind of external file it is, e.g. dropbox or gdoc.
	ExternalType string `json:"external_type"`
	Size         int    `json:"size"` // The size parameter is the filesize in bytes. Snippets are limited to a maximum file size of 1 megabyte.

	URL                string `json:"url"`          // Deprecated - never set
	URLDownload        string `json:"url_download"` // Deprecated - never set
	URLPrivate         string `json:"url_private"`
	URLPrivateDownload string `json:"url_private_download"`

	ImageExifRotation string `json:"image_exif_rotation"`
	OriginalWidth     int    `json:"original_w"`
	OriginalHeight    int    `json:"original_h"`
	Thumb64           string `json:"thumb_64"`
	Thumb80           string `json:"thumb_80"`
	Thumb160          string `json:"thumb_160"`
	Thumb360          string `json:"thumb_360"`
	Thumb360Gif       string `json:"thumb_360_gif"`
	Thumb360W         int    `json:"thumb_360_w"`
	Thumb360H         int    `json:"thumb_360_h"`
	Thumb480          string `json:"thumb_480"`
	Thumb480W         int    `json:"thumb_480_w"`
	Thumb480H         int    `json:"thumb_480_h"`
	Thumb720          string `json:"thumb_720"`
	Thumb720W         int    `json:"thumb_720_w"`
	Thumb720H         int    `json:"thumb_720_h"`
	Thumb960          string `json:"thumb_960"`
	Thumb960W         int    `json:"thumb_960_w"`
	Thumb960H         int    `json:"thumb_960_h"`
	Thumb1024         string `json:"thumb_1024"`
	Thumb1024W        int    `json:"thumb_1024_w"`
	Thumb1024H        int    `json:"thumb_1024_h"`

	Permalink       string `json:"permalink"`
	PermalinkPublic string `json:"permalink_public"`

	EditLink         string `json:"edit_link"`
	Preview          string `json:"preview"`
	PreviewHighlight string `json:"preview_highlight"`
	Lines            int    `json:"lines"`
	LinesMore        int    `json:"lines_more"`

	IsPublic        bool     `json:"is_public"`
	PublicURLShared bool     `json:"public_url_shared"`
	Channels        []string `json:"channels"`
	Groups          []string `json:"groups"`
	IMs             []string `json:"ims"`
	InitialComment  Comment  `json:"initial_comment"`
	CommentsCount   int      `json:"comments_count"`
	NumStars        int      `json:"num_stars"`
	IsStarred       bool     `json:"is_starred"`

	Title     string       `json:"title"`
	Reactions ReactionList `json:"reactions,omitempty"`
}

type Option struct {
	Text        string `json:"text"`
	Value       string `json:"value"`
	Description string `json:"description"`
}
type OptionList []*Option

type OptionGroup struct {
	Text    string     `json:"text"`
	Options OptionList `json:"options"`
}
type OptionGroupList []*OptionGroup

type ThreadInfo struct {
	Complete bool `json:"complete"`
	Count    int  `json:"count"`
}

// Usergroup represents a single UserGroup (like @accounting or @marketing).
// This should not be confused with private channels or multi-person messages.
// This is an alias that points to a collection of users and/or channels for
// notification purposes.
type Usergroup struct {
	AutoProvision       bool            `json:"auto_provision"`
	AutoType            string          `json:"auto_type"`
	CreatedBy           string          `json:"created_by"`
	DateCreate          EpochTime       `json:"date_create"`
	DateDelete          EpochTime       `json:"date_delete"`
	DateUpdate          EpochTime       `json:"date_update"`
	DeletedBy           string          `json:"deleted_by"`
	Description         string          `json:"description"`
	EnterpriseSubteamID string          `json:"enterprise_subteam_id"`
	Handle              string          `json:"handle"`
	ID                  string          `json:"id"`
	IsExternal          bool            `json:"is_external"`
	IsSubteam           bool            `json:"is_subteam"`
	IsUsergroup         bool            `json:"is_usergroup"`
	Name                string          `json:"name"`
	Prefs               *UsergroupPrefs `json:"prefs"`
	TeamID              string          `json:"team_id"`
	UpdatedBy           string          `json:"updated_by"`
	Users               []string        `json:"users"`
	UserCount           int             `json:"user_count"`
}

// UsergroupList is a list of UserGroup
type UsergroupList []*Usergroup

// UsergroupPrefs is the list of preferences for channels and groups for a given
// Usergroup.
type UsergroupPrefs struct {
	Channels []string `json:"channels"`
	Groups   []string `json:"groups"`
}

// UsergroupUsersList is the list of users in a given Usergroup.
type UsergroupUsersList []string

type Reaction struct {
	Count int      `json:"count"`
	Name  string   `json:"name"`
	Users []string `json:"users"`
}
type ReactionList []*Reaction

// ReactionsGetResponse represents the response obtained from
// reactions.get API (https://api.slack.com/methods/reactions.get)
type ReactionsGetResponse struct {
	Channel string   `json:"channel"`
	Comment string   `json:"comment"`
	File    *File    `json:"file"`
	Message *Message `json:"message"`
	Type    string   `json:"type"`
}

type ReactionsGetResponseList []ReactionsGetResponse
type ReactionsListResponse struct {
	Items  ReactionsGetResponseList `json:"items"`
	Paging Paging                   `json:"paging"`
}

type Reminder struct {
	ID                string    `json:"id"`
	Creator           string    `json:"creator"`
	User              string    `json:"user"`
	Text              string    `json:"text"`
	Recurring         bool      `json:"recurring"`
	Time              EpochTime `json:"time,omitempty"`
	CompleteTimestamp EpochTime `json:"complete_ts,omitempty"`
}
type ReminderList []*Reminder

type RTMResponse struct {
	URL      string       `json:"url"`
	Self     *UserDetails `json:"self"`
	Team     *Team        `json:"team"`
	Users    []*User      `json:"users"`
	Channels []*Channel   `json:"channels"`
	Groups   []*Group     `json:"groups"`
	Bots     []*Bot       `json:"bots"`
	IMs      []*IM        `json:"ims"`
}

type Paging struct {
	Count int `json:"count"`
	Total int `json:"total"`
	Page  int `json:"page"`
	Pages int `json:"pages"`
}

type PermalinkResponse struct {
	Channel   string `json:"channel"`
	Permalink string `json:"permalink"`
}

type StarsListResponse struct {
	Items  StarredItemList `json:"items"`
	Paging Paging          `json:"paging"`
}
type StarredItem interface{}
type StarredItemList []StarredItem
