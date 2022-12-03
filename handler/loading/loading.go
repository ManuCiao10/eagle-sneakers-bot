package loading

import (
	"embed"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

var Data Config

func Initialize() {
	Data = *Load()
}

func Load() *Config {
	return &Config{
		Settings: *loadSettings(),
		Env:      *loadEnv(),
	}
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
