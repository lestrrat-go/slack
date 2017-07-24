package rtm

import (
	"time"

	"github.com/cenkalti/backoff"
)

// WithPingInterval creates a new option that defines the length of
// intervals between each successive pings.
func WithPingInterval(t time.Duration) Option {
	return &option{
		name:  pingIntervalKey,
		value: t,
	}
}

// WithBackOffStrategy creates a new option that defines the
// backoff strategy to be used when connections to the slack
// RTM server fails.
func WithBackOffStrategy(b backoff.BackOff) Option {
	return &option{
		name:  backoffStrategyKey,
		value: b,
	}
}
