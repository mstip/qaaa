package web

import (
	"log"
	"net/http"
)

func indexController(w http.ResponseWriter, r *http.Request, ws *WebServer) {
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

func loginController(w http.ResponseWriter, r *http.Request, ws *WebServer) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	log.Println("hier", username, password)

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
