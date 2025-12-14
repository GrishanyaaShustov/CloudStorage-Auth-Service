package tokenRepo

import (
	"context"
	"fmt"
)

func (r *repo) DeleteRefreshToken(ctx context.Context, userID string) error {
	key := refreshKey(userID)

	if err := r.redisClient.Del(ctx, key).Err(); err != nil {
		return fmt.Errorf("%w: %v", ErrRedisFailed, err)
	}

	return nil
}
