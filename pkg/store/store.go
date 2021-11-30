package store

import (
	"github.com/mstip/qaaa/pkg/model"
)

type Store struct {
	projects      []model.Project
	suites        []model.Suite
	tasks         []model.Task
	nextProjectId uint64
	nextSuiteId   uint64
	nextTaskId    uint64
}

func NewStore() *Store {
	s := &Store{}
	s.seedWithDemoData()
	return s
}
