package slack

// See interface.go for general struct definitions
// See interface_objects.go for Slack object definitions
// See interface_services.go for Slack service definitions

import "golang.org/x/oauth2"

type EpochTime int64

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
	debug    bool
	slackURL string
	token    string
}
