package auth

import (
	"auth-service/internal/repository/tokenRepo"
	"auth-service/internal/service/auth/models"
	"auth-service/pkg/authctx"
	"context"
	"errors"
	"fmt"
)

// RefreshAccess requires authentication for silent refresh
// that checks in middleware
func (svc *authService) RefreshAccess(ctx context.Context, request models.RefreshAccessRequest) (models.RefreshAccessResponse, error) {
	// get userID from context
	userID, isContainsUserID := authctx.UserID(ctx)
	if !isContainsUserID {
		return models.RefreshAccessResponse{}, ErrInvalidCredentials
	}

	// get email from context
	email, isContainsEmail := authctx.Email(ctx)
	if !isContainsEmail {
		return models.RefreshAccessResponse{}, ErrInvalidCredentials
	}

	// check is refresh token exist in redis
	_, getRefreshErr := svc.tokenRepository.GetRefreshToken(ctx, userID)
	if getRefreshErr != nil {
		if errors.Is(getRefreshErr, tokenRepo.ErrTokenNotFound) {
			return models.RefreshAccessResponse{}, ErrInvalidRefreshToken
		}
		return models.RefreshAccessResponse{}, fmt.Errorf("%w: %v", ErrInternal, getRefreshErr)
	}

	// generate new access token
	accessToken, genAccessErr := svc.jwt.GenerateAccessToken(userID, email)
	if genAccessErr != nil {
		return models.RefreshAccessResponse{}, fmt.Errorf("%w: generate access token: %v", ErrInternal, genAccessErr)
	}

	// build response and return
	return models.RefreshAccessResponse{AccessToken: accessToken}, nil
}
