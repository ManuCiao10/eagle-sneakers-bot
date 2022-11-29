package create

import (
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/fatih/color"
)

var (
	JsonTemplate = []byte(`{
  "key": "EAGLE-LD9W-CJ3K-NAO7-KFOV",
  "webhook": "INSERT_YOUR_WEBHOOK",
		  
  "2captcha_key": "INSERT_YOUR_2CAPTCHA_KEY",
  "anticaptcha_key": "INSERT_YOUR_ANTICAPTCHA_KEY",
  "capmonster_key": "INSERT_YOUR_CAPMONSTER_KEY",
		  
  "solver": "SELECT_YOUR_SOLVER",
		  
  "delay": {
    "retry": "DELAY",
    "timeout": "DELAY"
  }
}`)

	CsvTemplate     = []byte(`Profile Name,First Name,Last Name,Mobile Number,Address,Address 2,House Number,City,State,ZIP,Country,Billing First Name,Billing Last Name,Billing Mobile Number,Billing Address,Billing Address 2,Billing Address 3,Billing House Number,Billing City,Billing State,Billing ZIP,Billing Country`)
	CsvTemplateTask = []byte(`Url / PID,Size,E-mail,Profile,Payment Method,Card Number,Month,Year,CVV,Proxy List`)
)

func Initialize() {
	color.Magenta("[" + time.Now().Format("15:04:05.000000") + "] " + "CHECKING FOLDERS...")

	if _, err := os.Stat("EagleBot"); os.IsNotExist(err) {
		err := os.Mkdir("EagleBot", 0755)
		if err != nil {
			log.Fatal(err)
		}
	}
	if _, err := os.Stat("EagleBot/Proxies"); os.IsNotExist(err) {
		err := os.Mkdir("EagleBot/Proxies", 0755)
		if err != nil {
			log.Fatal(err)
		}
		_, err = os.Create("EagleBot/Proxies/proxies.txt")
		if err != nil {
			log.Fatal(err)
		}
	}
	/* get current directory
		dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	    if err != nil {
	            log.Fatal(err)
	    }
	    fmt.Println(dir)
	*/

	if _, err := os.Stat("EagleBot/settings.json"); os.IsNotExist(err) {
		_, err := os.Create("EagleBot/settings.json")
		if err != nil {
			log.Fatal(err)
		}

		err = ioutil.WriteFile("EagleBot/settings.json", JsonTemplate, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}

	if _, err := os.Stat("EagleBot/profiles.csv"); os.IsNotExist(err) {
		_, err := os.Create("EagleBot/profiles.csv")
		if err != nil {
			log.Fatal(err)
		}
		err = ioutil.WriteFile("EagleBot/profiles.csv", CsvTemplate, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}

	if _, err := os.Stat("EagleBot/Zara"); os.IsNotExist(err) {
		err := os.Mkdir("EagleBot/Zara", 0755)
		if err != nil {
			log.Fatal(err)
		}

		_, err = os.Create("EagleBot/Zara/tasks.csv")
		if err != nil {
			log.Fatal(err)
		}

		err = ioutil.WriteFile("EagleBot/Zara/tasks.csv", CsvTemplateTask, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
	// time.Sleep(10 * time.Second)
}
