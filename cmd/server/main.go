package main

import (
	"fmt"
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

	apiCfg := apiConfig{
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

	startServer()
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
