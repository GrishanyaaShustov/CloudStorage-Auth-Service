package userRepo

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (r *repo) GetPasswordHashByEmail(ctx context.Context, email string) (string, error) {
	var passwordHash string

	err := r.pool.QueryRow(ctx, queryGetPasswordHashByEmail, email).Scan(&passwordHash)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", ErrUserNotFound
		}
		return "", fmt.Errorf("%w, %v", ErrUserQueryFailed, err)
	}

	return passwordHash, nil
}
