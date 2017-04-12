package slack

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
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

type option struct {
	name  string
	value interface{}
}

func (o *option) Name() string {
	return o.name
}
func (o *option) Value() interface{} {
	return o.value
}

const (
	debugkey    = "debug"
	httpclkey   = "httpclient"
	slackurlkey = "slackurl"
)

// WithClient allows you to specify an net/http.Client object to
// use to communicate with the Slack endpoints. For example, if you
// need to use this in Google App Engine, you can pass it the
// result of `urlfetch.Client`
func WithClient(cl *http.Client) Option {
	return &option{
		name:  httpclkey,
		value: cl,
	}
}

// WithAPIEndpoint allows you to specify an alternate API endpoint.
// The default is DefaultAPIEndpoint.
func WithAPIEndpoint(s string) Option {
	return &option{
		name:  slackurlkey,
		value: s,
	}
}

// WithDebug specifies that we want to run in debugging mode.
// You can set this value manually to override any existing global
// defaults.
//
// If one is not specified, the default value is false, or the
// value specified in SLACK_DEBUG environment variable
func WithDebug(b bool) Option {
	return &option{
		name:  debugkey,
		value: b,
	}
}

// New creates a new REST Slack API client. The `token` is
// required. Other optional parameters can be passed using the
// various `WithXXXX` functions
func New(token string, options ...Option) *Client {
	slackURL := DefaultAPIEndpoint
	httpcl := http.DefaultClient
	debug, _ := strconv.ParseBool(os.Getenv("SLACK_DEBUG"))
	for _, o := range options {
		switch o.Name() {
		case httpclkey:
			httpcl = o.Value().(*http.Client)
		case debugkey:
			debug = o.Value().(bool)
		case slackurlkey:
			slackURL = o.Value().(string)
		}
	}

	if !strings.HasSuffix(slackURL, "/") {
		slackURL = slackURL + "/"
	}

	wrappedcl := &httpClient{
		client:   httpcl,
		debug:    debug,
		slackURL: slackURL,
	}
	return &Client{
		auth:      &AuthService{client: wrappedcl, token: token},
		bots:      &BotsService{client: wrappedcl, token: token},
		channels:  &ChannelsService{client: wrappedcl, token: token},
		chat:      &ChatService{client: wrappedcl, token: token},
		oauth:     &OAuthService{client: wrappedcl},
		reactions: &ReactionsService{client: wrappedcl, token: token},
		rtm:       &RTMService{client: wrappedcl, token: token},
		users:     &UsersService{client: wrappedcl, token: token},
		debug:     debug,
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

// OAuth returns the Service object for `oauth.*` endpoints
func (c *Client) OAuth() *OAuthService {
	return c.oauth
}

// Reactions returns the Service object for `reactions.*` endpoints
func (c *Client) Reactions() *ReactionsService {
	return c.reactions
}

// RTM returns the Service object for `rtm.*` endpoints
func (c *Client) RTM() *RTMService {
	return c.rtm
}

// Users returns the Service object for `users.*` endpoints
func (c *Client) Users() *UsersService {
	return c.users
}

type httpClient struct {
	client   *http.Client
	debug    bool
	slackURL string
}

func (c *httpClient) makeSlackURL(path string) string {
	return c.slackURL + path
}

// path is only passed for debugging purposes
func (c *httpClient) parseResponse(path string, rdr io.Reader, res interface{}) error {
	if c.debug {
		var buf bytes.Buffer
		io.Copy(&buf, rdr)
		log.Printf("-----> %s", path)
		var m map[string]interface{}
		if err := json.Unmarshal(buf.Bytes(), &m); err != nil {
			log.Printf("failed to unmarshal payload: %s", err)
		} else {
			formatted, _ := json.MarshalIndent(m, "", "  ")
			log.Printf("%s", formatted)
		}
		log.Printf("<----- %s", path)
		rdr = &buf
	}
	return json.NewDecoder(rdr).Decode(res)
}

func (c *httpClient) postForm(ctx context.Context, path string, f url.Values, data interface{}) error {
	return c.post(ctx, path, "application/x-www-form-urlencoded", strings.NewReader(f.Encode()), data)
}

func (c *httpClient) post(octx context.Context, path, ct string, body io.Reader, data interface{}) error {
	u := c.makeSlackURL(path)
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
	return c.parseResponse(path, res.Body, data)
}

func (c *httpClient) get(octx context.Context, path string, f url.Values, data interface{}) error {
	ustr := c.makeSlackURL(path)
	u, err := url.Parse(ustr)
	if err != nil {
		return errors.Wrapf(err, `failed to parse get url: '%s'`, ustr)
	}
	u.RawQuery = f.Encode()

	if c.debug {
		log.Printf("> GET %s", u.String())
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
	return c.parseResponse(path, res.Body, data)
}

func (r SlackResponse) err() error {
	return errors.New(r.Error.String())
}

func (r SlackResponse) ok() bool {
	return r.OK
}

type genericResponse interface {
	ok() bool
	err() error
}

func genericPost(ctx context.Context, client *httpClient, endpoint string, v url.Values, res genericResponse) error {
	if err := client.postForm(ctx, endpoint, v, res); err != nil {
		return errors.Wrapf(err, `failed to post to %s`, endpoint)
	}

	if !res.ok() {
		return res.err()
	}
	return nil
}
