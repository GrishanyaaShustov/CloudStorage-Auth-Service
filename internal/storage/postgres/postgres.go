package postgres

import (
	"auth-service/internal/config"
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func New(ctx context.Context, cfg config.PostgresConfig) (*pgxpool.Pool, error) {
	poolcfg, err := pgxpool.ParseConfig("")
	if err != nil {
		return nil, fmt.Errorf("parse pool config: %w", err)
	}

	// Configure connection
	poolcfg.ConnConfig.Host = cfg.Host         // set DB host from config
	poolcfg.ConnConfig.Port = uint16(cfg.Port) // set DB port from config
	poolcfg.ConnConfig.User = cfg.User         // set DB user from config
	poolcfg.ConnConfig.Password = cfg.Password // set DB password from config
	poolcfg.ConnConfig.Database = cfg.Name     // set DB name from config

	if cfg.SslMode == "disable" {
		poolcfg.ConnConfig.TLSConfig = nil
	}

	// Configure connection pool
	poolcfg.MaxConns = int32(cfg.MaxOpenConns)
	poolcfg.MinConns = int32(cfg.MaxIdleConns)
	poolcfg.MaxConnLifetime = cfg.ConnMaxLifetime
	poolcfg.MaxConnIdleTime = cfg.ConnMaxIdleTime

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	pool, err := pgxpool.NewWithConfig(ctx, poolcfg)
	if err != nil {
		return nil, fmt.Errorf("pgxpool connect error: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("pgxpool ping error: %w", err)
	}

	return pool, nil
}

// MustNew calls New() and panics on error.
// Use it for services that *must* crash if DB is unavailable.
func MustNew(ctx context.Context, cfg config.PostgresConfig) *pgxpool.Pool {
	pool, err := New(ctx, cfg)
	if err != nil {
		panic("failed to initialize postgres: " + err.Error())
	}
	return pool
}
