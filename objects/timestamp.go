package objects

import (
	"encoding/json"
	"strconv"
	"unicode"
	"unicode/utf8"

	"github.com/pkg/errors"
)

func (t EpochTime) Int() int {
	return int(t)
}
func (t EpochTime) Add(seconds int64) EpochTime {
	return EpochTime(int64(t) + seconds)
}

func (t *Timestamp) UnmarshalJSON(data []byte) error {
	// check for the first non-whitespace character
	for len(data) > 0 {
		r, w := utf8.DecodeRune(data)
		if unicode.IsSpace(r) {
			data = data[w:]
			continue
		}
		break
	}

	if len(data) == 0 {
		return errors.New("invalid JSON")
	}

	switch data[0] {
	case '"':
		var s string
		if err := json.Unmarshal(data, &s); err != nil {
			return errors.Wrap(err, `failed to unmarshal string timestamp`)
		}
		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return errors.Wrap(err, `failed to convert string timestamp to float`)
		}
		*t = Timestamp(f)
	default:
		// otherwise it's a number
		var f float64
		if err := json.Unmarshal(data, &f); err != nil {
			return errors.Wrap(err, `failed to unmarshal numeric timestamp`)
		}
		*t = Timestamp(f)
	}
	return nil
}
