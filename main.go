package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/eagle/handler/auth"
	"github.com/eagle/handler/loading"
	"github.com/eagle/handler/sites/deadstock"
	"github.com/eagle/handler/utils"
	"github.com/eagle/handler/version"
	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

var (
	JsonTemplate = []byte(`{
  "key": "INSERT_YOUR_KEY_HERE",
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

func init() {
	err := godotenv.Load("config/.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func Menu() {
	mode := utils.SelectMode(color.WhiteString("[Eagle " + version.Version + "]" + " [" + time.Now().Format("15:04:05.000000") + "]" + " PLEASE SELECT SITE:"))
	if mode == "1" {
		print("GAMESTOP")
	} else if mode == "2" {
		deadstock.Menu_deadstock()
	} else if mode == "3" {
		print("DADSTOCK")
	} else {
		utils.ConsolePrint("INVALID OPTION!", "red")
	}
}

func Creating() {
	color.Magenta("[" + time.Now().Format("15:04:05.000000") + "] " + "CREATING FOLDERS...")

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

func main() {
	Creating()
	loading.Initialize()
	auth.Initialize()
	version.Updates()
	utils.GetVersionName()
	// rich_presence.Initialize() TO fiX
	utils.Banner()
	username := strings.ToUpper(auth.Auth.Integrations.Discord.Username)
	color.Red("WELCOME BACK " + color.WhiteString(username))
	println("\n")
	utils.Site_list()
	Menu()
}
