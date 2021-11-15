package store

type ProjectListData struct {
	Id          uint64
	Name        string
	Description string
}

func (s *Store) newProjectId() uint64 {
	id := s.nextProjectId
	s.nextProjectId = +1
	return id
}

func (s *Store) GetProjectsList() []ProjectListData {
	pld := []ProjectListData{}
	for _, v := range s.projects {
		pld = append(pld, ProjectListData{Id: v.Id, Name: v.Name, Description: v.Description})
	}
	return pld
}
