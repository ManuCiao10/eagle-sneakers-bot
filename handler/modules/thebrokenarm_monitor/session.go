package thebrokenarm_monitor

import (
	"time"

	"github.com/eagle/handler/logs"
	"github.com/eagle/handler/quicktask"
	"github.com/eagle/handler/utils"
)

func session(t *quicktask.Quicktask) quicktask.TaskState {
	logs.LogQuickSess(t, "getting session...")
	_, err := t.Client.NewRequest().
		SetURL("https://www.the-broken-arm.com/en/").
		SetMethod("GET").
		SetDefaultHeadersTBA().
		Do()

	if err != nil {
		// handle error and retry
		logs.LogQuickErr(t, "failed to get session, retrying...")
		return SESSION
	}

	return handleResponse(t)
}

func handleResponse(t *quicktask.Quicktask) quicktask.TaskState {
	if t.Client.LatestResponse.StatusCode() != 200 {
		time.Sleep(t.Delay)
		return SESSION
	}
	Id := utils.GetId(t.Client.LatestResponse.BodyAsString())
	if Id == "" {
		logs.LogQuickErr(t, "failed to get Id == -1")
		time.Sleep(t.Delay)
		return SESSION
	}
	saveCookie(t)
	// fmt.Println(t.Client.LatestResponse.CookiesAsString())

	TBAInternalQuick.ProductID = Id

	return GUEST
}
