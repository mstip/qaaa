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
	httpHandler        *mux.Router
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
		httpHandler:        mux.NewRouter(),
	}

	err = ws.registerTemplates()
	if err != nil {
		return nil, err
	}

	// ws.handler = mux.NewRouter()
	ws.httpHandler.PathPrefix("/static/").Handler(http.StripPrefix("/", http.FileServer(http.FS(staticFiles))))
	err = ws.registerRoutes(ws.httpHandler)
	if err != nil {
		return nil, err
	}

	ws.srv = &http.Server{
		Handler:      ws.httpHandler,
		Addr:         "0.0.0.0:3000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	return ws, nil
}

func (ws *WebServer) HttpHandler() http.Handler {
	return ws.httpHandler
}

func (ws *WebServer) RunAndServe() {
	log.Fatal(ws.srv.ListenAndServe())
}
