package handlers

import (
	"net/http"
	"web-with-go/pkg/config"
	"web-with-go/pkg/render"
)

//Repository structure is used when we connect the databases, it is used to share the connection pool

//Repo the repository used by the handlers
var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

//creates a new repository
func NewRepo(a *config.AppConfig) *Repository{
	return &Repository{
		App:a,
	}
}

//Sets the repository for new handlers
func NewHandlers(r *Repository) {
	Repo = r
}

//putting the reveivers will grant to the persmission to access everything inside the repository
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html")
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.html")
}
