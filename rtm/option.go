package rtm

import "time"

// WithPingInterval creates a new option that defines the length of
// intervals between each successive pings.
func WithPingInterval(t time.Duration) Option {
	return &option{
		name:  pingIntervalKey,
		value: t,
	}
}
