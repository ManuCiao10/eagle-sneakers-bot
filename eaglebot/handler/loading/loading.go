package loading

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var settings Settings

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	err = json.Unmarshal(byteValue, &settings.Settings)
	if err != nil {
		return nil
	}
	return &settings
}
