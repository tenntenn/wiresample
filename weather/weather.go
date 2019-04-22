package wether

import (
	"context"

	"github.com/morikuni/failure"
)

// Celsius represents temperature by the celsius.
type Celsius float64

// Weather represents weather infomation.
type Weather struct {
	Temperature       Celsius
	PrecipProbability float64
	Summary           string
}

// Reporter provides weather information.
type Reporter struct {
	Clock 	clock.Clock
	Geo   	geocoding.Geocoding
	Source	source.Source
}

// Now returns current weather information of specified address.
func (r *Reporter) Now(ctx context.Context, address string) (*Weather, error) {
	lat, lng, err := geo.AddrToLatLng(ctx, address)
	if err != nil {
		return nil, failure.Wrap(err)
	}

	t := clock.Now()
}
