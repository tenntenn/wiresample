// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package weather

import (
	"github.com/tenntenn/wiresample/clock"
	"github.com/tenntenn/wiresample/weather/geo/geomock"
	"github.com/tenntenn/wiresample/weather/source/mocksrc"
)

// Injectors from inject_default.go:

func newDefaultReporter() (*Reporter, func(), error) {
	clockClock := clock.Default()
	geoMock := &geomock.GeoMock{}
	source := &mocksrc.Source{}
	reporter := &Reporter{
		Clock:  clockClock,
		Geo:    geoMock,
		Source: source,
	}
	return reporter, func() {
	}, nil
}