package userRepo

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgconn"
)

func (r *repo) Create(ctx context.Context, email, passwordHash string) (string, error) {
	var userID string

	err := r.pool.QueryRow(ctx, queryCreateUser, email, passwordHash).
		Scan(&userID)

	if err != nil {
		// PostgreSQL error
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return "", ErrEmailAlreadyExists
			}
		}
		return "", fmt.Errorf("%w: %v", ErrUserCreateFailed, err)
	}

	return userID, nil
}
