package thebrokenarm

import (
	"time"

	"github.com/eagle/handler/client"
)

type Task struct {
	Mode        string        `json:"mode"`
	Pid         string        `json:"pid"`
	Size        string        `json:"size"`
	Email       string        `json:"email"`
	Profile     string        `json:"profile"`
	Method      string        `json:"method"`
	Card_Number string        `json:"card_number"`
	Month       string        `json:"month"`
	Year        string        `json:"year"`
	CVV         string        `json:"cvv"`
	Proxy_List  string        `json:"proxy_list"`
	Delay       time.Duration `json:"delay"` // delay (in ms)

	Client       *client.Client     `json:"-"` // http client
	CheckoutData CheckoutLogRequest `json:"-"` // checkout data
}

type CheckoutLogRequest struct {
	TaskStart   time.Time `json:"-"`            // auto defined
	TaskEnd     time.Time `json:"-"`            // auto defined
	CheckoutMs  int       `json:"checkout_ms"`  // auto defined
	Price       float64   `json:"price"`        // needs to be defined
	ProductName string    `json:"product_name"` // needs to be defined
	ProductMSKU string    `json:"product_msku"` // needs to be defined
	Mode        string    `json:"mode"`         // needs to be defined
	Size        string    `json:"size"`         // needs to be defined
	Status      string    `json:"status"`       // needs to be defined
	Website     string    `json:"website"`      // siteName, needs to be defined
	ImageUrl    string    `json:"image_url"`    // needs to be defined
	Proxy       string    `json:"proxy"`        // needs to be defined
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

var (
	DoneTaskState     TaskState = "done"
	ContinueTaskState TaskState = "continue"
	ErrorTaskState    TaskState = "error"
)

// type TaskType struct {
// 	firstHandlerState TaskState
// 	internalType      reflect.Type
// 	handlers          TaskReflectMap
// }

type TaskState string

// type TaskHandlerMap map[TaskState]interface{}
// type TaskReflectMap map[string]reflect.Value
