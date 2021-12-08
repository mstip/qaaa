package web

import (
	"net/http"

	"github.com/gorilla/mux"
)

const (
	routeIndex          = "/"
	routeAuthLogin      = "/login"
	routeDashboard      = "/dashboard"
	routeProjectsList   = "/projects/list"
	routeProjectsCreate = "/projects/create"
	routeProjectsStore  = "/projects/store"
	routeProjectDetails = "/projects/details/{projectId}"
	routeProjectDelete  = "/projects/details/{projectId}/delete"
	routeProjectEdit    = "/projects/details/{projectId}/edit"
	routeProjectUpdate  = "/projects/details/{projectId}/update"
	routeSuiteCreate    = "/projects/details/{projectId}/suites/create"
	routeSuiteStore     = "/projects/details/{projectId}/suites/create"
	routeSuiteDetail    = "/projects/details/{projectId}/suites/details/{suiteId}"
	routeTestRunsList   = "/testruns/list"
)

func (ws *WebServer) registerRoutes(r *mux.Router) error {
	r.HandleFunc(routeIndex, func(w http.ResponseWriter, r *http.Request) {
		indexController(w, r, ws)
	}).Methods(http.MethodGet)

	r.HandleFunc(routeAuthLogin, func(w http.ResponseWriter, r *http.Request) {
		loginController(w, r, ws)
	}).Methods(http.MethodPost)

	r.HandleFunc(routeDashboard, func(w http.ResponseWriter, r *http.Request) {
		data := struct {
			Breadcrumb []Breadcrumb
		}{
			Breadcrumb: []Breadcrumb{
				{Name: "Dashboard", Route: ""},
			},
		}
		ws.execTemplateHandler("dashboard", w, r, data)
	}).Methods(http.MethodGet)

	r.HandleFunc(routeProjectsList, func(w http.ResponseWriter, r *http.Request) {
		projectsListController(w, r, ws)
	}).Methods(http.MethodGet)

	r.HandleFunc(routeProjectsCreate, func(w http.ResponseWriter, r *http.Request) {
		projectCreateController(w, r, ws)
	}).Methods(http.MethodGet)

	r.HandleFunc(routeProjectsStore, func(w http.ResponseWriter, r *http.Request) {
		projectStoreController(w, r, ws)
	}).Methods(http.MethodPost)

	r.HandleFunc(routeProjectDetails, func(w http.ResponseWriter, r *http.Request) {
		projectDetailController(w, r, ws)
	}).Methods(http.MethodGet)

	r.HandleFunc(routeProjectDelete, func(w http.ResponseWriter, r *http.Request) {
		projectDeleteController(w, r, ws)
	}).Methods(http.MethodPost)

	r.HandleFunc(routeProjectEdit, func(w http.ResponseWriter, r *http.Request) {
		projectEditController(w, r, ws)
	}).Methods(http.MethodGet)

	r.HandleFunc(routeProjectUpdate, func(w http.ResponseWriter, r *http.Request) {
		projectUpdateController(w, r, ws)
	}).Methods(http.MethodPost)

	r.HandleFunc(routeSuiteCreate, func(w http.ResponseWriter, r *http.Request) {
		suiteCreateController(w, r, ws)
	}).Methods(http.MethodGet)

	r.HandleFunc(routeSuiteStore, func(w http.ResponseWriter, r *http.Request) {
		suiteStoreController(w, r, ws)
	}).Methods(http.MethodPost)

	r.HandleFunc(routeSuiteDetail, func(w http.ResponseWriter, r *http.Request) {
		suiteDetailController(w, r, ws)
	}).Methods(http.MethodGet)

	r.HandleFunc(routeTestRunsList, func(w http.ResponseWriter, r *http.Request) {
		data := struct {
			Breadcrumb []Breadcrumb
		}{
			Breadcrumb: []Breadcrumb{
				{Name: "Testruns", Route: ""},
			},
		}
		ws.execTemplateHandler("testruns_list", w, r, data)
	}).Methods(http.MethodGet)

	r.Use(ws.NoCacheMiddleware)
	r.Use(ws.authMiddleware)

	return nil
}
