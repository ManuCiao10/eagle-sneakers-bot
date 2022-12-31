package logs

import (
	"time"
)

func Time() string {
	return "[" + time.Now().Format("15:04:05.000000") + "] "
}
