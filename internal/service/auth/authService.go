package auth

import (
	"auth-service/internal/repository/tokenRepo"
	"auth-service/internal/repository/userRepo"
	"auth-service/internal/service/auth/models"
	"auth-service/pkg/hash"
	"auth-service/pkg/jwt"
	"context"
)

type Service interface {
	Register(ctx context.Context, request models.RegisterRequest) (models.RegisterResponse, error)
	Login(ctx context.Context, request models.LoginRequest) (models.LoginResponse, error)
	RefreshAccess(ctx context.Context, request models.RefreshAccessRequest) (models.RefreshAccessResponse, error)
	Logout(ctx context.Context, request models.LogoutRequest) (models.LogoutResponse, error)
	UserInformation(ctx context.Context, request models.UserInformationRequest) (models.UserInformationResponse, error)
}

type authService struct {
	userRepository  userRepo.UserRepo
	tokenRepository tokenRepo.TokenRepo

	jwt    *jwt.Manager
	hasher *hash.Hasher
}

func New(userRepository userRepo.UserRepo, tokenRepository tokenRepo.TokenRepo, jwt *jwt.Manager, hasher *hash.Hasher) Service {
	return &authService{
		userRepository:  userRepository,
		tokenRepository: tokenRepository,
		jwt:             jwt,
		hasher:          hasher,
	}
}
