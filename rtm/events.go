package rtm

import (
	"encoding/json"

	"github.com/pkg/errors"
)

func (e *Event) Type() EventType {
	return e.typ
}

func (e *Event) Data() interface{} {
	return e.data
}

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
	case ChannelCreatedTypeKey:
		typ = ChannelCreatedType
		item = &ChannelCreatedEvent{}
	case ChannelJoinedTypeKey:
		typ = ChannelJoinedType
		item = &ChannelJoinedEvent{}
	case HelloTypeKey:
		typ = HelloType
		item = &HelloEvent{} // XXX we should just skip unmarshaling after this
	case ImCreatedTypeKey:
		typ = ImCreatedType
		item = &ImCreatedEvent{}
	case MessageTypeKey:
		typ = MessageType
		item = &MessageEvent{}
	case PresenceChangeTypeKey:
		typ = PresenceChangeType
		item = &PresenceChangeEvent{}
	case ReconnectURLTypeKey:
		typ = ReconnectURLType
		item = &ReconnectURLEvent{}
	case UserTypingTypeKey:
		typ = UserTypingType
		item = &UserTypingEvent{}
	case PongTypeKey:
		typ = PongType
		item = &PongEvent{}
	case ErrorTypeKey:
		typ = ErrorType
		item = &ErrorEvent{}
	default:
		return errors.Errorf("unknown RTM event type: %s", m["type"])
	}

	if err := json.Unmarshal(data, &item); err != nil {
		return errors.Wrap(err, `failed to unmarshal data`)
	}

	*e = Event{typ: typ, data: item}
	return nil
}
