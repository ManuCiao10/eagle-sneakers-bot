package deadstock

// DEADSTOCK MODULE

import (
	"Eagle/utils"
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/fatih/color"
)

type Deadstock struct {
	Pid         string
	Size        string
	Emai        string
	profile     string
	method      string
	Card_Number string
	Month       string
	Year        string
	CVV         string
	Proxy_List  string
}

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
			Read_csv_info(f.Name())
		}
	}
}

func Read_csv_info(filename string) {
	// Read csv file
	csvFile, _ := os.Open("./deadstock_task/" + filename)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	data, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	data_list := Create_list(data)
	for i, each := range data_list {
		fmt.Println(each)
	}
}

func Create_list(data [][]string) []Deadstock {
	var list []Deadstock
	for i, each := range data {
		if i > 0 {
			list = append(list, Deadstock{
				Pid:         each[0],
				Size:        each[1],
				Emai:        each[2],
				profile:     each[3],
				method:      each[4],
				Card_Number: each[5],
				Month:       each[6],
				Year:        each[7],
				CVV:         each[8],
				Proxy_List:  each[9],
			})
		}
	}
	return list
}
