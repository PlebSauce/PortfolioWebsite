package config

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
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

func SetupDatabase(dbURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// perform any additional setup like migrations
	return db, nil
}

func SetupRedis(redisURL string) (*redis.Client, error) {
	// Implement Redis setup using a Redis client library (e.g., `go-redis`)
	// Example:
	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisURL,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// Test Redis connection
	_, err := redisClient.Ping( /*use context.TODO*/ nil).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	return redisClient, nil
}
