package model

import "github.com/mstip/qaaa/pkg/task"

const (
	TaskTypeJsonApi = "jsonapi"
	TaskTypeWeb     = "web"
)

type Task struct {
	Id          uint64
	Name        string
	Description string
	Type        string
	Task        task.Tasker
}

type Suite struct {
	Id          uint64
	Name        string
	Description string
	Tasks       []Task
	BeforeAll   Task
	BeforeEach  Task
	AfterAll    Task
	AfterEach   Task
}

type Project struct {
	Id          uint64
	Name        string
	Description string
	Suites      []Suite
}
