package config

import "time"

type TokensConfiguration struct {
	JwtSecret       string        `yaml:"jwt_secret" env:"JWT_SECRET"`
	AccessTokenTTL  time.Duration `yaml:"access_token_ttl" env:"ACCESS_TOKEN_TTL"`
	RefreshTokenTTL time.Duration `yaml:"refresh_token_ttl" env:"REFRESH_TOKEN_TTL"`
}
