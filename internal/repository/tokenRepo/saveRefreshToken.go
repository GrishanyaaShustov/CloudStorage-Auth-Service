package tokenRepo

import (
	"context"
	"fmt"
)

func (r *repo) SaveRefreshToken(ctx context.Context, userID, refreshToken string) error {
	key := refreshKey(userID)

	if err := r.redisClient.Set(ctx, key, refreshToken, r.tokenConfig.RefreshTokenTTL).Err(); err != nil {
		return fmt.Errorf("%w, %v", ErrRedisFailed, err)
	}

	return nil
}
