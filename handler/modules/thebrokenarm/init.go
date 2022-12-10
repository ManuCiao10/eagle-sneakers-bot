package thebrokenarm

import (
	"bufio"
	"encoding/csv"
	"os"
	"strconv"
	"sync"
)

var (
	taskMutex = sync.RWMutex{}
	tasks               = make(map[string]Task)
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

type Tasks struct {
	Tasks map[int][]string
}

func CvsInfo(filename string, name string) []Product {
	var tasks Tasks
	tasks.Tasks = make(map[int][]string)

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

	if len(data) <= 1 {
		err_("TASK FILE " + filename + " IS EMPTY")
	}
	//read csv file
	taskQuantity := len(data) - 1

	for i := 0; i < taskQuantity; i++ {
		if i != 0 {
			CreateTask(
				data[i][0],
				data[i][1],
				data[i][2],
				data[i][3],
				data[i][4],
				data[i][5],
				data[i][6],
				data[i][7],
				data[i][8],
				data[i][9],
			)
		}
	}

	return list
}

func CreateTask() {
	taskMutex.Lock()
	defer taskMutex.Unlock()

	tasks[index] = Task{

}

// func CvsProfile(filename string) []Info {

// 	csvFile, err := os.Open("./" + filename)
// 	if err != nil {
// 		err_("ERROR OPENING FILE")
// 	}
// 	reader := csv.NewReader(bufio.NewReader(csvFile))
// 	data, err := reader.ReadAll()
// 	if err != nil {
// 		err_("ERROR READING FILE")
// 	}

// 	defer csvFile.Close()

// 	for idx, each_line := range data {
// 		if idx != 0 {
// 			profile = append(profile, Info{
// 				Profile_name: each_line[0],
// 				First_name:   each_line[1],
// 				Last_name:    each_line[2],
// 				Phone:        each_line[3],
// 				Address:      each_line[4],
// 				Address_2:    each_line[5],
// 				House_Number: each_line[6],
// 				City:         each_line[7],
// 				State:        each_line[8],
// 				ZIP:          each_line[9],
// 				Country:      each_line[10],
// 			})
// 		}
// 	}

// 	if len(profile) == 0 {
// 		err_("PROFILE FILE " + filename + " IS EMPTY")
// 	}

// 	return profile
// }

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
