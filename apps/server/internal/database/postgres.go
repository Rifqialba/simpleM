package database

import (
	"context"
	"fmt"

	"github.com/Rifqialba/simplem/apps/server/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgres(
	cfg config.DatabaseConfig,
) (*pgxpool.Pool, error) {

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
		cfg.SSLMode,
	)

	dbpool, err := pgxpool.New(
		context.Background(),
		dsn,
	)

	if err != nil {
		return nil, err
	}

	if err := dbpool.Ping(
		context.Background(),
	); err != nil {

		return nil, err
	}

	return dbpool, nil
}