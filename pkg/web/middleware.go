package web

import (
	"net/http"
	"strings"
)

func (ws *WebServer) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" || r.URL.Path == "/login" || strings.HasPrefix(r.URL.Path, "/static/") {
			next.ServeHTTP(w, r)
			return
		}

		if !ws.sessionStore.isAuthenticated(r) {
			err := ws.sessionStore.addFlash("please login", w, r)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			http.Redirect(w, r, "/", http.StatusMovedPermanently)
			return
		}
		next.ServeHTTP(w, r)
	})
}