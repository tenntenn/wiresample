package location

import (
	"context"
)

type API struct {
}

func NewAPI() *API {
	return &API{}
}

func (api *API) Address(ctx context.Context, lat, lng float64) (string, error) {
}
