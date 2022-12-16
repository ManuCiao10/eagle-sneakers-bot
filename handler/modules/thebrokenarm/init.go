package thebrokenarm

import (
	"bufio"
	"encoding/csv"
	"os"
	"strconv"
	"strings"
	"sync"
)

var (
	taskMutex = sync.RWMutex{}
	tasks     = make(map[int]*Task)
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

func CvsInfo(filename string, name string) {
	csvFile, err := os.Open("./" + name + "/" + filename)
	if err != nil {
		err_("ERROR OPENING FILE")
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))
	task, err := reader.ReadAll()
	if err != nil {
		err_("ERROR READING FILE")
	}
	defer csvFile.Close()

	if len(task) <= 1 {
		err_("FILE " + strings.ToUpper(filename) + " IS EMPTY")
	}

	taskQuantity := len(task)
	for i := 0; i < taskQuantity; i++ {
		if i != 0 {
			CreateTask(i,
				task[i][0],
				task[i][1],
				task[i][2],
				task[i][3],
				task[i][4],
				task[i][5],
				task[i][6],
				task[i][7],
				task[i][8],
				task[i][9],
				task[i][10],
			)
		}
	}
}

func CreateTask(index int, mode, pid, size, mail, profile, payment, cardNumber, month, year, cvv, proxy_list string) {
	taskMutex.Lock()
	defer taskMutex.Unlock()

	tasks[index] = &Task{
		Mode:        strings.ToLower(mode),
		Pid:         pid,
		Size:        size,
		Email:       mail,
		Profile:     profile,
		Method:      payment,
		Card_Number: cardNumber,
		Month:       month,
		Year:        year,
		CVV:         cvv,
		Proxy_List:  proxy_list,
	}

}

func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
