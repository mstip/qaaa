package web

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/gorilla/mux"
	"github.com/mstip/qaaa/internal/tutils"
	"github.com/mstip/qaaa/pkg/store"
)

func TestSuiteCreateController(t *testing.T) {
	store := store.NewStore()
	ws, err := NewWebServer(store)
	if err != nil {
		log.Fatal(err)
	}
	wr := httptest.NewRecorder()

	req := tutils.FormPost(
		map[string]string{"name": "testsuite", "description": "description of suite"},
		map[string]string{"projectId": "3"},
	)

	suiteCreateController(wr, req, ws)

	tutils.EqualI(t, http.StatusMovedPermanently, wr.Code, "status code")

	url, _ := wr.Result().Location()
	tutils.EqualS(t, "/projects/details/3/suites/details/1", url.Path, "location")

	suites := store.GetSuitesByProjectId(3)

	tutils.EqualI(t, 1, len(suites), "suites count")
	tutils.EqualI(t, 1, int(suites[0].Id), "suite id")
	tutils.EqualS(t, "testsuite", suites[0].Name, "suite name")
	tutils.EqualS(t, "description of suite", suites[0].Description, "suite description")
}

func TestSuiteDetailController(t *testing.T) {
	store := store.NewStore()
	ws, err := NewWebServer(store)
	if err != nil {
		log.Fatal(err)
	}

	type test struct {
		suiteId        string
		expectedStatus int
	}

	tests := []test{
		{suiteId: "0", expectedStatus: http.StatusOK},
		{suiteId: "1337", expectedStatus: http.StatusNotFound},
	}

	for _, tc := range tests {
		wr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/projects/details/0/suites/details"+tc.suiteId, nil)
		req = mux.SetURLVars(req, map[string]string{"suiteId": tc.suiteId})

		suiteDetailController(wr, req, ws)

		tutils.EqualI(t, tc.expectedStatus, wr.Code, "status code")

		pId, err := strconv.ParseUint(tc.suiteId, 10, 64)
		if err != nil {
			t.Fatal(err)
		}
		suite := store.GetSuiteById(pId)
		if tc.expectedStatus == http.StatusOK && suite == nil {
			t.Fatalf("could not find suite %d in store", pId)
		}

		if tc.expectedStatus != http.StatusOK && suite != nil {
			t.Fatalf("found suite %d in store but it should not exist", pId)
		}
		if tc.expectedStatus != http.StatusOK {
			continue
		}
		doc, err := goquery.NewDocumentFromReader(wr.Body)
		if err != nil {
			t.Fatal(err)
		}

		tutils.EqualS(t, suite.Name, doc.Find("h4.card-title").First().Text(), "suite title")
		tutils.EqualS(t, suite.Description, doc.Find(".card p").First().Text(), "suite description")

		doc.Find(".card-body > table > tbody>tr ").Each(
			func(i int, row *goquery.Selection) {
				if i == 0 {
					return
				}
				sId, err := strconv.ParseUint(tc.suiteId, 10, 64)
				if err != nil {
					t.Fatal(err)
				}
				task := store.GetTasksBySuiteId(sId)[i-1]
				tutils.EqualS(t, task.Type, row.Children().Eq(0).Text(), "task type")
				tutils.EqualS(t, task.Name, row.Children().Eq(1).Text(), "task name")
				tutils.EqualS(t, task.Description, row.Children().Eq(2).Text(), "task description")
			})
	}
}
