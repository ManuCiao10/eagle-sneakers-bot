package eaglemonitor

import (
	"fmt"

	"github.com/eagle/handler/auth"
	"github.com/eagle/handler/utils"
)

func Initialize() {
	fmt.Print("\033[H\033[2J")
	utils.Banner()
	auth.Welcome()

	monitorInitialize()
}
