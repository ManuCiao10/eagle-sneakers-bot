package task

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/lithammer/shortuuid"
	"github.com/mitchellh/go-homedir"
)

var (
	taskMutex               = sync.RWMutex{}
	ErrTaskDoesNotExist     = errors.New("task does not exist")
	tasks                   = make(map[string]*Task)
	Dev                     = true
	taskTypes               = make(map[string]*TaskType)
	ErrTaskTypeDoesNotExist = errors.New("task type does not exist")
)

func CreateTask(tasktype, mode, pid, size, mail, Profile, payment, cardNum, month, year, cvv, proxy_list, type_ string) string {
	taskMutex.Lock()
	defer taskMutex.Unlock()

	id := shortuuid.New()

	size = strings.ToLower(size)
	if size == "random" {
		size = RandomSize()
	} else {
		size = SplitSize(size)
	}

	tasks[id] = &Task{
		TaskType:      tasktype,
		Mode:          strings.ToLower(mode),
		Pid:           strings.ToLower(pid),
		Size:          size,
		Email:         strings.ToLower(mail),
		Profile:       Profile,
		Method:        strings.ToLower(payment),
		Card_Number:   cardNum,
		Month:         month,
		Year:          year,
		CVV:           cvv,
		CheckoutProxy: strings.Split(proxy_list, ".")[0],
		Type:          strings.ToLower(type_),
	}
	return id
}

func PathTask() []string {
	path := findPath()

	var folder []string
	var paths []string

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if f.IsDir() && f.Name() != "proxies" {
			folder = append(folder, f.Name())
		}
	}

	for _, site := range folder {
		files, err := os.ReadDir(path + "/" + site)
		if err != nil {
			log.Fatal(err)
		}

		for _, fileName := range files {
			if fileName.Name() != "accounts.csv" && fileName.Name() != ".DS_Store" {
				paths = append(paths, site+"/"+fileName.Name())
			}
		}
	}

	return paths // return all the paths
}

func findPath() string {
	dir, err := homedir.Dir()
	if err != nil {
		log.Fatal(err)
	}

	path := dir + "/Desktop/EagleBot"

	return path
}
