package handlers

import (
	"html/template"
	"net/http"
	"path/filepath"
)

// RenderTemplate renders an HTML template with optional data
func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	tmplPath := filepath.Join("templates", tmpl+".html")
	t, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Unable to load template", http.StatusInternalServerError)
		return
	}
	t.Execute(w, data)
}

// HomeHandler handles requests to the home page and serves an HTML page
func MainscreenHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "index", nil)
}

// AboutHandler handles requests to the about page and serves an HTML page
func AboutMeHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "aboutme", nil)
}

func ProjectsHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "projects", nil)
}

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "contactandlinks", nil)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "login", nil)
}
