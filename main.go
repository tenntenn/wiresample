package main

import (
	"net"
	"net/http"
	"os"

	"github.com/tenntenn/wiresample/handler"
)

//go:generate wire ./...

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := net.JoinHostPort("", port)
	http.ListenAndServe(addr, handler.New())
}
