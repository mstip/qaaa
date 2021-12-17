package task

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func TaskTestRun(taskType string, method string, url string) (interface{}, error) {
	// TODO: other methods are missing
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	// TODO: move this to a helper or util
	defer resp.Body.Close()
	raw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	jsonMap := map[string]interface{}{}
	jsonArray := []interface{}{}
	err = json.Unmarshal(raw, &jsonMap)
	if err != nil {
		err = json.Unmarshal(raw, &jsonArray)
		if err != nil {
			return nil, err
		}

		return jsonArray, nil
	}

	return jsonMap, nil
}
