package modules

import (
	"github.com/eagle/handler/modules/thebrokenarm"
	"github.com/eagle/handler/utils"
)

func Initialize(site int) {

	if site == utils.ERROR {
		utils.ConsolePrint("INVALID OPTION!", "red")
	} else if site == utils.THEBROKENARM {
		thebrokenarm.Loading()
	}

}
