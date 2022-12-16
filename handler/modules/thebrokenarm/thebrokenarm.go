package thebrokenarm

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/eagle/handler/settings"
	"github.com/eagle/handler/utils"
	"github.com/eagle/handler/version"
	"github.com/fatih/color"

	"github.com/eagle/handler/loading"
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

func GetProxyList(t *Task) settings.Proxie {

	for _, proxy := range loading.Data.Proxies.Proxies {
		if proxy.ID == t.Proxy_List {
			return proxy
		}
	}

	return settings.Proxie{}

}

func Initialize(t *Task) TaskState {
	if !Contains([]string{"login", "normal"}, t.Mode) {
		err_("MODE IS NOT SUPPORTED FOR THIS SITE")
		return ErrorTaskState
	}

	taskProfile := GetProfile(t)
	if taskProfile.ID == "" {
		err_("PROFILE NOT FOUND")
		return ErrorTaskState
	}

	proxies := GetProxyList(t)
	if proxies.ID == "" {
		err_("PROXY LIST NOT FOUND" + strings.ToUpper(proxies.ID))
		return ErrorTaskState
	}
	fmt.Print(proxies.ID)
	fmt.Print(proxies.ProxyList)
	// proxy := loading.Data.Proxies.Proxies[1]

	// fmt.Println(proxy)
	// proxy := utils.GetProxy(t.Proxy)

	// client, err := client.NewClient()

	// if err != nil {
	// 	err_("CLIENT ERROR")
	// }

	return ContinueTaskState

}

func err_(msg string) {
	color.Red(msg)
	os.Exit(0)
}
