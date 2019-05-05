// +build wireinject

package weather

import (
	"github.com/google/wire"
	"github.com/tenntenn/wiresample/clock"
	"github.com/tenntenn/wiresample/weather/geo"
	"github.com/tenntenn/wiresample/weather/geo/geomock"
	"github.com/tenntenn/wiresample/weather/source"
	"github.com/tenntenn/wiresample/weather/source/mocksrc"
)

func newDefaultReporter() (*Reporter, func(), error) {
	wire.Build(
		wire.Struct(new(Reporter), "Clock", "Geo", "Source"),
		clock.Default,
		wire.Struct(new(geomock.GeoMock)),
		wire.Struct(new(mocksrc.Source)),
		wire.Bind((*geo.Geocoding)(nil), (*geomock.GeoMock)(nil)),
		wire.Bind((*source.Source)(nil), (*mocksrc.Source)(nil)),
	)
	return nil, nil, nil
}
