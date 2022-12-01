package main

import (
	"strings"

	"github.com/eagle/handler/auth"
	"github.com/eagle/handler/console"
	"github.com/eagle/handler/create"
	"github.com/eagle/handler/loading"
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
	//add all the favicon.icon to the bot embed
	create.Initialize()
	loading.Initialize()
	auth.Initialize()
	version.Updates()
	console.Initialize()
	utils.GetVersionName()
	// rich_presence.Initialize()

	utils.Banner()
	Welcome()
	utils.Site_list()
	utils.Menu()
}
