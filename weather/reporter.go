package wether

import "context"

// Reporter
type Reporter struct {
	Clock clock.Clock
	Geo   geocoding.Geocoding
	Src   source.Source
}

func (r *Reporter) Now(ctx context.Context, address string) (*Wether, error) {
	lat, lng, err := geo.AddrToLatlng(ctx, address)
	if err != nil {
		return failure.Wrap(err)
	}

	t := clock.Now()
}
