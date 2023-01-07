package logs

import (
	"fmt"
	"strings"

	"github.com/eagle/handler/quicktask"
	"github.com/eagle/handler/task"
)

func LogErr(t *task.Task, data ...interface{}) {
	TimeStamp := Time()
	siteName := strings.ToUpper(t.Type)
	taskMode := strings.ToUpper(t.Mode)
	taskSize := strings.ToUpper(t.Size)
	taskPid := strings.ToUpper(t.Pid)
	// add split url
	if len(taskPid) > 10 {
		taskPid = "URL"
	}
	stringData := strings.ToUpper(fmt.Sprint(data...))
	// authKey := loading.Data.Settings.Settings.AuthKey

	//log everything is doing a user (with key)
	// go LogLogTail(siteName, taskType, taskMode, taskSize, stringData, authKey)
	fmt.Println(colorRed + fmt.Sprintf("[%s %s] [%s] %s[%s] %s", siteName, taskMode, taskSize, TimeStamp, taskPid, stringData) + colorReset)
}

func LogInfo(t *task.Task, data ...interface{}) {
	TimeStamp := Time()
	siteName := strings.ToUpper(t.Type)
	taskMode := strings.ToUpper(t.Mode)
	taskSize := strings.ToUpper(t.Size)
	taskPid := strings.ToUpper(t.Pid)
	if len(taskPid) > 10 {
		taskPid = "URL"
	}
	stringData := strings.ToUpper(fmt.Sprint(data...))

	// authKey := loading.Data.Settings.Settings.AuthKey

	//log everything is doing a user (with key)
	// go LogLogTail(siteName, taskType, taskMode, taskSize, stringData, authKey)
	fmt.Println(colorWhite + fmt.Sprintf("[%s %s] [%s] %s[%s] %s", siteName, taskMode, taskSize, TimeStamp, taskPid, stringData) + colorReset)
}

func LogWarn(t *task.Task, data ...interface{}) {
	TimeStamp := Time()
	siteName := strings.ToUpper(t.Type)
	taskMode := strings.ToUpper(t.Mode)
	taskSize := strings.ToUpper(t.Size)
	taskPid := strings.ToUpper(t.Pid)
	if len(taskPid) > 10 {
		taskPid = "URL"
	}
	stringData := strings.ToUpper(fmt.Sprint(data...))

	// authKey := loading.Data.Settings.Settings.AuthKey

	//log everything is doing a user (with key)
	// go LogLogTail(siteName, taskType, taskMode, taskSize, stringData, authKey)
	fmt.Println(colorYellow + fmt.Sprintf("[%s %s] [%s] %s[%s] %s", siteName, taskMode, taskSize, TimeStamp, taskPid, stringData) + colorReset)
}

func LogSuccess(t *task.Task, data ...interface{}) {
	TimeStamp := Time()
	siteName := strings.ToUpper(t.Type)
	taskMode := strings.ToUpper(t.Mode)
	taskSize := strings.ToUpper(t.Size)
	taskPid := strings.ToUpper(t.Pid)
	if len(taskPid) > 10 {
		taskPid = "URL"
	}
	stringData := strings.ToUpper(fmt.Sprint(data...))

	// authKey := loading.Data.Settings.Settings.AuthKey

	//log everything is doing a user (with key)
	// go LogLogTail(siteName, taskType, taskMode, taskSize, stringData, authKey)
	fmt.Println(colorGreen + fmt.Sprintf("[%s %s] [%s] %s[%s] %s", siteName, taskMode, taskSize, TimeStamp, taskPid, stringData) + colorReset)
}

// LogLogTail is a function to log all the logs in a file

func LogCyan(t *task.Task, data ...interface{}) {
	TimeStamp := Time()
	siteName := strings.ToUpper(t.Type)
	taskMode := strings.ToUpper(t.Mode)
	taskSize := strings.ToUpper(t.Size)
	taskPid := strings.ToUpper(t.Pid)
	if len(taskPid) > 10 {
		taskPid = "URL"
	}
	stringData := strings.ToUpper(fmt.Sprint(data...))

	// authKey := loading.Data.Settings.Settings.AuthKey

	//log everything is doing a user (with key)
	// go LogLogTail(siteName, taskType, taskMode, taskSize, stringData, authKey)
	fmt.Println(colorCyan + fmt.Sprintf("[%s %s] [%s] %s[%s] %s", siteName, taskMode, taskSize, TimeStamp, taskPid, stringData) + colorReset)
}

func LogBlue(t *task.Task, data ...interface{}) {
	TimeStamp := Time()
	siteName := strings.ToUpper(t.Type)
	taskMode := strings.ToUpper(t.Mode)
	taskSize := strings.ToUpper(t.Size)
	taskPid := strings.ToUpper(t.Pid)
	if len(taskPid) > 10 {
		taskPid = "URL"
	}
	stringData := strings.ToUpper(fmt.Sprint(data...))

	// authKey := loading.Data.Settings.Settings.AuthKey

	//log everything is doing a user (with key)
	// go LogLogTail(siteName, taskType, taskMode, taskSize, stringData, authKey)
	fmt.Println(colorBlue + fmt.Sprintf("[%s %s] [%s] %s[%s] %s", siteName, taskMode, taskSize, TimeStamp, taskPid, stringData) + colorReset)
}

func LogPurple(t *task.Task, data ...interface{}) {
	TimeStamp := Time()
	siteName := strings.ToUpper(t.Type)
	taskMode := strings.ToUpper(t.Mode)
	taskSize := strings.ToUpper(t.Size)
	taskPid := strings.ToUpper(t.Pid)
	if len(taskPid) > 10 {
		taskPid = "URL"
	}
	stringData := strings.ToUpper(fmt.Sprint(data...))

	// authKey := loading.Data.Settings.Settings.AuthKey

	//log everything is doing a user (with key)
	// go LogLogTail(siteName, taskType, taskMode, taskSize, stringData, authKey)
	fmt.Println(colorPurple + fmt.Sprintf("[%s %s] [%s] %s[%s] %s", siteName, taskMode, taskSize, TimeStamp, taskPid, stringData) + colorReset)
}

// log for quickTask
func LogQuick(t *quicktask.Quicktask, data ...interface{}) {
	TimeStamp := Time()
	siteName := strings.ToUpper(t.Type)
	taskMode := strings.ToUpper(t.Mode)
	taskSize := strings.ToUpper(t.Size)
	taskPid := strings.ToUpper(t.Pid)
	if len(taskPid) > 10 {
		taskPid = "URL"
	}
	stringData := strings.ToUpper(fmt.Sprint(data...))

	// authKey := loading.Data.Settings.Settings.AuthKey

	//log everything is doing a user (with key)
	// go LogLogTail(siteName, taskType, taskMode, taskSize, stringData, authKey)
	fmt.Println(colorWhite + fmt.Sprintf("[%s %s] [%s] %s[%s] %s", siteName, taskMode, taskSize, TimeStamp, taskPid, stringData) + colorReset)
}
