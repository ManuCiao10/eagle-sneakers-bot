package thebrokenarm

import (
	"fmt"
	"strings"
	"time"

	"github.com/eagle/handler/utils"
	"github.com/eagle/handler/version"
)

func Loading() {
	fmt.Print("\033[H\033[2J")
	utils.Banner()
	utils.Directory("thebrokenarm")

	csv_index := utils.SelectMode("[Eagle " + version.Version + "] " + "[" + time.Now().Format("15:04:05.000000") + "]" + " PLEASE SELECT CSV:")
	task_name := CvsIndex(csv_index, "thebrokenarm")
	if task_name == "UNEXPECTED" {
		err_("UNEXPECTED ERROR")
	}

	CvsInfo(task_name, "thebrokenarm")

	for _, t := range tasks {
		Initialize(t)
	}

}

func proxyToProxyUrl(proxy string) string {
	proxySplit := strings.Split(proxy, ":")

	if len(proxySplit) == 2 {
		return fmt.Sprintf("http://%s:%s", proxySplit[0], proxySplit[1])
	} else if len(proxySplit) == 4 {
		return fmt.Sprintf("http://%s:%s@%s:%s", proxySplit[2], proxySplit[3], proxySplit[0], proxySplit[1])
	}

	return fmt.Sprintf("http://%s", proxy)
}

func Initialize(t *Task) TaskState {
	if !Contains([]string{"login", "normal"}, t.Mode) {
		err_("MODE IS NOT SUPPORTED FOR THIS SITE")
		return ErrorTaskState
	}

	taskProfile := GetProfile(t)
	if taskProfile.ID == "not_found" {
		err_("PROFILE NOT FOUND")
		return ErrorTaskState
	}

	proxies := GetProxyList(t)
	if proxies.ID == "not_found" {
		err_("PROXY LIST NOT FOUND")
		return ErrorTaskState
	}
	fmt.Print(proxies.ID)
	fmt.Print(proxies.ProxyList)

	// client, err := client.NewClient()

	// if err != nil {
	// 	err_("CLIENT ERROR")
	// }

	return ContinueTaskState

}
