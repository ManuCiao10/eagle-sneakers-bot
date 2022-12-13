package thebrokenarm

import "reflect"

type Task struct {
	Pid         string
	Size        string
	Email       string
	Profile     string
	Method      string
	Card_Number string
	Month       string
	Year        string
	CVV         string
	Proxy_List  string
}

var (
	INITIALIZE     TaskState = "initialize"
	GET_SESSION    TaskState = "get_session"
	LOGIN          TaskState = "login"
	CLEAR_CART     TaskState = "clear_cart"
	GET_ITEM       TaskState = "get_item"
	ADD_TO_CART    TaskState = "add_to_cart"
	CHECKOUT_ORDER TaskState = "checkout_order"
)

type TaskType struct {
	firstHandlerState TaskState
	internalType      reflect.Type
	handlers          TaskReflectMap
}

type TaskState string
type TaskHandlerMap map[TaskState]interface{}
type TaskReflectMap map[string]reflect.Value
