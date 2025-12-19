package config

type CORSConfig struct {
	AllowedOrigins []string `yaml:"allowed_origins"`
	AllowedHeaders []string `yaml:"allowed_headers"`
	AllowedMethods []string `yaml:"allowed_methods"`
}
