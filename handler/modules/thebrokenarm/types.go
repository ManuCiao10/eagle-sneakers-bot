package thebrokenarm

import (
	"github.com/eagle/handler/task"
)

var (
	INITIALIZE     task.TaskState = "initialize"
	GET_SESSION    task.TaskState = "get_session"
	LOGIN          task.TaskState = "login"
	CLEAR_CART     task.TaskState = "clear_cart"
	GET_ITEM       task.TaskState = "get_item"
	ADD_TO_CART    task.TaskState = "add_to_cart"
	CHECKOUT_ORDER task.TaskState = "checkout_order"
)


