package main

import (
	"github.com/eagle/handler/auth"
	"github.com/eagle/handler/loading"
	"github.com/eagle/handler/modules"
	"github.com/eagle/handler/utils"
)

//go:generate goversioninfo -skip-versioninfo=true -icon=handler/create/favicon.ico -manifest=handler/create/file.exe.manifest

func main() {
	// create.Initialize()
	loading.Initialize()
	// auth.Initialize()
	// version.Initialize()
	// console.Initialize()
	// presence.Initialize()

	utils.Banner()
	auth.Welcome()
	utils.Site()
	site := utils.Menu()
	modules.Initialize(site)

}
