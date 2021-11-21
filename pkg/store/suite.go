package store

func (s *Store) newSuiteId() uint64 {
	id := s.nextSuiteId
	s.nextSuiteId += 1
	return id
}
