package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Rifqialba/simplem/apps/server/internal/config"
	"github.com/Rifqialba/simplem/apps/server/internal/logger"
	"github.com/Rifqialba/simplem/apps/server/internal/routes"
	"github.com/Rifqialba/simplem/apps/server/internal/server"
)

func main() {
	cfg := config.Load()

	logg := logger.New()

	app := server.New()

	routes.Register(app)

	go func() {
		logg.Info().
			Str("port", cfg.AppPort).
			Msg("starting simpleM server")

		if err := app.Listen(":" + cfg.AppPort); err != nil {
			logg.Fatal().
				Err(err).
				Msg("failed to start server")
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(
		quit,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	<-quit

	logg.Info().Msg("shutting down server...")

	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)

	defer cancel()

	if err := app.ShutdownWithContext(ctx); err != nil {
		logg.Error().
			Err(err).
			Msg("graceful shutdown failed")
	}

	logg.Info().Msg("server stopped")
}