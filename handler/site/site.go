package site

import (
	"fmt"
	"os"
	"strconv"

	"github.com/eagle/handler/utils"
)

var (
	sites = map[int]string{
		1: "thebrokenarm",
		2: "monitor",
	}

	task_type string
)

func Validing(csv string, name string) string {
	intVar, err := strconv.Atoi(csv)
	if err != nil {
		err_("INVALID SELECTION")
	}
	files, err := os.ReadDir("./" + name)
	if err != nil {
		err_("INVALID TASK")
	}
	for i, f := range files {
		i = i + 1
		if i == intVar {
			return f.Name()
		}
	}
	return "UNEXPECTED"
}

func err_(err string) {
	utils.ConsolePrint(err, "red")
	os.Exit(0)
}

func Parsing(site int) string {
	fmt.Print("\033[H\033[2J")
	utils.Banner()
	utils.Directory(sites[site])

	csv_index := utils.SelectMode(utils.Version() + utils.Time() + "PLEASE SELECT CSV:")

	t_name := Validing(csv_index, sites[site])
	if t_name == "UNEXPECTED" {
		err_("INVALID SELECTION")
	}
	task_type = fmt.Sprint(sites[site], ",", csv_index)

	return task_type //--> site,index_csv
}
