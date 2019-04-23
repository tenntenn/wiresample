package geomock

import (
	"context"

	"github.com/tenntenn/wiresample/weather/geo"
)

// GeoMock is a mock for geo.Geo.
type GeoMock struct {
	LatLng func(ctx context.Context, addr string) (lat, lng float64, err error)
}

var _ geo.Geo = (*GeoMock)(nil)

// LatLng implements (geo.Geo).LatLng.
func (g *GeoMock) LatLng(ctx context.Context, addr string) (lat, lng float64, err error) {
	if g.LatLng != nil {
		return g.LatLng(ctx, addr)
	}
	return
}
