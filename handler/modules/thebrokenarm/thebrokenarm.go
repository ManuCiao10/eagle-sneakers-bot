package thebrokenarm

import (
	"fmt"
	"time"

	"github.com/eagle/handler/utils"
)

func Initialize() {
	fmt.Print("\033[H\033[2J")
	utils.Banner()
	utils.Directory("thebrokenarm")
	csv_index := utils.SelectMode("[Eagle 0.0.2]" + "[" + time.Now().Format("15:04:05.000000") + "]" + " PLEASE SELECT CSV:")
	Find_index_of_csv(csv_index, "thebrokenarm")

}
