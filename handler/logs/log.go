package logs

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func LogsMsgErr(data ...interface{}) {
	TimeStamp := Time()
	stringData := strings.ToUpper(fmt.Sprint(data...))
	fmt.Println(colorRed + fmt.Sprintf("%s%s", TimeStamp, stringData) + colorReset)
	time.Sleep(1 * time.Second)
	os.Exit(0)
}

func LogsMsgInfo(data ...interface{}) {
	TimeStamp := Time()
	stringData := strings.ToUpper(fmt.Sprint(data...))
	fmt.Println(colorWhite + fmt.Sprintf("%s%s", TimeStamp, stringData) + colorReset)
}

func LogsMsgWarn(data ...interface{}) {
	TimeStamp := Time()
	stringData := strings.ToUpper(fmt.Sprint(data...))
	fmt.Println(colorYellow + fmt.Sprintf("%s%s", TimeStamp, stringData) + colorReset)
}

func LogsMsgSuccess(data ...interface{}) {
	TimeStamp := Time()
	stringData := strings.ToUpper(fmt.Sprint(data...))
	fmt.Println(colorGreen + fmt.Sprintf("%s%s", TimeStamp, stringData) + colorReset)
}

func LogsMsg(data ...interface{}) {
	TimeStamp := Time()
	stringData := strings.ToUpper(fmt.Sprint(data...))
	fmt.Println(colorPurple + fmt.Sprintf("%s%s", TimeStamp, stringData) + colorReset)
}

func LogsMsgCyan(data ...interface{}) {
	TimeStamp := Time()
	stringData := strings.ToUpper(fmt.Sprint(data...))
	fmt.Println(colorCyan + fmt.Sprintf("%s%s", TimeStamp, stringData) + colorReset)
}

func Time() string {
	return "[" + time.Now().Format("15:04:05.000000") + "] "
}
