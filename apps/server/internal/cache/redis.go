package cache

import (
	"context"
	"fmt"

	"github.com/Rifqialba/simplem/apps/server/internal/config"
	"github.com/redis/go-redis/v9"
)

func NewRedis(cfg config.RedisConfig) (*redis.Client, error) {
	addr := fmt.Sprintf(
		"%s:%s",
		cfg.Host,
		cfg.Port,
	)

	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	ctx := context.Background()

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return client, nil
}