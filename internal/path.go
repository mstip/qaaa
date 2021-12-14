package internal

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

func PathValFromJson(path string, raw []byte) (interface{}, error) {
	val := map[string]interface{}{}
	jsonArray := []interface{}{}
	err := json.Unmarshal(raw, &val)
	if err != nil {
		err = json.Unmarshal(raw, &jsonArray)
		if err != nil {
			return nil, err
		}

		for k, v := range jsonArray {
			val[strconv.Itoa(k)] = v
		}
	}

	splitedPath := strings.Split(path, ".")
	if len(splitedPath) == 1 {
		return val[splitedPath[0]], nil
	}

	cpVal := val
	ok := true
	for i, sp := range splitedPath {
		if i == len(splitedPath)-1 {
			return cpVal[sp], nil
		}
		cpVal, ok = cpVal[sp].(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("fail to convert %#v", cpVal[sp])
		}
	}

	return nil, nil
}
