package rtm

import (
	"github.com/lestrrat/go-slack"
	"github.com/lestrrat/go-slack/objects"
)

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

type ChannelCreated struct {
	Created            objects.EpochTime `json:"created"`
	ID                 string            `json:"id"`
	IsIM               bool              `json:"is_im"`
	IsOpen             bool              `json:"is_open"`
	IsOrgShared        bool              `json:"is_org_shared"`
	LastRead           string            `json:"last_read"`
	Latest             *objects.Message  `json:"latest"`
	UnreadCount        int               `json:"unread_count"`
	UnreadCountDisplay int               `json:"unread_count_display"`
	User               string            `json:"user"`
}

type ChannelCreatedEvent struct {
	Channel        ChannelCreated `json:"channel"`
	EventTimestamp string         `json:"event_ts"`
}

type ChannelJoinedEvent struct {
	Channel        *objects.Channel `json:"channel"`
	EventTimestamp string         `json:"event_ts"`
}

// This event is... not in the events list... (as of Apr 3, 2017) https://api.slack.com/events
type DesktopNotificationEvent struct {
	AvatarImage    string `json:"avatar_image"`
	Channel        string `json:"channel"`
	Content        string `json:"content"`
	EventTimestamp string `json:"event_ts"`
	ImageURI       string `json:"image_uri"`
	IsShared       bool   `json:"is_shared"`
	LaunchURI      string `json:"launch_uri"`
	Msg            string `json:"msg"`
	Title          string `json:"title"`
	SsbFilename    string `json:"ssb_filename"`
	Subtitle       string `json:"subtitle"`
}

type HelloEvent struct{}

type ImCreatedEvent struct {
	User    string         `json:"user"`
	Channel ChannelCreated `json:"channel"`
}

type MessageEvent struct {
	Channel    string `json:"channel"`
	SourceTeam string `json:"source_team"`
	Team       string `json:"team"`
	Text       string `json:"text"`
	Timestamp  string `json:"ts"`
	User       string `json:"user"`
}

type PresenceChangeEvent struct {
	Presence string `json:"presence"`
	User     string `json:"user"`
}

type ReconnectURLEvent struct {
	URL string `json:"url"`
}

type UserTypingEvent struct {
	Channel string `json:"channel"`
	User    string `json:"user"`
}

// This event is... not in the events list... (as of Apr 3, 2017) https://api.slack.com/events
type MemberJoinedchannel struct {
	EventTimestamp string `json:"event_ts"`
	Channel        string `json:"channel"`
	ChannelType    string `json:"channel_type"`
	User           string `json:"user"`
}
