package main

import (
	"strings"

	"github.com/eagle/handler/auth"
	"github.com/eagle/handler/create"
	"github.com/eagle/handler/loading"
	"github.com/eagle/handler/rich_presence"
	"github.com/eagle/handler/utils"
	"github.com/eagle/handler/version"
	"github.com/fatih/color"
)

func Welcome() {
	username := strings.ToUpper(auth.Auth.Integrations.Discord.Username)
	color.Magenta("WELCOME BACK  \t" + color.WhiteString(username))
	println("\n")
}

func main() {
	//fix where the bot is checking for the Version of the bot for the update
	//add all the icon to the bot embed
	create.Initialize()
	loading.Initialize()
	auth.Initialize()
	version.Updates()
	utils.GetVersionName()
	rich_presence.Initialize()
	// console.Display()
	utils.Banner()
	Welcome()
	utils.Site_list()
	utils.Menu()
}
