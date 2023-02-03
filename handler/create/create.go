package create

import (
	"encoding/csv"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/eagle/handler/utils"
	"github.com/fatih/color"
)

func Initialize() {
	color.Magenta("Checking folders...")

	if _, err := os.Stat("proxies"); os.IsNotExist(err) {
		err := os.Mkdir("proxies", 0755)
		if err != nil {
			log.Fatal(err)
		}
		_, err = os.Create("proxies/proxies.txt")
		if err != nil {
			log.Fatal(err)
		}

	}
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

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

	if _, err := os.Stat("profiles.csv"); os.IsNotExist(err) {
		_, err := os.Create("profiles.csv")
		if err != nil {
			log.Fatal(err)
		}
		err = ioutil.WriteFile("profiles.csv", CsvTemplate, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}

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

		_, err = os.Create("thebrokenarm/accounts.csv")
		if err != nil {
			log.Fatal(err)
		}

		err = ioutil.WriteFile("thebrokenarm/accounts.csv", CsvTemplateAccount, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}

	if _, err := os.Stat("MQT.csv"); os.IsNotExist(err) {
		csvFile, err := os.Create("MQT.csv")
		if err != nil {
			log.Fatal(err)
		}

		// err = ioutil.WriteFile("MQT.csv", CsvTemplateMQT, 0644)

		// if err != nil {
		// 	log.Fatal(err)
		// }

		//write content to file
		csvwriter := csv.NewWriter(csvFile)

		for _, empRow := range empData {
			_ = csvwriter.Write(empRow)
		}

		csvwriter.Flush()

	}

	os.Remove(dir + "/.DS_Store")

}

var empData = [][]string{
	{"Site", "Tasks Quantity", "Profiles", "Accounts (guest/accounts)", "Email", "Proxylist", "Payment Method", "Credit Card", "Other"},
	{"thebrokenarm", "5", "Manu", "guest", "@cathall", "test", "CC", "5638378250-01-29-128", "PIDLV2;7866;9267;9311"},
}
