package source

import (
	"context"
	"time"
)

// Source represents a weather report source.
type Source interface {
	// ByLatLng fetch to source and get weather information by time, lat and lng.
	ByLatLng(ctx context.Context, t time.Time, lat, lng float64) (*Weather, error)
}
