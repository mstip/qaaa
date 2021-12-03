package store

import (
	"sync"

	"github.com/mstip/qaaa/pkg/model"
)

type Store struct {
	dataLock      sync.Mutex
	idLock        sync.Mutex
	projects      []model.Project
	suites        []model.Suite
	tasks         []model.Task
	nextProjectId uint64
	nextSuiteId   uint64
	nextTaskId    uint64
}

func NewStoreWithDemoData() *Store {
	s := &Store{}
	s.seedWithDemoData()
	return s
}

func NewStore() *Store {
	s := &Store{}
	return s
}
