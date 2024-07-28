package handlers

import (
	"net/http"

	"github.com/PlebSauce/PortfolioWebsite/internal/app/services"
)

func handlerErr(w http.ResponseWriter, r *http.Request) {
	services.RespondWithError(w, 400, "Something went wrong")
}
