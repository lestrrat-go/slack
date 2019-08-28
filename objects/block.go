package objects

import (
	"encoding/json"

	"github.com/pkg/errors"
)

type blockType string

const (
	blockTypeAction  blockType = "actions"
	blockTypeContext blockType = "context"
	blockTypeDivider blockType = "divider"
	blockTypeFile    blockType = "file"
	blockTypeImage   blockType = "image"
	blockTypeSection blockType = "section"
)

type ActionsBlock struct {
	BlockID  string        `json:"block_id,omitempty"`
	Elements []interface{} `json:"elements,omitempty"`
	Type     blockType     `json:"type,omitempty"`
}

func (b *ActionsBlock) Encode() (string, error) {
	if b.Type == "" {
		b.Type = b.blockType()
	}
	buf, err := json.Marshal(b)
	if err != nil {
		return "", errors.Wrap(err, `failed to encode actions block`)
	}
	return string(buf), nil
}

func (b *ActionsBlock) Decode(buf string) error {
	if buf == "" {
		return nil
	}
	return json.Unmarshal([]byte(buf), b)
}

type ContextBlock struct {
	BlockID  string        `json:"block_id,omitempty"`
	Elements []interface{} `json:"elements,omitempty"`
	Type     blockType     `json:"type,omitempty"`
}

func (b *ContextBlock) Encode() (string, error) {
	if b.Type == "" {
		b.Type = b.blockType()
	}
	buf, err := json.Marshal(b)
	if err != nil {
		return "", errors.Wrap(err, `failed to encode context block`)
	}
	return string(buf), nil
}

func (b *ContextBlock) Decode(buf string) error {
	if buf == "" {
		return nil
	}
	return json.Unmarshal([]byte(buf), b)
}

type DividerBlock struct {
	BlockID string    `json:"block_id,omitempty"`
	Type    blockType `json:"type,omitempty"`
}

func (b *DividerBlock) Encode() (string, error) {
	if b.Type == "" {
		b.Type = b.blockType()
	}
	buf, err := json.Marshal(b)
	if err != nil {
		return "", errors.Wrap(err, `failed to encode divider block`)
	}
	return string(buf), nil
}

func (b *DividerBlock) Decode(buf string) error {
	if buf == "" {
		return nil
	}
	return json.Unmarshal([]byte(buf), b)
}

type FileBlock struct {
	BlockID    string    `json:"block_id,omitempty"`
	ExternalID string    `json:"external_id,omitempty"`
	Source     string    `json:"source,omitempty"`
	Type       blockType `json:"type,omitempty"`
}

func (b *FileBlock) Encode() (string, error) {
	if b.Type == "" {
		b.Type = b.blockType()
	}
	buf, err := json.Marshal(b)
	if err != nil {
		return "", errors.Wrap(err, `failed to encode file block`)
	}
	return string(buf), nil
}

func (b *FileBlock) Decode(buf string) error {
	if buf == "" {
		return nil
	}
	return json.Unmarshal([]byte(buf), b)
}

type ImageBlock struct {
	AltText  string `json:"alt_text,omitempty"`
	BlockID  string `json:"block_id,omitempty"`
	ImageUrl string `json:"image_url,omitempty"`
	Title    struct {
		Emoji bool   `json:"emoji,omitempty"`
		Text  string `json:"text,omitempty"`
		Type  string `json:"type,omitempty"`
	} `json:"title"`
	Type blockType `json:"type,omitempty"`
}

func (b *ImageBlock) Encode() (string, error) {
	if b.Type == "" {
		b.Type = b.blockType()
	}
	buf, err := json.Marshal(b)
	if err != nil {
		return "", errors.Wrap(err, `failed to encode image block`)
	}
	return string(buf), nil
}

func (b *ImageBlock) Decode(buf string) error {
	if buf == "" {
		return nil
	}
	return json.Unmarshal([]byte(buf), b)
}

type SectionBlock struct {
	Accessory interface{}   `json:"accessory,omitempty"`
	BlockID   string        `json:"block_id,omitempty"`
	Fields    []interface{} `json:"fields,omitempty"`
	Text      interface{}   `json:"text,omitempty"`
	Type      blockType     `json:"type,omitempty"`
}

func (b *SectionBlock) Encode() (string, error) {
	if b.Type == "" {
		b.Type = b.blockType()
	}
	buf, err := json.Marshal(b)
	if err != nil {
		return "", errors.Wrap(err, `failed to encode section block`)
	}
	return string(buf), nil
}

func (b *SectionBlock) Decode(buf string) error {
	if buf == "" {
		return nil
	}
	return json.Unmarshal([]byte(buf), b)
}

func (ActionsBlock) blockType() blockType { return blockTypeAction }
func (ContextBlock) blockType() blockType { return blockTypeContext }
func (DividerBlock) blockType() blockType { return blockTypeDivider }
func (FileBlock) blockType() blockType    { return blockTypeFile }
func (ImageBlock) blockType() blockType   { return blockTypeImage }
func (SectionBlock) blockType() blockType { return blockTypeSection }
