package thebrokenarm_monitor

import (
	"fmt"
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
	//get a random account from the list only if the Used state is false
	if t.Accounts != "" {
		TBAInternalQuick.Account = GetRandomAccount(siteIdMap[t.Site])
		if TBAInternalQuick.Account.Email == "" {
			logs.LogQuick(t, "no account available")
			return quicktask.ErrorTaskState
		}
	} else {
		logs.LogQuick(t, "no account specified")
		return quicktask.ErrorTaskState
	}
	fmt.Println("Account: ", TBAInternalQuick.Account.Email, TBAInternalQuick.Account.Password)

	t.CheckoutProfile = GetProfile(t)
	if t.CheckoutProfile.ID == "not_found" {
		logs.LogQuick(t, "profile not found")
		return quicktask.ErrorTaskState
	}

	p := GetProxyList(t)
	if p.ID == "not_found" {
		logs.LogQuick(t, "proxy list not found")
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
	return LOGIN
}
