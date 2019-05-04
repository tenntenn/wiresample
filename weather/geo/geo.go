package geo

import "context"

// Geocoding provides location data.
type Geocoding interface {
	LatLng(ctx context.Context, addr string) (lat, lng float64, err error)
}
