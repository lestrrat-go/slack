package slack

import "golang.org/x/oauth2"

// DefaultSlackAPIEndpoint contains the prefix used for Slack REST API
const (
	DefaultAPIEndpoint         = "https://slack.com/api/"
	DefaultOAuth2AuthEndpoint  = "https://slack.com/oauth/authorize"
	DefaultOAuth2TokenEndpoint = "https://slack.com/api/oauth.access"
)

// Oauth2Endpoint contains the Slack OAuth2 endpoint configuration
var OAuth2Endpoint = oauth2.Endpoint{
	AuthURL:  DefaultOAuth2AuthEndpoint,
	TokenURL: DefaultOAuth2TokenEndpoint,
}

type Client struct {
	auth     *AuthService
	chat     *ChatService
	users    *UsersService
	slackURL string
	token    string
}

type SlackResponse struct {
	OK    bool   `json:"ok"`
	Error string `json:"error"`
}

type AuthService struct {
	client *httpClient
	token  string
}

type AuthTestResponse struct {
	URL    string `json:"url"`
	Team   string `json:"team"`
	User   string `json:"user"`
	TeamID string `json:"team_id"`
	UserID string `json:"user_id"`
}

type Attachment interface{} // TODO
type ChatService struct {
	client *httpClient
	token  string
}

type ChatResponse struct {
	Channel   string      `json:"channel"`
	Timestamp string      `json:"ts"`
	Message   interface{} `json:"message"` // TODO
}

type MessageParams struct {
	AsUser      bool
	Attachments []Attachment
	EscapeText  bool
	IconEmoji   string
	IconURL     string
	LinkNames   bool
	Markdown    bool `json:"mrkdwn,omitempty"`
	Parse       string
	UnfurlLinks bool
	UnfurlMedia bool
	Username    string
}

type UsersService struct {
	client *httpClient
	token  string
}

type UserProfile struct {
	AlwaysActive       bool   `json:"always_active"`
	AvatarHash         string `json:"avatar_hash"`
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

type UserList []*User

type EventType int

const (
	RTMConnectingEvent EventType = iota
	MaxEvent
)

type Event interface {
	Type() EventType
}

type RTM struct {
	outch chan Event
}

type RTMEvent struct {
	typ  EventType
	data interface{}
}
