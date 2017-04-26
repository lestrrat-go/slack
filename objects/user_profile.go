package objects

import (
	"encoding/json"

	"github.com/pkg/errors"
)

func (p *UserProfile) Encode() (string, error) {
	buf, err := json.Marshal(p)
	if err != nil {
		return "", errors.Wrap(err, `failed to encode user profile`)
	}
	return string(buf), nil
}
