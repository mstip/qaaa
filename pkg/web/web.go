package web

import (
	"embed"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/mstip/qaaa/pkg/store"
)

//go:embed static/*
var staticFiles embed.FS

//go:embed template/*
var templateFiles embed.FS

type WebServer struct {
	store              *store.Store
	srv                *http.Server
	templates          map[string]*template.Template
	templateComponents []string
	sessionStore       *sessionStore
}

func NewWebServer(s *store.Store) (*WebServer, error) {
	st, err := newSessionStore()
	if err != nil {
		return nil, err
	}

	ws := &WebServer{
		store:              s,
		templates:          map[string]*template.Template{},
		templateComponents: []string{},
		sessionStore:       st,
	}

	ws.registerTemplateComponent("navbar")

	ws.registerViewTemplate("index")
	ws.registerViewTemplate("dashboard")
	ws.registerViewTemplate("projects_list")
	ws.registerViewTemplate("project_details")
	ws.registerViewTemplate("testruns_list")

	r := mux.NewRouter()

	r.PathPrefix("/static/").Handler(http.StripPrefix("/", http.FileServer(http.FS(staticFiles))))

	r.HandleFunc("/", ws.index).Methods(http.MethodGet)
	r.HandleFunc("/login", ws.login).Methods(http.MethodPost)
	r.HandleFunc("/dashboard", func(w http.ResponseWriter, r *http.Request) {
		ws.execTemplateWithAuthentication("dashboard", w, r)
	}).Methods(http.MethodGet)
	r.HandleFunc("/projects/list", func(w http.ResponseWriter, r *http.Request) {
		ws.execTemplateWithAuthentication("projects_list", w, r)
	}).Methods(http.MethodGet)
	r.HandleFunc("/projects/details/{projectId}", func(w http.ResponseWriter, r *http.Request) {
		ws.execTemplateWithAuthentication("project_details", w, r)
	}).Methods(http.MethodGet)
	r.HandleFunc("/testruns/list", func(w http.ResponseWriter, r *http.Request) {
		ws.execTemplateWithAuthentication("testruns_list", w, r)
	}).Methods(http.MethodGet)

	ws.srv = &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:3000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	return ws, nil
}

func (ws *WebServer) RunAndServe() {
	log.Fatal(ws.srv.ListenAndServe())
}

func (ws *WebServer) registerTemplateComponent(templateName string) {
	ws.templateComponents = append(ws.templateComponents, "template/component/"+templateName+".html")
}

func (ws *WebServer) registerViewTemplate(templateName string) {
	tmpls := []string{"template/" + templateName + ".html", "template/base.html"}
	tmpls = append(tmpls, ws.templateComponents...)
	ws.templates[templateName] = template.Must(template.ParseFS(templateFiles, tmpls...))
}

func (ws *WebServer) index(w http.ResponseWriter, r *http.Request) {
	flashes, err := ws.sessionStore.flashes(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ws.templates["index"].Execute(w, struct {
		Flashes []interface{}
	}{
		Flashes: flashes,
	})
}

func (ws *WebServer) login(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "admin" && password == "admin" {
		err := ws.sessionStore.authenticate(1, w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/dashboard", http.StatusMovedPermanently)
		return
	}

	err := ws.sessionStore.addFlash("invalid credentials", w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func (ws *WebServer) execTemplateWithAuthentication(templateName string, w http.ResponseWriter, r *http.Request) {
	if !ws.sessionStore.isAuthenticated(r) {
		err := ws.sessionStore.addFlash("please login", w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
		return
	}

	ws.templates[templateName].Execute(w, nil)
}
