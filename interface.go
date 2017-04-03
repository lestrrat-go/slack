package slack

const DefaultSlackURL = "https://slack.com/api/"

type Client struct {
	auth     *AuthService
	chat     *ChatService
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
