package task_manager

import "github.com/eagle/handler/task"

func RunTask(t *task.Task) {
	t.Active = true
	// t.Run()
}
