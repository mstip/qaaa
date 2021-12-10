package web

import (
	"net/http"
	"strconv"

	"github.com/mstip/qaaa/pkg/model"
	"github.com/mstip/qaaa/pkg/waffel"
)

type ProjectForm struct {
	Title       string
	Action      string
	Name        string
	Description string
}

func projectsListController(wr http.ResponseWriter, r *http.Request, wf *waffel.Waffel) {
	projects := getStore(wf).GetProjects()
	wf.Render(wr, r, "projects_list",
		map[string]interface{}{
			"Projects":   projects,
			"Breadcrumb": []Breadcrumb{{Name: "Projects", Route: ""}},
		})

}

func projectCreateController(wr http.ResponseWriter, r *http.Request, wf *waffel.Waffel) {
	wf.Render(wr, r, "project_form", map[string]interface{}{
		"Breadcrumb": []Breadcrumb{
			{Name: "Projects", Route: wf.GetUrlForRoute("projectList")},
			{Name: "Create New Project", Route: ""},
		},
		"Form": ProjectForm{
			Title:       "Create new project",
			Action:      wf.GetUrlForRoute("projectStore"),
			Name:        "",
			Description: "",
		},
	})
}

func projectStoreController(w http.ResponseWriter, r *http.Request, wf *waffel.Waffel) {
	name := r.FormValue("name")
	Description := r.FormValue("description")
	p := getStore(wf).CreateProject(name, Description)
	http.Redirect(w, r, "/project/detail/"+strconv.FormatUint(p.Id, 10), http.StatusMovedPermanently)
}

func projectDetailController(wr http.ResponseWriter, r *http.Request, wf *waffel.Waffel) {
	project, err := getProjectFromUrlParams(wf, r)
	if err != nil {
		http.Error(wr, "Not found", http.StatusNotFound)
		return
	}
	suites := getStore(wf).GetSuitesByProjectId(project.Id)
	tasksBySuiteId := map[uint64][]model.Task{}
	for _, s := range suites {
		tasksBySuiteId[s.Id] = getStore(wf).GetTasksBySuiteId(s.Id)
	}

	wf.Render(wr, r, "project_details", map[string]interface{}{
		"Project":        project,
		"Suites":         suites,
		"TasksBySuiteId": tasksBySuiteId,
		"Breadcrumb": []Breadcrumb{
			{Name: "Projects", Route: wf.GetUrlForRoute("projectList")},
			{Name: project.Name, Route: ""},
		},
	})
}

func projectDeleteController(w http.ResponseWriter, r *http.Request, wf *waffel.Waffel) {
	pId, err := waffel.GetRouteParamAsUint64(r, "projectId")
	if err != nil {
		waffel.ErrorResponse(w, err)
		return
	}

	if p := getStore(wf).DeleteProjectById(pId); p == nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	if err = wf.GetSession().AddFlash("Project was deleted", "success", w, r); err != nil {
		waffel.ErrorResponse(w, err)
		return
	}

	http.Redirect(w, r, "/project/list", http.StatusMovedPermanently)
}

func projectUpdateController(wr http.ResponseWriter, r *http.Request, wf *waffel.Waffel) {
	params := waffel.RequestParams(r, "name", "description", "projectId")

	project := getStore(wf).UpdateProjectByIdParam(params["url_projectId"], params["form_name"], params["form_description"])
	if project == nil {
		http.Error(wr, "project not found", http.StatusNotFound)
		return
	}
	http.Redirect(wr, r, "/project/detail/"+params["url_projectId"], http.StatusMovedPermanently)
}

func projectEditController(wr http.ResponseWriter, r *http.Request, wf *waffel.Waffel) {
	params := waffel.RequestParams(r, "name", "description", "projectId")
	project, err := getProjectFromUrlParams(wf, r)

	if err != nil {
		http.Error(wr, "Not found", http.StatusNotFound)
		return
	}

	wf.Render(wr, r, "project_form", map[string]interface{}{
		"Breadcrumb": []Breadcrumb{
			{Name: "Projects", Route: wf.GetUrlForRoute("projectList")},
			{Name: project.Name, Route: wf.GetUrlForRoute("projectDetail", params["url_projectId"])},
			{Name: "Edit", Route: ""},
		},
		"Form": ProjectForm{
			Title:       "Edit " + project.Name,
			Action:      "/project/detail/" + params["url_projectId"] + "/update",
			Name:        project.Name,
			Description: project.Description,
		},
	})
}
