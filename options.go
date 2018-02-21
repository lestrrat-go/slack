package slack

import (
	"net/http"

	"github.com/lestrrat-go/slack/internal/option"
)

// WithClient allows you to specify an net/http.Client object to
// use to communicate with the Slack endpoints. For example, if you
// need to use this in Google App Engine, you can pass it the
// result of `urlfetch.Client`
func WithClient(cl *http.Client) Option {
	return option.New(httpclkey, cl)
}

// WithAPIEndpoint allows you to specify an alternate API endpoint.
// The default is DefaultAPIEndpoint.
func WithAPIEndpoint(s string) Option {
	return option.New(slackurlkey, s)
}

// WithDebug specifies that we want to run in debugging mode.
// You can set this value manually to override any existing global
// defaults.
//
// If one is not specified, the default value is false, or the
// value specified in SLACK_DEBUG environment variable
func WithDebug(b bool) Option {
	return option.New(debugkey, b)
}

// WithLogger specifies the logger object to be used.
// If not specified and `WithDebug` is enabled, then a default
// logger which writes to os.Stderr
func WithLogger(l Logger) Option {
	return option.New(loggerkey, l)
}
