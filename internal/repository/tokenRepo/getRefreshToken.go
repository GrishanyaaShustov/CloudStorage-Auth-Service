package tokenRepo

import (
	"context"
	"errors"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func (r *repo) GetRefreshToken(ctx context.Context, userID string) (string, error) {
	key := refreshKey(userID)

	token, err := r.redisClient.Get(ctx, key).Result()

	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", ErrTokenNotFound
		}
		return "", fmt.Errorf("%w, %v", ErrRedisFailed, err)
	}

	return token, nil
}
