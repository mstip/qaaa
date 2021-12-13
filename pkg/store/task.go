package store

import (
	"strconv"

	"github.com/mstip/qaaa/pkg/model"
)

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

func (s *Store) CreateTask(name string, description string, suiteId uint64) *model.Task {
	s.dataLock.Lock()
	defer s.dataLock.Unlock()
	task := model.Task{Id: s.newTaskId(), Name: name, Description: description, SuiteId: suiteId}
	s.tasks = append(s.tasks, task)
	return &task
}

func (s *Store) GetTaskById(tId uint64) *model.Task {
	s.dataLock.Lock()
	defer s.dataLock.Unlock()
	for _, t := range s.tasks {
		if t.Id == tId {
			return &t
		}
	}
	return nil
}

func (s *Store) UpdateTaskByIdParam(idParam string, name string, description string) *model.Task {
	tId, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return nil
	}

	return s.UpdateTaskById(tId, name, description)
}

func (s *Store) UpdateTaskById(tId uint64, name string, description string) *model.Task {
	s.dataLock.Lock()
	defer s.dataLock.Unlock()
	for i, v := range s.tasks {
		if v.Id == tId {
			s.tasks[i].Name = name
			s.tasks[i].Description = description
			return &s.tasks[i]
		}
	}
	return nil
}
func (s *Store) DeleteTaskByIdParm(idParam string) *model.Task {
	tId, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return nil
	}
	return s.DeleteTaskById(tId)
}

func (s *Store) DeleteTaskById(tId uint64) *model.Task {
	s.dataLock.Lock()
	defer s.dataLock.Unlock()
	for i, v := range s.tasks {
		if v.Id == tId {
			s.tasks = append(s.tasks[:i], s.tasks[i+1:]...)
			return &v
		}
	}
	return nil
}
