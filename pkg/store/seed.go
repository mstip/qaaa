package store

import (
	"github.com/mstip/qaaa/pkg/model"
	"github.com/mstip/qaaa/pkg/task"
)

func (s *Store) seedWithDemoData() {
	s.dataLock.Lock()
	defer s.dataLock.Unlock()
	arrayLengthCheck := task.JsonCheck{Type: task.JsonTaskTypeArrayLength, Value: 200}

	getAllJsonApiTask, _ := task.NewJsonapiTask("GET", "https://jsonplaceholder.typicode.com/todos", 200, []task.JsonCheck{arrayLengthCheck})
	s.tasks = append(s.tasks, model.Task{Id: s.newTaskId(), Name: "Get all todos", Description: "", Type: model.TaskTypeJsonApi, Task: getAllJsonApiTask})

	idCheck := task.JsonCheck{Type: task.JsonTaskTypeContains, Key: "id", Value: float64(1)}
	titleCheck := task.JsonCheck{Type: task.JsonTaskTypeContains, Key: "title", Value: "delectus aut autem"}
	completedCheck := task.JsonCheck{Type: task.JsonTaskTypeContains, Key: "completed", Value: false}
	completedNotCheck := task.JsonCheck{Type: task.JsonTaskTypeContainsNot, Key: "completed", Value: true}

	demoProject := model.Project{Id: s.newProjectId(), Name: "demo project 1", Description: "This is the test and demo project"}
	demoSuite := model.Suite{Id: s.newSuiteId(), Name: "Todoapi", Description: "Todo Rest api crud test", ProjectId: demoProject.Id}
	s.suites = append(s.suites, demoSuite)

	getJsonApiTask, _ := task.NewJsonapiTask("GET", "https://jsonplaceholder.typicode.com/todos/1", 200, []task.JsonCheck{idCheck, titleCheck, completedCheck, completedNotCheck})
	s.tasks = append(s.tasks, model.Task{Id: s.newTaskId(), SuiteId: demoSuite.Id, Name: "Get one todo", Description: "-", Type: model.TaskTypeJsonApi, Task: getJsonApiTask})

	s.projects = append(s.projects, demoProject)
	s.projects = append(s.projects, model.Project{Id: s.newProjectId(), Name: "demo project 2", Description: "This is the test and demo project 2 "})
	s.projects = append(s.projects, model.Project{Id: s.newProjectId(), Name: "demo project 3", Description: "This is the test and demo project 3 "})
	s.projects = append(s.projects, model.Project{Id: s.newProjectId(), Name: "demo project 4", Description: "This is the test and demo project 4"})
}
