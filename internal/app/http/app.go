package http

import (
	"auth-service/internal/config"
	"log/slog"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type App struct {
	Config *config.Config

	DBPool      *pgxpool.Pool
	RedisClient *redis.Client

	log *slog.Logger
	srv *http.Server
}
