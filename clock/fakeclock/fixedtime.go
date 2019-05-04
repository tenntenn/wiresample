package fakeclock

import (
	"context"
	"time"

	"github.com/tenntenn/wiresample/clock"
)

// Fixed implments clock.Clock which Now() returns Time field value.
type Fixed struct {
	Time time.Time
}

// Now implments clock.Clock.
func (t *Fixed) Now(ctx context.Context) time.Time {
	return t.Time
}

// NewFixed returns new Fixed which Time field is set as time.Now().
func NewFixed() clock.Clock {
	return &Fixed{Time: time.Now()}
}
