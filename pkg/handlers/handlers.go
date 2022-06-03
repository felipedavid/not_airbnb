package handlers

import (
	"net/http"

	"github.com/felipedavid/not_airbnb/pkg/config"
	"github.com/felipedavid/not_airbnb/pkg/models"
	"github.com/felipedavid/not_airbnb/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository si the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{App: a}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again. (this is data passed by the handler function.)"
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{StringMap: stringMap})
}
