package slack

import (
	"github.com/lestrrat/go-slack/objects"
	"golang.org/x/oauth2"
)

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
	rtm      *RTMService
	users    *UsersService
	debug    bool
	slackURL string
	token    string
}

// SlackResponse is the general response part given by all
// slack API response.
type SlackResponse struct {
	OK    bool   `json:"ok"`
	Error string `json:"error"`
}

// AuthService handles all `auth.*` API endpoints
type AuthService struct {
	client *httpClient
	token  string
}

// AuthTestResponse is the data structure response from auth.test
type AuthTestResponse struct {
	URL    string `json:"url"`
	Team   string `json:"team"`
	User   string `json:"user"`
	TeamID string `json:"team_id"`
	UserID string `json:"user_id"`
}

// ChatService handles all `chat.*` API endpoints
type ChatService struct {
	client *httpClient
	token  string
}

type ChatResponse struct {
	Channel   string      `json:"channel"`
	Timestamp string      `json:"ts"`
	Message   interface{} `json:"message"` // TODO
}

// RTMService handles all `rtm.*` API endpoints
type RTMService struct {
	client *httpClient
	token  string
}

type RTMResponse struct {
	URL      string               `json:"url"`
	Self     *objects.UserDetails `json:"self"`
	Team     *objects.Team        `json:"team"`
	Users    []*objects.User      `json:"users"`
	Channels []*objects.Channel   `json:"channels"`
	Groups   []*objects.Group     `json:"groups"`
	Bots     []*objects.Bot       `json:"bots"`
	IMs      []*objects.IM        `json:"ims"`
}

// UsersService handles all `users.*` API endpoints
type UsersService struct {
	client *httpClient
	token  string
}
