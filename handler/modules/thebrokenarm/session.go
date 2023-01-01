package thebrokenarm

import (
	"time"

	"github.com/eagle/handler/logs"
	"github.com/eagle/handler/task"
	"github.com/eagle/handler/utils"
)

/*
-- Modes --
if t.Mode == "login" {
 	return Login(t)
}
*/

func getSession(t *task.Task) task.TaskState {
	logs.LogWarn(t, "getting session")
	_, err := t.Client.NewRequest().
		SetURL("https://www.the-broken-arm.com/en/").
		SetMethod("GET").
		SetDefaultHeadersTBA().
		Do()

	if err != nil {
		// handle error and retry
		logs.LogErr(t, "failed to get session, retrying...")
		return GET_SESSION
	}

	return handleResponse(t)
}

func handleResponse(t *task.Task) task.TaskState {
	if t.Client.LatestResponse.StatusCode() != 200 {
		// handle error and retry
		time.Sleep(t.Delay)
		return GET_SESSION
	}
	orderId := utils.GetID(t.Client.LatestResponse.BodyAsString())
	if orderId == "" {
		// handle error and retry
		logs.LogErr(t, "failed to get orderId == -1")
		time.Sleep(t.Delay)
		return GET_SESSION
	}

	//get the ID from the response
	//get cf cookies --> https://www.the-broken-arm.com/cdn-cgi/challenge-platform/h/g/cv/result/ID

	return LOGIN
}
