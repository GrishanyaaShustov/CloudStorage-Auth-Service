package config

import "time"

type PostgresConfig struct {
	Host     string `yaml:"host"            env:"DB_HOST"`
	Port     int    `yaml:"port"            env:"DB_PORT"`
	User     string `yaml:"user"            env:"DB_USER"`
	Password string `yaml:"password"        env:"DB_PASSWORD"`
	Name     string `yaml:"name"            env:"DB_NAME"`
	SslMode  string `yaml:"ssl-mode"        env:"DB_SSL_MODE"`

	MaxOpenConns    int           `yaml:"max_open_conns"     env:"DB_MAX_OPEN_CONNS" env-default:"20"`
	MaxIdleConns    int           `yaml:"max_idle_conns"     env:"DB_MAX_IDLE_CONNS" env-default:"5"`
	ConnMaxLifetime time.Duration `yaml:"conn_max_lifetime"  env:"DB_CONN_MAX_LIFETIME" env-default:"30m"`
	ConnMaxIdleTime time.Duration `yaml:"conn_max_idle_time" env:"DB_CONN_MAX_IDLE_TIME" env-default:"10m"`
}
