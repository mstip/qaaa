package store

import (
	"github.com/mstip/qaaa/pkg/model"
)

type Store struct {
	counter       int
	projects      []model.Project
	nextProjectId uint64
	nextSuiteId   uint64
	nextTaskId    uint64
}

func NewStore() *Store {
	s := &Store{}
	s.seedWithDemoData()
	return s
}
