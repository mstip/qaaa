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
			if val != tC.expectedVal {
				t.Fatalf("val - expected (%T)%#v got (%T)%#v", tC.expectedVal, tC.expectedVal, val, val)
			}
		})
	}
}
