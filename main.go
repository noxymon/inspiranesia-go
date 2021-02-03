package main

import (
	"context"
	"go.uber.org/fx"
	homeController "inspiranesia/internal/home/controller"
	homeService "inspiranesia/internal/home/service"
	"inspiranesia/system/config"
	"inspiranesia/system/http"
	"inspiranesia/system/logging"
)

func main() {
	app := fx.New(
		fx.Provide(logging.ProvideLogging),
		fx.Provide(config.ProvideViper),
		fx.Provide(http.ProvideHttpHandler),
		fx.Provide(http.ProvideHttpAdapter),
		fx.Invoke(homeController.Provide),
		fx.Invoke(homeService.Provide),
		fx.Invoke(start),
	)

	app.Run()
}

func start(lifecycle fx.Lifecycle, logger logging.NougatLoggingProvider, handler http.NougatHttpHandler) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go handler.Start()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return logger.Sync()
		},
	})
}
