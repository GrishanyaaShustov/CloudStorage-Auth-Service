package redis

import (
	"auth-service/internal/config"
	"context"

	"github.com/redis/go-redis/v9"
)

func New(ctx context.Context, cfg config.RedisConfig) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	ctx, cancel := context.WithTimeout(ctx, cfg.Timeout)
	defer cancel()

	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return rdb, nil
}

func MustNew(ctx context.Context, cfg config.RedisConfig) *redis.Client {
	rdb, err := New(ctx, cfg)
	if err != nil {
		panic("failed to connect to redis: " + err.Error())
	}
	return rdb
}
