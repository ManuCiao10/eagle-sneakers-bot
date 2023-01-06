package main

import (
	"github.com/eagle/handler/auth"
	"github.com/eagle/handler/cmd"
	"github.com/eagle/handler/loading"
	"github.com/eagle/handler/logs"
	"github.com/eagle/handler/modules/thebrokenarm"
	"github.com/eagle/handler/utils"
)

//go:generate goversioninfo -skip-versioninfo=true -icon=handler/create/favicon.ico -manifest=handler/create/file.exe.manifest

func main() {
	thebrokenarm.Initialize()
	loading.Initialize()
	logs.LogtailInitialize()
	// create.Initialize()
	// auth.Initialize()
	// version.Initialize()
	// console.Initialize() Only Windows
	// presence.Initialize()

	utils.Banner()
	auth.Welcome()
	utils.Site()

	cmd.Initialize()

}
