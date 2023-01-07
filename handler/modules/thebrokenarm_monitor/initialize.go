package thebrokenarm_monitor

import (
	"fmt"

	"github.com/eagle/handler/quicktask"
)

func initialize(t *quicktask.Quicktask) quicktask.TaskState {
	//print struct quicktask.Quicktask
	fmt.Printf("%+v", t)
	//read account
	//set profile info
	//set proxy info
	//crate the login
	return quicktask.DoneTaskState
}
