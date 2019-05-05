package handler

import (
	"net/http"

	"github.com/tenntenn/wiresample/weather"
)

// Handler provides handlers for the service.
type Handler struct {
	router   http.ServeMux
	Reporter *weather.Reporter
}

var _ http.Handler = (*Handler)(nil)

// New creates a handler which implments http.Handler.
func New() *Handler {
	h, _, err := newDefaultHandler()
	if err != nil {
		// error must not be occurred
		panic(err)
	}
	h.init()
	return h
}

// ServeHTTP implments http.Handler.
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}

func (h *Handler) init() {
	h.router.HandleFunc("/", h.Weather)
}
