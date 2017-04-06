package slack

import (
	"bytes"
	"log"
	"strings"

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

func (l *ExternalLink) String() string {
	return stringifyLink(l.URL, l.Text)
}

func (l *UserLink) String() string {
	return stringifyLink(l.ID, l.Username)
}

func (l *ChannelLink) String() string {
	return stringifyLink(l.ID, l.Channel)
}

// https://api.slack.com/docs/message-formatting#control_sequences
func ExtractControlSequences(s string) ([]ControlSequence, error) {
	const ltmark = '<'
	const gtmark = '>'

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
		rawseq := s[start+1 : end+start]
		s = s[end+start:]

		log.Printf("rawseq: %s", rawseq)
		seq, err := ParseControlSequence(rawseq)
		if err != nil {
			return nil, errors.Wrapf(err, `failed to parse '%s'`, rawseq)
		}
		list = append(list, seq)
	}

	return list, nil
}

var invalidLink = errors.New("invalid link format")

func ParseControlSequence(s string) (ControlSequence, error) {
	if len(s) == 0 {
		return nil, invalidLink
	}

	switch s[0] {
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
