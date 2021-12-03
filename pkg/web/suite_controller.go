package web

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mstip/qaaa/pkg/model"
)

func suiteCreateController(w http.ResponseWriter, r *http.Request, ws *WebServer) {
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

	ws.templates["suite_create"].Execute(w, struct {
		Breadcrumb []Breadcrumb
	}{
		Breadcrumb: []Breadcrumb{
			{Name: "Projects", Route: "/projects/list"},
			{Name: project.Name, Route: "/projects/details/" + params["projectId"]},
			{Name: "Create New Suite", Route: ""},
		},
	})
}
func suiteStoreController(w http.ResponseWriter, r *http.Request, ws *WebServer) {
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

	s := ws.store.CreateSuite(name, Description, project.Id)
	http.Redirect(
		w, r,
		"/projects/details/"+strconv.FormatUint(project.Id, 10)+"/suites/details/"+strconv.FormatUint(s.Id, 10),
		http.StatusMovedPermanently,
	)
}

func suiteDetailController(w http.ResponseWriter, r *http.Request, ws *WebServer) {
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
		Suite      *model.Suite
		Breadcrumb []Breadcrumb
	}{
		Suite: suite,
		Breadcrumb: []Breadcrumb{
			{Name: "Projects", Route: "/projects/list"},
			{Name: project.Name, Route: "/projects/details/" + params["projectId"]},
			{Name: "Suites", Route: ""},
			{Name: suite.Name, Route: ""},
		},
	})
}