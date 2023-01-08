package quicktask

import (
	"context"
	"reflect"
	"time"

	"github.com/eagle/client"
	"github.com/eagle/handler/profile"
)

type Quicktask struct {
	Id             string
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

	mode    string
	size    string
	pid     string
	Client  *client.Client     `json:"-"` // client
	Context context.Context    `json:"-"`
	Cancel  context.CancelFunc `json:"-"` // cancel function
	Delay   time.Duration      `json:"-"` // delay (in ms)

	Type            string             `json:"type"` // site name + monitor
	CheckoutData    CheckoutLogRequest `json:"-"`    // checkout data
	CheckoutProxy   string             // proxy data
	CheckoutProfile profile.Profile    `json:"-"` // profile data
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
