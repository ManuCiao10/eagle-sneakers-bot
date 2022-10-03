package utils

import (
	"os"
	"time"

	"github.com/fatih/color"
)

func Profile_error() {
	color.Red("[ Eagle 0.0.2 ]" + "[ " + time.Now().Format("15:04:05.000000") + " ] " + "FIELD WRONG IN FILE.CSV")
	os.Exit(1)
}
