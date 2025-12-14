package auth

import "errors"

var (
	ErrEmailAlreadyExists  = errors.New("email already exists")
	ErrInvalidInput        = errors.New("invalid input")
	ErrInvalidCredentials  = errors.New("invalid credentials")
	ErrInternal            = errors.New("internal error")
	ErrInvalidRefreshToken = errors.New("invalid refresh token")
	ErrUserNotFound        = errors.New("user not found")
)
