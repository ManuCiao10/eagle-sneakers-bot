package thebrokenarm

import (
	"fmt"
	"os"
	"time"

	"github.com/eagle/handler/utils"
	"github.com/eagle/handler/version"
	"github.com/fatih/color"
)

func Loading() {
	fmt.Print("\033[H\033[2J")
	utils.Banner()
	utils.Directory("thebrokenarm")

	csv_index := utils.SelectMode("[Eagle " + version.Version + "] " + "[" + time.Now().Format("15:04:05.000000") + "]" + " PLEASE SELECT CSV:")
	task_name := CvsIndex(csv_index, "thebrokenarm")
	if task_name == "UNEXPECTED" {
		err_("UNEXPECTED ERROR")
	}

	CvsInfo(task_name, "thebrokenarm")

	for _, t := range tasks {
		Initialize(t)
	}

}

func Initialize(t *Task) TaskState {
	if !Contains([]string{"login", "normal"}, t.Mode) {
		err_("MODE IS NOT SUPPORTED FOR THIS SITE")
		return ErrorTaskState
	}

	//proxies

	// client, err := client.NewClient()

	// if err != nil {
	// 	err_("CLIENT ERROR")
	// }

	return ContinueTaskState

}

func err_(msg string) {
	color.Red(msg)
	os.Exit(0)
}
