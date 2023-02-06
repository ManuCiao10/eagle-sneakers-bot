package fiver

import (
	"math/rand"

	"github.com/eagle/client"
	"github.com/eagle/handler/logs"
	"github.com/eagle/handler/task"
	"github.com/eagle/handler/utils"
)

func initialize(t *task.Task) task.TaskState {
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

	client, err := client.NewClient(t.CheckoutProxy)

	if err != nil {
		return task.ErrorTaskState
	}

	logs.LogPurple(t, "getting session...")

	req, err := client.NewRequest().
		SetURL("https://www.fiverr.com/").
		SetMethod("GET").
		SetHeadersFiver().
		Do()

	if err != nil {
		logs.LogWarn(t, "error getting session: %s", err)
		return INITIALIZE
	}

	if req.StatusCode() != 200 {
		logs.LogWarn(t, "error getting session: ", req.StatusCode())
		return INITIALIZE
	}
	logs.LogBlue(t, "got session..")

	t.Client = client
	return FILL_DATA

}
