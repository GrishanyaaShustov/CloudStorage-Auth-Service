package auth

import (
	"auth-service/internal/config"
	svc "auth-service/internal/service/auth"
	"auth-service/internal/service/auth/models"
	"encoding/json"
	"log/slog"
	"net/http"
	"time"
)

type Handler struct {
	Service    svc.Service
	CookiesCfg config.CookiesConfig
	Logger     *slog.Logger
}

func New(svc svc.Service, cookiesCfg config.CookiesConfig, logger *slog.Logger) *Handler {
	return &Handler{
		Service:    svc,
		CookiesCfg: cookiesCfg,
		Logger:     logger,
	}
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req registerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	resp, err := h.Service.Register(r.Context(), models.RegisterRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		h.writeAuthError(w, err)
		return
	}

	h.setAccessCookie(w, resp.AccessToken)
	writeJSON(w, http.StatusOK, struct{}{})
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req loginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	resp, err := h.Service.Login(r.Context(), models.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		h.writeAuthError(w, err)
		return
	}

	h.setAccessCookie(w, resp.AccessToken)
	writeJSON(w, http.StatusOK, struct{}{})
}

// Refresh - protected (middleware уже положил userID/email в ctx)
func (h *Handler) Refresh(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	resp, err := h.Service.RefreshAccess(r.Context(), models.RefreshAccessRequest{})
	if err != nil {
		h.writeAuthError(w, err)
		return
	}

	h.setAccessCookie(w, resp.AccessToken)
	writeJSON(w, http.StatusOK, struct{}{})
}

// Logout - protected
func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	_, err := h.Service.Logout(r.Context(), models.LogoutRequest{})
	if err != nil {
		h.writeAuthError(w, err)
		return
	}

	h.clearAccessCookie(w)
	writeJSON(w, http.StatusOK, struct{}{})
}

// Me - protected
func (h *Handler) Me(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	resp, err := h.Service.UserInformation(r.Context(), models.UserInformationRequest{})
	if err != nil {
		h.writeAuthError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, meResponse{
		UserID:       resp.UserID,
		Email:        resp.Email,
		RegisterDate: resp.RegisterDate.UTC().Format(time.RFC3339),
	})
}
