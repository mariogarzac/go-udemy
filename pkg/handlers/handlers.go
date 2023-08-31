package handlers

import (
	"net/http"

	"github.com/mariogarzac/go-learngo/pkg/config"
	"github.com/mariogarzac/go-learngo/pkg/models"
	"github.com/mariogarzac/go-learngo/pkg/render"
)


var Repo *Repository

type Repository struct {
    App *config.AppConfig
}

// NewRepos creates new repository
func NewRepo(a *config.AppConfig) *Repository {
    return &Repository{
        App: a,
    }
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r* Repository){
    Repo = r
}

func (m *Repository)Home(w http.ResponseWriter, r *http.Request) {
    render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

func (m *Repository)About(w http.ResponseWriter, r *http.Request) {


    stringMap := make(map[string]string)
    stringMap["test"] = "Hello again"

    render.RenderTemplate(w, "about.page.html", &models.TemplateData{})
}

