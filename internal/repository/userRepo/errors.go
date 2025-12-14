package userRepo

import "errors"

var (
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrUserCreateFailed   = errors.New("failed to create user")

	ErrUserNotFound    = errors.New("user not found")
	ErrUserQueryFailed = errors.New("failed to query user")
)
