package config

import "github.com/caarlos0/env"

var (
	Env = new(EnvConfig)
)

type EnvConfig struct {
	DBHost     string `env:"DB_HOST"`
	DBPort     int    `env:"DB_PORT"`
	DBUser     string `env:"DB_USER"`
	DBPassword string `env:"DB_PASSWORD"`
	DBDatabase string `env:"DB_DATABASE"`

	ServerAllowOrigins string `env:"SERVER_ALLOW_ORIGINS"`
}

func Init() error {
	// Initialize environment variables config
	if err := env.Parse(Env); err != nil {
		return err
	}

	return nil
}
