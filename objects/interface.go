package objects

const (
	ButtonActionType = "button"
)
type EpochTime int64
type Timestamp float64

// ErrorResponse wraps errors returned by Slack. It's usually a string,
// but it could be a structure.
// https://api.slack.com/rtm#handling_responses
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

type DialogResponse struct {
	ResponseMetadata struct {
		Messages []string `json:"messages"`
	} `json:"response_metadata"`
}

type EmojiListResponse map[string]string

// UsergroupUsersList is the list of users in a given Usergroup.
type UsergroupUsersList []string

type Presence string

const (
	Presencective Presence = "away"
	PresenceAway  Presence = "away"
)


