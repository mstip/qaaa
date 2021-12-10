package waffel

import (
	"embed"
	"log"
	"net/http"
	"testing"

	"github.com/mstip/qaaa/internal/tutils"
)

func mockHandler(rw http.ResponseWriter, r *http.Request, w *Waffel) {}

func TestGetUrlForRoute(t *testing.T) {
	routes := []Route{
		{Name: "suiteCreate", Method: http.MethodGet, Url: "/project/detail/{projectId}/suite/create", Handler: mockHandler},
		{Name: "suiteDetail", Method: http.MethodGet, Url: "/project/detail/{projectId}/suite/detail/{suiteId}", Handler: mockHandler},
	}

	waffel, err := NewWaffel("", routes, []Middleware{}, embed.FS{}, embed.FS{})
	if err != nil {
		log.Fatal(err)
		t.Fail()
	}

	tutils.EqualS(t, "/project/detail/0/suite/create", waffel.GetUrlForRoute("suiteCreate", "0"), "GetUrlForRoute suiteCreate")
	tutils.EqualS(t, "/project/detail/1/suite/detail/2", waffel.GetUrlForRoute("suiteDetail", "1", "2"), "GetUrlForRoute suiteCreate")
}
