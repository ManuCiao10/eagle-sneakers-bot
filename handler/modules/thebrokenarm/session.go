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
		time.Sleep(t.Delay)
		return GET_SESSION
	}
	Id := utils.GetId(t.Client.LatestResponse.BodyAsString())
	if Id == "" {
		logs.LogErr(t, "failed to get Id == -1")
		time.Sleep(t.Delay)
		return GET_SESSION
	}
	t.Client.SaveCookies()

	//add t.Client.LatestResponse.Cookies() to the request
	// for _, v := range t.Client.LatestResponse.Cookies() {
	// 	fmt.Println(v)
	// 	t.Client.NewRequest().AddCookie(v)
	// }

	//save id to struct TBAInternal
	TBAInternal.ProductID = Id

	return GET_CLOUD
}
