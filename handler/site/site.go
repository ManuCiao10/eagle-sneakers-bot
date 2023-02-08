package site

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/eagle/handler/logs"
	"github.com/eagle/handler/utils"
	"github.com/eagle/handler/version"
	"github.com/mitchellh/go-homedir"
)

var (
	sites = map[int]string{
		1: "thebrokenarm",
		2: "fiver",
		3: "nike",
		4: "monitor",
	}
)

func Validing(csv string, name string) string {
	path := Path()
	// fmt.Println(csv, name)

	intVar, err := strconv.Atoi(csv)
	if err != nil {
		err_("invalid task")
	}
	// fmt.Println(intVar)

	var i = 1
	files, err := os.ReadDir(path + "/" + name)
	if err != nil {
		err_("invalid task")
	}

	for _, f := range files {
		if f.Name() != "accounts.csv" && f.Name() != ".DS_Store" {
			if i == intVar {
				return f.Name()
			}
			i++
		}

	}

	// files, err := os.ReadDir(path + "/" + name)
	// if err != nil {
	// 	err_("invalid task")
	// }
	//remove account.csv and .DS_Store from files

	// for i, f := range files {
	// 	i = i + 1

	// 	if i == intVar {
	// 		return f.Name()
	// 	}

	// }
	return "UNEXPECTED"
}

func Parsing(site int) string {
	if site == utils.ERROR {
		logs.LogsMsgErr("invalid option")
	}

	if site != utils.MONITOR {
		fmt.Print("\033[H\033[2J")
		utils.Banner()
		utils.Directory(sites[site])
		csv_index := utils.SelectMode(version.GetVersion() + logs.Time() + "PLEASE SELECT CSV:")

		t_name := Validing(csv_index, sites[site])
		fmt.Println("t_name ==>" + t_name)

		if t_name == "UNEXPECTED" {
			err_("INVALID SELECTION")
		}
		task_type := fmt.Sprint(sites[site], ",", csv_index)
		fmt.Println("task_type ==>" + task_type)

		return task_type //--> site,index_csv
	}

	return "monitor"

}

func err_(err string) {
	logs.LogsMsgErr(err)
}

func Path() string {
	dir, err := homedir.Dir()
	if err != nil {
		log.Fatal(err)
	}

	path := dir + "/Desktop/EagleBot"
	return path

}
