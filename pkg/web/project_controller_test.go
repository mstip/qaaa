package web

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/gorilla/mux"
	"github.com/mstip/qaaa/internal/tutils"
	"github.com/mstip/qaaa/pkg/store"
)

func TestProjectsListController(t *testing.T) {
	ws, err := NewWebServer(store.NewStore())
	if err != nil {
		log.Fatal(err)
	}

	wr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	projectsListController(wr, req, ws)

	tutils.EqualI(t, http.StatusOK, wr.Code, "status code")

	doc, err := goquery.NewDocumentFromReader(wr.Body)
	if err != nil {
		t.Error(err)
	}

	tutils.EqualS(t, "Projects", doc.Find("h4").First().Text(), "header")

	tableBody := doc.Find("tbody").First().Text()

	projectNames := []string{"demo project 1", "demo project 2", "demo project 3", "demo project 4"}

	for _, projectName := range projectNames {
		if !strings.Contains(tableBody, projectName) {
			t.Error("could not find project " + projectName)
		}
	}

	doc.Find("tbody a").Each(func(i int, el *goquery.Selection) {
		href, _ := el.Attr("href")
		tutils.EqualS(t, "/projects/details/"+strconv.Itoa(i), href, "project link")
	})
}

func TestProjectCreateController(t *testing.T) {
	store := store.NewStore()
	ws, err := NewWebServer(store)
	if err != nil {
		log.Fatal(err)
	}
	wr := httptest.NewRecorder()

	req := tutils.FormPost(map[string]string{"name": "testproject", "description": "description of proj"}, nil)

	projectCreateController(wr, req, ws)

	tutils.EqualI(t, http.StatusMovedPermanently, wr.Code, "status code")

	url, _ := wr.Result().Location()
	tutils.EqualS(t, "/projects/details/4", url.Path, "location")

	p := store.GetProjectById(4)
	if p == nil {
		t.Error("could not find new project")
	}

	tutils.EqualI(t, 4, int(p.Id), "project id")
	tutils.EqualS(t, "testproject", p.Name, "project name")
	tutils.EqualS(t, "description of proj", p.Description, "project description")
}

func TestProjectDetailController(t *testing.T) {
	store := store.NewStore()
	ws, err := NewWebServer(store)
	if err != nil {
		log.Fatal(err)
	}

	type test struct {
		projectId      string
		expectedStatus int
	}

	tests := []test{
		{projectId: "0", expectedStatus: http.StatusOK},
		{projectId: "1337", expectedStatus: http.StatusNotFound},
	}

	for _, tc := range tests {
		wr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/projects/details/"+tc.projectId, nil)
		req = mux.SetURLVars(req, map[string]string{"projectId": tc.projectId})

		projectDetailController(wr, req, ws)

		tutils.EqualI(t, tc.expectedStatus, wr.Code, "status code")

		pId, err := strconv.ParseUint(string(tc.projectId), 10, 64)
		if err != nil {
			t.Error(err)
		}
		project := store.GetProjectById(pId)
		if tc.expectedStatus == http.StatusOK && project == nil {
			t.Errorf("could not find project %d in store", pId)
		}

		if tc.expectedStatus != http.StatusOK && project != nil {
			t.Errorf("found project %d in store but it should not exist", pId)
		}
		if tc.expectedStatus != http.StatusOK {
			continue
		}

		doc, err := goquery.NewDocumentFromReader(wr.Body)
		if err != nil {
			t.Error(err)
		}

		tutils.EqualS(t, project.Name, doc.Find("h4.card-title").First().Text(), "project title")
		tutils.EqualS(t, project.Description, doc.Find(".card p").First().Text(), "project title")

		doc.Find(".card-body > table > tbody > tr ").Each(
			func(suiteIndex int, suiteRow *goquery.Selection) {
				if suiteIndex%2 == 0 {
					tutils.EqualS(
						t,
						project.Suites[suiteIndex].Name,
						suiteRow.Children().Eq(0).Text(),
						fmt.Sprintf("suite %d name", suiteIndex),
					)

					tutils.EqualS(
						t,
						project.Suites[suiteIndex].Description,
						suiteRow.Children().Eq(1).Text(),
						fmt.Sprintf("suite %d description", suiteIndex),
					)

					suiteDetailHref := fmt.Sprintf(
						"/projects/details/%s/suites/details/%s",
						tc.projectId,
						strconv.FormatUint(project.Suites[suiteIndex].Id, 10),
					)
					href, _ := suiteRow.Find("a").First().Attr("href")
					tutils.EqualS(t, suiteDetailHref, href, fmt.Sprintf("suite %d link", suiteIndex))
				} else {
					suiteRow.Find("tbody > tr ").Each(func(taskIndex int, taskRow *goquery.Selection) {
						if taskIndex == 0 {
							return
						}
						task := project.Suites[suiteIndex-1].Tasks[taskIndex-1]
						tutils.EqualS(t, task.Type, taskRow.Children().Eq(0).Text(), "task type")
						tutils.EqualS(t, task.Name, taskRow.Children().Eq(1).Text(), "task name")
						tutils.EqualS(t, task.Description, taskRow.Children().Eq(2).Text(), "task description")
					})
				}
			})
	}
}
