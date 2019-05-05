package mocksrc

import (
	"context"
	"time"

	"github.com/tenntenn/wiresample/weather/source"
)

// Source is a mock for source.Source.
type Source struct {
	ByLatLngFunc func(ctx context.Context, t time.Time, lat, lng float64) (*source.Weather, error)
}

var _ source.Source = (*Source)(nil)

// ByLatLng implments source.Source.
func (s *Source) ByLatLng(ctx context.Context, t time.Time, lat, lng float64) (*source.Weather, error) {
	if s.ByLatLngFunc != nil {
		return s.ByLatLngFunc(ctx, t, lat, lng)
	}
	return &source.Weather{
		Temperature:       20,
		PrecipProbability: 10,
		Summary:           "good weather",
	}, nil
}
