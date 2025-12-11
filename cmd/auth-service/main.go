package main

import (
	"auth-service/cmd/auth-service/logger"
	"auth-service/internal/config"
)

func main() {

	// initialize config
	cfg := config.MustLoad()

	//initialize logger
	log := logger.SetupLogger(cfg.Env)

	log.Info("Hello")

}
