package eaglemonitor

import "github.com/eagle/handler/task"

var (
	WAIT_FOR_MONITOR task.TaskState = "wait_for_monitor"
	GET_SESSION      task.TaskState = "get_session"
	SUBMIT_ORDER     task.TaskState = "submit_order"
	CHECKOUT_ORDER   task.TaskState = "checkout_order"
)
