package task_manager

type TaskManager interface {
	CreateTask(name, description string) (Task, error)
}

type Task interface {
	GetKey() string
}
