package store

import "github.com/mstip/qaaa/pkg/model"

func (s *Store) newProjectId() uint64 {
	s.idLock.Lock()
	defer s.idLock.Unlock()
	id := s.nextProjectId
	s.nextProjectId += 1
	return id
}

func (s *Store) GetProjects() []model.Project {
	s.dataLock.Lock()
	defer s.dataLock.Unlock()
	return s.projects
}

func (s *Store) CreateProject(name string, description string) *model.Project {
	s.dataLock.Lock()
	defer s.dataLock.Unlock()
	project := model.Project{Id: s.newProjectId(), Name: name, Description: description}
	s.projects = append(s.projects, project)
	return &project
}

func (s *Store) GetProjectById(pId uint64) *model.Project {
	s.dataLock.Lock()
	defer s.dataLock.Unlock()
	for _, v := range s.projects {
		if v.Id == pId {
			return &v
		}
	}
	return nil
}
