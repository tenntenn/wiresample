package server

import (
	"net/http"

	"github.com/tenntenn/wiresample/weather"
)

// Server
type Server struct {
	router http.ServeMux
	wr     *weather.Reporter
}

var _ http.Handler = (*Server)(nil)

// New creates a handler which implments http.Handler.
func New(wr *weather.Reporter) *Server {
	s := &Server{wr: wr}
	s.init()
	return s
}

// ServeHTTP implments http.Handler.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) init() {
	s.router.HandleFunc("/", s.HandleWeather)
}
