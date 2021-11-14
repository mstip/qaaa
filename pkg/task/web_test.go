package task

import "testing"

func TestWebRequest(t *testing.T) {
	success, err := WebTask(WebTaskRequest{Method: "GET", Url: "http://example.com"})
	if err != nil {
		t.Error(err)
	}
	if !success {
		t.Error("success was false")
	}
}

func TestWebRequestWithStatusAndContains(t *testing.T) {
	success, err := WebTask(WebTaskRequest{Method: "GET", Url: "http://example.com", StatusCode: 200, Contains: "<h1>Example Domain</h1>"})
	if err != nil {
		t.Error(err)
	}
	if !success {
		t.Error("success was false")
	}
}

func TestWebRequestWithWrongStatus(t *testing.T) {
	success, err := WebTask(WebTaskRequest{Method: "GET", Url: "http://example.com", StatusCode: 404})
	if err != nil {
		t.Error(err)
	}
	if success {
		t.Error(success)
	}
}

func TestWebRequestWithWrongContains(t *testing.T) {
	success, err := WebTask(WebTaskRequest{Method: "GET", Url: "http://example.com", Contains: "das gibts hier nicht"})
	if err != nil {
		t.Error(err)
	}
	if success {
		t.Error(success)
	}
}
