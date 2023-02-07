package main

import (
	"github.com/eagle/handler/auth"
	"github.com/eagle/handler/cmd"
	"github.com/eagle/handler/create"
	"github.com/eagle/handler/loading"
	"github.com/eagle/handler/logs"

	"github.com/eagle/handler/modules/fiver"
	"github.com/eagle/handler/modules/nike"
	"github.com/eagle/handler/modules/thebrokenarm"
	"github.com/eagle/handler/modules/thebrokenarm_monitor"

	"github.com/eagle/handler/presence"
	"github.com/eagle/handler/quicktasking"
	"github.com/eagle/handler/utils"
)

//go:generate goversioninfo -skip-versioninfo=true -icon=handler/create/favicon.ico -manifest=handler/create/file.exe.manifest

func main() {
	create.Initialize()
	loading.Initialize()
	// console.Initialize() //only windows
	quicktasking.Initialize()
	nike.Initialize()
	thebrokenarm.Initialize()
	thebrokenarm_monitor.Initialize()
	fiver.Initialize()
	logs.LogtailInitialize()
	// nike.Initialize()

	// ws_quicktasking.Initialize() //to be implemented
	// auth.Initialize() //to be fixed
	// version.Initialize() //to be fixed

	presence.Initialize()
	utils.Banner()
	auth.Welcome()
	utils.Site()

	cmd.Initialize()
}
