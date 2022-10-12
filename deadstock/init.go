package deadstock

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

type Product struct {
	Pid         string
	Size        string
	Email       string
	profile     string
	method      string
	Card_Number string
	Month       string
	Year        string
	CVV         string
	Proxy_List  string
}

type Info struct {
	Profile_name string
	First_name   string
	Last_name    string
	Phone        string
	Address      string
	Address_2    string
	House_Number string
	City         string
	State        string
	ZIP          string
	Country      string
}

func Menu_deadstock() {
	fmt.Print("\033[H\033[2J")
	utils.Logo()
	Read_file()
	mode := utils.SelectMode("[Eagle 0.0.2]" + "[" + time.Now().Format("15:04:05.000000") + "]" + " PLEASE SELECT CSV:")
	Find_index_of_csv(mode)
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

func Create_list(data [][]string) []Product {
	var list []Product
	for i, each := range data {
		if i != 0 {
			list = append(list, Product{
				Pid:         each[0],
				Size:        each[1],
				Email:       each[2],
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
	Check_product(list)
	return list
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
	csvFile, _ := os.Open("./deadstock_task/" + filename)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	data, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	data_list := Create_list(data)
	for _, each_line := range data_list {
		Run_Module(each_line)
	}
	defer csvFile.Close()
}

func Run_Module(each Product) {
	var profile []Info

	csvFile, _ := os.Open("./profiles.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	data, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	for _, each_line := range data {
		if each_line[0] == each.profile {
			profile = append(profile, Info{
				Profile_name: each_line[0],
				First_name:   each_line[1],
				Last_name:    each_line[2],
				Phone:        each_line[3],
				Address:      each_line[4],
				Address_2:    each_line[5],
				House_Number: each_line[6],
				City:         each_line[7],
				State:        each_line[8],
				ZIP:          each_line[9],
				Country:      each_line[10],
			})
		}
	}

	defer csvFile.Close()
	Check_profile(profile)
	Module_deadstock(profile)

}

