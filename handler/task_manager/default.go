package task_manager

import (
	"context"
	"log"
	"time"

	"github.com/eagle/handler/task"
	"github.com/getsentry/sentry-go"
)

func RunTask(t *task.Task) {
	t.Context, t.Cancel = context.WithCancel(context.Background())
	t.Active = true

	defer func() {
		if r := recover(); r != nil {
			log.Println("Task error:", r)

			sentry.RecoverWithContext(t.Context)
			sentry.Flush(time.Second * 5)
		}
	}()

	if !task.DoesTaskTypeExist(t.Type) {
		return
	}
}

// StopTask stops a task
func StopTask(t *task.Task) {
	t.Cancel()
}
