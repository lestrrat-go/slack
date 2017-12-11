package slack

import (
	"context"

	"golang.org/x/oauth2"
)

// Logger is an interface for logging/tracing the client's
// execution.
//
// In particular, `Debugf` will only be called if `WithDebug`
// is provided to the constructor.
type Logger interface {
	Debugf(context.Context, string, ...interface{})
	Infof(context.Context, string, ...interface{})
}

const (
	ParseFull = "full"
	ParseNone = "none"
)

type ControlSequence interface {
	Data() string
	Surface() string
	String() string
}

type ChannelLink struct {
	ID      string
	Channel string
}

type UserLink struct {
	ID       string
	Username string
}

type ExternalLink struct {
	URL  string
	Text string
}

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
	auth            *AuthService
	bots            *BotsService
	channels        *ChannelsService
	chat            *ChatService
	dialog          *DialogService
	emoji           *EmojiService
	groups          *GroupsService
	oauth           *OAuthService
	reactions       *ReactionsService
	reminders       *RemindersService
	rtm             *RTMService
	users           *UsersService
	usersProfile    *UsersProfileService
	usergroups      *UsergroupsService
	usergroupsUsers *UsergroupsUsersService
	debug           bool
	slackURL        string
	token           string
}
