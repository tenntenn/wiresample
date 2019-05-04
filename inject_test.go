//+build wireinject

package main

import (
	"context"

	"github.com/google/wire"
	"github.com/tenntenn/wiresample/clock/fakeclock"
	"github.com/tenntenn/wiresample/handler"
	"github.com/tenntenn/wiresample/weather/geo/geomock"
)

func setupTest(ctx context.Context, flags *cliFlags) (*handler.Handler, func(), error) {
	wire.Build(
		appSet,
		geomock.New,
		fakeclock.NewFixed,
	)
	return nil, nil, nil
}
