package store

import "github.com/mstip/qaaa/pkg/model"

func (s *Store) newTaskId() uint64 {
	s.idLock.Lock()
	defer s.idLock.Unlock()
	id := s.nextTaskId
	s.nextTaskId += 1
	return id
}

func (s *Store) GetTasksBySuiteId(sId uint64) []model.Task {
	s.dataLock.Lock()
	defer s.dataLock.Unlock()
	tasks := []model.Task{}
	for _, t := range s.tasks {
		if t.SuiteId == sId {
			tasks = append(tasks, t)
		}
	}
	return tasks
}
