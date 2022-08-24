package handlers

import (
	"learn-golang/pkg/config"
	"learn-golang/pkg/models"
	"learn-golang/pkg/render"
	"net/http"
)

// Repo the repository used by handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (r *Repository) Home(w http.ResponseWriter, _ *http.Request) {
	render.Template(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (r *Repository) About(w http.ResponseWriter, _ *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again!"

	render.Template(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
