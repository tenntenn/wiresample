package main

import (
	"github.com/google/wire"
	"github.com/tenntenn/wiresample/handler"
	"github.com/tenntenn/wiresample/weather"
)

var appSet = wire.NewSet(
	weather.NewReporter,
	handler.New,
)

func main() {
}
