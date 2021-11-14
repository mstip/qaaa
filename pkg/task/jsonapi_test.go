package task

import "testing"

func TestJsonApiTask(t *testing.T) {
	success, err := JsonApiTask(
		JsonApiTaskRequest{
			Method: "GET",
			Url:    "https://jsonplaceholder.typicode.com/todos/1",
			Contains: map[string]interface{}{"title": "delectus aut autem", "completed": false},
		},
	)
	if err != nil {
		t.Error(err)
	}
	if !success {
		t.Error(success)
	}
}

func TestJsonApiTaskWrongResponse(t *testing.T) {
	success, err := JsonApiTask(
		JsonApiTaskRequest{
			Method:   "GET",
			Url:      "https://jsonplaceholder.typicode.com/todos/1",
			Contains: map[string]interface{}{"title": "delectus aut autem", "completed": true, "unknown": 1337},
		},
	)
	if err != nil {
		t.Error(err)
	}
	if success {
		t.Error(success)
	}
}
