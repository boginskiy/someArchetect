package app

import (
	"context"
	"goAggreg/cmd/config"
	"goAggreg/cmd/server"
	"goAggreg/internal/handler"
	"goAggreg/internal/logger"
)

type App struct {
	config config.Config
	logger logger.Logger
	server server.Server
}

func NewApp(ctx context.Context) (*App, error) {
	tmp := &App{}

	err := tmp.InitAttrs(ctx)
	if err != nil {
		return nil, err
	}

	return tmp, nil
}

func (a *App) Run() error {
	// // Converter
	// converter :=

	// // Service
	// service :=

	// Handler
	handler := handler.NewHandle()

	return a.server.Run(handler)

}

func (a *App) InitAttrs(ctx context.Context) error {
	attrsFunc := []func(ctx context.Context) error{
		a.initConfig,
		a.initLogger,
		a.initServer,
	}

	for _, f := range attrsFunc {
		err := f(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *App) initConfig(ctx context.Context) error {
	cfg, err := config.NewCfg(ctx)
	if err != nil {
		return err
	}
	a.config = cfg
	return nil
}

func (a *App) initLogger(ctx context.Context) error {
	lgr, err := logger.NewLog(ctx)
	if err != nil {
		return err
	}
	a.logger = lgr
	return nil
}

func (a *App) initServer(ctx context.Context) error {
	srv, err := server.NewServerHTTP(ctx, a.config)
	if err != nil {
		return err
	}
	a.server = srv
	return nil
}
