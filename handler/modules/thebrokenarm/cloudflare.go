package thebrokenarm

import (
	"time"

	"github.com/eagle/handler/logs"
	"github.com/eagle/handler/task"
)

// probably getCloud not needed
func getCloud(t *task.Task) task.TaskState {

	link := "https://www.the-broken-arm.com/cdn-cgi/challenge-platform/h/g/cv/result/" + TBAInternal.ProductID

	_, err := t.Client.NewRequest().
		SetURL(link).
		SetMethod("POST").
		SetDefaultHeadersTBA().
		Do()

	if err != nil {
		logs.LogErr(t, "failed to get cloudflare token, retrying...")
		return GET_CLOUD
	}

	return handlecloudflare(t)
}

func handlecloudflare(t *task.Task) task.TaskState {
	if t.Client.LatestResponse.StatusCode() != 200 && t.Client.LatestResponse.BodyAsString() != "ok" {
		logs.LogWarn(t, "cloudflare token failed, retrying...")
		time.Sleep(t.Delay)
		return GET_CLOUD
	}

	return ADD_TO_CART
}
