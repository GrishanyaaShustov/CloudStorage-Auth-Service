package middleware

import (
	"auth-service/internal/config"
	"net/http"
	"strings"
)

func CORS(cfg config.CORSConfig) Middleware {
	allowedOrigins := make(map[string]struct{}, len(cfg.AllowedOrigins))
	for _, o := range cfg.AllowedOrigins {
		allowedOrigins[o] = struct{}{}
	}

	allowedHeaders := cfg.AllowedHeaders
	allowedMethods := cfg.AllowedMethods

	allowHeadersStr := strings.Join(allowedHeaders, ", ")
	allowMethodsStr := strings.Join(allowedMethods, ", ")

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			origin := r.Header.Get("Origin")

			// не CORS запрос
			if origin == "" {
				next.ServeHTTP(w, r)
				return
			}

			// Проверяем origin
			allowed := true
			if len(allowedOrigins) > 0 {
				_, allowed = allowedOrigins[origin]
			}

			if allowed {
				// Важно: с cookie нельзя ставить "*"
				w.Header().Set("Access-Control-Allow-Origin", origin)
				w.Header().Set("Access-Control-Allow-Credentials", "true")

				// чтобы кеш не путал ответы для разных origin / preflight
				w.Header().Add("Vary", "Origin")
				w.Header().Add("Vary", "Access-Control-Request-Method")
				w.Header().Add("Vary", "Access-Control-Request-Headers")

				// Preflight
				if r.Method == http.MethodOptions {
					reqHeaders := r.Header.Get("Access-Control-Request-Headers")
					if reqHeaders != "" {
						// браузер прислал свои заголовки — отвечаем тем, что разрешаем
						w.Header().Set("Access-Control-Allow-Headers", allowHeadersStr)
					} else {
						w.Header().Set("Access-Control-Allow-Headers", allowHeadersStr)
					}
					w.Header().Set("Access-Control-Allow-Methods", allowMethodsStr)
					w.WriteHeader(http.StatusNoContent)
					return
				}

				// обычный запрос
				w.Header().Set("Access-Control-Allow-Headers", allowHeadersStr)
				w.Header().Set("Access-Control-Allow-Methods", allowMethodsStr)
				next.ServeHTTP(w, r)
				return
			}

			// Origin не разрешён:
			// - preflight: можно вернуть 204 без CORS заголовков (браузер сам заблокирует)
			// - обычный: тоже можно просто отдать ответ без CORS заголовков
			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusNoContent)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
