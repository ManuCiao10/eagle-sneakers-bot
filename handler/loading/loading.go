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
		// Version:  *loadVersion(), //starting loading version reading from bin folder
	}
}

// func loadVersion() *Version {
// 	//read in bin folder

// }

func loadSettings() *Settings {
	jsonFile, err := os.Open("bin/setting.json")
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
