package thebrokenarm

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/eagle/handler/task"
)

var (
	taskMutex = sync.RWMutex{}
	tasks     = make(map[int]*task.Task)
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

func CreateTask(index int, mode, pid, size, mail, profile, payment, cardNum, month, year, cvv, proxy_list string) {
	taskMutex.Lock()
	defer taskMutex.Unlock()

	tasks[index] = &task.Task{
		Mode:        strings.ToLower(mode),
		Pid:         pid,
		Size:        strings.ToLower(strings.TrimSpace(size)),
		Email:       strings.ToLower(mail),
		Profile:     profile,
		Method:      strings.ToLower(payment),
		Card_Number: cardNum,
		Month:       month,
		Year:        year,
		CVV:         cvv,
		Proxy_List:  strings.Split(proxy_list, ".")[0],
	}

}

func ProxyToUrl(proxy string) string {
	proxySplit := strings.Split(proxy, ":")

	if len(proxySplit) == 2 {
		return fmt.Sprintf("http://%s:%s", proxySplit[0], proxySplit[1])
	} else if len(proxySplit) == 4 {
		return fmt.Sprintf("http://%s:%s@%s:%s", proxySplit[2], proxySplit[3], proxySplit[0], proxySplit[1])
	}

	return fmt.Sprintf("http://%s", proxy)
}
