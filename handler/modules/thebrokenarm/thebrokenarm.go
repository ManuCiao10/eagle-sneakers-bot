package thebrokenarm

import (
	"github.com/eagle/handler/task"
)

func Initialize() {
	monitorType := task.RegisterTaskType("thebrokenarm")

	monitorType.SetFirstHandlerState(INITIALIZE)

	monitorType.AddHandlers(task.TaskHandlerMap{
		INITIALIZE: initialize,
		// GET_SESSION: getSession,
		// LOGIN:       login,
		// CLEAR_CART:  clearCart,
		// GET_ITEM:    getItem,
		// ADD_TO_CART: addToCart,
	})
}
