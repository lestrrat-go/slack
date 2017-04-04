package rtm

import (
	"encoding/json"

	"github.com/pkg/errors"
)

func (e *Event) UnmarshalJSON(data []byte) error {
	// mental note: RTM events are not really related to their
	// counterparts in events API in terms of structure.

	// here, we must unpack the JSON bytes into a map[string]interface{}
	// check for the "type" field, and then re-parse
	var m map[string]interface{}
	if err := json.Unmarshal(data, &m); err != nil {
		return errors.Wrap(err, `failed to unmarshal data`)
	}

	var item interface{}
	var typ EventType
	switch m["type"] {
	case HelloTypeKey:
		typ = HelloType
		item = &HelloEvent{} // XXX we should just skip unmarshaling after this
	case PresenceChangeTypeKey:
		typ = PresenceChangeType
		item = &PresenceChangeEvent{}
	case ReconnectURLTypeKey:
		typ = ReconnectURLType
		item = &ReconnectURLEvent{}
	default:
		return errors.Errorf("unknown RTM event type: %s", m["type"])
	}

	if err := json.Unmarshal(data, &item); err != nil {
		return errors.Wrap(err, `failed to unmarshal data`)
	}

	*e = Event{typ: typ, data: item}
	return nil
}
