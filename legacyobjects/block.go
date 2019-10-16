package objects

import (
	"bytes"
	"encoding/json"

	"github.com/pkg/errors"
)

type blockType string

const (
	blockTypeSection blockType = "section"
	blockTypeAction  blockType = "actions"
	blockTypeDivider blockType = "divider"
	blockTypeFile    blockType = "file"
	blockTypeImage   blockType = "image"
	blockTypeContext blockType = "context"
)

func (SectionBlock) blockType() blockType { return blockTypeSection }

func (b *SectionBlock) Encode() (string, error) {
	block := struct {
		*SectionBlock
		Type blockType `json:"type"`
	}{
		SectionBlock: &SectionBlock{
			Text:      b.Text,
			BlockID:   b.BlockID,
			Fields:    b.Fields,
			Accessory: b.Accessory,
		},
		Type: b.blockType(),
	}
	buf, err := json.Marshal(&block)
	if err != nil {
		return "", errors.Wrap(err, `failed to encode section block`)
	}
	return string(buf), nil
}

func (b *SectionBlock) Decode(buf string) error {
	if buf == "" {
		return nil
	}
	r := bytes.NewReader([]byte(buf))
	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()

	return dec.Decode(b)
}

func (ActionsBlock) blockType() blockType { return blockTypeAction }

func (b *ActionsBlock) Encode() (string, error) {
	block := struct {
		*ActionsBlock
		Type blockType `json:"type"`
	}{
		ActionsBlock: &ActionsBlock{
			BlockID:  b.BlockID,
			Elements: b.Elements,
		},
		Type: b.blockType(),
	}
	buf, err := json.Marshal(&block)
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

func (DividerBlock) blockType() blockType { return blockTypeDivider }

func (b *DividerBlock) Encode() (string, error) {
	block := struct {
		*DividerBlock
		Type blockType `json:"type"`
	}{
		DividerBlock: &DividerBlock{
			BlockID: b.BlockID,
		},
		Type: b.blockType(),
	}
	buf, err := json.Marshal(&block)
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

func (FileBlock) blockType() blockType { return blockTypeFile }

func (b *FileBlock) Encode() (string, error) {
	block := struct {
		*FileBlock
		Type blockType `json:"type"`
	}{
		FileBlock: &FileBlock{
			BlockID:    b.BlockID,
			ExternalID: b.ExternalID,
			Source:     b.Source,
		},
		Type: b.blockType(),
	}
	buf, err := json.Marshal(&block)
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

func (ImageBlock) blockType() blockType { return blockTypeImage }

func (b *ImageBlock) Encode() (string, error) {
	block := struct {
		*ImageBlock
		Type blockType `json:"type"`
	}{
		ImageBlock: &ImageBlock{
			AltText:  b.AltText,
			BlockID:  b.BlockID,
			ImageURL: b.ImageURL,
			Title:    b.Title,
		},
		Type: b.blockType(),
	}
	buf, err := json.Marshal(&block)
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

func (ContextBlock) blockType() blockType { return blockTypeContext }

func (b *ContextBlock) Encode() (string, error) {
	block := struct {
		*ContextBlock
		Type blockType `json:"type"`
	}{
		ContextBlock: &ContextBlock{
			BlockID:  b.BlockID,
			Elements: b.Elements,
		},
		Type: b.blockType(),
	}
	buf, err := json.Marshal(&block)
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
