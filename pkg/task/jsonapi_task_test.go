package task

import "testing"

func TestJsonapiTask(t *testing.T) {
	task, _ := NewJsonapiTask("GET", "https://jsonplaceholder.typicode.com/todos/1", 200, []JsonCheck{})
	if err := task.Arrange(); err != nil {
		t.Error(err)
	}
	if err := task.Act(); err != nil {
		t.Error(err)
	}

	success, err := task.Assert()
	if err != nil {
		t.Error(err)
	}
	if !success {
		t.Error("was not successful")
	}
}

func TestJsonapiTaskInterface(t *testing.T) {
	var task Tasker
	task, _ = NewJsonapiTask("GET", "https://jsonplaceholder.typicode.com/todos/1", 200, []JsonCheck{})
	if err := task.Arrange(); err != nil {
		t.Error(err)
	}
	if err := task.Act(); err != nil {
		t.Error(err)
	}

	success, err := task.Assert()
	if err != nil {
		t.Error(err)
	}
	if !success {
		t.Error("was not successful")
	}
}

func TestJsonapiTaskJson(t *testing.T) {
	idCheck := JsonCheck{Type: JsonTaskTypeContains, Key: "id", Value: float64(1)}
	titleCheck := JsonCheck{Type: JsonTaskTypeContains, Key: "title", Value: "delectus aut autem"}
	completedCheck := JsonCheck{Type: JsonTaskTypeContains, Key: "completed", Value: false}
	completedNotCheck := JsonCheck{Type: JsonTaskTypeContainsNot, Key: "completed", Value: true}

	task, _ := NewJsonapiTask("GET", "https://jsonplaceholder.typicode.com/todos/1", 200, []JsonCheck{idCheck, titleCheck, completedCheck, completedNotCheck})
	if err := task.Arrange(); err != nil {
		t.Error(err)
	}
	if err := task.Act(); err != nil {
		t.Error(err)
	}

	success, err := task.Assert()
	if err != nil {
		t.Error(err)
	}
	if !success {
		t.Error("was not successful")
	}
}

func TestJsonapiTaskJsonArray(t *testing.T) {
	arrayLengthCheck := JsonCheck{Type: JsonTaskTypeArrayLength, Value: 200}

	task, _ := NewJsonapiTask("GET", "https://jsonplaceholder.typicode.com/todos", 200, []JsonCheck{arrayLengthCheck})
	if err := task.Arrange(); err != nil {
		t.Error(err)
	}
	if err := task.Act(); err != nil {
		t.Error(err)
	}

	success, err := task.Assert()
	if err != nil {
		t.Error(err)
	}
	if !success {
		t.Error("was not successful")
	}
}
