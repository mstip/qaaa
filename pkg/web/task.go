package web

import (
	"net/http"
	"strconv"

	"github.com/mstip/qaaa/pkg/task"
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
	task, err := getTaskFromUrlParams(wf, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	suite, err := getSuiteFromUrlParams(wf, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	project, err := getProjectFromUrlParams(wf, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	wf.Render(w, r, "task_detail", map[string]interface{}{
		"Breadcrumb": []Breadcrumb{
			{Name: "Projects", Route: wf.GetUrlForRoute("projectList")},
			{Name: project.Name, Route: wf.GetUrlForRouteUInt64("projectDetail", project.Id)},
			{Name: suite.Name, Route: wf.GetUrlForRouteUInt64("suiteDetail", project.Id, suite.Id)},
			{Name: task.Name, Route: ""},
		},
		"Task":    task,
		"Suite":   suite,
		"Project": project,
	})
}

func taskEditController(w http.ResponseWriter, r *http.Request, wf *waffel.Waffel) {
	task, err := getTaskFromUrlParams(wf, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
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
			{Name: task.Name, Route: wf.GetUrlForRouteUInt64("taskDetail", project.Id, suite.Id, task.Id)},
			{Name: "Edit", Route: ""},
		},
		"Form": TaskForm{
			Title:       task.Name,
			Action:      wf.GetUrlForRouteUInt64("taskUpdate", project.Id, suite.Id, task.Id),
			Name:        task.Name,
			Description: task.Description,
			Type:        task.Type,
		},
	})
}

func taskUpdateController(w http.ResponseWriter, r *http.Request, wf *waffel.Waffel) {
	params := waffel.RequestParams(r, "projectId", "suiteId", "taskId", "name", "description", "type")
	task := getStore(wf).UpdateTaskByIdParam(params["url_taskId"], params["form_name"], params["form_description"])
	if task == nil {
		http.Error(w, "task not found", http.StatusNotFound)
		return
	}
	wf.RedirectToRoute(w, r, "taskDetail", params["url_projectId"], params["url_suiteId"], params["url_taskId"])
}

func taskDeleteController(w http.ResponseWriter, r *http.Request, wf *waffel.Waffel) {
	params := waffel.RequestParams(r, "projectId", "suiteId", "taskId", "name", "description", "type")
	task := getStore(wf).DeleteTaskByIdParm(params["url_taskId"])
	if task == nil {
		http.Error(w, "task not found", http.StatusNotFound)
		return
	}
	wf.RedirectToRoute(w, r, "suiteDetail", params["url_projectId"], params["url_suiteId"])
}

func taskTestRunController(w http.ResponseWriter, r *http.Request, wf *waffel.Waffel) {
	body, err := waffel.JsonBody(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	method, ok := body["method"].(string)
	if !ok {
		http.Error(w, "method is not a string", http.StatusBadRequest)
		return
	}
	url, ok := body["url"].(string)
	if !ok {
		http.Error(w, "url is not a string", http.StatusBadRequest)
		return
	}

	result, err := task.TaskTestRun("json", method, url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = waffel.JsonResponse(w, result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
