// +build wireinject

package handler

import (
	"github.com/google/wire"
	"github.com/tenntenn/wiresample/weather"
)

func newDefaultHandler() (*Handler, func(), error) {
	wire.Build(
		wire.Struct(new(Handler), "Reporter"),
		weather.NewReporter,
	)
	return nil, nil, nil
}
