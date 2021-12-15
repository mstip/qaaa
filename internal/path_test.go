package internal

import (
	"testing"

	"github.com/mstip/qaaa/internal/tutils"
)

func TestPathValFromJson(t *testing.T) {
	testCases := []struct {
		desc        string
		success     bool
		json        string
		path        string
		expectedVal interface{}
	}{
		{
			desc:        "simple number",
			success:     true,
			json:        `{"userId": 1,"id": 1,"title": "delectus aut autem","completed": false}`,
			path:        "userId",
			expectedVal: float64(1),
		},
		{
			desc:        "simple text",
			success:     true,
			json:        `{"userId": 1,"id": 1,"title": "delectus aut autem","completed": false}`,
			path:        "title",
			expectedVal: "delectus aut autem",
		},
		{
			desc:        "simple bool",
			success:     true,
			json:        `{"userId": 1,"id": 1,"title": "delectus aut autem","completed": false}`,
			path:        "completed",
			expectedVal: false,
		},
		{
			desc:        "simple array",
			success:     true,
			json:        `["1", "2", "3"]`,
			path:        "0",
			expectedVal: "1",
		},
		{
			desc:        "nested",
			success:     true,
			json:        `{"deep": {"deep": {"nested":"ok"}}}`,
			path:        "deep.deep.nested",
			expectedVal: "ok",
		},
		{
			desc:        "nested",
			success:     true,
			json:        `{"deep": {"deep": {"nested":"ok"}}}`,
			path:        "deep.deep.nested",
			expectedVal: "ok",
		},
		{
			desc:        "array with multiple types",
			success:     true,
			json:        `[1, "2", false]`,
			path:        "0",
			expectedVal: float64(1),
		},
		{
			desc:        "nested arrays",
			success:     true,
			json:        `[1, "2", {"arr": ["3","4",{"deep": ["5",6]}, "7"]}]`,
			path:        "2.arr.2.deep.1",
			expectedVal: float64(6),
		},
		{
			desc:        "nested arrays not last value",
			success:     true,
			json:        `[1, "2", {"arr": ["3","4",{"deep": ["5",6]}, "7"]}]`,
			path:        "2.arr.2.deep",
			expectedVal: []interface{}{"5", float64(6)},
		},
		{
			desc:        "get root",
			success:     true,
			json:        `{"root":"root"}`,
			path:        "",
			expectedVal: map[string]interface{}{"root": "root"},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			val, err := PathValFromJson(tC.path, []byte(tC.json))
			if tC.success {
				tutils.EqualNil(t, err, "success")
			} else {
				tutils.EqualNotNil(t, err, "success")
				return
			}

			switch exVal := tC.expectedVal.(type) {
			case []interface{}:
				for i := range exVal {
					val, ok := val.([]interface{})
					if !ok {
						t.Fatalf("could not convert %v", val)
					}
					if val[i] != exVal[i] {
						t.Fatalf("val - expected (%T)%#v got (%T)%#v", exVal[i], exVal[i], val[i], val[i])
					}
				}
			case map[string]interface{}:
				for i := range exVal {
					val, ok := val.(map[string]interface{})
					if !ok {
						t.Fatalf("could not convert %v", val)
					}
					if val[i] != exVal[i] {
						t.Fatalf("val - expected (%T)%#v got (%T)%#v", exVal[i], exVal[i], val[i], val[i])
					}
				}
			default:
				if val != tC.expectedVal {
					t.Fatalf("val - expected (%T)%#v got (%T)%#v", tC.expectedVal, tC.expectedVal, val, val)
				}
			}

		})
	}
}
