package web

import (
	"net/http"
	"strconv"

	"github.com/mstip/qaaa/pkg/waffel"
)

func suiteCreateController(wr http.ResponseWriter, r *http.Request, wf *waffel.Waffel) {
	project, err := getProjectFromUrlParams(wf, r)
	if err != nil {
		http.Error(wr, "Not found", http.StatusNotFound)
		return
	}
	params := waffel.RequestParams(r, "projectId")

	wf.Render(wr, r, "suite_form", map[string]interface{}{
		"Breadcrumb": []Breadcrumb{
			{Name: "Projects", Route: "/project/list"},
			{Name: project.Name, Route: "/project/detail/" + params["url_projectId"]},
			{Name: "Create New Suite", Route: ""},
		},
		"Form": ProjectForm{
			Title:       "Create new suite",
			Action:      wf.GetUrlForRoute("suiteStore", params["url_projectId"]),
			Name:        "",
			Description: "",
		},
	})
}

func suiteStoreController(wr http.ResponseWriter, r *http.Request, wf *waffel.Waffel) {
	project, err := getProjectFromUrlParams(wf, r)
	if err != nil {
		http.Error(wr, err.Error(), http.StatusNotFound)
		return
	}

	params := waffel.RequestParams(r, "name", "description")

	s := getStore(wf).CreateSuite(params["form_name"], params["form_description"], project.Id)

	http.Redirect(
		wr, r,
		"/project/detail/"+strconv.FormatUint(project.Id, 10)+"/suite/detail/"+strconv.FormatUint(s.Id, 10),
		http.StatusMovedPermanently,
	)
}

func suiteDetailController(wr http.ResponseWriter, r *http.Request, wf *waffel.Waffel) {
	project, err := getProjectFromUrlParams(wf, r)
	if err != nil {
		http.Error(wr, err.Error(), http.StatusNotFound)
		return
	}

	suite, err := getSuiteFromUrlParams(wf, r)
	if err != nil {
		http.Error(wr, err.Error(), http.StatusNotFound)
		return
	}

	tasks := getStore(wf).GetTasksBySuiteId(suite.Id)

	wf.Render(wr, r, "suite_details", map[string]interface{}{
		"Suite": suite,
		"Tasks": tasks,
		"Breadcrumb": []Breadcrumb{
			{Name: "Projects", Route: "/project/list"},
			{Name: project.Name, Route: "/project/detail/" + strconv.FormatUint(project.Id, 10)},
			{Name: "Suites", Route: ""},
			{Name: suite.Name, Route: ""},
		},
	})
}
