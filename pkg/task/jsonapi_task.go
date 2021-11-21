package task

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	JsonTaskTypeContains    = "contains"
	JsonTaskTypeContainsNot = "containsNot"
	JsonTaskTypeArrayLength = "arrayLength"
)

type JsonCheck struct {
	Type  string
	Key   string
	Value interface{}
}

type JsonapiTask struct {
	Method             string
	Url                string
	ExpectedStatusCode int
	ExpectedJson       []JsonCheck
	statusCode         int
	jsonBody           map[string]interface{}
	jsonBodyArray      []interface{}
}

func NewJsonapiTask(method string, url string, expectedStatusCode int, expectedJson []JsonCheck) (*JsonapiTask, error) {
	return &JsonapiTask{Method: method, Url: url, ExpectedStatusCode: expectedStatusCode, ExpectedJson: expectedJson}, nil
}

func (t *JsonapiTask) Arrange() error {
	return nil
}

func (t *JsonapiTask) Act() error {
	client := &http.Client{}
	req, err := http.NewRequest(t.Method, t.Url, nil)
	if err != nil {
		log.Println(err)
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return err
	}

	t.statusCode = resp.StatusCode

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return err
	}

	err = json.Unmarshal(body, &t.jsonBody)
	if err != nil {
		errArray := json.Unmarshal(body, &t.jsonBodyArray)
		if errArray != nil {
			log.Println(err)
			log.Println(errArray)
			return errors.New("can not unmarshal invalid json")
		}
	}
	return nil
}

func (t *JsonapiTask) Assert() (bool, error) {
	if t.ExpectedStatusCode != 0 {
		if t.statusCode != t.ExpectedStatusCode {
			log.Println("status code missmatch", t.statusCode, t.ExpectedStatusCode)
			return false, nil
		}
	}

	for _, expected := range t.ExpectedJson {
		if expected.Type == JsonTaskTypeArrayLength {
			if len(t.jsonBodyArray) != expected.Value {
				log.Printf("length is %v expected %v", len(t.jsonBodyArray), expected.Value)
				return false, nil
			}
		}

		if len(t.jsonBodyArray) != 0 {
			continue
		}
		val := t.jsonBody[expected.Key]
		if expected.Type == JsonTaskTypeContains {
			if val != expected.Value {
				log.Printf("key %s got %T %v  expected %T %v", expected.Key, val, val, expected.Value, expected.Value)
				return false, nil
			}
		}

		if expected.Type == JsonTaskTypeContainsNot {
			if val == expected.Value {
				log.Printf("key %s got %T %v  not expected %T %v", expected.Key, val, val, expected.Value, expected.Value)
				return false, nil
			}
		}
	}

	return true, nil
}
