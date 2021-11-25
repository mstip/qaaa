package store

import (
	"github.com/mstip/qaaa/pkg/model"
)

func (s *Store) newSuiteId() uint64 {
	id := s.nextSuiteId
	s.nextSuiteId += 1
	return id
}

func (s *Store) CreateSuite(name string, description string, project *model.Project) *model.Suite {
	suite := model.Suite{Id: s.newSuiteId(), Name: name, Description: description}
	project.Suites = append(project.Suites, suite)
	return &suite
}

func (s *Store) GetSuiteById(sId uint64) *model.Suite {
	// TODO: this needs to be refactored
	for _, p := range s.projects {
		for _, s := range p.Suites {
			if s.Id == sId {
				// TODO: can this leak ?
				return &s
			}
		}
	}
	return nil
}
