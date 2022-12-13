package thebrokenarm

import (
	"fmt"
	"os"
	"time"

	"github.com/eagle/handler/utils"
	"github.com/eagle/handler/version"
	"github.com/fatih/color"
)

func Initialize() {
	fmt.Print("\033[H\033[2J")
	utils.Banner()
	utils.Directory("thebrokenarm")

	csv_index := utils.SelectMode("[Eagle " + version.Version + "] " + "[" + time.Now().Format("15:04:05.000000") + "]" + " PLEASE SELECT CSV:")
	task_name := CvsIndex(csv_index, "thebrokenarm")
	if task_name == "UNEXPECTED" {
		err_("UNEXPECTED ERROR")
	}

	CvsInfo(task_name, "thebrokenarm")

	Thebrokenarm()

}

func Thebrokenarm() {

	//use range to loop through map
	for _, i := range tasks {
		go test(i)
	}

}

func test(i *Task) {
	fmt.Println("task ["+"]", i)
	time.Sleep(1 * time.Second)

}

func err_(msg string) {
	color.Red(msg)
	os.Exit(0)
}
