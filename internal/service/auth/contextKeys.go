package auth

type ctxKey string

const (
	ctxUserIDKey ctxKey = "userID"
	ctxEmailKey  ctxKey = "email"
)
