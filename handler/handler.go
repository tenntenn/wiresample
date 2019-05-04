package handler

import "net/http"

// Handler provides handlers for the service.
type Handler struct {
	router *http.ServeMux
	wr     *weather.Reporter
}

var _ http.Handler = (*Handler)(nil)

// New creates a handler which implments http.Handler.
func New(wr *weather.Reporter) *Handler {
	var h Handler
	h.init()
	return &h
}

// ServeHTTP implments http.Handler.
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}

func (h *Handler) init() {
	h.router.HandleFunc("/", h.Weather)
}
