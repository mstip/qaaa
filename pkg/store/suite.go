package store

import (
	"github.com/mstip/qaaa/pkg/model"
)

func (s *Store) newSuiteId() uint64 {
	s.idLock.Lock()
	defer s.idLock.Unlock()
	id := s.nextSuiteId
	s.nextSuiteId += 1
	return id
}

func (s *Store) GetSuites() []model.Suite {
	s.dataLock.Lock()
	defer s.dataLock.Unlock()
	return s.suites
}

func (s *Store) CreateSuite(name string, description string, projectId uint64) *model.Suite {
	s.dataLock.Lock()
	defer s.dataLock.Unlock()
	suite := model.Suite{Id: s.newSuiteId(), Name: name, Description: description, ProjectId: projectId}
	s.suites = append(s.suites, suite)
	return &suite
}

func (s *Store) GetSuiteById(sId uint64) *model.Suite {
	s.dataLock.Lock()
	defer s.dataLock.Unlock()
	for _, s := range s.suites {
		if s.Id == sId {
			return &s
		}
	}
	return nil
}

func (s *Store) GetSuitesByProjectId(pId uint64) []model.Suite {
	s.dataLock.Lock()
	defer s.dataLock.Unlock()
	suites := []model.Suite{}
	for _, s := range s.suites {
		if s.ProjectId == pId {
			suites = append(suites, s)
		}
	}
	return suites
}
