package main

import (
	"fmt"

	"github.com/eagle/handler/auth"
	"github.com/eagle/handler/create"
	"github.com/eagle/handler/eaglemonitor"
	"github.com/eagle/handler/loading"
	"github.com/eagle/handler/modules/thebrokenarm"
	"github.com/eagle/handler/presence"
	"github.com/eagle/handler/site"
	"github.com/eagle/handler/task"
	"github.com/eagle/handler/task_manager"
	"github.com/eagle/handler/utils"
	"github.com/eagle/handler/version"
)

var (
	Run bool
	Dev = true
)

//go:generate goversioninfo -skip-versioninfo=true -icon=handler/create/favicon.ico -manifest=handler/create/file.exe.manifest

func main() {
	eaglemonitor.Initialize()
	thebrokenarm.Initialize()
	loading.Initialize()
	//print quicktask
	fmt.Println(loading.Data.Quicktask.Quicktask)
	fmt.Println(loading.Data.Tasks.Tasks)
	if !Dev {
		create.Initialize()
		auth.Initialize()
		version.Initialize()
		// console.Initialize() Only for windows
		presence.Initialize()
	}

	utils.Banner()
	auth.Welcome()
	utils.Site()

	for {
		if !Run {
			index := utils.Menu()
			data := site.Parsing(index)
			for _, taskUUID := range loading.Data.Tasks.Tasks[data] {
				fmt.Println("Task UUID: ", taskUUID)
				taskObject, err := task.GetTask(taskUUID)

				if err != nil {
					fmt.Println("Failed to get task: ", err.Error())
					continue
				}

				if !taskObject.Active {
					go task_manager.RunTask(taskObject)
				} else if taskObject.Done {
					task_manager.StopTask(taskObject)
				}
			}
			Run = true
		}
	}
}
