package auth

import "net/http"

func (h *Handler) setAccessCookie(w http.ResponseWriter, accessToken string) {
	http.SetCookie(w, &http.Cookie{
		Name:     h.CookiesCfg.AccessCookieName,
		Value:    accessToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   h.CookiesCfg.SecureCookies,
		SameSite: http.SameSiteLaxMode,
	})
}

func (h *Handler) clearAccessCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     h.CookiesCfg.AccessCookieName,
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   h.CookiesCfg.SecureCookies,
		SameSite: http.SameSiteLaxMode,
	})
}
