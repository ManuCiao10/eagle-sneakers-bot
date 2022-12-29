package main

import (
	"fmt"
	"os"

	"github.com/eagle/handler/auth"
	"github.com/eagle/handler/loading"
	"github.com/eagle/handler/modules/thebrokenarm"
	"github.com/eagle/handler/site"
	"github.com/eagle/handler/task"
	"github.com/eagle/handler/task_manager"
	"github.com/eagle/handler/utils"
)

var (
	sites = map[int]string{
		1: "thebrokenarm",
		2: "monitor",
	}
)

//go:generate goversioninfo -skip-versioninfo=true -icon=handler/create/favicon.ico -manifest=handler/create/file.exe.manifest

func main() {
	thebrokenarm.Init()

	// create.Initialize()
	loading.Initialize()
	// auth.Initialize()
	// version.Initialize()
	// console.Initialize()
	// presence.Initialize()

	utils.Banner()
	auth.Welcome()
	utils.Site()

	for {
		index := utils.Menu()
		if index == utils.ERROR {
			utils.ConsolePrint("INVALID OPTION!", "red")
			os.Exit(0)
		}
		data := site.Parsing(index) //--> thebrokenarm,1

		for _, taskUUID := range loading.Data.Tasks.Tasks[data] {
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
	}

}
