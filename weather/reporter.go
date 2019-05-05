package weather

import (
	"context"

	"github.com/morikuni/failure"
	"github.com/tenntenn/wiresample/clock"
	"github.com/tenntenn/wiresample/weather/geo"
	"github.com/tenntenn/wiresample/weather/source"
)

// Reporter provides weather information.
type Reporter struct {
	Clock  clock.Clock
	Geo    geo.Geocoding
	Source source.Source
}

// Now returns current weather information of specified address.
func (r *Reporter) Now(ctx context.Context, address string) (*source.Weather, error) {
	lat, lng, err := r.Geo.LatLng(ctx, address)
	if err != nil {
		return nil, failure.Wrap(err)
	}

	t := r.Clock.Now(ctx)
	w, err := r.Source.ByLatLng(ctx, t, lat, lng)
	if err != nil {
		return nil, failure.Wrap(err)
	}

	return w, nil
}
