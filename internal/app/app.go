package app

import (
	"context"
	"log/slog"
	"time"

	httpapp "auth-service/internal/app/http"
	"auth-service/internal/config"
	"auth-service/internal/repository/tokenRepo"
	"auth-service/internal/repository/userRepo"
	"auth-service/internal/service/auth"
	"auth-service/internal/storage/postgres"
	rdb "auth-service/internal/storage/redis"
	authhandler "auth-service/internal/transport/http/handlers/auth"
	"auth-service/internal/transport/http/router"
	"auth-service/pkg/hash"
	"auth-service/pkg/jwt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type App struct {
	log *slog.Logger
	cfg config.Config

	pg    *pgxpool.Pool
	redis *redis.Client

	http *httpapp.App
}

func New(ctx context.Context, log *slog.Logger, cfg config.Config) *App {
	// --- infra ---
	pgPool := postgres.MustNew(ctx, cfg.Postgres)

	redisClient := rdb.MustNew(ctx, cfg.Redis)

	// --- core deps ---
	jwtManager := jwt.NewManagerFromConfig(cfg.Tokens)
	hasher := hash.New(10)

	// --- repos ---
	userRepository := userRepo.New(pgPool)
	tokenRepository := tokenRepo.New(cfg.Tokens, redisClient)

	// --- services ---
	authService := auth.New(userRepository, tokenRepository, jwtManager, hasher)

	// --- handlers ---
	authHandler := authhandler.New(authService, cfg.Cookies, log)

	// --- http router ---
	mainHandler := router.NewRouter(cfg.CORS, jwtManager, authHandler)

	// --- http app ---
	addr := ":8080"

	httpSrv := httpapp.New(httpapp.Deps{
		Log:     log,
		Addr:    addr,
		Handler: mainHandler,

		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	})

	return &App{
		log:   log,
		cfg:   cfg,
		pg:    pgPool,
		redis: redisClient,
		http:  httpSrv,
	}
}

func (a *App) Run() error {
	return a.http.Run()
}

func (a *App) Shutdown(ctx context.Context) error {
	// 1) stop accepting new HTTP requests, finish inflight
	_ = a.http.Shutdown(ctx)

	// 2) close infra
	if a.redis != nil {
		_ = a.redis.Close()
	}
	if a.pg != nil {
		a.pg.Close()
	}
	return nil
}
