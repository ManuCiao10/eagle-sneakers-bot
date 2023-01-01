package thebrokenarm

import (
	"github.com/eagle/handler/task"
)

var (
	INITIALIZE  task.TaskState = "initialize"
	GET_SESSION task.TaskState = "get_session"
	GET_CLOUD   task.TaskState = "get_cloud"
	PRODUCT     task.TaskState = "product"
)

var TBAInternal = struct {
	ProductID string
}{}

// type TBAInternal struct {
// 	ProductID string
// }
