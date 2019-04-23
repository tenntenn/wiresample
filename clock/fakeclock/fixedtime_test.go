package fakeclock_test

import (
	"context"
	"testing"
	"time"

	"github.com/tenntenn/wiresample/clock/fakeclock"
)

const layout = "2006/01/02 15:04"

func TestFixedTime(t *testing.T) {
	tests := []struct {
		t   time.Time
		ctx context.Context
	}{
		{T(t, "2019/04/23 12:00"), context.Background()},
	}

	for _, tt := range tests {
		tt := tt
		name := tt.t.Format(layout)
		t.Run(name, func(t *testing.T) {
			now := fakeclock.FixedTime(tt.t).Now(tt.ctx)
			if got := now.Format(layout); now != tt.t {
				t.Errorf("expect %s but got %s", got, name)
			}
		})
	}
}

func T(t *testing.T, s string) time.Time {
	t.Helper()
	tm, err := time.Parse(layout, s)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	return tm
}
