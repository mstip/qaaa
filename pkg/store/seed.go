package store

import "github.com/mstip/qaaa/pkg/model"

func (s *Store) seedWithDemoData() {
	s.projects = append(s.projects, model.Project{Id: s.newProjectId(), Name: "demo project", Description: "This is the test and demo project"})
}
