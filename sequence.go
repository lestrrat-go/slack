package slack

import (
	"bytes"
	"strings"

	pdebug "github.com/lestrrat-go/pdebug"
	"github.com/pkg/errors"
)

func stringifyLink(href, text string) string {
	var buf bytes.Buffer
	buf.WriteByte('<')
	buf.WriteString(href)
	if len(text) > 0 {
		buf.WriteByte('|')
		buf.WriteString(text)
	}
	buf.WriteByte('>')
	return buf.String()
}

func (l *ExternalLink) Data() string {
	return l.URL
}

func (l *ExternalLink) Surface() string {
	return l.Text
}

func (l *ExternalLink) String() string {
	return stringifyLink(l.URL, l.Text)
}

func (l *UserLink) Data() string {
	return l.ID
}

func (l *UserLink) Surface() string {
	return l.Username
}

func (l *UserLink) String() string {
	return stringifyLink(l.ID, l.Username)
}

func (l *ChannelLink) Data() string {
	return l.ID
}

func (l *ChannelLink) Surface() string {
	return l.Channel
}

func (l *ChannelLink) String() string {
	return stringifyLink(l.ID, l.Channel)
}

const ltmark = '<'
const gtmark = '>'

// https://api.slack.com/docs/message-formatting#control_sequences
func ExtractControlSequences(s string) ([]ControlSequence, error) {
	var list []ControlSequence
	for len(s) > 0 {
		start := strings.IndexByte(s, ltmark)
		if start == -1 {
			break
		}

		end := strings.IndexByte(s[start:], gtmark)
		if end == -1 {
			break
		}

		// Note that end is relative to start, so we need to
		// do `end + start` in order to move the string start
		// to the end of the control sequence
		rawseq := s[start : end+start+1]
		s = s[end+start:]

		seq, err := ParseControlSequence(rawseq)
		if err != nil {
			return nil, errors.Wrapf(err, `failed to parse '%s'`, rawseq)
		}
		list = append(list, seq)
	}

	return list, nil
}

var invalidSequence = errors.New("invalid control sequence format")

func ParseControlSequence(s string) (ControlSequence, error) {
	if pdebug.Enabled {
		pdebug.Printf("parsing control sequence %s", s)
	}

	if len(s) < 3 {
		return nil, invalidSequence
	}

	if s[0] != ltmark {
		return nil, invalidSequence
	}

	s = s[1:]

	if s[len(s)-1] != gtmark {
		return nil, invalidSequence
	}
	s = s[:len(s)-1]

	switch s[0] {
	case ltmark, gtmark:
		return nil, invalidSequence
	case '@':
		return parseUser(s)
	case '#':
		return parseChannel(s)
	case '!':
		return parseBang(s)
	default:
		return parseExternalLink(s)
	}
}

func parseBang(s string) (ControlSequence, error) {
	return nil, errors.New("parsing '!' sequences unimplemented")
}

func parseExternalLink(s string) (ControlSequence, error) {
	if i := strings.IndexByte(s, '|'); i != -1 {
		return &ExternalLink{
			URL:  s[:i],
			Text: s[i+1:],
		}, nil
	}
	return &ExternalLink{
		URL:  s,
		Text: s,
	}, nil
}

func parseUser(s string) (ControlSequence, error) {
	if i := strings.IndexByte(s, '|'); i != -1 {
		return &UserLink{
			ID:       s[:i],
			Username: s[i+1:],
		}, nil
	}
	return &UserLink{
		ID:       s,
		Username: s,
	}, nil
}

func parseChannel(s string) (ControlSequence, error) {
	if i := strings.IndexByte(s, '|'); i != -1 {
		return &ChannelLink{
			ID:      s[:i],
			Channel: s[i+1:],
		}, nil
	}
	return &ChannelLink{
		ID:      s,
		Channel: s,
	}, nil
}
