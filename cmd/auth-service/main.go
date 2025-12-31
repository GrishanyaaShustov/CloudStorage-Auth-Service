package main

import (
	"auth-service/cmd/auth-service/logger"
	"auth-service/internal/config"
)

func main() {
	cfg := config.MustLoad()

	log := logger.SetupLogger(cfg.Env)

	log.Info("Hello world!")
}
