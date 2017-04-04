package slack

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
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
// XXX I wrote it, but haven't actually implemented debugging yet
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
	var debug bool
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
		auth:  &AuthService{client: wrappedcl, token: token},
		chat:  &ChatService{client: wrappedcl, token: token},
		rtm:   &RTMService{client: wrappedcl, token: token},
		users: &UsersService{client: wrappedcl, token: token},
		debug: debug,
	}
}

// Auth returns the Service object for `auth.*` endpoints
func (c *Client) Auth() *AuthService {
	return c.auth
}

// Chat returns the Service object for `chat.*` endpoints
func (c *Client) Chat() *ChatService {
	return c.chat
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

func (c *httpClient) parseResponse(rdr io.Reader, res interface{}) error {
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

	return c.parseResponse(res.Body, data)
}
