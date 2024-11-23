package handlers

import (
	"html/template"
	"net/http"
	"path/filepath"
)

// RenderTemplate renders an HTML template with optional data
func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	tmplPath := filepath.Join("../..", "templates", tmpl+".html")
	t, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Unable to load template", http.StatusInternalServerError)
		return
	}
	t.Execute(w, data)
}

// HomeHandler handles requests to the home page and serves an HTML page
func MainscreenHandler(apiCfg *APIConfig) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		RenderTemplate(w, "index", nil)
	})
}

// AboutHandler handles requests to the about page and serves an HTML page
func AboutMeHandler(apiCfg *APIConfig) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		RenderTemplate(w, "aboutme", nil)
	})
}

func ProjectsHandler(apiCfg *APIConfig) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		RenderTemplate(w, "projects", nil)
	})
}

func ContactHandler(apiCfg *APIConfig) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		RenderTemplate(w, "contactandlinks", nil)
	})
}

func LoginHandler(apiCfg *APIConfig) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		RenderTemplate(w, "login", nil)
	})
}
