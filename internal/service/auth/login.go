package auth

import (
	"auth-service/internal/repository/userRepo"
	"auth-service/internal/service/auth/models"
	"context"
	"errors"
	"fmt"
)

func (svc *authService) Login(ctx context.Context, request models.LoginRequest) (models.LoginResponse, error) {
	// get password hash by email
	userID, hash, getCredentialsByEmailErr := svc.userRepository.GetCredentialsByEmail(ctx, request.Email)
	if getCredentialsByEmailErr != nil {
		if errors.Is(getCredentialsByEmailErr, userRepo.ErrUserNotFound) {
			return models.LoginResponse{}, ErrInvalidCredentials
		}
		return models.LoginResponse{}, fmt.Errorf("%w: get password hash: %v", ErrInternal, getCredentialsByEmailErr)
	}

	// check password hash
	if !svc.hasher.CheckPassword(request.Password, hash) {
		return models.LoginResponse{}, ErrInvalidCredentials
	}

	// generate access token
	accessToken, generateAccessErr := svc.jwt.GenerateAccessToken(userID, request.Email)
	if generateAccessErr != nil {
		return models.LoginResponse{}, fmt.Errorf("%w: generate access token: %v", ErrInternal, generateAccessErr)
	}

	// generate refresh token
	refreshToken, generateRefreshErr := svc.jwt.GenerateRefreshToken()
	if generateRefreshErr != nil {
		return models.LoginResponse{}, fmt.Errorf("%w: generate refresh token: %v", ErrInternal, generateRefreshErr)
	}

	// save refresh token
	saveRefreshErr := svc.tokenRepository.SaveRefreshToken(ctx, userID, refreshToken)
	if saveRefreshErr != nil {
		return models.LoginResponse{}, fmt.Errorf("%w: save refresh token: %v", ErrInternal, saveRefreshErr)
	}

	// build and return response
	return models.LoginResponse{
		UserID:       userID,
		Email:        request.Email,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
