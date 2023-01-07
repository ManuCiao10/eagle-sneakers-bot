package thebrokenarm_monitor

import (
	"time"

	"github.com/eagle/handler/quicktask"
)

func initialize(t *quicktask.Quicktask) quicktask.TaskState {
	//print struct quicktask.Quicktask
	// fmt.Printf("%+v", t)
	time.Sleep(500 * time.Second)
	//read account
	//set profile info
	//set proxy info
	//crate the login
	return quicktask.DoneTaskState
}
