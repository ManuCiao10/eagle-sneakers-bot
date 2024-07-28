package create

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/eagle/handler/utils"
	"github.com/fatih/color"
)

func Initialize() {
	path := utils.Path()

	color.Magenta("Checking folders...")

	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, 0755)
		if err != nil {
			log.Fatal(err)
		}

		err = os.Mkdir(path+"/proxies", 0755)
		if err != nil {
			log.Fatal(err)
		}

		_, err = os.Create(path + "/proxies/proxies.txt")
		if err != nil {
			log.Fatal(err)
		}

		_, err = os.Create(path + "/settings.json")
		if err != nil {
			log.Fatal(err)
		}

		if utils.Dev {
			os.WriteFile(path+"/settings.json", JsonTemplateDEV, 0644)
		} else {
			err = os.WriteFile(path+"/settings.json", JsonTemplate, 0644)
			if err != nil {
				log.Fatal(err)
			}
		}

		_, err = os.Create(path + "/profiles.csv")
		if err != nil {
			log.Fatal(err)
		}

		err = os.WriteFile(path+"/profiles.csv", CsvTemplate, 0644)
		if err != nil {
			log.Fatal(err)
		}

		csvFile, err := os.Create(path + "/MQT.csv")
		if err != nil {
			log.Fatal(err)
		}
		csvwriter := csv.NewWriter(csvFile)

		for _, empRow := range empData {
			_ = csvwriter.Write(empRow)
		}

		csvwriter.Flush()

		err = os.Mkdir(path+"/thebrokenarm", 0755)
		if err != nil {
			log.Fatal(err)
		}
		_, err = os.Create(path + "/thebrokenarm/tasks.csv")
		if err != nil {
			log.Fatal(err)
		}

		err = os.WriteFile(path+"/thebrokenarm/tasks.csv", CsvTemplateTask, 0644)
		if err != nil {
			log.Fatal(err)
		}

		_, err = os.Create(path + "/thebrokenarm/accounts.csv")
		if err != nil {
			log.Fatal(err)
		}

		err = os.WriteFile(path+"/thebrokenarm/accounts.csv", CsvTemplateAccount, 0644)
		if err != nil {
			log.Fatal(err)
		}

		_ = os.Remove(path + "/thebrokenarm/.DS_Store")

		color.Green("EagleBot folder created...")

	}

}

var empData = [][]string{
	{"Site", "Tasks Quantity", "Profiles", "Accounts (guest/accounts)", "Email", "Proxylist", "Payment Method", "Credit Card", "Other"},
	{"thebrokenarm", "5", "Manu", "guest", "@cathall", "test", "CC", "563-01-29-128", "test"},
}
