package rtm

import "github.com/lestrrat/go-slack"

type EventType int

const (
	RTMConnectingEvent EventType = iota
	MaxEvent
)

type Client struct {
	client   *slack.Client
	eventsCh chan *Event
}

type Event struct {
	typ  EventType
	data interface{}
}

type HelloEvent struct{}
type PresenceChangeEvent struct {
	Presence string `json:"presence"`
	User     string `json:"user"`
}
type ReconnectURLEvent struct {
	URL string `json:"url"`
}
