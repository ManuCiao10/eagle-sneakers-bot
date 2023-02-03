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

	Pid     string
	Client  *client.Client     `json:"-"` // client
	Context context.Context    `json:"-"`
	Cancel  context.CancelFunc `json:"-"` // cancel function
	Delay   time.Duration      `json:"-"` // delay (in ms)
	Size    string             `json:"-"` // size

	Type            string          `json:"type"` // site name + monitor
	CheckoutData    CheckoutLog     `json:"-"`    // checkout data
	CheckoutProfile profile.Profile `json:"-"`    // profile data
}

type CheckoutLog struct {
	TaskStart time.Time `json:"-"`
	TaskEnd   time.Time `json:"-"`
	Proxy     string    `json:"-"`
	Price     string    `json:"-"`
	Profile   string    `json:"-"`

	Link        string `json:"-"`
	Size        string `json:"-"`
	Image_url   string `json:"-"`
	ProductName string `json:"-"`
	ProductMSKU string `json:"-"`

	CheckoutMs int    `json:"checkout_ms"`
	Status     string `json:"status"`

	Website string `json:"website"`
	Mode    string `json:"mode"`
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
