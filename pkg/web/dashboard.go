package web

import (
	"net/http"

	"github.com/mstip/qaaa/pkg/waffel"
)

func dashboardController(wr http.ResponseWriter, r *http.Request, wf *waffel.Waffel) {
	wf.Render(wr, r, "dashboard", map[string]interface{}{
		"Breadcrumb": []Breadcrumb{
			{Name: "Dashboard", Route: ""},
		},
	})
}
