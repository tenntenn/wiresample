package weather

import (
	"context"

	"github.com/morikuni/failure"
	"github.com/tenntenn/wiresample/clock"
)

// Reporter provides weather information.
type Reporter struct {
	c clock.Clock
	g geo.Geocoding
	s source.Source
}

// NewReporter creates a Reporter.
func NewReporter(c clock.Clock, g geo.Geocoding, s source.Source) *Reporter {
	return &Reporter{c: c, g: g, s: s}
}

// Now returns current weather information of specified address.
func (r *Reporter) Now(ctx context.Context, address string) (*Weather, error) {
	lat, lng, err := r.g.LatLng(ctx, address)
	if err != nil {
		return nil, failure.Wrap(err)
	}

	t := r.c.Now(ctx)
	w, err := r.s.ByLatLng(ctx, t, lat, lng)
	if err != nil {
		return nil, failue.Wrap(err)
	}

	return w, nil
}
