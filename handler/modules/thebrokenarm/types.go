package thebrokenarm

import (
	"github.com/eagle/handler/task"
)

var (
	INITIALIZE  task.TaskState = "initialize"
	GET_SESSION task.TaskState = "get_session"
	GET_CLOUD   task.TaskState = "get_cloud"
	ADD_TO_CART task.TaskState = "product"
	CHECKOUT    task.TaskState = "checkout"
)

var TBAInternal = struct {
	ProductID string
}{}
