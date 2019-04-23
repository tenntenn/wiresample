package fakeclock

import (
	"context"
	"time"

	"github.com/tenntenn/wiresample/clock"
)

// FixedTime returns a clock which provides always fixed time.
func FixedTime(t time.Time) clock.Clock {
	return clock.ClockFunc(func(ctx context.Context) time.Time {
		return t
	})
}
