package clock

import (
	"context"
	"time"
)

// Clock is an interface which provides current time.
type Clock interface {
	Now(ctx context.Context) time.Time
}

// ClockFunc implements Clock.
type ClockFunc func(ctx context.Context) time.Time

// Now implements Clock.Now.
func (f ClockFunc) Now(ctx context.Context) time.Time {
	return f(ctx)
}

// Default returns current time by using time.Now.
func Default() Clock {
	return ClockFunc(func(ctx context.Context) time.Time {
		return time.Now()
	})
}
