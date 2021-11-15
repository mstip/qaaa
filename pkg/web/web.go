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

	err = ws.registerTemplates()
	if err != nil {
		return nil, err
	}

	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/", http.FileServer(http.FS(staticFiles))))
	err = ws.registerRoutes(r)
	if err != nil {
		return nil, err
	}

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
