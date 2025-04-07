package stale_flag_cleaner

type Template struct {
	taskName        string
	taskDescription string
}

func NewTemplate(taskName string, taskDescription string) *Template {
	return &Template{
		taskName:        taskName,
		taskDescription: taskDescription,
	}
}
