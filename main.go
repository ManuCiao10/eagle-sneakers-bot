package main

import (
	"github.com/eagle/handler/auth"
	"github.com/eagle/handler/cmd"
	"github.com/eagle/handler/cmd/console"
	"github.com/eagle/handler/create"
	"github.com/eagle/handler/loading"
	"github.com/eagle/handler/logs"
	"github.com/eagle/handler/modules/thebrokenarm"
	"github.com/eagle/handler/modules/thebrokenarm_monitor"
	"github.com/eagle/handler/version"
	"github.com/eagle/handler/ws_quicktasking"

	"github.com/eagle/handler/presence"
	"github.com/eagle/handler/quicktasking"
	"github.com/eagle/handler/utils"
)

//go:generate goversioninfo -skip-versioninfo=true -icon=handler/create/favicon.ico

func main() {
	create.Initialize()
	loading.Initialize()
	console.Initialize() //only windows
	quicktasking.Initialize()
	thebrokenarm.Initialize()
	thebrokenarm_monitor.Initialize()
	logs.LogtailInitialize()

	ws_quicktasking.Initialize()
	auth.Initialize()
	version.Initialize()

	presence.Initialize()
	utils.Banner()
	auth.Welcome()
	utils.Site()

	cmd.Initialize()
}
