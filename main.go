package main

import (
	"strings"

	"github.com/eagle/handler/auth"
	"github.com/eagle/handler/console"
	"github.com/eagle/handler/create"
	"github.com/eagle/handler/loading"
	"github.com/eagle/handler/rich_presence"
	"github.com/eagle/handler/utils"
	"github.com/eagle/handler/version"
	"github.com/fatih/color"
)

//go:generate goversioninfo -skip-versioninfo=true -icon=handler/create/favicon.ico -manifest=handler/create/file.exe.manifest

func Welcome() {
	username := strings.ToUpper(auth.Auth.Integrations.Discord.Username)
	color.Magenta("WELCOME BACK  \t" + color.WhiteString(username))
	println("\n")
}

func main() {
	create.Initialize()
	loading.Initialize()
	auth.Initialize()
	version.Updates()
	console.Initialize()
	rich_presence.Initialize()

	utils.Banner()
	Welcome()
	utils.Site_list()
	utils.Menu()
}
