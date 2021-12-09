package web

import (
	"net/http"
	"strings"

	"github.com/mstip/qaaa/pkg/waffel"
)

func noCacheMiddleware(wr http.ResponseWriter, r *http.Request, next http.Handler, wf *waffel.Waffel) {
	wr.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate") // HTTP 1.1.
	wr.Header().Set("Pragma", "no-cache")                                   // HTTP 1.0.
	wr.Header().Set("Expires", "0")                                         // Proxies.
	next.ServeHTTP(wr, r)
}

func authMiddleware(wr http.ResponseWriter, r *http.Request, next http.Handler, wf *waffel.Waffel) {
	if r.URL.Path == "/" || r.URL.Path == "/login" || strings.HasPrefix(r.URL.Path, "/static/") {
		next.ServeHTTP(wr, r)
		return
	}

	s, err := wf.GetService("session")
	if err != nil {
		http.Error(wr, err.Error(), http.StatusInternalServerError)
		return
	}
	session, ok := s.(*waffel.Session)
	if !ok {
		http.Error(wr, "service cast failed", http.StatusInternalServerError)
		return
	}

	if !session.IsAuthenticated(r) {
		err := session.AddFlash("please login", "danger", wr, r)
		if err != nil {
			http.Error(wr, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(wr, r, "/", http.StatusMovedPermanently)
		return
	}
	next.ServeHTTP(wr, r)
}
