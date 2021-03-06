package web

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/mstip/qaaa/internal/tutils"
)

func TestGetIndex(t *testing.T) {
	wr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	indexController(wr, req, CreateWaffel())

	tutils.EqualI(t, http.StatusOK, wr.Code, "status code")

	doc, err := goquery.NewDocumentFromReader(wr.Body)
	if err != nil {
		t.Fatal(err)
	}

	tutils.EqualS(t, "Login", doc.Find("h4").First().Text(), "header")
	tutils.EqualI(t, 1, doc.Find("input#username").Length(), "username input")
	tutils.EqualI(t, 1, doc.Find("input#password").Length(), "password input")
}

func TestLogin(t *testing.T) {
	type test struct {
		username     string
		password     string
		statusCode   int
		locationPath string
	}

	tests := []test{
		{username: "admin", password: "admin", statusCode: http.StatusMovedPermanently, locationPath: "/dashboard"},
		{username: "admin", password: "wrongpass", statusCode: http.StatusMovedPermanently, locationPath: "/"},
		{username: "wronguser", password: "admin", statusCode: http.StatusMovedPermanently, locationPath: "/"},
	}

	for _, tc := range tests {
		wr := httptest.NewRecorder()
		req := tutils.FormPost(map[string]string{"username": tc.username, "password": tc.password}, nil)

		loginController(wr, req, CreateWaffel())

		tutils.EqualI(t, tc.statusCode, wr.Code, "status code")

		url, _ := wr.Result().Location()
		tutils.EqualS(t, tc.locationPath, url.Path, "location")
	}
}
