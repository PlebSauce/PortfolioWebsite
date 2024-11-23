package handlers

import "github.com/PlebSauce/PortfolioWebsite/internal/database"

// APIConfig holds dependencies for handlers, such as the database.
type APIConfig struct {
	DB *database.Queries
}
