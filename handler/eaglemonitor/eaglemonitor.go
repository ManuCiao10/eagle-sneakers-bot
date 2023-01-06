package eaglemonitor

import "github.com/eagle/handler/task"

func Initialize() {
	taskType := task.RegisterTaskType("monitor")

	taskType.SetFirstHandlerState(WAIT_FOR_MONITOR)

	taskType.AddHandlers(task.TaskHandlerMap{
		WAIT_FOR_MONITOR: waitForMonitor,
		// GET_SESSION:      getSession,
		// SUBMIT_ORDER:     submitOrder,
		// CHECKOUT_ORDER:   checkout,
	})
}
