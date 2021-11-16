package store

import (
	"github.com/mstip/qaaa/pkg/model"
)

func (s *Store) seedWithDemoData() {
	s.projects = append(s.projects, model.Project{Id: s.newProjectId(), Name: "demo project", Description: "This is the test and demo project"})
	s.projects = append(s.projects, model.Project{Id: s.newProjectId(), Name: "demo project 2", Description: "This is the test and demo project 2 "})
	s.projects = append(s.projects, model.Project{Id: s.newProjectId(), Name: "demo project 3", Description: "This is the test and demo project 3 "})
	s.projects = append(s.projects, model.Project{Id: s.newProjectId(), Name: "demo project 4", Description: "This is the test and demo project 4"})
}
