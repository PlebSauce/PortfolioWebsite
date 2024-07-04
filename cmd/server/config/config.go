package config

import (
	"os"
)

type Config struct {
	DatabaseURL string
	RedisURL    string
}

func LoadConfig() (*Config, error) {
	dbURL := os.Getenv("DATABASE_URL")
	redisURL := os.Getenv("REDIS_URL")

	config := &Config{
		DatabaseURL: dbURL,
		RedisURL:    redisURL,
	}

	return config, nil
}
