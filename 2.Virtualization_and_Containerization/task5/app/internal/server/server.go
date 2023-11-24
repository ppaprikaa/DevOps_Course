package server

import (
	"context"
	"net/http"
)

type server struct {
	*http.Server
}

func New(addr string, handler http.Handler) *server {
	srv := http.Server{
		Addr:    addr,
		Handler: handler,
	}

	return &server{
		Server: &srv,
	}
}

func (s *server) Run() error {
	return s.ListenAndServe()
}

func (s *server) Shutdown(ctx context.Context) error {
	return s.Server.Shutdown(ctx)
}
