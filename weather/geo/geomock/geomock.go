package geomock

import (
	"context"

	"github.com/tenntenn/wiresample/weather/geo"
)

// GeoMock is a mock for geo.Geocoding.
type GeoMock struct {
	LatLngFunc func(ctx context.Context, addr string) (lat, lng float64, err error)
}

var _ geo.Geocoding = (*GeoMock)(nil)

// LatLng implements (geo.Geo).LatLng.
func (g *GeoMock) LatLng(ctx context.Context, addr string) (lat, lng float64, err error) {
	if g.LatLngFunc != nil {
		return g.LatLngFunc(ctx, addr)
	}
	return
}
