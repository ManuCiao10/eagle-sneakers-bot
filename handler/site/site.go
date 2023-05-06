package site

import (
	"fmt"
	"os"
	"strconv"

	"github.com/eagle/handler/logs"
	"github.com/eagle/handler/utils"
	"github.com/eagle/handler/version"
)

var (
	sites = map[int]string{
		1: "thebrokenarm",
		2: "adidas",
		3: "nike",
	}
)

func Validing(csv string, name string) string {
	path := utils.Path()

	intVar, err := strconv.Atoi(csv)
	if err != nil {
		logs.LogsMsgErr("invalid task")
	}

	var i = 1
	files, err := os.ReadDir(path + "/" + name)
	if err != nil {
		logs.LogsMsgErr("invalid task")
	}

	for _, f := range files {
		if f.Name() != "accounts.csv" && f.Name() != ".DS_Store" {
			if i == intVar {
				return f.Name()
			}
			i++
		}
	}
	return "unexpected"
}

func Parsing(site int) string {
	sitename, err := sites[site]
	if !err {
		logs.LogsMsgErr("invalid option")
	}

	fmt.Print("\033[H\033[2J")
	utils.Banner()
	utils.Directory(sitename)
	csv_index := utils.SelectMode(version.GetVersion() + logs.Time() + "PLEASE SELECT CSV:")

	t_name := Validing(csv_index, sitename)

	if t_name == "unexpected" {
		logs.LogsMsgErr("invalid selection")
	}
	task_type := fmt.Sprint(sitename, ",", csv_index)

	return task_type

}
