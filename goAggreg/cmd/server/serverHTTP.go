package server

import (
	"context"
	"goAggreg/cmd/config"
	"net/http"
)

type Server interface {
	Run() error
}

type ServerHTTP struct {
	config config.Config
	S      http.Server
}

func NewServerHTTP(ctx context.Context, cfg config.Config, handler http.Handler) (*ServerHTTP, error) {
	return &ServerHTTP{
		config: cfg,
		S: http.Server{
			Addr:    cfg.GetAddress(),
			Handler: handler},
	}, nil
}

func (s *ServerHTTP) Run() error {
	return s.S.ListenAndServe()
}
