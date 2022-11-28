package main

import (
	"log"
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

func init() {
	err := godotenv.Load("config/.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	loading.Initialize()
	auth.Initialize()
	version.Updates()

	// rich_presence.Initialize()
	//launch EageBot.exe from bin reposotory

	utils.Banner()
	username := strings.ToUpper(auth.Auth.Integrations.Discord.Username)
	color.Red("WELCOME BACK " + color.WhiteString(username))
	println("\n")
	utils.Site_list()
	// utils.Menu()
	mode := utils.SelectMode(color.HiWhiteString("[Eagle " + version.Version + "]" + "[" + time.Now().Format("15:04:05.000000") + "]" + " PLEASE SELECT SITE:"))
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
