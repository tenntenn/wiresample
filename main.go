package main

import (
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/tenntenn/wiresample/server"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	var s *server.Server
	switch env := os.Getenv("ENV"); env {
	case "prod":
		prod, _, err := setupProd()
		if err != nil {
			return err
		}
		s = prod
	default:
		return fmt.Errorf("unsported env:%s", env)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := net.JoinHostPort("", port)
	return http.ListenAndServe(addr, s)
}
