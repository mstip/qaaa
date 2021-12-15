package task

import (
	"testing"

	"github.com/mstip/qaaa/internal/tutils"
)

func TestJsonTaskChecks(t *testing.T) {
	testCases := []struct {
		desc    string
		url     string
		success bool
		checks  []WebTaskCheck
	}{
		{
			desc:    "get one todo userid",
			success: true,
			url:     "https://jsonplaceholder.typicode.com/todos/1",
			checks: []WebTaskCheck{
				{Type: WebTaskSelectorCheck_Equals, Selector: "userId", Value: float64(1)},
			},
		},
		{
			desc:    "get one todo userid - fail",
			success: false,
			url:     "https://jsonplaceholder.typicode.com/todos/1",
			checks: []WebTaskCheck{
				{Type: WebTaskSelectorCheck_Equals, Selector: "userId", Value: float64(2)},
			},
		},
		{
			desc:    "get from all todos specific title",
			success: true,
			url:     "https://jsonplaceholder.typicode.com/todos",
			checks: []WebTaskCheck{
				{Type: WebTaskSelectorCheck_Equals, Selector: "5.title", Value: "qui ullam ratione quibusdam voluptatem quia omnis"},
				{Type: WebTaskSelectorCheck_Equals, Selector: "5.id", Value: float64(6)},
				{Type: WebTaskSelectorCheck_Equals, Selector: "5.completed", Value: false},
			},
		},
		{
			desc:    "not get one todo userid",
			success: true,
			url:     "https://jsonplaceholder.typicode.com/todos/1",
			checks: []WebTaskCheck{
				{Type: WebTaskSelectorCheck_EqualsNot, Selector: "userId", Value: float64(5)},
			},
		},
		{
			desc:    "get from all todos specific title contains",
			success: true,
			url:     "https://jsonplaceholder.typicode.com/todos",
			checks: []WebTaskCheck{
				{Type: WebTaskSelectorCheck_Contains, Selector: "5.title", Value: "quibusdam voluptatem"},
			},
		},
		{
			desc:    "fail get from all todos specific title contains",
			success: false,
			url:     "https://jsonplaceholder.typicode.com/todos",
			checks: []WebTaskCheck{
				{Type: WebTaskSelectorCheck_Contains, Selector: "5.title", Value: "not here"},
			},
		},
		{
			desc:    "fail get from all todos specific title contains not",
			success: false,
			url:     "https://jsonplaceholder.typicode.com/todos",
			checks: []WebTaskCheck{
				{Type: WebTaskSelectorCheck_ContainsNot, Selector: "5.title", Value: "quibusdam voluptatem"},
			},
		},
		{
			desc:    "get from all todos specific title contains not",
			success: true,
			url:     "https://jsonplaceholder.typicode.com/todos",
			checks: []WebTaskCheck{
				{Type: WebTaskSelectorCheck_ContainsNot, Selector: "5.title", Value: "not here"},
			},
		},
		{
			desc:    "get count from all todos ",
			success: true,
			url:     "https://jsonplaceholder.typicode.com/todos",
			checks: []WebTaskCheck{
				{Type: WebTaskSelectorCheck_Count, Selector: "", Value: 200},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result, err := JsonApiTask(
				WebTaskRequest{
					Checks: tC.checks,
					Method: "GET", Url: tC.url, StatusCode: 200,
				},
			)
			if err != nil {
				t.Fatal(err)
			}
			tutils.EqualB(t, tC.success, result.Success, result.Detail)
		})
	}
}
