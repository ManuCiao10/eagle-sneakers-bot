package thebrokenarm_monitor

import (
	"net/url"
	"strings"
	"time"

	"github.com/eagle/handler/logs"
	"github.com/eagle/handler/quicktask"
)

func login(t *quicktask.Quicktask) quicktask.TaskState {
	requestBody := url.Values{}
	requestBody.Add("back", "")
	requestBody.Add("email", TBAInternalQuick.Account.Email)
	requestBody.Add("password", TBAInternalQuick.Account.Password)
	requestBody.Add("submitLogin", "1")

	_, err := t.Client.NewRequest().
		SetURL("https://www.the-broken-arm.com/en/connexion").
		SetMethod("POST").
		SetLoginHeadersTBA().
		SetFormBody(requestBody).
		Do()

	if err != nil {
		logs.LogQuickErr(t, "failed to login, retrying...")
		time.Sleep(t.Delay)
		return LOGIN
	}
	saveCookie(t)

	return handleLoginResponse(t)
}

func handleLoginResponse(t *quicktask.Quicktask) quicktask.TaskState {
	if strings.Contains(t.Client.LatestResponse.BodyAsString(), "Authentication failed.") {
		logs.LogQuickErr(t, "failed to login...")
		return quicktask.ErrorTaskState
	}

	logs.LogQuickSuccess(t, "logged in")

	return quicktask.ErrorTaskState
}

func saveCookie(t *quicktask.Quicktask) {
	cookiesjar := strings.Split(t.Client.LatestResponse.CookiesAsString(), ";")
	sessionID := cookiesjar[0]
	PrestaShop := cookiesjar[2]
	TBAInternalQuick.Cookies = sessionID + ";" + PrestaShop
}
