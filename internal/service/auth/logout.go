package auth

import (
	"auth-service/internal/service/auth/models"
	"auth-service/pkg/authctx"
	"context"
	"fmt"
)

func (svc *authService) Logout(ctx context.Context, request models.LogoutRequest) (models.LogoutResponse, error) {
	// get userID from context (set by auth middleware)
	userID, isContainsUserID := authctx.UserID(ctx)
	if !isContainsUserID {
		return models.LogoutResponse{}, ErrInvalidCredentials
	}

	// delete refresh token in redis
	delRefreshErr := svc.tokenRepository.DeleteRefreshToken(ctx, userID)
	if delRefreshErr != nil {
		return models.LogoutResponse{}, fmt.Errorf("%w: delete refresh token: %v", ErrInternal, delRefreshErr)
	}

	return models.LogoutResponse{}, nil
}
