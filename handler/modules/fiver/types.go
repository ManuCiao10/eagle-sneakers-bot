package fiver

import (
	"github.com/eagle/handler/loading"
	"github.com/eagle/handler/settings"
	"github.com/eagle/handler/task"
)

var (
	INITIALIZE task.TaskState = "initialize"
	FILL_DATA  task.TaskState = "fill_data"
	FILL_USER  task.TaskState = "fill_user"
)

var FiverInternal = struct {
	Cookies []string
}{}

func GetProxyList(t *task.Task) settings.Proxie {
	for _, proxy := range loading.Data.Proxies.Proxies {
		if proxy.ID == t.CheckoutProxy {
			return proxy
		}
	}

	return settings.Proxie{
		ID: "not_found",
	}
}
