package thebrokenarm

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/eagle/client"
	"github.com/eagle/handler/loading"
	"github.com/eagle/handler/logs"
	"github.com/eagle/handler/task"
	"github.com/eagle/handler/utils"
)

func initialize(t *task.Task) task.TaskState {
	//TODO --> check if all the variable in the task are set (size, mode, profile existing etc..)
	rand.Seed(time.Now().UnixNano())

	if !utils.Contains([]string{"login", "normal"}, t.Mode) {
		logs.LogErr(t, "mode is not supported for this site.")
		return task.ErrorTaskState
	}

	t.CheckoutProfile = GetProfile(t)
	if t.CheckoutProfile.ID == "not_found" {
		logs.LogErr(t, "profile not found")
		return task.ErrorTaskState
	}

	p := GetProxyList(t)
	if p.ID == "not_found" {
		logs.LogErr(t, "proxy list not found")
		return task.ErrorTaskState
	}

	if len(p.ProxyList) == 0 {
		logs.LogErr(t, "proxy list is empty...")
		return task.ErrorTaskState
	}

	t.CheckoutProxy = utils.ProxyToUrl(p.ProxyList[rand.Intn(len(p.ProxyList))])

	delay, err := strconv.Atoi(loading.Data.Settings.Settings.Delay.Retry)
	if err != nil {
		logs.LogsMsgErr("Check the delay in the settings.json file.")
	}
	t.Delay = time.Duration(delay) * time.Millisecond

	client, err := client.NewClient() //t.CheckoutProxy

	if err != nil {
		return task.ErrorTaskState
	}

	t.CheckoutData.Proxy = t.CheckoutProxy
	t.CheckoutData.Website = t.Type
	t.CheckoutData.Mode = t.Mode
	t.CheckoutData.ProductMSKU = t.Pid
	t.CheckoutData.Size = t.Size
	t.CheckoutData.Profile = t.CheckoutProfile.ID

	t.Client = client

	return GET_SESSION
}
