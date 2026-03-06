package server

import (
	"aggregatorProject/cmd/config"
	"aggregatorProject/internal/logg"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server interface {
	GracefulShutdown(ctx context.Context)
	Run(http.Handler) error
}

type HTTPServer struct {
	S      *http.Server
	cfg    config.Config
	logger logg.Logger
	done   chan struct{}
}

func NewHTTPServer(
	ctx context.Context,
	config config.Config,
	logger logg.Logger) *HTTPServer {

	srv := &HTTPServer{
		S:      &http.Server{Addr: config.GetAddress()},
		cfg:    config,
		logger: logger,
		done:   make(chan struct{}),
	}

	srv.GracefulShutdown(ctx)
	return srv
}

func (s *HTTPServer) Run(handler http.Handler) error {
	s.S.Handler = handler

	if err := s.S.ListenAndServe(); err != http.ErrServerClosed {
		s.logger.RaiseFatal("http server has not started")
		return err
	}

	<-s.done
	fmt.Fprint(os.Stdout, "\nServer has been successfully stopped\n")
	return nil
}

func (s *HTTPServer) GracefulShutdown(ctx context.Context) {
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		<-ctx.Done()
		defer stop()
		defer close(s.done)

		shutdownCtx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		if err := s.S.Shutdown(shutdownCtx); err != nil {
			s.logger.RaiseFatal("http server has bad Shutdown")
		}
	}()
}
