package thebrokenarm

import (
	"bufio"
	"encoding/csv"
	"os"
	"strconv"
)

func CvsIndex(csv string, name string) string {
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

func CvsInfo(filename string, name string) []Product {
	csvFile, err := os.Open("./" + name + "/" + filename)
	if err != nil {
		err_("ERROR OPENING FILE")
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))
	data, err := reader.ReadAll()
	if err != nil {
		err_("ERROR READING FILE")
	}
	defer csvFile.Close()

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

	if len(list) == 0 {
		err_("TASK FILE " + filename + " IS EMPTY")
	}

	return list
}

func CvsProfile(filename string) []Info {

	csvFile, err := os.Open("./" + filename)
	if err != nil {
		err_("ERROR OPENING FILE")
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))
	data, err := reader.ReadAll()
	if err != nil {
		err_("ERROR READING FILE")
	}

	defer csvFile.Close()

	for idx, each_line := range data {
		if idx != 0 {
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
	if len(profile) == 0 {
		err_("PROFILE FILE " + filename + " IS EMPTY")
	}

	return profile

}

/*

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

func Write_data_to_file(data string, filename string) {
	f, err := os.Create(filename)
	if err != nil {
		// Print_err("FILE CREATION ERROR")
	}
	defer f.Close()
	f.WriteString(data)
}

*/
