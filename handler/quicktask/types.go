package quicktask

import (
	"context"
	"reflect"
	"time"

	"github.com/eagle/client"
)

type Quicktask struct {
	Site           string
	Tasks_Quantity string
	Profiles       string
	Accounts       string
	Email          string
	Proxylist      string
	Payment_Method string
	Credit_Card    string
	Other          string
	Active         bool
	Done           bool

	Mode    string
	Size    string
	Pid     string
	Client  *client.Client     `json:"-"` // client
	Context context.Context    `json:"-"`
	Cancel  context.CancelFunc `json:"-"` // cancel function

	Type         string             `json:"type"` // registered task type aka site name
	CheckoutData CheckoutLogRequest `json:"-"`    // checkout data
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
	Profile     string    `json:"profile"`      // needs to be defined
}

type TaskType struct {
	firstHandlerState TaskState
	internalType      reflect.Type
	handlers          TaskReflectMap
}

type TaskState string
type TaskHandlerMap map[TaskState]interface{}
type TaskReflectMap map[string]reflect.Value

var (
	DoneTaskState  TaskState = "done"
	ErrorTaskState TaskState = "error"
)
