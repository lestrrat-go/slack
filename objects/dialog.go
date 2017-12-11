package objects

import (
	"encoding/json"

	"github.com/pkg/errors"
)

func (d *Dialog) Encode() (string, error) {
	buf, err := json.Marshal(d)
	if err != nil {
		return "", errors.Wrap(err, `failed to encode user profile`)
	}
	return string(buf), nil
}

func (d *Dialog) Decode(buf string) error {
	if buf == "" {
		return nil
	}
	return json.Unmarshal([]byte(buf), d)
}
