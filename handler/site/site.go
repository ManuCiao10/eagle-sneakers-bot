package site

import (
	"fmt"

	"github.com/eagle/handler/utils"
)

var (
	sites = map[int]string{
		1: "thebrokenarm",
		2: "monitor",
	}

	task_type string
)

func Parsing(site int) string {
	fmt.Print("\033[H\033[2J")
	utils.Banner()
	utils.Directory(sites[site])

	csv_index := utils.SelectMode(utils.Version() + utils.Time() + "PLEASE SELECT CSV:")
	task_type = fmt.Sprint(sites[site], ",", csv_index)

	return task_type
}
