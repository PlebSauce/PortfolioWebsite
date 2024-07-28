package handlers

import (
	"fmt"
	"net/http"

	"github.com/PlebSauce/PortfolioWebsite/internal/app/auth"
	"github.com/PlebSauce/PortfolioWebsite/internal/app/services"
	"github.com/PlebSauce/PortfolioWebsite/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

type apiConfig struct {
	DB *database.Queries
}

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			services.RespondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			services.RespondWithError(w, 400, fmt.Sprintf("Couldn't get user: %v", err))
		}

		handler(w, r, user)
	}
}
