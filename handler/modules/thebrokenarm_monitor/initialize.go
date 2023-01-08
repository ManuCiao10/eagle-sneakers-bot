package thebrokenarm_monitor

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/eagle/client"
	"github.com/eagle/handler/loading"
	"github.com/eagle/handler/logs"
	"github.com/eagle/handler/quicktask"
	"github.com/eagle/handler/utils"
)

func initialize(t *quicktask.Quicktask) quicktask.TaskState {
	if !utils.Contains([]string{"guest", "accouts"}, t.Accounts) {
		logs.LogQuickErr(t, "mode not supported for this site.")
		return quicktask.ErrorTaskState
	}

	t.CheckoutProfile = GetProfile(t)
	if t.CheckoutProfile.ID == "not_found" {
		logs.LogQuickErr(t, "profile not found")
		return quicktask.ErrorTaskState
	}

	p := GetProxyList(t)
	if p.ID == "not_found" {
		logs.LogQuickErr(t, "proxy list not found")
		return quicktask.ErrorTaskState
	}

	t.CheckoutProxy = utils.ProxyToUrl(p.ProxyList[rand.Intn(len(p.ProxyList))])

	delay, err := strconv.Atoi(loading.Data.Settings.Settings.Delay.Retry)
	if err != nil {
		logs.LogsMsgErr("Check the delay in the settings.json file.")
	}
	t.Delay = time.Duration(delay) * time.Millisecond

	client, err := client.NewClient() //t.CheckoutProxy

	if err != nil {
		return quicktask.ErrorTaskState
	}
	t.Client = client

	if t.Accounts == "accounts" {
		TBAInternalQuick.Account = GetRandomAccount(siteIdMap[t.Site])
		if TBAInternalQuick.Account.Email == "" {
			logs.LogQuickErr(t, "no account available")
			return quicktask.ErrorTaskState
		}
		return LOGIN
	}

	return SESSION
}
