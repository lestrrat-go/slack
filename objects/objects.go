// Package objects provide the building blocks to creating the various
// objects used within the Slack API. It provides Builder objects
// to cleanly create objects to be consumed by API calls
package objects

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

func PlainText(txt string) *Text {
	// ignore errors, b/c we only check for the type, and we're darn sure
	// what the type is
	o := BuildText(PlainTextType, txt).MustBuild()
	return o
}

func MarkdownText(txt string) *Text {
	// ignore errors, b/c we only check for the type, and we're darn sure
	// what the type is
	o := BuildText(MarkdownTextType, txt).MustBuild()
	return o
}

func (b *TextBuilder) Validate() error {
	switch b.typ {
	case MarkdownTextType, PlainTextType:
		return nil
	}

	return fmt.Errorf(`text object must have type of either %s or %s: got %s`, MarkdownTextType, PlainTextType, b.typ)
}

func (b *ContextBlockBuilder) Validate() error {
	if len(b.elements) > 10 {
		return fmt.Errorf(`maximum number of elements in context block is 10: got %d`, len(b.elements))
	}

	for _, e := range b.elements {
		switch e.(type) {
		case *ImageElement, *Text:
		default:
			return fmt.Errorf(`elements in context block can only be image elements or text objects: got %T`, e)
		}
	}
	return nil
}

func (u MultiLevelJSONUnmarshaler) UnmarshalJSON(data []byte) error {
	// This is really really really really inefficient...

	for _, v := range u {
		if err := json.Unmarshal(data, v); err != nil {
			return errors.Wrapf(err, `failed to unmarshal over %T`, v)
		}
	}


	return nil
}