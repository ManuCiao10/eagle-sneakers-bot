package eaglemonitor

import (
	"github.com/eagle/handler/logs"
	"github.com/eagle/handler/task"
)

func waitForMonitor(t *task.Task) task.TaskState {

	// t.MonitorData = make(chan interface{})
	logs.LogsMsgWarn(t, "Waiting for monitor...")
	//reading MQT data and store all of them
	//if the pid received is the same as the pid in the task, then return GET_SESSION

	return WAIT_FOR_MONITOR
}
