package task_manager

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/eagle/handler/loading"
	"github.com/eagle/handler/logs"
	"github.com/eagle/handler/quicktask"
	"github.com/getsentry/sentry-go"
)

var Debug = true

func handleQuickTaskState(taskState quicktask.TaskState, taskType *quicktask.TaskType, t *quicktask.Quicktask) quicktask.TaskState {
	nextTaskHandler, err := taskType.GetHandler(taskState)

	if err != nil {
		log.Println("Task handler error: ", err)
		return quicktask.ErrorTaskState
	}

	nextQuickTaskHandlerFunc := nextTaskHandler.Interface().(func(*quicktask.Quicktask) quicktask.TaskState)

	return nextQuickTaskHandlerFunc(t)
}

// RunQuickTask runs a QuickTask
func RunQuickTask(t *quicktask.Quicktask) {
	fmt.Println("Running task", "of type", t.Type)
	t.Context, t.Cancel = context.WithCancel(context.Background())
	t.Active = true

	defer func() {
		if r := recover(); r != nil {
			log.Println("Task error:", r)

			sentry.RecoverWithContext(t.Context)
			sentry.Flush(time.Second * 5)
		}
	}()

	if !quicktask.DoesTaskTypeExist(t.Type) {
		return
	}

	taskType, err := quicktask.GetTaskType(t.Type)

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
	logs.LogQuick(t, "Starting task...")
	t.CheckoutData.TaskStart = time.Now()

	// t.Internal = reflect.New(taskType.GetInternalType().Elem()).Interface()

	// loop the task states
	for {
		nextState = handleQuickTaskState(nextState, taskType, t)
		if Debug {
			fmt.Println(t, nextState)
		}

		if nextState == quicktask.DoneTaskState || t.Context.Err() != nil {
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
			// you can report that the task stopped here
			t.Active = false
			break
		} else if nextState == quicktask.ErrorTaskState {
			// report errors
			t.Active = false
			break
		}

		time.Sleep(1 * time.Millisecond)
	}
}

// StopQuickTask stops a QuickTask
func StopQuickTask(t *quicktask.Quicktask) {
	t.Cancel()
}
