package thebrokenarm

import (
	"github.com/eagle/handler/task"
)

// probably getCloud not needed GET_CLOUD:   getCloud,
func Initialize() {
	monitorType := task.RegisterTaskType("thebrokenarm")

	monitorType.SetFirstHandlerState(INITIALIZE)

	monitorType.AddHandlers(task.TaskHandlerMap{
		INITIALIZE:  initialize,
		GET_SESSION: getSession,
		ADD_TO_CART: addToCart,
		CHECKOUT:    checkout,
	})
}
