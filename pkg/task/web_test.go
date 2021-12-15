package task

import (
	"testing"

	"github.com/mstip/qaaa/internal/tutils"
)

func TestWebRequestStatusCode(t *testing.T) {
	testCases := []struct {
		desc       string
		success    bool
		url        string
		statusCode int
	}{
		{
			desc:       "success",
			success:    true,
			url:        "http://example.com",
			statusCode: 200,
		},
		{
			desc:       "wrong status",
			success:    false,
			url:        "http://example.com",
			statusCode: 404,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result, err := WebTask(
				WebTaskRequest{
					Method: "GET", Url: tC.url, StatusCode: tC.statusCode,
				},
			)
			if err != nil {
				t.Fatal(err)
			}
			tutils.EqualB(t, tC.success, result.Success, "request success")
		})
	}

}

func TestWebTaskChecks(t *testing.T) {
	testCases := []struct {
		desc    string
		success bool
		checks  []WebTaskCheck
	}{
		{
			desc:    "header equals",
			success: true,
			checks: []WebTaskCheck{
				{Type: WebTaskSelectorCheck_Equals, Selector: "h1", Value: "Example Domain", Key: "text"},
			},
		},
		{
			desc:    "header equals fail",
			success: false,
			checks: []WebTaskCheck{
				{Type: WebTaskSelectorCheck_Equals, Selector: "h1", Value: "wooo", Key: "text"},
			},
		},
		{
			desc:    "header equals not",
			success: true,
			checks: []WebTaskCheck{
				{Type: WebTaskSelectorCheck_EqualsNot, Selector: "h1", Value: "wooob", Key: "text"},
			},
		},
		{
			desc:    "header equals not fail",
			success: false,
			checks: []WebTaskCheck{
				{Type: WebTaskSelectorCheck_EqualsNot, Selector: "h1", Value: "Example Domain", Key: "text"},
			},
		},
		{
			desc:    "header contains",
			success: true,
			checks: []WebTaskCheck{
				{Type: WebTaskSelectorCheck_Contains, Selector: "h1", Value: "ample", Key: "text"},
			},
		},
		{
			desc:    "header contains fail",
			success: false,
			checks: []WebTaskCheck{
				{Type: WebTaskSelectorCheck_Contains, Selector: "h1", Value: "wooo", Key: "text"},
			},
		},
		{
			desc:    "header contains not",
			success: true,
			checks: []WebTaskCheck{
				{Type: WebTaskSelectorCheck_ContainsNot, Selector: "h1", Value: "wooob", Key: "text"},
			},
		},
		{
			desc:    "header contains not fail",
			success: false,
			checks: []WebTaskCheck{
				{Type: WebTaskSelectorCheck_ContainsNot, Selector: "h1", Value: "ample", Key: "text"},
			},
		},
		{
			desc:    "paragraph count",
			success: true,
			checks: []WebTaskCheck{
				{Type: WebTaskSelectorCheck_Count, Selector: "p", Count: 2},
			},
		},
		{
			desc:    "paragraph count fail",
			success: false,
			checks: []WebTaskCheck{
				{Type: WebTaskSelectorCheck_Count, Selector: "p", Count: 5},
			},
		},
		{
			desc:    "link text",
			success: true,
			checks: []WebTaskCheck{
				{Type: WebTaskSelectorCheck_Equals, Selector: "a", Value: "More information...", Key: "text"},
			},
		},
		{
			desc:    "link text fail",
			success: false,
			checks: []WebTaskCheck{
				{Type: WebTaskSelectorCheck_Equals, Selector: "a", Value: "wwww", Key: "text"},
			},
		},
		{
			desc:    "link href equals",
			success: true,
			checks: []WebTaskCheck{
				{Type: WebTaskSelectorCheck_Equals, Selector: "a", Value: "https://www.iana.org/domains/example", Key: "href"},
			},
		},
		{
			desc:    "link href equals fail",
			success: false,
			checks: []WebTaskCheck{
				{Type: WebTaskSelectorCheck_Equals, Selector: "a", Value: "https://www.iana.org/domains/examp", Key: "href"},
			},
		},
		{
			desc:    "link href equals not",
			success: true,
			checks: []WebTaskCheck{
				{Type: WebTaskSelectorCheck_EqualsNot, Selector: "a", Value: "https://heise.de", Key: "href"},
			},
		},
		{
			desc:    "link href equals not fail",
			success: false,
			checks: []WebTaskCheck{
				{Type: WebTaskSelectorCheck_EqualsNot, Selector: "a", Value: "https://www.iana.org/domains/example", Key: "href"},
			},
		},
		{
			desc:    "link href contains",
			success: true,
			checks: []WebTaskCheck{
				{Type: WebTaskSelectorCheck_Contains, Selector: "a", Value: "domains/ex", Key: "href"},
			},
		},
		{
			desc:    "link href contains fail",
			success: false,
			checks: []WebTaskCheck{
				{Type: WebTaskSelectorCheck_Contains, Selector: "a", Value: "heise", Key: "href"},
			},
		},
		{
			desc:    "link href contains not",
			success: true,
			checks: []WebTaskCheck{
				{Type: WebTaskSelectorCheck_ContainsNot, Selector: "a", Value: "heise", Key: "href"},
			},
		},
		{
			desc:    "link href contains not fail",
			success: false,
			checks: []WebTaskCheck{
				{Type: WebTaskSelectorCheck_ContainsNot, Selector: "a", Value: "domains/ex", Key: "href"},
			},
		},
		{
			desc:    "link unknown attr",
			success: false,
			checks: []WebTaskCheck{
				{Type: WebTaskSelectorCheck_ContainsNot, Selector: "a", Value: "domains/ex", Key: "unknown"},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result, err := WebTask(
				WebTaskRequest{
					Checks: tC.checks,
					Method: "GET", Url: "http://example.com", StatusCode: 200,
				},
			)
			if err != nil {
				t.Fatal(err)
			}
			tutils.EqualB(t, tC.success, result.Success, result.Detail)
		})
	}
}
