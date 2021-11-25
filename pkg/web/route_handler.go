package web

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mstip/qaaa/pkg/model"
	"github.com/mstip/qaaa/pkg/store"
)

func (ws *WebServer) registerRoutes(r *mux.Router) error {
	r.HandleFunc("/", ws.index).Methods(http.MethodGet)
	r.HandleFunc("/login", ws.login).Methods(http.MethodPost)

	r.HandleFunc("/dashboard", func(w http.ResponseWriter, r *http.Request) {
		ws.execTemplate("dashboard", w, r)
	}).Methods(http.MethodGet)

	r.HandleFunc("/projects/list", ws.projectsList).Methods(http.MethodGet)

	r.HandleFunc("/projects/create", func(w http.ResponseWriter, r *http.Request) {
		ws.execTemplate("project_create", w, r)
	}).Methods(http.MethodGet)
	r.HandleFunc("/projects/create", ws.projectCreate).Methods(http.MethodPost)

	r.HandleFunc("/projects/details/{projectId}", ws.projectDetail).Methods(http.MethodGet)

	r.HandleFunc("/projects/details/{projectId}/suites/create", ws.suiteCreate).Methods(http.MethodPost)
	r.HandleFunc("/projects/details/{projectId}/suites/details/{suiteId}", ws.suiteDetail).Methods(http.MethodGet)

	r.HandleFunc("/testruns/list", func(w http.ResponseWriter, r *http.Request) {
		ws.execTemplate("testruns_list", w, r)
	}).Methods(http.MethodGet)

	r.Use(ws.authMiddleware)

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

func (ws *WebServer) execTemplate(templateName string, w http.ResponseWriter, r *http.Request) {
	ws.templates[templateName].Execute(w, nil)
}

func (ws *WebServer) projectsList(w http.ResponseWriter, r *http.Request) {
	projects := ws.store.GetProjectsList()

	ws.templates["projects_list"].Execute(w, struct {
		Projects []store.ProjectListData
	}{
		Projects: projects,
	})
}

func (ws *WebServer) projectCreate(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	Description := r.FormValue("description")
	p := ws.store.CreateProject(name, Description)
	http.Redirect(w, r, "/projects/details/"+strconv.FormatUint(p.Id, 10), http.StatusMovedPermanently)
}

func (ws *WebServer) projectDetail(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	pId, err := strconv.ParseUint(string(params["projectId"]), 10, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	project := ws.store.GetProjectById(pId)
	if project == nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	ws.templates["project_details"].Execute(w, struct {
		Project *model.Project
	}{
		Project: project,
	})
}

func (ws *WebServer) suiteCreate(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	Description := r.FormValue("description")

	params := mux.Vars(r)
	pId, err := strconv.ParseUint(string(params["projectId"]), 10, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	project := ws.store.GetProjectById(pId)
	if project == nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	s := ws.store.CreateSuite(name, Description, project)
	http.Redirect(w, r, "/projects/details/"+strconv.FormatUint(project.Id, 10)+"/suites/details/"+strconv.FormatUint(s.Id, 10), http.StatusMovedPermanently)
}

func (ws *WebServer) suiteDetail(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	sId, err := strconv.ParseUint(string(params["suiteId"]), 10, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	suite := ws.store.GetSuiteById(sId)
	if suite == nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	ws.templates["suite_details"].Execute(w, struct {
		Suite *model.Suite
	}{
		Suite: suite,
	})
}
