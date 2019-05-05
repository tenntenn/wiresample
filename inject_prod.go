//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/tenntenn/wiresample/clock"
	"github.com/tenntenn/wiresample/server"
	"github.com/tenntenn/wiresample/weather"
	"github.com/tenntenn/wiresample/weather/geo"
	"github.com/tenntenn/wiresample/weather/geo/geomock"
	"github.com/tenntenn/wiresample/weather/source"
	"github.com/tenntenn/wiresample/weather/source/mocksrc"
)

func setupProd() (*server.Server, func(), error) {
	wire.Build(
		server.New,
		wire.Struct(new(weather.Reporter), "Clock", "Geo", "Source"),
		clock.Default,
		wire.Struct(new(geomock.GeoMock)),
		wire.Struct(new(mocksrc.Source)),
		wire.Bind((*geo.Geocoding)(nil), (*geomock.GeoMock)(nil)),
		wire.Bind((*source.Source)(nil), (*mocksrc.Source)(nil)),
	)
	return nil, nil, nil
}
