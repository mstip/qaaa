package store

func (s *Store) newTaskId() uint64 {
	id := s.nextTaskId
	s.nextTaskId += 1
	return id
}
