package objects

import (
	"encoding/json"

	"github.com/pkg/errors"
)

func (l *ActionList) Append(a *Action) *ActionList {
	*l = append(*l, a)
	return l
}

func (l *AttachmentList) Append(a *Attachment) *AttachmentList {
	*l = append(*l, a)
	return l
}

func (l *AttachmentList) Encode() (string, error) {
	buf, err := json.Marshal(l)
	if err != nil {
		return "", errors.Wrap(err, `failed to encode attachment list`)
	}
	return string(buf), nil
}

func (l *AttachmentFieldList) Append(a *AttachmentField) *AttachmentFieldList {
	*l = append(*l, a)
	return l
}
