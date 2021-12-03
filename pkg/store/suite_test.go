package store

import (
	"strconv"
	"testing"

	"github.com/mstip/qaaa/internal/tutils"
)

func TestCreateAndGetSuites(t *testing.T) {
	type testCase struct {
		suiteNames        []string
		suiteDescriptions []string
	}

	hugeTc := testCase{}
	for i := 0; i < 1000; i++ {
		hugeTc.suiteNames = append(hugeTc.suiteNames, "tp-"+strconv.Itoa(i))
		hugeTc.suiteDescriptions = append(hugeTc.suiteNames, "tp-"+strconv.Itoa(i)+"description")
	}

	tests := []testCase{
		{
			suiteNames:        []string{"testsuite 1"},
			suiteDescriptions: []string{"testdescription 1"},
		},
		{
			suiteNames:        []string{},
			suiteDescriptions: []string{},
		},
		{
			suiteNames:        []string{"tp", "tp", "tp", "tp"},
			suiteDescriptions: []string{"td", "td", "td", "td"},
		},
		hugeTc,
	}

	for _, tc := range tests {
		s := NewStore()
		tutils.EqualI(t, 0, len(s.GetSuites()), "suite count")
		for i, suiteName := range tc.suiteNames {
			tutils.EqualNil(t, s.GetSuiteById(uint64(i)), "get suite by id before creation")
			suite := s.CreateSuite(suiteName, tc.suiteDescriptions[i], 0)

			tutils.EqualI(t, i, int(suite.Id), "suite id")
			tutils.EqualS(t, suiteName, suite.Name, "suite name")
			tutils.EqualS(t, tc.suiteDescriptions[i], suite.Description, "suite name")

			tutils.EqualNotNil(t, s.GetSuiteById(uint64(i)), "get suite by id after creation")
			suiteFromStore := s.GetSuiteById(uint64(i))
			tutils.EqualI(t, i, int(suiteFromStore.Id), "suiteById id")
			tutils.EqualS(t, suiteName, suiteFromStore.Name, "suiteById name")
			tutils.EqualS(t, tc.suiteDescriptions[i], suiteFromStore.Description, "suiteById name")
		}
		tutils.EqualI(t, len(tc.suiteNames), len(s.GetSuites()), "suite count")
	}
}
