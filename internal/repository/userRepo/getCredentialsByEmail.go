package userRepo

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (r *repo) GetCredentialsByEmail(ctx context.Context, email string) (string, string, error) {
	var userID string
	var passwordHash string

	err := r.pool.QueryRow(ctx, queryGetCredentialsByEmail, email).Scan(&userID, &passwordHash)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", "", ErrUserNotFound
		}
		return "", "", fmt.Errorf("%w: %v", ErrUserQueryFailed, err)
	}

	return userID, passwordHash, nil
}
