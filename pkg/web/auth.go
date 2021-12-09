package web

import (
	"net/http"

	"github.com/mstip/qaaa/pkg/waffel"
)

func indexController(wr http.ResponseWriter, r *http.Request, wf *waffel.Waffel) {
	wf.Render(wr, r, "index", nil)
}

func loginController(wr http.ResponseWriter, r *http.Request, wf *waffel.Waffel) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	session := wf.GetSession()

	if username == "admin" && password == "admin" {
		err := session.Authenticate(1, wr, r)
		if err != nil {
			http.Error(wr, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(wr, r, "/dashboard", http.StatusMovedPermanently)
		return
	}

	err := session.AddFlash("invalid credentials", "danger", wr, r)
	if err != nil {
		http.Error(wr, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(wr, r, "/", http.StatusMovedPermanently)
}
