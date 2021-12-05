package web

import (
	"net/http"
	"strconv"

	"github.com/mstip/qaaa/pkg/model"
)

func projectsListController(w http.ResponseWriter, r *http.Request, ws *WebServer) {
	flashes, err := ws.sessionStore.Flashes(w, r)
	if err != nil {
		errorResponse(w, err)
		return
	}
	projects := ws.store.GetProjects()
	ws.templates["projects_list"].Execute(w, struct {
		Projects   []model.Project
		Flashes    []Flash
		Breadcrumb []Breadcrumb
	}{
		Projects:   projects,
		Flashes:    flashes,
		Breadcrumb: []Breadcrumb{{Name: "Projects", Route: ""}},
	})
}

func projectCreateController(w http.ResponseWriter, r *http.Request, ws *WebServer) {
	name := r.FormValue("name")
	Description := r.FormValue("description")
	p := ws.store.CreateProject(name, Description)
	http.Redirect(w, r, "/projects/details/"+strconv.FormatUint(p.Id, 10), http.StatusMovedPermanently)
}

func projectDetailController(w http.ResponseWriter, r *http.Request, ws *WebServer) {
	pId, err := getRouteParamAsUint64(r, "projectId")
	if err != nil {
		errorResponse(w, err)
		return
	}
	project := ws.store.GetProjectById(pId)
	if project == nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	suites := ws.store.GetSuitesByProjectId(pId)

	ws.templates["project_details"].Execute(w, struct {
		Project    *model.Project
		Suites     []model.Suite
		Breadcrumb []Breadcrumb
	}{
		Project: project,
		Suites:  suites,
		Breadcrumb: []Breadcrumb{
			{Name: "Projects", Route: "/projects/list"},
			{Name: project.Name, Route: ""},
		},
	})
}

func projectDeleteController(w http.ResponseWriter, r *http.Request, ws *WebServer) {
	pId, err := getRouteParamAsUint64(r, "projectId")
	if err != nil {
		errorResponse(w, err)
		return
	}

	if p := ws.store.DeleteProjectById(pId); p == nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	if err = ws.sessionStore.AddFlash("Project was deleted", "success", w, r); err != nil {
		errorResponse(w, err)
		return
	}

	http.Redirect(w, r, "/projects/list", http.StatusMovedPermanently)
}
