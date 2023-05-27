package task_manager

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/eagle/handler/loading"
	"github.com/eagle/handler/logs"
	"github.com/eagle/handler/task"
	"github.com/eagle/handler/utils"
	"github.com/getsentry/sentry-go"
)

func handleTaskState(taskState task.TaskState, taskType *task.TaskType, t *task.Task) task.TaskState {
	nextTaskHandler, err := taskType.GetHandler(taskState)

	if err != nil {
		log.Println("Task handler error: ", err)
		return task.ErrorTaskState
	}

	nextTaskHandlerFunc := nextTaskHandler.Interface().(func(*task.Task) task.TaskState)

	return nextTaskHandlerFunc(t)
}

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

	taskType, err := task.GetTaskType(t.Type)

	if err != nil {
		log.Println("Task type error: ", err)
		t.Active = false
		return
	}

	hasHandlers := taskType.HasHandlers()

	if !hasHandlers {
		fmt.Printf("Task type %s has no handlers\n", t.Type)
		t.Done = true
		return
	}

	nextState := taskType.GetFirstHandlerState()

	if len(nextState) == 0 {
		fmt.Printf("Task type %s has no first handler state\n", t.Type)
		t.Done = true
		return
	}
	logs.LogInfo(t, "Starting task...")
	t.CheckoutData.TaskStart = time.Now()

	for {
		nextState = handleTaskState(nextState, taskType, t)
		if utils.Debug {
			fmt.Println(t, nextState)
		}

		if nextState == task.DoneTaskState || t.Context.Err() != nil {
			t.CheckoutData.TaskEnd = time.Now()
			t.CheckoutData.CheckoutMs = int(t.CheckoutData.TaskEnd.Sub(t.CheckoutData.TaskStart).Milliseconds())
			logs.LogCheckout(&logs.CheckoutLogRequest{
				TaskStart:   t.CheckoutData.TaskStart,
				TaskEnd:     t.CheckoutData.TaskEnd,
				Price:       t.CheckoutData.Price,
				ProductName: t.CheckoutData.ProductName,
				ProductMSKU: t.CheckoutData.ProductMSKU,
				Mode:        t.CheckoutData.Mode,
				CheckoutMs:  t.CheckoutData.CheckoutMs,
				Size:        t.CheckoutData.Size,
				Status:      t.CheckoutData.Status,
				Website:     t.CheckoutData.Website,
				ImageUrl:    t.CheckoutData.ImageUrl,
			}, loading.Data.Settings.Settings.DiscordWebhook)

			t.Active = false
			break
		} else if nextState == task.ErrorTaskState {
			fmt.Println("Task error")
			t.Active = false
			break
		}

		time.Sleep(1 * time.Millisecond)
	}
}

// StopTask stops a task
func StopTask(t *task.Task) {
	t.Cancel()
}
