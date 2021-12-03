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

func TestDeleteById(t *testing.T) {
	type testCase struct {
		projectNames        []string
		projectDescriptions []string
		idsToDelete         []uint64
		expectedSuccess     []bool
	}

	hugeTc := testCase{}
	for i := 0; i < 1000; i++ {
		hugeTc.projectNames = append(hugeTc.projectNames, "tp-"+strconv.Itoa(i))
		hugeTc.projectDescriptions = append(hugeTc.projectNames, "tp-"+strconv.Itoa(i)+"description")
		hugeTc.idsToDelete = append(hugeTc.idsToDelete, uint64(i))
		hugeTc.expectedSuccess = append(hugeTc.expectedSuccess, true)
	}

	tests := []testCase{
		{
			projectNames:        []string{"testproject 1"},
			projectDescriptions: []string{"testdescription 1"},
			idsToDelete:         []uint64{0},
			expectedSuccess:     []bool{true},
		},
		{
			projectNames:        []string{},
			projectDescriptions: []string{},
			idsToDelete:         []uint64{0},
			expectedSuccess:     []bool{false},
		},
		{
			projectNames:        []string{"tp", "tp", "tp", "tp"},
			projectDescriptions: []string{"td", "td", "td", "td"},
			idsToDelete:         []uint64{3, 1, 0, 2},
			expectedSuccess:     []bool{true, true, true, true},
		},
		{
			projectNames:        []string{"tp", "tp", "tp", "tp"},
			projectDescriptions: []string{"td", "td", "td", "td"},
			idsToDelete:         []uint64{4, 1000},
			expectedSuccess:     []bool{false, false},
		},
		hugeTc,
	}

	for _, tc := range tests {
		s := NewStore()
		tutils.EqualI(t, 0, len(s.GetProjects()), "projects count")

		for i, projectName := range tc.projectNames {
			s.CreateProject(projectName, tc.projectDescriptions[i])
			tutils.EqualNotNil(t, s.GetProjectById(uint64(i)), "get project by id after creation")
		}

		for i, idToDelete := range tc.idsToDelete {
			deletedProject := s.DeleteProjectById(idToDelete)
			if tc.expectedSuccess[i] {
				tutils.EqualNotNil(t, deletedProject, "project to delete was not found")
				tutils.EqualNil(t, s.GetProjectById(idToDelete), "deleted project still could be found")
			} else {
				tutils.EqualNil(t, deletedProject, "project to delete was found but it should not exist")
			}
		}
	}
}

func TestUpdateById(t *testing.T) {
	type testCase struct {
		projectNames               []string
		projectDescriptions        []string
		updatedProjectNames        []string
		updatedProjectDescriptions []string
		idsToUpdate                []uint64
		expectedSuccess            []bool
	}

	hugeTc := testCase{}
	for i := 0; i < 1000; i++ {
		hugeTc.projectNames = append(hugeTc.projectNames, "tp-"+strconv.Itoa(i))
		hugeTc.updatedProjectNames = append(hugeTc.projectNames, "updated-tp-"+strconv.Itoa(i))
		hugeTc.projectDescriptions = append(hugeTc.projectNames, "tp-"+strconv.Itoa(i)+"description")
		hugeTc.updatedProjectDescriptions = append(hugeTc.projectNames, "updated-tp-"+strconv.Itoa(i)+"description")
		hugeTc.idsToUpdate = append(hugeTc.idsToUpdate, uint64(i))
		hugeTc.expectedSuccess = append(hugeTc.expectedSuccess, true)
	}

	tests := []testCase{
		{
			projectNames:               []string{"testproject 1"},
			updatedProjectNames:        []string{"updatedname"},
			projectDescriptions:        []string{"testdescription 1"},
			updatedProjectDescriptions: []string{"updateddesc"},
			idsToUpdate:                []uint64{0},
			expectedSuccess:            []bool{true},
		},
		{
			projectNames:               []string{},
			updatedProjectNames:        []string{"updatedname"},
			projectDescriptions:        []string{},
			updatedProjectDescriptions: []string{"updateddesc"},
			idsToUpdate:                []uint64{0},
			expectedSuccess:            []bool{false},
		},
		{
			projectNames:               []string{"tp", "tp", "tp", "tp"},
			updatedProjectNames:        []string{"utp", "uutp", "uuutp", "uuuutp"},
			projectDescriptions:        []string{"td", "td", "td", "td"},
			updatedProjectDescriptions: []string{"utd", "uutd", "uuutd", "uuutd"},
			idsToUpdate:                []uint64{3, 1, 0, 2},
			expectedSuccess:            []bool{true, true, true, true},
		},
		{
			projectNames:               []string{"tp", "tp", "tp", "tp"},
			updatedProjectNames:        []string{"utp", "uutp", "uuutp", "uuuutp"},
			projectDescriptions:        []string{"td", "td", "td", "td"},
			updatedProjectDescriptions: []string{"utd", "uutd", "uuutd", "uuutd"},
			idsToUpdate:                []uint64{4, 1000},
			expectedSuccess:            []bool{false, false},
		},
		hugeTc,
	}

	for _, tc := range tests {
		s := NewStore()
		tutils.EqualI(t, 0, len(s.GetProjects()), "projects count")

		for i, projectName := range tc.projectNames {
			s.CreateProject(projectName, tc.projectDescriptions[i])
			tutils.EqualNotNil(t, s.GetProjectById(uint64(i)), "get project by id after creation")
		}

		for i, idToUpdate := range tc.idsToUpdate {
			updatedProject := s.UpdateProjectById(idToUpdate, tc.updatedProjectNames[i], tc.updatedProjectDescriptions[i])
			if tc.expectedSuccess[i] {
				tutils.EqualI(t, int(idToUpdate), int(updatedProject.Id), "project id")
				tutils.EqualS(t, tc.updatedProjectNames[i], updatedProject.Name, "project name")
				tutils.EqualS(t, tc.updatedProjectDescriptions[i], updatedProject.Description, "project description")

				tutils.EqualNotNil(t, s.GetProjectById(uint64(idToUpdate)), "get project by id after update")
				pp := s.GetProjectById(uint64(idToUpdate))
				tutils.EqualI(t, int(idToUpdate), int(pp.Id), "updated project id")
				tutils.EqualS(t, tc.updatedProjectNames[i], pp.Name, "updated project name")
				tutils.EqualS(t, tc.updatedProjectDescriptions[i], pp.Description, "updated project description")
			} else {
				tutils.EqualNil(t, updatedProject, "project to update was found but it should not exist")
			}
		}
	}
}
