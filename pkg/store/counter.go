package store

func (s *Store) IncCounter() {
	s.counter += 1
}

func (s *Store) Counter() int {
	return s.counter
}
