package server

import (
	"context"
	"goAggreg/cmd/config"
	"net/http"
)

type Server interface {
	Run(h http.Handler) error
}

type ServerHTTP struct {
	config config.Config
	S      http.Server
}

func NewServerHTTP(ctx context.Context, cfg config.Config) (*ServerHTTP, error) {
	return &ServerHTTP{
		config: cfg,
		S:      http.Server{Addr: cfg.GetAddress()},
	}, nil
}

func (s *ServerHTTP) Run(h http.Handler) error {
	s.S.Handler = h
	return s.S.ListenAndServe()
}
