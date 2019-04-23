package geo

import "context"

// Geo provides location data.
type Geo interface {
	LatLng(ctx context.Context, addr string) (lat, lng float64, err error)
}
