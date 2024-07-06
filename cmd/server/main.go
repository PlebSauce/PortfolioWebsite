package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/PlebSauce/PortfolioWebsite/cmd/server/config"
	"github.com/PlebSauce/PortfolioWebsite/internal/app/handlers"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
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

func setupRedis(redisURL string) (*redis.Client, error) {
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

func startServer() {
	// Implement server initialization (HTTP handlers, middleware setup, etc.)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("static/css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("static/js"))))

	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/", handlers.MainscreenHandler)
	http.HandleFunc("/aboutme", handlers.AboutMeHandler)
	http.HandleFunc("/projects", handlers.ProjectsHandler)
	http.HandleFunc("/contactandlinks", handlers.ContactHandler)
	//http.HandleFunc("/mainscreen", handlers.MainScreenHandler)
	// start server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server error:", err)
	}

}
