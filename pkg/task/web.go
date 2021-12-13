package task

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	WebTaskSelectorCheck_Equals      = "equals"
	WebTaskSelectorCheck_EqualsNot   = "equals_not"
	WebTaskSelectorCheck_Contains    = "contains"
	WebTaskSelectorCheck_ContainsNot = "contains_not"
	WebTaskSelectorCheck_Count       = "count"
)

type WebTaskCheck struct {
	Type     string
	Selector string
	Count    int
	Key      string
	Value    string
}

type WebTaskRequest struct {
	Method      string
	Url         string
	StatusCode  int
	Checks      []WebTaskCheck
	ExpectedUrl string
	FormParams  map[string]interface{}
}

type WebTaskResult struct {
	Success bool
	Detail  string
}

func WebTask(request WebTaskRequest) (WebTaskResult, error) {
	client := &http.Client{}
	req, err := http.NewRequest(request.Method, request.Url, nil)
	if err != nil {
		return WebTaskResult{Success: false, Detail: err.Error()}, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return WebTaskResult{Success: false, Detail: err.Error()}, err
	}

	if request.StatusCode != 0 {
		if request.StatusCode != resp.StatusCode {
			log.Println("status code missmatch", request.StatusCode, resp.StatusCode)
			return WebTaskResult{Success: false, Detail: fmt.Sprintf("status code - expected %d got %d", request.StatusCode, resp.StatusCode)}, nil
		}
	}

	if len(request.Checks) > 0 {
		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			return WebTaskResult{Success: false, Detail: err.Error()}, err
		}
		for _, check := range request.Checks {
			if check.Type == WebTaskSelectorCheck_Equals || check.Type == WebTaskSelectorCheck_EqualsNot ||
				check.Type == WebTaskSelectorCheck_Contains || check.Type == WebTaskSelectorCheck_ContainsNot {
				if err = checkKeyValGeneric(check, doc); err != nil {
					return WebTaskResult{Success: false, Detail: err.Error()}, nil
				}
			} else if check.Type == WebTaskSelectorCheck_Count {
				if err = checkCount(check, doc); err != nil {
					return WebTaskResult{Success: false, Detail: err.Error()}, nil
				}
			} else {
				err = fmt.Errorf("unkown web check type %s", check.Type)
				return WebTaskResult{Success: false, Detail: err.Error()}, err
			}
		}
	}
	return WebTaskResult{Success: true, Detail: "success"}, nil
}

func checkKeyValGeneric(check WebTaskCheck, doc *goquery.Document) error {
	var selectorAttrValue string
	var exists bool
	if check.Key == "text" {
		selectorAttrValue = doc.Find(check.Selector).First().Text()
	} else {
		selectorAttrValue, exists = doc.Find(check.Selector).First().Attr(check.Key)
		if !exists {
			return fmt.Errorf("selector: %s - equals - attr %s could not be found ", check.Selector, check.Key)
		}
	}

	switch check.Type {
	case WebTaskSelectorCheck_Equals:
		if selectorAttrValue != check.Value {
			return fmt.Errorf("selector: %s - equals - expected attr %s %s got %s ", check.Selector, check.Key, check.Value, selectorAttrValue)
		}
	case WebTaskSelectorCheck_EqualsNot:
		if selectorAttrValue == check.Value {
			return fmt.Errorf("selector: %s - equals not - expected attr %s %s got %s ", check.Selector, check.Key, check.Value, selectorAttrValue)
		}
	case WebTaskSelectorCheck_Contains:
		if !strings.Contains(selectorAttrValue, check.Value) {
			return fmt.Errorf("selector: %s - contains - expected attr %s %s contains %s ", check.Selector, check.Key, check.Value, selectorAttrValue)
		}
	case WebTaskSelectorCheck_ContainsNot:
		if strings.Contains(selectorAttrValue, check.Value) {
			return fmt.Errorf("selector: %s - contains not - expected attr %s %s contains %s ", check.Selector, check.Key, check.Value, selectorAttrValue)
		}
	}

	return nil
}

func checkCount(check WebTaskCheck, doc *goquery.Document) error {
	selectorCount := doc.Find(check.Selector).Length()
	if check.Count != selectorCount {
		return fmt.Errorf("selector: %s - count - expected %d got %d ", check.Selector, check.Count, selectorCount)
	}
	return nil
}
