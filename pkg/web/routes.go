package web

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (ws *WebServer) registerRoutes(r *mux.Router) error {
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		indexController(w, r, ws)
	}).Methods(http.MethodGet)

	r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		loginController(w, r, ws)
	}).Methods(http.MethodPost)

	r.HandleFunc("/dashboard", func(w http.ResponseWriter, r *http.Request) {
		ws.execTemplateHandler("dashboard", w, r)
	}).Methods(http.MethodGet)

	r.HandleFunc("/projects/list", func(w http.ResponseWriter, r *http.Request) {
		projectsListController(w, r, ws)
	}).Methods(http.MethodGet)

	r.HandleFunc("/projects/create", func(w http.ResponseWriter, r *http.Request) {
		ws.execTemplateHandler("project_create", w, r)
	}).Methods(http.MethodGet)

	r.HandleFunc("/projects/create", func(w http.ResponseWriter, r *http.Request) {
		projectCreateController(w, r, ws)
	}).Methods(http.MethodPost)

	r.HandleFunc("/projects/details/{projectId}", func(w http.ResponseWriter, r *http.Request) {
		projectDetailController(w, r, ws)
	}).Methods(http.MethodGet)

	r.HandleFunc("/projects/details/{projectId}/suites/create", func(w http.ResponseWriter, r *http.Request) {
		ws.execTemplateHandler("suite_create", w, r)
	}).Methods(http.MethodGet)

	r.HandleFunc("/projects/details/{projectId}/suites/create", func(w http.ResponseWriter, r *http.Request) {
		suiteCreateController(w, r, ws)
	}).Methods(http.MethodPost)

	r.HandleFunc("/projects/details/{projectId}/suites/details/{suiteId}", func(w http.ResponseWriter, r *http.Request) {
		suiteDetailController(w, r, ws)
	}).Methods(http.MethodGet)

	r.HandleFunc("/testruns/list", func(w http.ResponseWriter, r *http.Request) {
		ws.execTemplateHandler("testruns_list", w, r)
	}).Methods(http.MethodGet)

	r.Use(ws.authMiddleware)

	return nil
}
