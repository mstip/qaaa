package web

import (
	"embed"
	"errors"
	"log"
	"net/http"

	"github.com/mstip/qaaa/pkg/model"
	"github.com/mstip/qaaa/pkg/store"
	"github.com/mstip/qaaa/pkg/waffel"
)

//go:embed static/*
var staticFiles embed.FS

//go:embed templates/*
var templateFiles embed.FS

func CreateWaffel() *waffel.Waffel {
	routes := []waffel.Route{
		{Name: "index", Method: http.MethodGet, Url: "/", Handler: indexController},
		{Name: "login", Method: http.MethodPost, Url: "/login", Handler: loginController},
		{Name: "dashboard", Method: http.MethodGet, Url: "/dashboard", Handler: dashboardController},
		{Name: "projectList", Method: http.MethodGet, Url: "/project/list", Handler: projectsListController},
		{Name: "projectCreate", Method: http.MethodGet, Url: "/project/create", Handler: projectCreateController},
		{Name: "projectStore", Method: http.MethodPost, Url: "/project/store", Handler: projectStoreController},
		{Name: "projectDetail", Method: http.MethodGet, Url: "/project/detail/{projectId}", Handler: projectDetailController},
		{Name: "projectDelete", Method: http.MethodPost, Url: "/project/detail/{projectId}/delete", Handler: projectDeleteController},
		{Name: "projectUpdate", Method: http.MethodPost, Url: "/project/detail/{projectId}/update", Handler: projectUpdateController},
		{Name: "projectEdit", Method: http.MethodGet, Url: "/project/detail/{projectId}/edit", Handler: projectEditController},
		// TODO Suites
	}

	middlewares := []waffel.Middleware{
		{Name: "nocache", Handler: noCacheMiddleware},
		{Name: "auth", Handler: authMiddleware},
	}

	waffel, err := waffel.NewWaffel("0.0.0.0:3000", routes, middlewares, staticFiles, templateFiles)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	store := store.NewStoreWithDemoData()
	waffel.RegisterService("store", store)

	return waffel
}

func RunAndServe() {
	waffel := CreateWaffel()
	waffel.RunAndServe()
}

func getStore(wf *waffel.Waffel) *store.Store {
	s, err := wf.GetService("store")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	store, _ := s.(*store.Store)
	return store
}

func getProjectFromUrlParams(wf *waffel.Waffel, r *http.Request) (*model.Project, error) {
	pId, err := waffel.GetRouteParamAsUint64(r, "projectId")
	if err != nil {
		return nil, err
	}
	project := getStore(wf).GetProjectById(pId)
	if project == nil {
		return nil, errors.New("project not found")
	}
	return project, nil
}
