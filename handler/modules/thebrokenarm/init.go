package thebrokenarm

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

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
	//check if the porfiles or the info are meplty with a range for loop
	// Check_product(list)
	return list
}

// used
func Find_index_of_csv(csv string, name string) {
	intVar, err := strconv.Atoi(csv)
	if err != nil {
		fmt.Println(err)
	}
	files, err := os.ReadDir("./" + name)
	if err != nil {
		log.Fatal(err)
	}
	for i, f := range files {
		i = i + 1
		if i == intVar {
			Read_csv_info(f.Name(), "thebrokenarm")
		}
	}
}

func Read_csv_info(filename string, name string) {
	csvFile, err := os.Open("./" + name + "/" + filename)
	if err != nil {
		log.Fatal(err)
	}
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

	csvFile, err := os.Open("./profiles.csv")
	if err != nil {
		log.Fatal(err)
	}
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
	// Check_profile(profile)
	// Module_deadstock(profile)

}

// func timer(name string) func() {
// 	start := time.Now()
// 	return func() {
// 		fmt.Printf("%s took %v\n", name, time.Since(start))
// 	}
// }

// func Write_data_to_file(data string, filename string) {
// 	f, err := os.Create(filename)
// 	if err != nil {
// 		// Print_err("FILE CREATION ERROR")
// 	}
// 	defer f.Close()
// 	f.WriteString(data)
// }

// func RandomString(n int) string {
// 	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
// 	sb := strings.Builder{}
// 	sb.Grow(n)
// 	for i := 0; i < n; i++ {
// 		sb.WriteByte(charset[rand.Intn(len(charset))])
// 	}
// 	return sb.String()
// }
