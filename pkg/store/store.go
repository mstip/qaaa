package store

type Store struct {
	counter int
}

func (s *Store) IncCounter() {
	s.counter += 1
}

func (s *Store) Counter() int {
	return s.counter
}

func NewStore() *Store {
	return &Store{}
}