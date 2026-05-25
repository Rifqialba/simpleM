package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Rifqialba/simplem/apps/server/internal/app"
	"github.com/Rifqialba/simplem/apps/server/internal/cache"
	"github.com/Rifqialba/simplem/apps/server/internal/config"
	"github.com/Rifqialba/simplem/apps/server/internal/database"
	"github.com/Rifqialba/simplem/apps/server/internal/logger"
	"github.com/Rifqialba/simplem/apps/server/internal/routes"
	"github.com/Rifqialba/simplem/apps/server/internal/server"
	"github.com/Rifqialba/simplem/apps/server/internal/realtime"
)

func main() {
	cfg := config.Load()

	logg := logger.New()

	db, err := database.NewPostgres(cfg.Database)
	if err != nil {
		logg.Fatal().
			Err(err).
			Msg("failed to connect postgres")
	}

	redisClient, err := cache.NewRedis(cfg.Redis)
	if err != nil {
		logg.Fatal().
			Err(err).
			Msg("failed to connect redis")
	}

	appContainer := &app.App{
		Config: cfg,
		Logger: logg,
		DB:     db,
		Redis:  redisClient,
		RealtimeManager: realtime.NewManager(),
	}

	appFiber := server.New(appContainer)

	routes.Register(
		appFiber,
		appContainer,
	)

	go func() {
		logg.Info().
			Str("port", cfg.App.Port).
			Msg("starting simpleM server")

		if err := appFiber.Listen(
			":" + cfg.App.Port,
		); err != nil {
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

	if err := appFiber.ShutdownWithContext(ctx); err != nil {
		logg.Error().
			Err(err).
			Msg("graceful shutdown failed")
	}

	db.Close()

	if err := redisClient.Close(); err != nil {
		logg.Error().
			Err(err).
			Msg("failed closing redis")
	}

	logg.Info().Msg("server stopped")
}