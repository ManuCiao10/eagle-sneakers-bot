package cmd

import (
	"fmt"

	"github.com/eagle/handler/loading"
	"github.com/eagle/handler/site"
	"github.com/eagle/handler/task"
	"github.com/eagle/handler/task_manager"
	"github.com/eagle/handler/utils"
)

var Run bool

func Initialize() {
	for {
		if !Run {
			index := utils.Menu()
			data := site.Parsing(index)
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
			Run = true
		}
	}
}
