package task

import (
	"encoding/json"
	"log"
	"io/ioutil"
	"net/http"
)

type JsonApiTaskRequest struct {
	Method     string
	Url        string
	StatusCode int
	Contains   map[string]interface{}
}

func JsonApiTask(request JsonApiTaskRequest) (bool, error) {
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

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	var jsonResponse map[string]interface{}

	err = json.Unmarshal(body, &jsonResponse)
	if err != nil {
		return false, err
	}

	for k, v := range request.Contains {
		if jsonResponse[k] != v {
			log.Println("json missmatch", k, v)
			return false, nil
		}
	}

	return true, nil
}
