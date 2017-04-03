package slack

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/pkg/errors"
)

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
	httpclkey   = "httpclient"
	slackurlkey = "slackurl"
)

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

func New(token string, options ...Option) *Client {
	slackURL := DefaultAPIEndpoint
	httpcl := http.DefaultClient
	for _, o := range options {
		switch o.Name() {
		case httpclkey:
			httpcl = o.Value().(*http.Client)
		case slackurlkey:
			slackURL = o.Value().(string)
		}
	}

	if !strings.HasSuffix(slackURL, "/") {
		slackURL = slackURL + "/"
	}
	fmt.Println(slackURL)

	wrappedcl := &httpClient{
		client:   httpcl,
		slackURL: slackURL,
	}
	return &Client{
		auth:  &AuthService{client: wrappedcl, token: token},
		chat:  &ChatService{client: wrappedcl, token: token},
		users: &UsersService{client: wrappedcl, token: token},
	}
}

func (c *Client) Auth() *AuthService {
	return c.auth
}

func (c *Client) Chat() *ChatService {
	return c.chat
}

func (c *Client) Users() *UsersService {
	return c.users
}

type httpClient struct {
	client   *http.Client
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

func (c *httpClient) post(ctx context.Context, path, ct string, body io.Reader, data interface{}) error {
	u := c.makeSlackURL(path)
	req, err := http.NewRequest(http.MethodPost, u, body)
	if err != nil {
		return errors.New(`failed to create new POST request`)
	}

	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", ct)
	res, err := c.client.Do(req)
	if err != nil {
		return errors.Wrap(err, `failed to post to slack`)
	}
	defer res.Body.Close()

	return c.parseResponse(res.Body, data)
}
