package tutils

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"

	"github.com/gorilla/mux"
)

func FormPost(formParams map[string]string, urlParams map[string]string) *http.Request {
	form := url.Values{}

	for k, v := range formParams {
		form.Add(k, v)
	}

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if urlParams != nil {
		req = mux.SetURLVars(req, urlParams)
	}
	return req
}
