package handlers

import (
	"net/http"
)

func InitializeRoutes(mux *http.ServeMux, apiCfg *APIConfig) {
	mux.Handle("/", MainscreenHandler(apiCfg))
	mux.Handle("/login", LoginHandler(apiCfg))
	mux.Handle("/aboutme", AboutMeHandler(apiCfg))
	mux.Handle("/projects", ProjectsHandler(apiCfg))
	mux.Handle("/contactandlinks", ContactHandler(apiCfg))

	mux.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("../../static/css"))))
	mux.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("../../static/js"))))

}
