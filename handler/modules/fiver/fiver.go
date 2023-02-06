package fiver

import (
	"github.com/eagle/handler/task"
)

// probably getCloud not needed
func Initialize() {
	monitorType := task.RegisterTaskType("fiver")

	monitorType.SetFirstHandlerState(INITIALIZE)

	monitorType.AddHandlers(task.TaskHandlerMap{
		INITIALIZE: initialize,
		FILL_DATA:  fillData,
		FILL_USER:  fillUser,
	})
}
