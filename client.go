package slack

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// Option defines an interface of optional parameters to the
// `slack.New` constructor.
type Option interface {
	Name() string
	Value() interface{}
}

const (
	debugkey    = "debug"
	httpclkey   = "httpclient"
	loggerkey   = "logger"
	slackurlkey = "slackurl"
)

// New creates a new REST Slack API client. The `token` is
// required. Other optional parameters can be passed using the
// various `WithXXXX` functions
func New(token string, options ...Option) *Client {
	slackURL := DefaultAPIEndpoint
	httpcl := http.DefaultClient
	debug, _ := strconv.ParseBool(os.Getenv("SLACK_DEBUG"))
	var logger Logger = nilLogger{}
	for _, o := range options {
		switch o.Name() {
		case httpclkey:
			httpcl = o.Value().(*http.Client)
		case debugkey:
			debug = o.Value().(bool)
		case loggerkey:
			logger = o.Value().(Logger)
		case slackurlkey:
			slackURL = o.Value().(string)
		}
	}

	if !strings.HasSuffix(slackURL, "/") {
		slackURL = slackURL + "/"
	}

	if _, ok := logger.(nilLogger); debug && ok {
		logger = traceLogger{dst: os.Stderr}
	}

	wrappedcl := &httpClient{
		client:   httpcl,
		debug:    debug,
		slackURL: slackURL,
		logger:   logger,
	}
	return &Client{
		auth:            &AuthService{client: wrappedcl, token: token},
		bots:            &BotsService{client: wrappedcl, token: token},
		channels:        &ChannelsService{client: wrappedcl, token: token},
		chat:            &ChatService{client: wrappedcl, token: token},
		dialog:          &DialogService{client: wrappedcl, token: token},
		emoji:           &EmojiService{client: wrappedcl, token: token},
		groups:          &GroupsService{client: wrappedcl, token: token},
		oauth:           &OAuthService{client: wrappedcl},
		reminders:       &RemindersService{client: wrappedcl, token: token},
		reactions:       &ReactionsService{client: wrappedcl, token: token},
		rtm:             &RTMService{client: wrappedcl, token: token},
		users:           &UsersService{client: wrappedcl, token: token},
		usersProfile:    &UsersProfileService{client: wrappedcl, token: token},
		usergroups:      &UsergroupsService{client: wrappedcl, token: token},
		usergroupsUsers: &UsergroupsUsersService{client: wrappedcl, token: token},
		debug:           debug,
	}
}

// Auth returns the Service object for `auth.*` endpoints
func (c *Client) Auth() *AuthService {
	return c.auth
}

// Bots returns the Service object for `bots.*` endpoints
func (c *Client) Bots() *BotsService {
	return c.bots
}

// Channels returns the Service object for `channels.*` endpoints
func (c *Client) Channels() *ChannelsService {
	return c.channels
}

// Chat returns the Service object for `chat.*` endpoints
func (c *Client) Chat() *ChatService {
	return c.chat
}

// Dialog returns the Service object for `dialog.*` endpoints
func (c *Client) Dialog() *DialogService {
	return c.dialog
}

// Emoji returns the Service object for `emoji.*` endpoints
func (c *Client) Emoji() *EmojiService {
	return c.emoji
}

// Groups returns the Service object for `emoji.*` endpoints
func (c *Client) Groups() *GroupsService {
	return c.groups
}

// OAuth returns the Service object for `oauth.*` endpoints
func (c *Client) OAuth() *OAuthService {
	return c.oauth
}

// Reactions returns the Service object for `reactions.*` endpoints
func (c *Client) Reactions() *ReactionsService {
	return c.reactions
}

// Reminders returns the Service object for `reminders.*` endpoints
func (c *Client) Reminders() *RemindersService {
	return c.reminders
}

// RTM returns the Service object for `rtm.*` endpoints
func (c *Client) RTM() *RTMService {
	return c.rtm
}

