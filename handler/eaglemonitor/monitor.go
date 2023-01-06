package eaglemonitor

import (
	"github.com/eagle/handler/logs"
	"github.com/eagle/handler/task"
)

func waitForMonitor(t *task.Task) task.TaskState {

	t.MonitorData = make(chan interface{})
	logs.LogsMsgWarn(t, "Waiting for monitor...")

	return WAIT_FOR_MONITOR
}
