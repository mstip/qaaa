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
	store     *store.Store
	srv       *http.Server
	templates map[string]*template.Template
}

func NewWebServer(s *store.Store) (*WebServer, error) {
	ws := &WebServer{
		store:     s,
		templates: map[string]*template.Template{},
	}

	ws.templates["index"] = template.Must(template.ParseFS(templateFiles, "template/index.html", "template/base.html"))

	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/", http.FileServer(http.FS(staticFiles))))
	r.HandleFunc("/", ws.index).Methods("GET")

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

func (ws *WebServer) index(w http.ResponseWriter, r *http.Request) {
	ws.store.IncCounter()
	counter := ws.store.Counter()
	ws.templates["index"].Execute(w, struct{ Counter int }{Counter: counter})
}
