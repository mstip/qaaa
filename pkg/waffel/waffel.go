package waffel

import (
	"embed"
	"errors"
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

type Route struct {
	Name    string
	Url     string
	Method  string
	Handler func(http.ResponseWriter, *http.Request, *Waffel)
}

type Middleware struct {
	Name    string
	Handler func(http.ResponseWriter, *http.Request, http.Handler, *Waffel)
}

type Waffel struct {
	services    map[string]interface{}
	routes      []Route
	router      *mux.Router
	srv         *http.Server
	middlewares []Middleware
	templates   map[string]*template.Template
}

func NewWaffel(addr string, routes []Route, middlewares []Middleware, staticFiles, templateFiles embed.FS) (*Waffel, error) {
	// session
	s, err := newSession()
	if err != nil {
		return nil, err
	}

	// service
	waffel := &Waffel{
		services: map[string]interface{}{
			"session": s,
		},
		routes:      routes,
		middlewares: middlewares,
		router:      mux.NewRouter(),
		templates:   map[string]*template.Template{},
	}
	// templates
	if err = waffel.registerTemplates(templateFiles); err != nil {
		return nil, err
	}

	// routes
	waffel.router.PathPrefix("/static/").Handler(http.StripPrefix("/", http.FileServer(http.FS(staticFiles))))
	if err = waffel.registerRoutes(); err != nil {
		return nil, err
	}
	if err = waffel.registerMiddlewares(); err != nil {
		return nil, err
	}

	// http srv
	waffel.srv = &http.Server{
		Handler:      waffel.router,
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return waffel, nil
}

func (w *Waffel) HttpHandler() http.Handler {
	return w.router
}

func (w *Waffel) RunAndServe() {
	log.Fatal(w.srv.ListenAndServe())
}

func (w *Waffel) RegisterService(name string, service interface{}) {
	w.services[name] = service
}

func (wf *Waffel) GetService(name string) (interface{}, error) {
	if service, ok := wf.services[name]; ok {
		return service, nil
	}
	log.Printf("Service: %s not found", name)
	return nil, errors.New("service not found")
}

func (wf *Waffel) GetSession() *Session {
	s := wf.services["session"]
	session, _ := s.(*Session)
	return session
}

func (wf *Waffel) Render(wr http.ResponseWriter, r *http.Request, tmplName string, data map[string]interface{}) {
	tmpl, ok := wf.templates[tmplName]
	if !ok {
		http.Error(wr, "could not find template "+tmplName, http.StatusInternalServerError)
		return
	}

	if data == nil {
		data = map[string]interface{}{}
	}

	session := wf.GetSession()
	flashes, err := session.Flashes(wr, r)
	if err != nil {
		http.Error(wr, err.Error(), http.StatusInternalServerError)
		return
	}
	data["Flashes"] = flashes

	if err := tmpl.Execute(wr, data); err != nil {
		http.Error(wr, err.Error(), http.StatusInternalServerError)
	}
}

func (wf *Waffel) GetUrlForRoute(name string) string {
	for _, route := range wf.routes {
		if route.Name == name {
			return route.Url
		}
	}
	return ""
}

func (wf *Waffel) registerRoutes() error {
	for _, route := range wf.routes {
		route := route
		wf.router.HandleFunc(route.Url, func(w http.ResponseWriter, r *http.Request) {
			route.Handler(w, r, wf)
		}).Methods(route.Method)
	}

	return nil
}

func (wf *Waffel) registerMiddlewares() error {
	for _, mw := range wf.middlewares {
		mw := mw
		wf.router.Use(func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
				mw.Handler(wr, r, next, wf)
			})
		})
	}

	return nil
}

func (wf *Waffel) registerTemplates(templateFiles embed.FS) error {
	componentFiles, err := templateFiles.ReadDir("templates/components")
	if err != nil {
		return err
	}
	templateComponents := []string{}
	for _, cf := range componentFiles {
		templateComponents = append(templateComponents, "templates/components/"+cf.Name())
	}

	viewFiles, err := templateFiles.ReadDir("templates/views")
	if err != nil {
		return err
	}

	for _, vf := range viewFiles {
		tmpls := []string{"templates/views/" + vf.Name(), "templates/base.html"}
		tmpls = append(tmpls, templateComponents...)

		tmplKey := vf.Name()[:strings.LastIndex(vf.Name(), ".")]

		wf.templates[tmplKey] = template.Must(template.ParseFS(templateFiles, tmpls...))
	}

	return nil
}