// Users returns the Service object for `users.*` endpoints
func (c *Client) Users() *UsersService {
	return c.users
}

// UsersProfile returns the Service object for `users.profile.*` endpoints
func (c *Client) UsersProfile() *UsersProfileService {
	return c.usersProfile
}

// Usergroups returns the Service object for `usergroups.*` endpoints
func (c *Client) Usergroups() *UsergroupsService {
	return c.usergroups
}

// UsergroupsUsers returns the Service object for `usergroups.users.*` endpoints
func (c *Client) UsergroupsUsers() *UsergroupsUsersService {
	return c.usergroupsUsers
}

type httpClient struct {
	client   *http.Client
	debug    bool
	slackURL string
	logger   Logger
}

func (c *httpClient) makeSlackURL(path string) string {
	return c.slackURL + path
}

// path is only passed for debugging purposes
func (c *httpClient) parseResponse(ctx context.Context, path string, rdr io.Reader, res interface{}) error {
	if c.debug {
		var buf bytes.Buffer
		io.Copy(&buf, rdr)

		c.logger.Debugf(ctx, "-----> %s (response)", path)
		var m map[string]interface{}
		if err := json.Unmarshal(buf.Bytes(), &m); err != nil {
			c.logger.Debugf(ctx, "failed to unmarshal payload: %s", err)
			c.logger.Debugf(ctx, buf.String())
		} else {
			formatted, _ := json.MarshalIndent(m, "", "  ")
			c.logger.Debugf(ctx, "%s", formatted)
		}
		c.logger.Debugf(ctx, "<----- %s (response)", path)
		rdr = &buf
	}
	return json.NewDecoder(rdr).Decode(res)
}

func (c *httpClient) postForm(ctx context.Context, path string, f url.Values, data interface{}) error {
	if c.debug {
		c.logger.Debugf(ctx, "-----> %s (request)", path)
		for k, list := range f {
			var buf bytes.Buffer
			if k == "token" {
				buf.WriteString("xxxxxxxxxxxx (redacted)")
			} else {
				for i, v := range list {
					buf.WriteString(v)
					if i < len(list)-1 {
						buf.WriteString(", ")
					}
				}
			}
			c.logger.Debugf(ctx, "%s = %s", k, buf.String())
		}
		c.logger.Debugf(ctx, "<----- %s (request)", path)
	}
	return c.post(ctx, path, "application/x-www-form-urlencoded", strings.NewReader(f.Encode()), data)
}

func (c *httpClient) post(octx context.Context, path, ct string, body io.Reader, data interface{}) error {
	u := c.makeSlackURL(path)
	if c.debug {
		c.logger.Debugf(octx, "posting to %s", u)
	}
	req, err := http.NewRequest(http.MethodPost, u, body)
	if err != nil {
		return errors.New(`failed to create new POST request`)
	}

	ctx, cancel := context.WithCancel(octx)
	defer cancel()

	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", ct)
	res, err := c.client.Do(req)
	if err != nil {
		return errors.Wrap(err, `failed to post to slack`)
	}
	defer res.Body.Close()
	return c.parseResponse(octx, path, res.Body, data)
}

func (c *httpClient) get(octx context.Context, path string, f url.Values, data interface{}) error {
	ustr := c.makeSlackURL(path)
	u, err := url.Parse(ustr)
	if err != nil {
		return errors.Wrapf(err, `failed to parse get url: '%s'`, ustr)
	}
	u.RawQuery = f.Encode()

	if c.debug {
		c.logger.Debugf(octx, "> GET %s", u.String())
	}
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return errors.New(`failed to create new GET request`)
	}

	ctx, cancel := context.WithCancel(octx)
	defer cancel()

	req = req.WithContext(ctx)
	res, err := c.client.Do(req)
	if err != nil {
		return errors.Wrap(err, `failed to get slack`)
	}
	defer res.Body.Close()
	return c.parseResponse(octx, path, res.Body, data)
}
