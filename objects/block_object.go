package objects

import (
	"bytes"
	"encoding/json"

	"github.com/pkg/errors"
)

type textBlockObjectType string

const (
	textBlockMarkdown  textBlockObjectType = "mrkdwn"
	textBlockPlainText textBlockObjectType = "plain_text"
)

type blockObjectType string

const (
	objectsConfirmation       blockObjectType = "confirm"
	objectsOptionObjects      blockObjectType = "option"
	objectsOptionGroupObjects blockObjectType = "option_group"
)

func (b *PlainTextBlock) Encode() (string, error) {
	block := struct {
		*PlainTextBlock
		Type textBlockObjectType `json:"type"`
	}{
		PlainTextBlock: &PlainTextBlock{
			Text:  b.Text,
			Emoji: b.Emoji,
		},
		Type: b.textBlockObjectType(),
	}
	buf, err := json.Marshal(&block)
	if err != nil {
		return "", errors.Wrap(err, `failed to encode plain text block`)
	}
	return string(buf), nil
}

func (b *PlainTextBlock) Decode(buf string) error {
	if buf == "" {
		return nil
	}
	r := bytes.NewReader([]byte(buf))
	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()

	return dec.Decode(b)
}

func (b *MarkdownTextBlock) Encode() (string, error) {
	block := struct {
		*MarkdownTextBlock
		Type textBlockObjectType `json:"type"`
	}{
		MarkdownTextBlock: &MarkdownTextBlock{
			Text:     b.Text,
			Verbatim: b.Verbatim,
		},
		Type: b.textBlockObjectType(),
	}
	buf, err := json.Marshal(&block)
	if err != nil {
		return "", errors.Wrap(err, `failed to encode markdown text block`)
	}
	return string(buf), nil
}

func (b *MarkdownTextBlock) Decode(buf string) error {
	if buf == "" {
		return nil
	}
	r := bytes.NewReader([]byte(buf))
	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()

	return dec.Decode(b)
}

func (b *ConfirmationDialogBlockObject) Encode() (string, error) {
	block := struct {
		*ConfirmationDialogBlockObject
		Type blockObjectType `json:"type"`
	}{
		ConfirmationDialogBlockObject: &ConfirmationDialogBlockObject{
			Title:   b.Title,
			Text:    b.Text,
			Confirm: b.Confirm,
			Deny:    b.Deny,
		},
		Type: b.blockObjectType(),
	}
	buf, err := json.Marshal(&block)
	if err != nil {
		return "", errors.Wrap(err, `failed to encode confirmation dialog block object`)
	}
	return string(buf), nil
}

func (b *ConfirmationDialogBlockObject) Decode(buf string) error {
	if buf == "" {
		return nil
	}
	r := bytes.NewReader([]byte(buf))
	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()

	return dec.Decode(b)
}

func (b *OptionBlockObject) Encode() (string, error) {
	block := struct {
		*OptionBlockObject
		Type blockObjectType `json:"type"`
	}{
		OptionBlockObject: &OptionBlockObject{
			Text:  b.Text,
			Value: b.Value,
			URL:   b.URL,
		},
		Type: b.blockObjectType(),
	}
	buf, err := json.Marshal(&block)
	if err != nil {
		return "", errors.Wrap(err, `failed to encode option block object`)
	}
	return string(buf), nil
}

func (b *OptionBlockObject) Decode(buf string) error {
	if buf == "" {
		return nil
	}
	r := bytes.NewReader([]byte(buf))
	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()

	return dec.Decode(b)
}

func (b *OptionGroupBlockObject) Encode() (string, error) {
	block := struct {
		*OptionGroupBlockObject
		Type blockObjectType `json:"type"`
	}{
		OptionGroupBlockObject: &OptionGroupBlockObject{
			Label:   b.Label,
			Options: b.Options,
		},
		Type: b.blockObjectType(),
	}
	buf, err := json.Marshal(&block)
	if err != nil {
		return "", errors.Wrap(err, `failed to encode option group block object`)
	}
	return string(buf), nil
}

func (b *OptionGroupBlockObject) Decode(buf string) error {
	if buf == "" {
		return nil
	}
	r := bytes.NewReader([]byte(buf))
	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()

	return dec.Decode(b)
}
