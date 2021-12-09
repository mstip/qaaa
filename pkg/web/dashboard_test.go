package web

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mstip/qaaa/internal/tutils"
)

func TestDashboardController(t *testing.T) {
	wr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	dashboardController(wr, req, CreateWaffel())

	tutils.EqualI(t, http.StatusOK, wr.Code, "status code")

}
