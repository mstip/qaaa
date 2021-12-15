package model

import "github.com/mstip/qaaa/pkg/task"

const (
	TaskTypeJsonApi = "jsonapi"
	TaskTypeWeb     = "web"
)

type Project struct {
	Id          uint64
	Name        string
	Description string
}

type Suite struct {
	Id               uint64
	ProjectId        uint64
	Name             string
	Description      string
	BeforeAllTaskId  uint64
	BeforeEachTaskId uint64
	AfterAllTaskId   uint64
	AfterEachTaskId  uint64
}

type Task struct {
	Id          uint64
	SuiteId     uint64
	Name        string
	Description string
	Type        string
	Task        task.WebTaskRequest
}
