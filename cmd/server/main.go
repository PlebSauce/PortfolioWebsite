package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/PlebSauce/PortfolioWebsite/cmd/server/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	db, err := setupDatabase(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Error setting up database: %v", err)
	}
	defer db.Close()

	// Initialize Redis connection (if applicable)
	if cfg.RedisURL != "" {
		redisClient, err := setupRedis(cfg.RedisURL)
		if err != nil {
			log.Fatalf("Error setting up Redis: %v", err)
		}
		defer redisClient.Close()

		// Pass `redisClient` to services or handlers as needed
	}

	startServer()
}

func setupDatabase(dbURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// perform any additional setup like migrations
	return db, nil
}

func setupRedis(redisURL string) (*redis.redisClientClient, error) {
	// Implement Redis setup using a Redis client library (e.g., `go-redis`)
	// Example:
	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisURL,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// Test Redis connection
	_, err := redisClient.Ping().Result()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	return redisClient, nil
}

func startServer() {
	// Implement server initialization (HTTP handlers, middleware setup, etc.)
	// Example:
	// router := setupRouter()
	// http.ListenAndServe(":8080", router)
}
