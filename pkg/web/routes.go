package web

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (ws *WebServer) registerRoutes(r *mux.Router) error {
	r.HandleFunc("/", ws.indexController).Methods(http.MethodGet)
	r.HandleFunc("/login", ws.loginController).Methods(http.MethodPost)

	r.HandleFunc("/dashboard", func(w http.ResponseWriter, r *http.Request) {
		ws.execTemplateHandler("dashboard", w, r)
	}).Methods(http.MethodGet)

	r.HandleFunc("/projects/list", ws.projectsListController).Methods(http.MethodGet)

	r.HandleFunc("/projects/create", func(w http.ResponseWriter, r *http.Request) {
		ws.execTemplateHandler("project_create", w, r)
	}).Methods(http.MethodGet)
	r.HandleFunc("/projects/create", ws.projectCreateController).Methods(http.MethodPost)

	r.HandleFunc("/projects/details/{projectId}", ws.projectDetailController).Methods(http.MethodGet)

	r.HandleFunc("/projects/details/{projectId}/suites/create", func(w http.ResponseWriter, r *http.Request) {
		ws.execTemplateHandler("suite_create", w, r)
	}).Methods(http.MethodGet)
	r.HandleFunc("/projects/details/{projectId}/suites/create", ws.suiteCreateController).Methods(http.MethodPost)
	r.HandleFunc("/projects/details/{projectId}/suites/details/{suiteId}", ws.suiteDetailController).Methods(http.MethodGet)

	r.HandleFunc("/testruns/list", func(w http.ResponseWriter, r *http.Request) {
		ws.execTemplateHandler("testruns_list", w, r)
	}).Methods(http.MethodGet)

	r.Use(ws.authMiddleware)

	return nil
}
