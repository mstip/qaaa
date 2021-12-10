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
			{Name: "Projects", Route: wf.GetUrlForRoute("projectList")},
			{Name: project.Name, Route: wf.GetUrlForRoute("projectDetail", params["url_projectId"])},
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

	params := waffel.RequestParams(r, "name", "description", "projectId", "suiteId")

	s := getStore(wf).CreateSuite(params["form_name"], params["form_description"], project.Id)
	if s == nil {
		http.Error(wr, "not found", http.StatusNotFound)
		return
	}

	wf.RedirectToRoute(wr, r, "suiteDetail", params["url_projectId"], strconv.FormatUint(s.Id, 10))
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
			{Name: "Projects", Route: wf.GetUrlForRoute("projectList")},
			{Name: project.Name, Route: wf.GetUrlForRoute("projectDetail", strconv.FormatUint(project.Id, 10))},
			{Name: "Suites", Route: ""},
			{Name: suite.Name, Route: ""},
		},
	})
}

func suiteEditController(w http.ResponseWriter, r *http.Request, wf *waffel.Waffel) {
	project, err := getProjectFromUrlParams(wf, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	suite, err := getSuiteFromUrlParams(wf, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	params := waffel.RequestParams(r, "projectId", "suiteId")

	wf.Render(w, r, "suite_form", map[string]interface{}{
		"Breadcrumb": []Breadcrumb{
			{Name: "Projects", Route: wf.GetUrlForRoute("projectList")},
			{Name: project.Name, Route: wf.GetUrlForRoute("projectDetail", strconv.FormatUint(project.Id, 10))},
			{Name: "Suites", Route: ""},
			{Name: suite.Name, Route: wf.GetUrlForRoute("suite_detail", params["url_projectId"], params["url_suiteId"])},
			{Name: "Edit", Route: ""},
		},
		"Form": ProjectForm{
			Title:       "Edit " + suite.Name,
			Action:      wf.GetUrlForRoute("suiteUpdate", params["url_projectId"], params["url_suiteId"]),
			Name:        suite.Name,
			Description: suite.Description,
		},
	})
}

func suiteUpdateController(w http.ResponseWriter, r *http.Request, wf *waffel.Waffel) {
	params := waffel.RequestParams(r, "projectId", "suiteId", "name", "description")
	s := getStore(wf).UpdateSuiteByIdParam(params["url_suiteId"], params["form_name"], params["form_description"])
	if s == nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	wf.RedirectToRoute(w, r, "suiteDetail", params["url_projectId"], params["url_suiteId"])
}

func suiteDeleteController(w http.ResponseWriter, r *http.Request, wf *waffel.Waffel) {
	params := waffel.RequestParams(r, "projectId", "suiteId")
	s := getStore(wf).DeleteSuiteByIdParm(params["url_suiteId"])
	if s == nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	wf.RedirectToRoute(w, r, "projectDetail", params["url_projectId"])
}
