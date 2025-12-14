package auth

import (
	"auth-service/internal/service/auth/models"
	"context"
	"fmt"
)

func (svc *authService) Logout(ctx context.Context, request models.LogoutRequest) (models.LogoutResponse, error) {
	// get userID from context (set by auth middleware)
	userID, isContainsUserID := ctx.Value(ctxUserIDKey).(string)
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
