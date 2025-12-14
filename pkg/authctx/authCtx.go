package authctx

import "context"

type key int

const (
	userIDKey key = iota
	emailKey
)

func WithUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, userIDKey, userID)
}

func UserID(ctx context.Context) (string, bool) {
	v, ok := ctx.Value(userIDKey).(string)
	return v, ok
}

func WithEmail(ctx context.Context, email string) context.Context {
	return context.WithValue(ctx, emailKey, email)
}

func Email(ctx context.Context) (string, bool) {
	v, ok := ctx.Value(emailKey).(string)
	return v, ok
}
