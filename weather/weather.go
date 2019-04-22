package wether

import "context"

// Celsius represents temperature by the celsius.
type Celsius float64

type Wether struct {
	Temperature       Celsius
	PrecipProbability float64
	Summary           string
}

type Reporter struct {
	clock clock.Clock
	geo   geocoding.Geocoding
	src   source.Source
}

type ReporterConfig func(*Reporter)

func NewReporter(conf ...ReporterConifg) {
}

func (r *Reporter) Now(ctx context.Context, address string) (*Wether, error) {
	lat, lng, err := geo.AddrToLatlng(ctx, address)
	if err != nil {
		return failure.Wrap(err)
	}

	t := clock.Now()
}
