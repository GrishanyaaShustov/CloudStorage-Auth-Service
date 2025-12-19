package middleware

import (
	"net/http"

	"auth-service/pkg/authctx"
	"auth-service/pkg/jwt"
)

func RequireAuth(jwtm *jwt.Manager, accessCookieName string) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c, err := r.Cookie(accessCookieName)
			if err != nil || c == nil || c.Value == "" {
				http.Error(w, "unauthorized", http.StatusUnauthorized)
				return
			}

			claims, err := jwtm.ParseAndValidate(c.Value)
			if err != nil {
				http.Error(w, "unauthorized", http.StatusUnauthorized)
				return
			}

			ctx := authctx.WithUserID(r.Context(), claims.UserID)
			ctx = authctx.WithEmail(ctx, claims.Email)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
