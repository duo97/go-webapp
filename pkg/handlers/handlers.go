package handlers

import (
	"fmt"
	"net/http"

	"github.com/duo97/go-webapp/pkg/config"
	"github.com/duo97/go-webapp/pkg/models"
	"github.com/duo97/go-webapp/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

//NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

//NewHandlers set the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

//Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{}, r)

}

//About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	//perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP
	//send the data to the template
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	}, r)

}

//Home is the home page handler
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "majors.page.tmpl", &models.TemplateData{}, r)

}

//Home is the home page handler
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "generals.page.tmpl", &models.TemplateData{}, r)

}

//Home is the home page handler
func (m *Repository) SearchAvailability(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "search-availability.page.tmpl", &models.TemplateData{}, r)

}
func (m *Repository) PostSearchAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")

	w.Write([]byte(fmt.Sprintf("start date is %s and end date is %s", start, end)))

}

//Home is the home page handler
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "reservation.page.tmpl", &models.TemplateData{}, r)

}
