package thebrokenarm_monitor

import (
	"github.com/eagle/handler/quicktask"
)

func Initialize() {
	monitorType := quicktask.RegisterTaskType("thebrokenarmmonitor")

	monitorType.SetFirstHandlerState(INITIALIZE)

	monitorType.AddHandlers(quicktask.TaskHandlerMap{
		INITIALIZE: initialize,
		LOGIN:      login,
		SESSION:    session,
		GUEST:      guest,
		CHECKOUT:   checkout,
	})
}
