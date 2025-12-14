package userRepo

import (
	"auth-service/internal/domain"
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (r *repo) GetUserByID(ctx context.Context, userID string) (domain.User, error) {
	var user domain.User

	err := r.pool.QueryRow(ctx, queryGetUserByID).
		Scan(&user.UserID, &user.Email, &user.RegisterDate)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.User{}, ErrUserNotFound
		}
		return domain.User{}, fmt.Errorf("%w, %v", ErrUserQueryFailed, err)
	}

	return user, nil
}
