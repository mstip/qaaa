package task

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/mstip/qaaa/internal"
)

func JsonApiTask(request WebTaskRequest) (WebTaskResult, error) {
	resp, err := httpRequest(request)
	if err != nil {
		return WebTaskResult{Success: false, Detail: err.Error()}, nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return WebTaskResult{Success: false, Detail: err.Error()}, err
	}
	respBody := string(body)

	if len(request.Checks) > 0 {
		for _, check := range request.Checks {
			val, err := internal.PathValFromJson(check.Selector, body)
			if err != nil {
				return WebTaskResult{Success: false, Detail: err.Error(), ResponseBody: respBody}, err
			}

			if check.Type == WebTaskSelectorCheck_Equals {
				if val != check.Value {
					return WebTaskResult{Success: false, Detail: fmt.Sprintf("selector: %s - equals - expected %v got %v ", check.Selector, check.Value, val), ResponseBody: respBody}, nil
				}
			} else if check.Type == WebTaskSelectorCheck_EqualsNot {
				if val == check.Value {
					return WebTaskResult{Success: false, Detail: fmt.Sprintf("selector: %s - equals not - expected %v got %v ", check.Selector, check.Value, val), ResponseBody: respBody}, nil
				}
			} else if check.Type == WebTaskSelectorCheck_Contains || check.Type == WebTaskSelectorCheck_ContainsNot {
				valString, ok := val.(string)
				if !ok {
					return WebTaskResult{Success: false, Detail: fmt.Sprintf("val %v is not a string ", val), ResponseBody: respBody}, nil
				}
				checkValString, ok := check.Value.(string)
				if !ok {
					return WebTaskResult{Success: false, Detail: fmt.Sprintf("val %v is not a string ", check.Value), ResponseBody: respBody}, nil
				}
				if check.Type == WebTaskSelectorCheck_Contains {
					if !strings.Contains(valString, checkValString) {
						return WebTaskResult{Success: false, Detail: fmt.Sprintf("selector: %s - contains - expected %s contains %s ", check.Selector, valString, checkValString), ResponseBody: respBody}, nil
					}
				} else {
					if strings.Contains(valString, checkValString) {
						return WebTaskResult{Success: false, Detail: fmt.Sprintf("selector: %s - contains not - expected %s contains %s ", check.Selector, valString, checkValString), ResponseBody: respBody}, nil
					}
				}

			} else if check.Type == WebTaskSelectorCheck_Count {
				arrVal, ok := val.(map[string]interface{})
				if !ok {
					return WebTaskResult{Success: false, Detail: fmt.Sprintf("val %v is not an array ", val), ResponseBody: respBody}, nil
				}

				if len(arrVal) != check.Value {
					return WebTaskResult{Success: false, Detail: fmt.Sprintf("selector: %s - count - expected %d got %d", check.Selector, check.Value, len(arrVal)), ResponseBody: respBody}, nil
				}
			} else {
				err = fmt.Errorf("unkown check type %s", check.Type)
				return WebTaskResult{Success: false, Detail: err.Error(), ResponseBody: respBody}, err
			}
		}
	}
	return WebTaskResult{Success: true, Detail: "success", ResponseBody: respBody}, nil
}
