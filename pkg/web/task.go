package web

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/mstip/qaaa/pkg/waffel"
)

type TaskForm struct {
	Title       string
	Action      string
	Name        string
	Description string
	Type        string
}

func taskCreateController(w http.ResponseWriter, r *http.Request, wf *waffel.Waffel) {
	project, err := getProjectFromUrlParams(wf, r)
	if err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	suite, err := getSuiteFromUrlParams(wf, r)
	if err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	wf.Render(w, r, "task_form", map[string]interface{}{
		"Breadcrumb": []Breadcrumb{
			{Name: "Projects", Route: wf.GetUrlForRoute("projectList")},
			{Name: project.Name, Route: wf.GetUrlForRouteUInt64("projectDetail", project.Id)},
			{Name: suite.Name, Route: wf.GetUrlForRouteUInt64("suiteDetail", project.Id, suite.Id)},
			{Name: "Create New Task", Route: ""},
		},
		"Form": TaskForm{
			Title:       "Create new task",
			Action:      wf.GetUrlForRouteUInt64("taskStore", project.Id, suite.Id),
			Name:        "",
			Description: "",
			Type:        "",
		},
	})
}

func taskStoreController(w http.ResponseWriter, r *http.Request, wf *waffel.Waffel) {
	suite, err := getSuiteFromUrlParams(wf, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	params := waffel.RequestParams(r, "name", "description", "type", "projectId", "suiteId")
	t := getStore(wf).CreateTask(params["form_name"], params["form_description"], suite.Id)

	wf.RedirectToRoute(w, r, "taskDetail", params["url_projectId"], params["url_suiteId"], strconv.FormatUint(t.Id, 10))
}

func taskDetailController(w http.ResponseWriter, r *http.Request, wf *waffel.Waffel) {
	// TODO: implement
	t, err := getTaskFromUrlParams(wf, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "%#v", t)
}
