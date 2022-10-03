package deadstock

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
	if mode == "1" {
		print("GAMESTOP")
	}
}
