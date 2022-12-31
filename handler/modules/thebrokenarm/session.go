package thebrokenarm

import (
	"time"

	"github.com/eagle/handler/logs"
	"github.com/eagle/handler/task"
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
		logs.LogWarn(t, "failed to get session, retrying...")
		return GET_SESSION
	}

	return handleResponse(t)
}

func handleResponse(t *task.Task) task.TaskState {
	if t.Client.LatestResponse.StatusCode() != 200 {
		// retry
		time.Sleep(t.Delay)
		return GET_SESSION
	}
	//get cookies and set them in the client

	return LOGIN //add LOGIN
}
