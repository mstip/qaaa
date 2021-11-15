package model

type Task struct {
	Id          uint64
	Name        string
	Description string
	Command     string
	Check       string
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
