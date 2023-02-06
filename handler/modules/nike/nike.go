package fiver

import (
	"github.com/eagle/handler/task"
)

func Initialize() {
	monitorType := task.RegisterTaskType("nike")

	monitorType.SetFirstHandlerState(INITIALIZE)

	monitorType.AddHandlers(task.TaskHandlerMap{
		INITIALIZE: initialize,
	})
}
