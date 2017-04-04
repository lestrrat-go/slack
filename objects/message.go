package objects

import (
	"bytes"
	"encoding/json"
	"net/url"

	"github.com/pkg/errors"
)

// NewMessageParams -- XXX This method and the MessageParams things
// just feels so out of place within the context of the rest of the
// API. perhaps we could do something better.
func NewMessageParams() *MessageParams {
	return &MessageParams{
		// everything else should be initialzed to the zero value
		EscapeText:  true,
		Markdown:    true,
		UnfurlMedia: true,
	}
}

func (p *MessageParams) Values(v url.Values) error {
	if p.AsUser {
		v.Set("as_user", "true")
	}
	if len(p.Attachments) > 0 {
		var buf bytes.Buffer
		if err := json.NewEncoder(&buf).Encode(p.Attachments); err != nil {
			return errors.Wrap(err, `failed to serialize attachments`)
		}
		v.Set("attachments", buf.String())
	}
	if len(p.IconEmoji) > 0 {
		v.Set("icon_emoji", p.IconEmoji)
	}
	if len(p.IconURL) > 0 {
		v.Set("icon_url", p.IconURL)
	}
	if p.LinkNames {
		v.Set("link_names", "true")
	}

	if !p.Markdown {
		v.Set("mrkdwn", "false")
	}

	if len(p.Parse) > 0 {
		v.Set("parse", p.Parse)
	}

	// taken from github.com/nlopes/slack:
	//    I want to send a message with explicit `as_user` `true` and
	//    `unfurl_links` `false` in request. Because setting `as_user` to
	//    `true` will change the default value for `unfurl_links` to `true`
	//    on Slack API side.
	if p.AsUser && !p.UnfurlLinks {
		v.Set("unfurl_link", "false")
	} else if p.UnfurlLinks {
		v.Set("unfurl_link", "true")
	}

	if p.UnfurlMedia {
		v.Set("unfurl_media", "true")
	}
	if len(p.Username) > 0 {
		v.Set("username", p.Username)
	}
	return nil
}
