package main

import (
	"log"
	"net/http"

	"github.com/PlebSauce/PortfolioWebsite/cmd/server/config"
	"github.com/PlebSauce/PortfolioWebsite/internal/app/handlers"
	"github.com/PlebSauce/PortfolioWebsite/internal/database"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	db, err := config.SetupDatabase(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Error setting up database: %v", err)
	}
	defer db.Close()

	queries := database.New(db)

	apiCfg := &handlers.APIConfig{
		DB: queries,
	}

	// Initialize Redis connection (if applicable)
	if cfg.RedisURL != "" {
		redisClient, err := config.SetupRedis(cfg.RedisURL)
		if err != nil {
			log.Fatalf("Error setting up Redis: %v", err)
		}
		defer redisClient.Close()

		// Pass `redisClient` to services or handlers as needed
	}

	mux := http.NewServeMux()

	handlers.InitializeRoutes(mux, apiCfg)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
