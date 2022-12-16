package loading

import (
	"embed"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/eagle/handler/profile"
	"github.com/eagle/handler/settings"
)

var (
	Data       Config
	proxyMutex = sync.RWMutex{}
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
	}
}

func loadProxies() *Proxies {
	proxyMutex.RLock()
	defer proxyMutex.RUnlock()

	files, err := os.ReadDir("./proxies")
	if err != nil {
		log.Fatal(err)
	}

	var proxies Proxies

	for _, file := range files {
		content, err := os.ReadFile("proxies/" + file.Name())
		if err != nil {
			log.Fatal(err)
		}

		proxies.Proxies = append(proxies.Proxies, settings.Proxie{
			ID:        strings.Split(file.Name(), ".")[0],
			ProxyList: string(content),
		})

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
