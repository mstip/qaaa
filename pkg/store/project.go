package store

import "github.com/mstip/qaaa/pkg/model"

type ProjectListData struct {
	Id          uint64
	Name        string
	Description string
}

func (s *Store) newProjectId() uint64 {
	id := s.nextProjectId
	s.nextProjectId += 1
	return id
}

func (s *Store) GetProjectsList() []ProjectListData {
	pld := []ProjectListData{}
	for _, v := range s.projects {
		pld = append(pld, ProjectListData{Id: v.Id, Name: v.Name, Description: v.Description})
	}
	return pld
}

func (s *Store) CreateProject(name string, description string) *model.Project {
	project := model.Project{Id: s.newProjectId(), Name: name, Description: description}
	s.projects = append(s.projects, project)
	return &project
}

func (s *Store) GetProjectById(pId uint64) *model.Project {
	for _, v := range s.projects {
		if v.Id == pId {
			return &v
		}
	}
	return nil
}
