package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"auth-service/cmd/auth-service/logger"
	"auth-service/internal/app"
	"auth-service/internal/config"
)

func main() {
	cfg := config.MustLoad()
	log := logger.SetupLogger(cfg.Env)

	// context which will be cancelled on SIGINT/SIGTERM
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	application := app.New(ctx, log, *cfg)

	// run server in goroutine
	errCh := make(chan error, 1)
	go func() {
		errCh <- application.Run()
	}()

	select {
	case <-ctx.Done():
		// graceful shutdown
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := application.Shutdown(shutdownCtx); err != nil {
			log.Error("shutdown failed", "err", err)
			os.Exit(1)
		}

		log.Info("shutdown complete")
		os.Exit(0)
	case err := <-errCh:
		if err != nil {
			log.Error("app stopped with error", "err", err)
			os.Exit(1)
		}
	}
}
