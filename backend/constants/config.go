package constants

import (
	"errors"
	"sync"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	Host       string `env:"DB_HOST"`
	DBPort     int    `env:"DB_PORT" envDefault:"5432"`
	DBUser     string `env:"DB_USER"`
	DBPassword string `env:"DB_PASSWORD"`
	DBName     string `env:"DB_NAME"`
	DBSSLMode  string `env:"DB_SSLMODE" envDefault:"disable"`
	ServerPort int    `env:"SERVER_PORT" envDefault:"8080"`
}

var (
	configOnce sync.Once
	config     Config
	configErr  error
)

func GetEnv() (Config, error) {
	configOnce.Do(func() {
		_ = godotenv.Load()

		if err := env.Parse(&config); err != nil {
			configErr = err
			return
		}

		if config.Host == "" || config.DBUser == "" || config.DBName == "" {
			configErr = errors.New("missing required database environment variables")
		}
	})

	return config, configErr
}
