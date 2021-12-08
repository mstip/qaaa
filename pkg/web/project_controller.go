package web

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mstip/qaaa/pkg/model"
)

type ProjectForm struct {
	Title       string
	Action      string
	Name        string
	Description string
}

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
	data := struct {
		Breadcrumb []Breadcrumb
		Form       ProjectForm
	}{
		Breadcrumb: []Breadcrumb{
			{Name: "Projects", Route: "/projects/list"},
			{Name: "Create New Project", Route: ""},
		},
		Form: ProjectForm{
			Title:       "Create new project",
			Action:      routeProjectsStore,
			Name:        "",
			Description: "",
		},
	}
	ws.execTemplateHandler("project_form", w, r, data)
}

func projectStoreController(w http.ResponseWriter, r *http.Request, ws *WebServer) {
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

func projectUpdateController(w http.ResponseWriter, r *http.Request, ws *WebServer) {
	params := mux.Vars(r)
	name := r.FormValue("name")
	description := r.FormValue("description")
	pId, err := getRouteParamAsUint64(r, "projectId")
	if err != nil {
		errorResponse(w, err)
		return
	}
	project := ws.store.UpdateProjectById(pId, name, description)
	if project == nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, "/projects/details/"+params["projectId"], http.StatusMovedPermanently)
}

func projectEditController(w http.ResponseWriter, r *http.Request, ws *WebServer) {
	params := mux.Vars(r)
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

	data := struct {
		Breadcrumb []Breadcrumb
		Form       ProjectForm
	}{
		Breadcrumb: []Breadcrumb{
			{Name: "Projects", Route: "/projects/list"},
			{Name: project.Name, Route: "/projects/details/" + params["projectId"]},
			{Name: "Edit", Route: ""},
		},
		Form: ProjectForm{
			Title:       "Update " + project.Name,
			Action:      "/projects/details/" + params["projectId"] + "/update",
			Name:        project.Name,
			Description: project.Description,
		},
	}
	ws.execTemplateHandler("project_form", w, r, data)
}
