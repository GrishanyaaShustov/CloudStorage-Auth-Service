package auth

import (
	"auth-service/internal/repository/userRepo"
	"auth-service/internal/service/auth/models"
	"context"
	"errors"
	"fmt"
)

func (svc *authService) UserInformation(ctx context.Context, request models.UserInformationRequest) (models.UserInformationResponse, error) {
	// get userID from context (set by auth middleware)
	userID, isContainsUserID := ctx.Value(ctxUserIDKey).(string)
	if !isContainsUserID {
		return models.UserInformationResponse{}, ErrInvalidCredentials
	}

	// get user from repository
	user, getUserErr := svc.userRepository.GetUserByID(ctx, userID)
	if getUserErr != nil {
		if errors.Is(getUserErr, userRepo.ErrUserNotFound) {
			return models.UserInformationResponse{}, ErrUserNotFound
		}
		return models.UserInformationResponse{}, fmt.Errorf("%w: get user by id: %v", ErrInternal, getUserErr)
	}

	// build response
	return models.UserInformationResponse{
		UserID:       user.UserID,
		Email:        user.Email,
		RegisterDate: user.RegisterDate,
	}, nil
}
