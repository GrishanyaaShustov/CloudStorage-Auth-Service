package tokenRepo

import (
	"auth-service/internal/config"
	"context"

	"github.com/redis/go-redis/v9"
)

type TokenRepo interface {
	SaveRefreshToken(ctx context.Context, userID, refreshToken string) error
	GetRefreshToken(ctx context.Context, userID string) (string, error)
	DeleteRefreshToken(ctx context.Context, userID string) error
}

type repo struct {
	tokenConfig config.TokensConfiguration
	redisClient *redis.Client
}

func New(tokenConfig config.TokensConfiguration, redisClient *redis.Client) TokenRepo {
	return &repo{tokenConfig: tokenConfig, redisClient: redisClient}
}

func refreshKey(userID string) string {
	return "refresh:" + userID
}
