package objects

type EpochTime int64
type Timestamp float64

const (
	ButtonActionType = "button"
)

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

type GroupConversation struct {
	Conversation
	Creator       string   `json:"creator"`
	IsArchived    bool     `json:"is_archived"`
	Members       []string `json:"members"`
	Name          string   `json:"name"`
	NumMembers    int      `json:"num_members,omitempty"`
	PreviousNames []string `json:"previous_names"`
	Purpose       Purpose  `json:"purpose"`
	Topic         Topic    `json:"topic"`
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

type Channel struct {
	GroupConversation
	IsChannel   bool `json:"is_channel"`
	IsGeneral   bool `json:"is_general"`
	IsMember    bool `json:"is_member"`
	IsOrgShared bool `json:"is_org_shared"`
	IsShared    bool `json:"is_shared"`
}

type ChannelList []*Channel

// Confirmation is used in conjunction with message buttons
type Confirmation struct {
	Title       string `json:"title"`
	Text        string `json:"text"`
	OkText      string `json:"ok_text"`
	DismissText string `json:"dismiss_text"`
}

type Comment struct {
	ID        string    `json:"id,omitempty"`
	Created   EpochTime `json:"created,omitempty"`
	Timestamp EpochTime `json:"timestamp,omitempty"`
	User      string    `json:"user,omitempty"`
	Comment   string    `json:"comment,omitempty"`
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
	Reactions []ItemReaction `json:"reactions,omitempty"`
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

type ItemReaction struct {
	Name  string   `json:"name"`
	Count int      `json:"count"`
	Users []string `json:"users"`
}

type UserProfile struct {
	AlwaysActive       bool   `json:"always_active"`
	AvatarHash         string `json:"avatar_hash"`
	Email              string `json:"email"`
	FirstName          string `json:"first_name"`
	Image24            string `json:"image_24"`
	Image32            string `json:"image_32"`
	Image48            string `json:"image_48"`
	Image72            string `json:"image_72"`
	Image192           string `json:"image_192"`
	Image512           string `json:"image_512"`
	LastName           string `json:"last_name"`
	RealName           string `json:"real_name"`
	RealNameNormalized string `json:"real_name_normalized"`
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
	Update            int         `json:"updated"`
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

type Group interface{}
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
