package server

import "github.com/lestrrat-go/slack/internal/option"

// WithPrefix allows you to specify a custom prefix in the URL path
// for the server. By default, "/api" is used
func WithPrefix(s string) Option {
	return option.New(optkeyPrefix, s)
}
