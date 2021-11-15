package web

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mstip/qaaa/pkg/store"
)

func (ws *WebServer) registerRoutes(r *mux.Router) error {
	r.HandleFunc("/", ws.index).Methods(http.MethodGet)
	r.HandleFunc("/login", ws.login).Methods(http.MethodPost)
	r.HandleFunc("/dashboard", func(w http.ResponseWriter, r *http.Request) {
		ws.execTemplateWithAuthentication("dashboard", w, r)
	}).Methods(http.MethodGet)
	r.HandleFunc("/projects/list", ws.projectsList).Methods(http.MethodGet)
	r.HandleFunc("/projects/details/{projectId}", func(w http.ResponseWriter, r *http.Request) {
		ws.execTemplateWithAuthentication("project_details", w, r)
	}).Methods(http.MethodGet)
	r.HandleFunc("/testruns/list", func(w http.ResponseWriter, r *http.Request) {
		ws.execTemplateWithAuthentication("testruns_list", w, r)
	}).Methods(http.MethodGet)
	return nil
}

func (ws *WebServer) index(w http.ResponseWriter, r *http.Request) {
	flashes, err := ws.sessionStore.flashes(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ws.templates["index"].Execute(w, struct {
		Flashes []interface{}
	}{
		Flashes: flashes,
	})
}

func (ws *WebServer) login(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "admin" && password == "admin" {
		err := ws.sessionStore.authenticate(1, w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/dashboard", http.StatusMovedPermanently)
		return
	}

	err := ws.sessionStore.addFlash("invalid credentials", w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func (ws *WebServer) execTemplateWithAuthentication(templateName string, w http.ResponseWriter, r *http.Request) {
	if !ws.sessionStore.isAuthenticated(r) {
		err := ws.sessionStore.addFlash("please login", w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
		return
	}

	ws.templates[templateName].Execute(w, nil)
}

func (ws *WebServer) projectsList(w http.ResponseWriter, r *http.Request) {
	if !ws.sessionStore.isAuthenticated(r) {
		err := ws.sessionStore.addFlash("please login", w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
		return
	}

	projects := ws.store.GetProjectsList()
	log.Println(projects)
	err := ws.templates["projects_list"].Execute(w, struct {
		Projects []store.ProjectListData
	}{
		Projects: projects,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
