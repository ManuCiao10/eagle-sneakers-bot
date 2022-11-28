package loading

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	// "fmt"
)

var Data Config

func Initialize() {
	Data = *Load()
}

func Load() *Config {
	return &Config{
		Settings: *loadSettings(),
	}
}

func loadSettings() *Settings {
	jsonFile, err := os.Open("eaglebot/bin/release/setting.json")
	// if we os.Open returns an error then handle it
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
