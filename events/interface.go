package events

import "encoding/json"

// Evenlope is the "outer event" as described in
// https://api.slack.com/events-api#callback_field_overview
type Envelope struct {
	APIAppID    string   `json:"api_app_id"`
	AuthedUsers []string `json:"authed_users"`
	Event       Event    `json:"event"`
	EventID     string   `json:"event_id"`
	EventTime   int      `json:"event_time"`
	TeamID      string   `json:"team_id"`
	Token       string   `json:"token"`
	Type        string   `json:"type"`
}

type Event struct {
	EventTimestamp string
	Item           interface{}
	Timestamp      string
	Type           string
	User           string
}

type eventUnmarshalProxy struct {
	EventTimestamp string
	Item           json.RawMessage
	Timestamp      string
	Type           string
	User           string
}
