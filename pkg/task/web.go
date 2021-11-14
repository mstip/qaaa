package task

import (
	"log"
	"io/ioutil"
	"net/http"
	"strings"
)

type WebTaskRequest struct {
	Method     string
	Url        string
	StatusCode int
	Contains   string
}

func WebTask(request WebTaskRequest) (bool, error) {
	client := &http.Client{}
	req, err := http.NewRequest(request.Method, request.Url, nil)
	if err != nil {
		return false, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}

	if request.StatusCode != 0 {
		if request.StatusCode != resp.StatusCode {
			log.Println("status code missmatch", request.StatusCode, resp.StatusCode)
			return false, nil
		}
	}

	if request.Contains != "" {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return false, err
		}
		if !strings.Contains(string(body), request.Contains) {
			log.Println("body doesnt contain string", request.Contains)
			return false, nil
		}
	}

	return true, nil
}
