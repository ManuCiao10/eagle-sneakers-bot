package loading

import (
	"bufio"
	"embed"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/eagle/handler/account"
	"github.com/eagle/handler/profile"
	quicktask_handler "github.com/eagle/handler/quicktask"
	"github.com/eagle/handler/settings"
	task_handler "github.com/eagle/handler/task"
)

var (
	Data       Config
	proxyMutex = sync.RWMutex{}

	array []string
)

func Initialize() {
	Data = *Load()
}

func Load() *Config {
	return &Config{
		Settings:  *loadSettings(),
		Env:       *loadEnv(),
		Profiles:  *loadProfiles(),
		Proxies:   *loadProxies(),
		Tasks:     *loadTask(),
		Quicktask: *loadQuicktasks(),
		Accounts:  *loadAccounts(),
	}
}

func loadAccounts() *Accounts {
	paths := []string{
		"thebrokenarm/accounts.csv",
	}

	var accounts Accounts
	accounts.Accounts = make(map[int][]account.Account)

	for siteId, path := range paths {
		f, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}

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

			if rec[0] == "" || rec[1] == "" {
				continue
			}

			account.CreateAccount(&account.Account{
				SiteId:   siteId,
				Email:    rec[0],
				Password: rec[1],
			})
			accountObject, _ := account.GetAccount(siteId, rec[0])
			accounts.Accounts[siteId] = append(accounts.Accounts[siteId], *accountObject)

		}
		f.Close()
	}

	return &accounts
}

func loadQuicktasks() *Quicktask {
	var quicktask Quicktask
	quicktask.Quicktask = make(map[string][]string)

	csvFile, err := os.Open("MQT.csv")
	if err != nil {
		log.Fatal("error opening file MQT.csv")
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))

	task, err := reader.ReadAll()
	if err != nil {
		log.Fatal("error reading file: MQT.csv")
	}
	defer csvFile.Close()

	numberSite := len(task)
	for i := 0; i < numberSite; i++ {
		if i != 0 {
			taskQuantity, err := strconv.Atoi(task[i][1])
			if err != nil {
				log.Fatal("error reading taskQuantity: MQT.csv")
			}
			siteName := task[i][0]
			for t := 0; t < taskQuantity; t++ {
				quickUUID := quicktask_handler.CreateQuicktask(
					Trim(task[i][0]),
					Trim(task[i][1]),
					Trim(task[i][2]),
					Trim(task[i][3]),
					Trim(task[i][4]),
					Trim(task[i][5]),
					Trim(task[i][6]),
					Trim(task[i][7]),
					Trim(task[i][8]),
				)
				quicktask.Quicktask[siteName] = append(quicktask.Quicktask[siteName], quickUUID)
			}
		}
	}
	return &quicktask
}

func loadTask() *Tasks {
	paths := task_handler.PathTask()

	var tasks Tasks
	tasks.Tasks = make(map[string][]string)

	for _, path := range paths {
		index := 1

		type_ := strings.Split(path, "/")[0]
		if task_handler.Contains(array, type_) {
			index = index + 1
		} else {
			array = append(array, type_)
		}

		csvFile, err := os.Open(path)
		if err != nil {
			log.Fatal("error opening file")
		}
		reader := csv.NewReader(bufio.NewReader(csvFile))

		task, err := reader.ReadAll()
		if err != nil {
			log.Fatal("error reading file:", type_)
		}
		defer csvFile.Close()

		tasktype := fmt.Sprint(type_, ",", strconv.Itoa(index))
		taskQuantity := len(task)

		for i := 0; i < taskQuantity; i++ {
			if i != 0 {
				taskUUID := task_handler.CreateTask(
					strings.ToLower(tasktype),
					Trim(task[i][0]),
					Trim(task[i][1]),
					Trim(task[i][2]),
					Trim(task[i][3]),
					Trim(task[i][4]),
					Trim(task[i][5]),
					Trim(task[i][6]),
					Trim(task[i][7]),
					Trim(task[i][8]),
					Trim(task[i][9]),
					Trim(task[i][10]),
					Trim(type_),
				)
				tasks.Tasks[tasktype] = append(tasks.Tasks[tasktype], taskUUID)
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
