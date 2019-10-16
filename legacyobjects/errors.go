package objects

import (
	"encoding/json"
	"unicode"
	"unicode/utf8"

	"github.com/pkg/errors"
)

func (e *ErrorResponse) String() string {
	return e.Message
}

func (e *ErrorResponse) UnmarshalJSON(data []byte) error {
	// check for the first non-whitespace character
	for len(data) > 0 {
		r, w := utf8.DecodeRune(data)
		if unicode.IsSpace(r) {
			data = data[w:]
			continue
		}
		break
	}

	if len(data) > 0 {
		switch data[0] {
		case '"':
			*e = ErrorResponse{}
			return json.Unmarshal(data, &e.Message)
		case '{':
			var dummy struct {
				Code    int    `json:"code"`
				Message string `json:"msg"`
			}
			if err := json.Unmarshal(data, &dummy); err != nil {
				return errors.Wrap(err, `failed to unmarshal structured error`)
			}
			e.Code = dummy.Code
			e.Message = dummy.Message
			return nil
		}
	}

	return errors.New("invalid JSON")
}
