package create

import (
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/eagle/handler/utils"
	"github.com/fatih/color"
)

func Initialize() {
	color.Magenta("[" + time.Now().Format("15:04:05.000000") + "] " + "CHECKING FOLDERS...")

	// if _, err := os.Stat("Proxies"); os.IsNotExist(err) {
	// 	err := os.Mkdir("Proxies", 0755)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	_, err = os.Create("Proxies/proxies.txt")
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }
	// dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	if _, err := os.Stat("settings.json"); os.IsNotExist(err) {
		_, err := os.Create("settings.json")

		if err != nil {
			log.Fatal(err)
		}

		if utils.Dev {
			ioutil.WriteFile("settings.json", JsonTemplateDEV, 0644)
		} else {
			err = ioutil.WriteFile("settings.json", JsonTemplate, 0644)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	// if _, err := os.Stat("profiles.csv"); os.IsNotExist(err) {
	// 	_, err := os.Create("profiles.csv")
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	err = ioutil.WriteFile("profiles.csv", CsvTemplate, 0644)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }

	if _, err := os.Stat("thebrokenarm"); os.IsNotExist(err) {
		err := os.Mkdir("thebrokenarm", 0755)
		if err != nil {
			log.Fatal(err)
		}

		_, err = os.Create("thebrokenarm/tasks.csv")
		if err != nil {
			log.Fatal(err)
		}

		err = ioutil.WriteFile("thebrokenarm/tasks.csv", CsvTemplateTask, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}

	// os.Remove(dir + "/.DS_Store")

	// time.Sleep(10 * time.Second)
}
