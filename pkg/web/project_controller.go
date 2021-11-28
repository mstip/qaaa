package web

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mstip/qaaa/pkg/model"
	"github.com/mstip/qaaa/pkg/store"
)

func (ws *WebServer) projectsListController(w http.ResponseWriter, r *http.Request) {
	projects := ws.store.GetProjectsList()

	ws.templates["projects_list"].Execute(w, struct {
		Projects []store.ProjectListData
	}{
		Projects: projects,
	})
}

func (ws *WebServer) projectCreateController(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	Description := r.FormValue("description")
	p := ws.store.CreateProject(name, Description)
	http.Redirect(w, r, "/projects/details/"+strconv.FormatUint(p.Id, 10), http.StatusMovedPermanently)
}

func (ws *WebServer) projectDetailController(w http.ResponseWriter, r *http.Request) {
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
