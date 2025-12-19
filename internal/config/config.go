package config

import (
	"flag"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env      string              `yaml:"env" env-default:"prod"`
	Postgres PostgresConfig      `yaml:"postgres"`
	Redis    RedisConfig         `yaml:"redis"`
	Tokens   TokensConfiguration `yaml:"tokens"`
	CORS     CORSConfig          `yaml:"cors"`
	Cookies  CookiesConfig       `yaml:"cookies"`
}

func MustLoad() *Config {
	path := getConfigPath()
	if path == "" {
		panic("CONFIG_PATH or --config flag is required but empty")
	}
	return mustLoadPath(path)
}

func mustLoadPath(path string) *Config {
	// Проверяем наличие файла
	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file does not exist: " + path)
	}

	var cfg Config

	// Читаем YAML + env-переменные
	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("cannot read config: " + err.Error())
	}

	return &cfg
}

func getConfigPath() string {
	var path string
	flag.StringVar(&path, "config", "", "path to config file")
	flag.Parse()

	if path != "" {
		return path
	}

	return os.Getenv("CONFIG_PATH")
}
