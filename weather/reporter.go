package wether

import (
	"context"

	"github.com/morikuni/failure"
)

// Reporter provides weather information.
type Reporter struct {
	Clock  clock.Clock
	Geo    geocoding.Geocoding
	Source source.Source
}

// Now returns current weather information of specified address.
func (r *Reporter) Now(ctx context.Context, address string) (*Weather, error) {
	lat, lng, err := r.Geo.LatLng(ctx, address)
	if err != nil {
		return nil, failure.Wrap(err)
	}

	t := r.Clock.Now(ctx)
	w, err := r.Source.ByLatLng(ctx, t, lat, lng)
	if err != nil {
		return nil, failue.Wrap(err)
	}

	return w, nil
}
