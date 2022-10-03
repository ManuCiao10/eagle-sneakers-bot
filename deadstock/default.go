package deadstock

// DEADSTOCK MODULE

import (
	"Eagle/utils"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/fatih/color"
)

func Read_file() {
	files, err := os.ReadDir("./deadstock_task")
	if err != nil {
		log.Fatal(err)
	}
	for i, f := range files {
		i = i + 1
		s := strconv.Itoa(i)
		color.Red("[ Eagle 0.0.2 ]" + "[ " + time.Now().Format("15:04:05.000000") + " ] " + s + ". " + f.Name())
	}
	println("\n")
}

func Menu_deadstock() {
	fmt.Print("\033[H\033[2J")
	utils.Logo()
	Read_file()
	mode := utils.SelectMode("[ Eagle 0.0.2 ]" + "[ " + time.Now().Format("15:04:05.000000") + " ]" + " PLEASE SELECT CSV:")
	Find_index_of_csv(mode)
}

func Find_index_of_csv(mode string) {
	intVar, err := strconv.Atoi(mode)
	if err != nil {
		fmt.Println(err)
	}
	files, err := os.ReadDir("./deadstock_task")
	if err != nil {
		log.Fatal(err)
	}
	for i, f := range files {
		i = i + 1
		if i == intVar {
			Run_task(f.Name())
		}
	}
}

func Read_csv_info()  {
	
	
}

func Run_task(filename string) {
	fmt.Println(filename)
	Read_csv_info(filename)
}
