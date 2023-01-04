package thebrokenarm

import (
	"github.com/eagle/handler/task"
)

// probably getCloud not needed
func Initialize() {
	monitorType := task.RegisterTaskType("thebrokenarm")

	monitorType.SetFirstHandlerState(INITIALIZE)

	monitorType.AddHandlers(task.TaskHandlerMap{
		INITIALIZE:  initialize,
		GET_SESSION: getSession,
		GET_CLOUD:   getCloud,
		ADD_TO_CART: addToCart,
		CHECKOUT:    checkout,
	})
}
