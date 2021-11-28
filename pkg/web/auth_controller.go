package web

import (
	"net/http"
)

func (ws *WebServer) indexController(w http.ResponseWriter, r *http.Request) {
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

func (ws *WebServer) loginController(w http.ResponseWriter, r *http.Request) {
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
