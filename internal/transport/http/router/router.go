package router

import (
	"auth-service/internal/config"
	authh "auth-service/internal/transport/http/handlers/auth"
	"auth-service/internal/transport/http/middleware"
	"auth-service/pkg/jwt"
	"net/http"
)

func NewRouter(CORS config.CORSConfig, JWT *jwt.Manager, auth *authh.Handler) http.Handler {
	mux := http.NewServeMux()

	// public
	mux.Handle("/api/v1/auth/register", http.HandlerFunc(auth.Register))
	mux.Handle("/api/v1/auth/login", http.HandlerFunc(auth.Login))

	// protected wrapper
	protected := func(h http.Handler) http.Handler {
		return middleware.Chain(
			h,
			middleware.RequireAuth(JWT, "access_token"),
		)
	}

	// protected
	mux.Handle("/api/v1/auth/refresh", protected(http.HandlerFunc(auth.Refresh)))
	mux.Handle("/api/v1/auth/logout", protected(http.HandlerFunc(auth.Logout)))
	mux.Handle("/api/v1/auth/me", protected(http.HandlerFunc(auth.Me)))

	// global middleware
	return middleware.Chain(
		mux,
		middleware.CORS(CORS),
	)
}
