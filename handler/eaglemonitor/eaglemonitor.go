package eaglemonitor

import (
	"fmt"
	"strings"
	"time"

	"github.com/eagle/handler/auth"
	"github.com/eagle/handler/loading"
	"github.com/eagle/handler/logs"
	"github.com/eagle/handler/quicktask"
	"github.com/eagle/handler/utils"
)

func Initialize() {
	fmt.Print("\033[H\033[2J")
	utils.Banner()
	auth.Welcome()

	logs.LogsMsgCyan("waiting for monitor...")

	// getValues()
	pid := "PIDLV1"

	dataMonitor := "thebrokenarm"
	for {
		for _, taskUUID := range loading.Data.Quicktask.Quicktask[dataMonitor] {
			taskObject, err := quicktask.GetQuicktask(taskUUID)

			if err != nil {
				fmt.Println("Failed to get task: ", err.Error())
				continue
			}
			pidMqt := strings.Split(taskObject.Other, ";")
			allPidMqt = append(allPidMqt, pidMqt...)

			if Contains(allPidMqt, pid) {
				logs.LogsMsgCyan("restock detected!")
				monitorWebhook(&MonitorDetected{
					pid:          pid,
					size:         "L",
					taskQuantity: 15,
					proxy:        "proxy",
					taskFile:     "taskfile",
					delay:        1000,
					store:        dataMonitor,
				}, loading.Data.Settings.Settings.DiscordWebhook)

				//run task
				// if !taskObject.Active {
				// 	go task_manager.RunTask(taskObject)
				// } else if taskObject.Done {
				// 	task_manager.StopTask(taskObject)
				// send webhook terminate task
				// }
			}

		}
		time.Sleep(50 * time.Millisecond)

	}

}

func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
