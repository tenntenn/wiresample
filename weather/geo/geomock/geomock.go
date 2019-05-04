package geomock

import (
	"context"

	"github.com/tenntenn/wiresample/weather/geo"
)

// GeoMock is a mock for geo.Geocoding.
type GeoMock struct {
	LatLngFunc func(ctx context.Context, addr string) (lat, lng float64, err error)
}

// New returns mocking geocoding.
func New() geo.Geocoding {
	return &GeoMock{LatLngFunc: tokyo}
}

// LatLng implements (geo.Geo).LatLng.
func (g *GeoMock) LatLng(ctx context.Context, addr string) (lat, lng float64, err error) {
	if g.LatLng != nil {
		return g.LatLng(ctx, addr)
	}
	return
}

func tokyo(ctx context.Context, addr string) (lat, lng float64, err error) {
	return 35.689634, 139.692101, nil
}
