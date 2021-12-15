package store

import (
	"net/http"

	"github.com/mstip/qaaa/pkg/model"
	"github.com/mstip/qaaa/pkg/task"
)

func (s *Store) seedWithDemoData() {
	s.dataLock.Lock()
	defer s.dataLock.Unlock()

	s.tasks = append(s.tasks, model.Task{
		Id:          s.newTaskId(),
		Name:        "Get all todos",
		Description: "",
		Type:        model.TaskTypeJsonApi,
		Task: task.WebTaskRequest{
			Method:     http.MethodGet,
			Url:        "https://jsonplaceholder.typicode.com/todos",
			StatusCode: 200,
			Checks: []task.WebTaskCheck{
				{Type: task.WebTaskSelectorCheck_Count, Selector: "", Value: 200},
			},
		},
	})

	demoProject := model.Project{Id: s.newProjectId(), Name: "demo project 1", Description: "This is the test and demo project"}
	demoSuite := model.Suite{Id: s.newSuiteId(), Name: "Todoapi", Description: "Todo Rest api crud test", ProjectId: demoProject.Id}
	s.suites = append(s.suites, demoSuite)

	s.tasks = append(s.tasks, model.Task{
		Id:          s.newTaskId(),
		Name:        "Get one todo",
		Description: "-",
		Type:        model.TaskTypeJsonApi,
		Task: task.WebTaskRequest{
			Method:     http.MethodGet,
			Url:        "https://jsonplaceholder.typicode.com/todos/1",
			StatusCode: 200,
			Checks: []task.WebTaskCheck{
				{Type: task.WebTaskSelectorCheck_Equals, Selector: "id", Value: float64(1)},
				{Type: task.WebTaskSelectorCheck_Equals, Selector: "title", Value: "delectus aut autem"},
				{Type: task.WebTaskSelectorCheck_Equals, Selector: "completed", Value: false},
				{Type: task.WebTaskSelectorCheck_EqualsNot, Selector: "completed", Value: true},
			},
		},
	})

	s.projects = append(s.projects, demoProject)
	s.projects = append(s.projects, model.Project{Id: s.newProjectId(), Name: "demo project 2", Description: "This is the test and demo project 2 "})
	s.projects = append(s.projects, model.Project{Id: s.newProjectId(), Name: "demo project 3", Description: "This is the test and demo project 3 "})
	s.projects = append(s.projects, model.Project{Id: s.newProjectId(), Name: "demo project 4", Description: "This is the test and demo project 4"})
}
