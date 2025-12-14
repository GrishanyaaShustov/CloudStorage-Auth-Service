package tokenRepo

import "errors"

var (
	ErrTokenNotFound = errors.New("refresh token not found")
	ErrRedisFailed   = errors.New("redis operation failed")
)
