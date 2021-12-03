package store

import (
	"strconv"
	"testing"

	"github.com/mstip/qaaa/internal/tutils"
)

func TestCreateAndGetProjects(t *testing.T) {
	type testCase struct {
		projectNames        []string
		projectDescriptions []string
	}

	hugeTc := testCase{}
	for i := 0; i < 1000; i++ {
		hugeTc.projectNames = append(hugeTc.projectNames, "tp-"+strconv.Itoa(i))
		hugeTc.projectDescriptions = append(hugeTc.projectNames, "tp-"+strconv.Itoa(i)+"description")
	}

	tests := []testCase{
		{
			projectNames:        []string{"testproject 1"},
			projectDescriptions: []string{"testdescription 1"},
		},
		{
			projectNames:        []string{},
			projectDescriptions: []string{},
		},
		{
			projectNames:        []string{"tp", "tp", "tp", "tp"},
			projectDescriptions: []string{"td", "td", "td", "td"},
		},
		hugeTc,
	}

	for _, tc := range tests {
		s := NewStore()
		tutils.EqualI(t, 0, len(s.GetProjects()), "projects count")
		for i, projectName := range tc.projectNames {
			tutils.EqualNil(t, s.GetProjectById(uint64(i)), "get project by id before creation")

			p := s.CreateProject(projectName, tc.projectDescriptions[i])
			tutils.EqualI(t, i, int(p.Id), "project id")
			tutils.EqualS(t, projectName, p.Name, "project name")
			tutils.EqualS(t, tc.projectDescriptions[i], p.Description, "project name")

			tutils.EqualNotNil(t, s.GetProjectById(uint64(i)), "get project by id after creation")
			pp := s.GetProjectById(uint64(i))
			tutils.EqualI(t, i, int(pp.Id), "projectById id")
			tutils.EqualS(t, projectName, pp.Name, "projectById name")
			tutils.EqualS(t, tc.projectDescriptions[i], pp.Description, "projectById name")
		}
		tutils.EqualI(t, len(tc.projectNames), len(s.GetProjects()), "projects count")
	}
}
