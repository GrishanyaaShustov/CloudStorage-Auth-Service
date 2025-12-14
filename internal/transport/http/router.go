package http

import (
	"auth-service/internal/transport/http/handlers"
	"net/http"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	healthHandler := handlers.NewHealthHandler()

	mux.HandleFunc("/health", healthHandler.Health)

	return mux
}
