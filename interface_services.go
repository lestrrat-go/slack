package slack

// See interface.go for general struct definitions
// See interface_objects.go for Slack object definitions
// See interface_services.go for Slack service definitions

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

type StartRTMResponse struct {
	URL      string       `json:"url"`
	Self     *UserDetails `json:"self"`
	Team     *Team        `json:"team"`
	Users    []*User      `json:"users"`
	Channels []Channel    `json:"channels"`
	Groups   []Group      `json:"groups"`
	Bots     []Bot        `json:"bots"`
	IMs      []IM         `json:"ims"`
}

// UsersService handles all `users.*` API endpoints
type UsersService struct {
	client *httpClient
	token  string
}

