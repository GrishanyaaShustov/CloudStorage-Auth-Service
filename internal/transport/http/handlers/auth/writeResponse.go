package auth

import (
	svc "auth-service/internal/service/auth"
	"encoding/json"
	"errors"
	"net/http"
)

func (h *Handler) writeAuthError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, svc.ErrInvalidInput):
		http.Error(w, err.Error(), http.StatusBadRequest)
	case errors.Is(err, svc.ErrEmailAlreadyExists):
		http.Error(w, err.Error(), http.StatusConflict)
	case errors.Is(err, svc.ErrInvalidCredentials):
		http.Error(w, err.Error(), http.StatusUnauthorized)
	case errors.Is(err, svc.ErrInvalidRefreshToken):
		http.Error(w, err.Error(), http.StatusUnauthorized)
	case errors.Is(err, svc.ErrUserNotFound):
		http.Error(w, err.Error(), http.StatusNotFound)
	case errors.Is(err, svc.ErrInternal):
		h.Logger.Error("auth handler internal error", "err", err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
	default:
		h.Logger.Error("auth handler internal error", "err", err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
