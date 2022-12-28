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

	"github.com/eagle/handler/profile"
	"github.com/eagle/handler/settings"
	task_ "github.com/eagle/handler/task"
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
		// Settings: *loadSettings(),
		Env:      *loadEnv(),
		Profiles: *loadProfiles(),
		Proxies:  *loadProxies(),
		Tasks:    *loadTask(),
	}
}

func Trim(s string) string {
	return strings.TrimSpace(s)
}

func loadTask() *Tasks {
	paths := task_.PathTask()

	var tasks Tasks
	tasks.Tasks = make(map[string][]string)

	for _, path := range paths {
		index := 1

		folder := strings.Split(path, "/")[0]
		if task_.Contains(array, folder) {
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

		tasktype := fmt.Sprint(folder, ",", strconv.Itoa(index))
		taskQuantity := len(task)

		for i := 0; i < taskQuantity; i++ {
			if i != 0 {
				taskUUID := task_.CreateTask(
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

func CreateSliceProxy(scanner *bufio.Scanner) []string {
	var proxies []string

	for scanner.Scan() {
		proxies = append(proxies, scanner.Text())
	}

	return proxies
}
