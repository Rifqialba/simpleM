package app

import (
	"github.com/Rifqialba/simplem/apps/server/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
)

type App struct {
	Config *config.Config
	Logger zerolog.Logger
	DB     *pgxpool.Pool
	Redis  *redis.Client
}