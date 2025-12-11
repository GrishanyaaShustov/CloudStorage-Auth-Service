package config

import "time"

type RedisConfig struct {
	Addr     string        `yaml:"addr" env:"REDIS_ADDR"`
	Password string        `yaml:"password" env:"REDIS_PASSWORD"`
	DB       int           `yaml:"db" env:"REDIS_DB"`
	Timeout  time.Duration `yaml:"timeout" env:"REDIS_TIMEOUT"`
}
