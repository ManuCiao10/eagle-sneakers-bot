package loading

import (
	"bufio"
	"embed"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/eagle/handler/profile"
	"github.com/eagle/handler/settings"
	"github.com/eagle/handler/task"
)

var (
	Data       Config
	proxyMutex = sync.RWMutex{}

	taskMutex = sync.RWMutex{}
	tasks     = make(map[int]*task.Task)
	Dev       = true
	array     []string
)

func Initialize() {
	Data = *Load()
}

func Load() *Config {
	return &Config{
		// Settings: *loadSettings(),
		Env:      *loadEnv(),
		Profiles: *loadProfiles(),
		Proxies:  *loadProxies(),
		Tasks:    *loadTask(),
	}
}

func CreateTask(index int, tasktype, mode, pid, size, mail, Profile, payment, cardNum, month, year, cvv, proxy_list string) {
	taskMutex.Lock()
	defer taskMutex.Unlock()

	tasks[index] = &task.Task{
		TaskType:    strings.ToLower(tasktype),
		Mode:        strings.ToLower(mode),
		Pid:         pid,
		Size:        strings.ToLower(strings.TrimSpace(size)),
		Email:       strings.ToLower(mail),
		Profile:     Profile,
		Method:      strings.ToLower(payment),
		Card_Number: cardNum,
		Month:       month,
		Year:        year,
		CVV:         cvv,
		Proxy_List:  strings.Split(proxy_list, ".")[0],
	}
	// fmt.Println(tasks[index])
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// func Getprofile(name string) profile.Profile {
// 	for _, p := range Data.Profiles.Profiles {
// 		fmt.Println(p)
// 		if p.ID == name {
// 			return p
// 		}
// 	}

// 	return profile.Profile{
// 		ID: "not_found",
// 	}
// }

func loadTask() *Tasks {
	paths := PathTask()

	var tasks Tasks
	tasks.Tasks = make(map[int][]string)

	for _, path := range paths {
		index := 1

		folder := strings.Split(path, "/")[0]
		if contains(array, folder) {
			index = index + 1
		} else {
			array = append(array, folder)
		}

		csvFile, err := os.Open(path)
		if err != nil {
			log.Fatal("ERROR OPENING FILE")
		}
		reader := csv.NewReader(bufio.NewReader(csvFile))
		task, err := reader.ReadAll()
		if err != nil {
			log.Fatal("ERROR READING FILE")
		}
		defer csvFile.Close()

		// if len(task) <= 1 {
		// 	log.Fatal("FILE " + strings.ToUpper(path) + " IS EMPTY")
		// }

		tasktype := fmt.Sprint(folder, ",", strconv.Itoa(index))
		taskQuantity := len(task)
		for i := 0; i <= taskQuantity-1; i++ {
			if i != 0 {
				CreateTask(i,
					tasktype,
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
					// Getprofile(task[i][4]),
				)
			}
		}

		csvFile.Close()
	}

	return &tasks
}

func loadProxies() *Proxies {
	proxyMutex.RLock()
	defer proxyMutex.RUnlock()

	files, err := os.ReadDir("./proxies")
	if err != nil {
		log.Fatal(err)
	}

	var proxies Proxies

	for _, fileName := range files {
		file, err := os.Open("proxies/" + fileName.Name())
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		contenet := CreateSliceProxy(scanner)
		proxies.Proxies = append(proxies.Proxies, settings.Proxie{
			ID:        strings.Split(fileName.Name(), ".")[0],
			ProxyList: contenet,
		})

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

	}
	return &proxies

}

//go:embed config.json
var JsonTemplate embed.FS

func loadEnv() *Env {
	env, _ := JsonTemplate.ReadFile("config.json")

	var envs Env

	err := json.Unmarshal(env, &envs.Env)
	if err != nil {
		fmt.Println(err)
	}
	return &envs

}

func loadSettings() *Settings {
	jsonFile, err := os.Open("settings.json")
	// jsonFile, err := os.Open("handler/loading/settings.json")

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()
	byteValue, _ := io.ReadAll(jsonFile)

	var settings Settings

	err = json.Unmarshal(byteValue, &settings.Settings)
	if err != nil {
		return nil
	}
	return &settings
}

func loadProfiles() *Profiles {
	f, err := os.Open("profiles.csv")
	if err != nil {
		log.Fatal(err)
	}

	var profiles Profiles

	defer f.Close()

	csvReader := csv.NewReader(f)
	c := 0
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if c == 0 {
			c += 1
			continue
		}

		profiles.Profiles = append(profiles.Profiles, profile.Profile{
			ID:           rec[0],
			FirstName:    rec[1],
			LastName:     rec[2],
			MobileNumber: rec[3],
			Address:      rec[4],
			Address2:     rec[5],
			HouseNumber:  rec[6],
			City:         rec[7],
			State:        rec[8],
			Zip:          rec[9],
			Country:      rec[10],
		})
	}
	return &profiles
}

func PathTask() []string {
	var folder []string
	var paths []string

	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	if Dev {
		for _, f := range files {
			if f.IsDir() && f.Name() != ".git" && f.Name() != "proxies" && f.Name() != "handler" {
				folder = append(folder, f.Name())
			}
		}
	} else {
		for _, f := range files {
			if f.IsDir() && f.Name() != "proxies" {
				folder = append(folder, f.Name())
			}
		}
	}

	for _, site := range folder {
		files, err := os.ReadDir(site)
		if err != nil {
			log.Fatal(err)
		}

		for _, fileName := range files {
			paths = append(paths, site+"/"+fileName.Name())
		}
	}

	return paths // return all the paths
}

func CreateSliceProxy(scanner *bufio.Scanner) []string {
	var proxies []string

	for scanner.Scan() {
		proxies = append(proxies, scanner.Text())
	}

	return proxies
}
