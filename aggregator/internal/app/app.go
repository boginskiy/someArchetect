package app

import (
	"aggregator/cmd/config"
	"aggregator/cmd/server"
	"aggregator/internal/converter"
	"aggregator/internal/db"
	"aggregator/internal/handlers"
	"aggregator/internal/logg"
	"aggregator/internal/model"
	"aggregator/internal/repository"
	"aggregator/internal/response"
	"aggregator/internal/service"
	"aggregator/pkg/router"
	"context"
)

type App struct {
	Server server.Server
	ctx    context.Context
	cfg    config.Config
	logger logg.Logger
}

func NewApp(ctx context.Context) (*App, error) {
	app := &App{ctx: ctx}

	err := app.InitDeps(ctx)
	if err != nil {
		return nil, err
	}
	return app, nil
}

func (a *App) Run() error {
	// DB
	db := db.NewDB(a.ctx, a.cfg, a.logger)
	defer db.Close()

	// Repository
	userRepo := repository.NewUserRepo(a.cfg, a.logger)
	eventRepo := repository.NewEventRepo(a.ctx, a.cfg, a.logger, db)

	// Converter
	userConverter := converter.NewUserConvert()
	eventConverter := converter.NewEventConvert()

	// Channels
	eventCh := make(chan *model.Event)

	// Service
	userService := service.NewUserServi(a.ctx, a.cfg, a.logger, userRepo)
	_ = service.NewEventServi(a.ctx, a.cfg, a.logger, eventRepo, eventCh, eventConverter)

	// Response
	response := response.NewResp()

	// Handlers
	userHandlers := handlers.NewUserHandle(userService, userConverter, response)
	eventHandlers := handlers.NewEventHandle(a.ctx, eventCh, eventConverter)

	// Router
	router := a.initRoutes(router.NewRoute(), userHandlers, eventHandlers)

	// Server
	return a.Server.Run(router)
}

func (a *App) InitDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initLogger,
		a.initServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *App) initServer(ctx context.Context) error {
	a.Server = server.NewHTTPServer(ctx, a.cfg, a.logger)
	return nil
}

func (a *App) initConfig(_ context.Context) error {
	envConf, err := config.NewEnvConf("")
	if err != nil {
		return err
	}
	a.cfg = config.NewConf(envConf)
	return nil
}

func (a *App) initLogger(_ context.Context) error {
	a.logger = logg.NewLogg()
	return nil
}

func (a *App) initRoutes(
	r router.Router,
	userH handlers.UserHandler,
	eventH handlers.EventHandler) router.Router {

	r.Handle("GET", "/user", userH.Read)
	r.Handle("POST", "/user", userH.Create)
	r.Handle("POST", "/event", eventH.Listen)

	return r
}
