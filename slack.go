//go:generate go run internal/cmd/genevents/genevents.go
//go:generate go run internal/cmd/genscopes/genscopes.go
//go:generate go run internal/cmd/genmethods/genmethods.go

// Package slack implements a REST client for Slack services.
package slack

import (
	"encoding/json"
	"log"

	"github.com/lestrrat/go-slack/objects"
	"github.com/pkg/errors"
)

type starredItem struct {
	Type string `json:"type"`
}

func (c *StarredItemList) UnmarshalJSON(data []byte) error {
	log.Printf("data -> %s", data)
	var list []json.RawMessage
	if err := json.Unmarshal(data, &list); err != nil {
		return errors.Wrap(err, `failed to decode starred items`)
	}

	*c = make([]StarredItem, len(list))

	for i, itemData := range list {
		var item starredItem
		if err := json.Unmarshal(itemData, &item); err != nil {
			return errors.Wrap(err, `failed to detect item type`)
		}

		var v interface{}
		switch item.Type {
		case "message":
			v = &struct {
				Channel string           `json:"channel"`
				Message *objects.Message `json:"message"`
			}{}
		case "file":
			fallthrough
		case "file_comment":
			fallthrough
		case "channel":
			fallthrough
		case "im":
			fallthrough
		case "group":
			fallthrough
		default:
			return errors.Errorf("unimplemented item type: %s", item.Type)
		}

		if err := json.Unmarshal(itemData, v); err != nil {
			return errors.Wrap(err, `failed to decode item`)
		}

		(*c)[i] = v
	}
	return nil
}
