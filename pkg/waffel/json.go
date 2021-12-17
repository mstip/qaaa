package waffel

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// JsonBody parase the json body of the request, if the request body is an array it fails
func JsonBody(r *http.Request) (map[string]interface{}, error) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	data := map[string]interface{}{}
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return data, nil
}

func JsonResponse(w http.ResponseWriter, data interface{}) error {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
	return nil
}
