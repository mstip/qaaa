package task

type Tasker interface {
	Arrange() error
	Act() error
	Assert() (bool, error)
}
