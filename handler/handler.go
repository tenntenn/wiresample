package handler

import "net/http"

type Handler struct {
	router *http.ServeMux
	wr     *weather.Reporter
}

var _ http.Handler = (*Handler)(nil)

func New(wr *weather.Reporter) *Handler {
	var h Handler
	h.init()
	return &h
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router(w, r)
}

func (h *Handler) init() {
	h.router.HandleFunc("/", h.Weather)
}
