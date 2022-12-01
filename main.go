package main

import (
	"strings"
	"time"

	"github.com/eagle/handler/auth"
	"github.com/eagle/handler/create"
	"github.com/eagle/handler/loading"
	"github.com/eagle/handler/rich_presence"
	"github.com/eagle/handler/sites/deadstock"
	"github.com/eagle/handler/utils"
	"github.com/eagle/handler/version"
	"github.com/fatih/color"
)

func Menu() {
	mode := utils.SelectMode(color.MagentaString("[Eagle " + version.Version + "]" + " [" + time.Now().Format("15:04:05.000000") + "]" + color.WhiteString(" PLESE SELECT A SITE:")))
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

func main() {
	//delete files config.json from mods
	//fix where the bot is checking for the Version of the bot for the update
	//add all the icon and the img or banner to the bot
	create.Initialize()
	loading.Initialize()
	auth.Initialize()
	version.Updates()
	utils.GetVersionName()
	rich_presence.Initialize()
	// console.Display()
	utils.Banner()
	username := strings.ToUpper(auth.Auth.Integrations.Discord.Username)
	//add tab to the menu
	color.Magenta("WELCOME BACK  \t" + color.WhiteString(username))
	println("\n")
	utils.Site_list()
	Menu()
}
