package thebrokenarm

import (
	"fmt"
	"os"
	"reflect"
	"time"

	"github.com/eagle/handler/utils"
	"github.com/eagle/handler/version"
	"github.com/fatih/color"
)

var (
	taskTypes         = make(map[string]*TaskType)
	SiteConversionSTI = make(map[string]int)
	SiteConversionITS = make(map[int][]string)
)

func Loading() {
	monitorType := RegisterTaskType("thebrokenarm", 0)

	monitorType.SetFirstHandlerState(INITIALIZE)

	monitorType.AddHandlers(TaskHandlerMap{
		INITIALIZE: initialize,
		// GET_SESSION: waitForMonitor,
		// LOGIN:          getOrderId,
		// CLEAR_CART:     submitOrder,
		// GET_ITEM:       getItem,
		// CHECKOUT_ORDER: checkout,
	})
	fmt.Println("Loading")
}
func initialize() TaskState {
	fmt.Print("\033[H\033[2J")
	utils.Banner()
	utils.Directory("thebrokenarm")

	csv_index := utils.SelectMode("[Eagle " + version.Version + "] " + "[" + time.Now().Format("15:04:05.000000") + "]" + " PLEASE SELECT CSV:")
	task_name := CvsIndex(csv_index, "thebrokenarm")
	if task_name == "UNEXPECTED" {
		err_("UNEXPECTED ERROR")
	}

	CvsInfo(task_name, "thebrokenarm")

	return GET_SESSION

}

// SetFirstHandlerState sets the first handler state
func (t *TaskType) SetFirstHandlerState(firstHandlerState TaskState) {
	t.firstHandlerState = firstHandlerState
}

func RegisterTaskType(registerSiteName string, siteId int) *TaskType {
	taskTypes[registerSiteName] = &TaskType{
		handlers: make(TaskReflectMap),
	}

	SiteConversionSTI[registerSiteName] = siteId
	SiteConversionITS[siteId] = append(SiteConversionITS[siteId], registerSiteName)

	return taskTypes[registerSiteName]
}

func (t *TaskType) addHandler(handlerName TaskState, handler interface{}) {
	t.handlers[string(handlerName)] = reflect.ValueOf(handler)
}

// AddHandlers adds multiple handles to a task type
func (t *TaskType) AddHandlers(handlers TaskHandlerMap) {
	for handlerName, handler := range handlers {
		if t.internalType == nil {
			handleTypes := reflect.TypeOf(handler)
			// func (t *task.Task, internal *SiteInternal) task.TaskState

			// we want the second one because the first one (0 index) will be task.Task type
			handleType := handleTypes.In(1)

			t.internalType = handleType
		}

		t.addHandler(handlerName, handler)
	}
}

// func Thebrokenarm() {

// 	//use range to loop through map
// 	for _, i := range tasks {
// 		go test(i)
// 	}

// }

// func test(i *Task) {
// 	fmt.Println("task ["+"]", i)
// 	time.Sleep(1 * time.Second)

// }

func err_(msg string) {
	color.Red(msg)
	os.Exit(0)
}
