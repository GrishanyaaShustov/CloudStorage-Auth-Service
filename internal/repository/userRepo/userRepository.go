package userRepo

import (
	"auth-service/internal/domain"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepo interface {
	Create(ctx context.Context, email, passwordHash string) (string, error)
	GetCredentialsByEmail(ctx context.Context, email string) (userID string, passwordHash string, err error)
	GetUserByID(ctx context.Context, userID string) (domain.User, error)
}

type repo struct {
	pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) UserRepo {
	return &repo{pool: pool}
}
