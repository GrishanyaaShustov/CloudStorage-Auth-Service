package auth

import (
	"auth-service/internal/repository/userRepo"
	"auth-service/internal/service/auth/models"
	"context"
	"errors"
	"fmt"
)

func (svc *authService) Register(ctx context.Context, request models.RegisterRequest) (models.RegisterResponse, error) {
	// hash password
	passwordHash, hashPasswordErr := svc.hasher.HashPassword(request.Password)
	if hashPasswordErr != nil {
		return models.RegisterResponse{}, fmt.Errorf("%w: hash password: %v", ErrInternal, hashPasswordErr)
	}

	// create user
	userID, createUserErr := svc.userRepository.Create(ctx, request.Email, passwordHash)
	if createUserErr != nil {
		if errors.Is(createUserErr, userRepo.ErrEmailAlreadyExists) {
			return models.RegisterResponse{}, ErrEmailAlreadyExists
		}
		return models.RegisterResponse{}, fmt.Errorf("%w: create user: %v", ErrInternal, createUserErr)
	}

	// generate access token
	accessToken, generateAccessErr := svc.jwt.GenerateAccessToken(userID, request.Email)
	if generateAccessErr != nil {
		return models.RegisterResponse{}, fmt.Errorf("%w: generate access token: %v", ErrInternal, generateAccessErr)
	}

	// generate refresh token
	refreshToken, generateRefreshErr := svc.jwt.GenerateRefreshToken()
	if generateRefreshErr != nil {
		return models.RegisterResponse{}, fmt.Errorf("%w: generate refresh token: %v", ErrInternal, generateRefreshErr)
	}

	// save refresh token
	saveRefreshErr := svc.tokenRepository.SaveRefreshToken(ctx, userID, refreshToken)
	if saveRefreshErr != nil {
		return models.RegisterResponse{}, fmt.Errorf("%w: save refresh token: %v", ErrInternal, saveRefreshErr)
	}

	// build and return response
	return models.RegisterResponse{
		UserID:       userID,
		Email:        request.Email,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
