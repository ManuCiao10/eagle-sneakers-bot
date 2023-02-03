package task

import (
	"context"
	"reflect"
	"time"

	"github.com/eagle/client"

	"github.com/eagle/handler/profile"
)

type Task struct {
	TaskType    string `json:"task_type"` // Site name, Index file name choosen
	Mode        string `json:"mode"`
	Pid         string `json:"pid"`
	Size        string `json:"size"`
	Email       string `json:"email"`
	Profile     string `json:"profile"` // profile name
	Method      string `json:"method"`
	Card_Number string `json:"card_number"`
	Month       string `json:"month"`
	Year        string `json:"year"`
	CVV         string `json:"cvv"`

	Active      bool             `json:"-"`    // active status
	Done        bool             `json:"-"`    // done status
	Delay       time.Duration    `json:"-"`    // delay (in ms)
	Type        string           `json:"type"` // registered task type aka site name
	Internal    interface{}      `json:"-"`    // internal data, gotten from second func argument
	MonitorData chan interface{} `json:"-"`    // monitor data, only used in checkout tasks

	Client          *client.Client     `json:"-"` // client
	Context         context.Context    `json:"-"`
	Cancel          context.CancelFunc `json:"-"` // cancel function
	CheckoutProfile profile.Profile    `json:"-"` // profile data
	CheckoutProxy   string             // proxy data
	CheckoutData    CheckoutLogRequest `json:"-"` // checkout data
}

type CheckoutLogRequest struct {
	TaskStart   time.Time `json:"-"`            // auto defined
	TaskEnd     time.Time `json:"-"`            // auto defined
	CheckoutMs  int       `json:"checkout_ms"`  // auto defined
	Price       string    `json:"price"`        // needs to be defined
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

var (
	DoneTaskState  TaskState = "done"
	ErrorTaskState TaskState = "error"
)

type TaskType struct {
	firstHandlerState TaskState
	internalType      reflect.Type
	handlers          TaskReflectMap
}

type TaskState string
type TaskHandlerMap map[TaskState]interface{}
type TaskReflectMap map[string]reflect.Value
