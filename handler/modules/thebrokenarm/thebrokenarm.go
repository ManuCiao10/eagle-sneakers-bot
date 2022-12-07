package thebrokenarm

import (
	"fmt"
	"os"
	"time"

	"github.com/eagle/handler/utils"
	"github.com/fatih/color"
)

func Initialize() {
	fmt.Print("\033[H\033[2J")
	utils.Banner()
	utils.Directory("thebrokenarm")

	csv_index := utils.SelectMode("[Eagle 0.0.2]" + "[" + time.Now().Format("15:04:05.000000") + "]" + " PLEASE SELECT CSV:")
	task_name := CvsIndex(csv_index, "thebrokenarm")
	if task_name == "UNEXPECTED" {
		err_("UNEXPECTED ERROR")
	}

	CvsInfo(task_name, "thebrokenarm")
	CvsProfile("profiles.csv")

	//RUN TASK MODULES

}

func err_(msg string) {
	color.Red(msg)
	os.Exit(0)
}

// func print_struct() {
// 	fmt.Print(list)
// 	fmt.Print(profile)
// 	fmt.Print(profile[0].Profile_name)
// 	fmt.Println("PID: ", list[0].Pid)
// }
