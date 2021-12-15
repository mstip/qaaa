package internal

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

func convertArrToMap(arr []interface{}) map[string]interface{} {
	m := map[string]interface{}{}
	for k, v := range arr {
		m[strconv.Itoa(k)] = v
	}
	return m
}

func PathValFromJson(path string, raw []byte) (interface{}, error) {
	val := map[string]interface{}{}
	jsonArray := []interface{}{}
	err := json.Unmarshal(raw, &val)
	if err != nil {
		err = json.Unmarshal(raw, &jsonArray)
		if err != nil {
			return nil, err
		}

		val = convertArrToMap(jsonArray)
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
		origin := cpVal[sp]
		cpVal, ok = cpVal[sp].(map[string]interface{})
		if !ok {
			arrVal, ok := origin.([]interface{})
			if !ok {
				return nil, fmt.Errorf("fail to convert %#v", cpVal[sp])
			}

			cpVal = convertArrToMap(arrVal)
		}
	}

	return nil, fmt.Errorf("fail %#v", cpVal)
}
