package thebrokenarm_monitor

import (
	"net/url"
	"time"

	"github.com/eagle/handler/quicktask"
)

func login(t *quicktask.Quicktask) quicktask.TaskState {
	requestBody := url.Values{}
	requestBody.Add("login_email", TBAInternalQuick.Account.Email)
	requestBody.Add("login_password", TBAInternalQuick.Account.Password)
	requestBody.Add("back_url", "https://www.buzzsneakers.gr/oloklirosi-parangelias")
	requestBody.Add("ajax", "yes")
	requestBody.Add("task", "login")

	_, err := t.Client.NewRequest().
		SetURL("https://www.buzzsneakers.gr/eisodos").
		SetMethod("POST").
		// SetDefaultHeadersBuzz().
		SetFormBody(requestBody).
		Do()

	if err != nil {
		// logs.Log(c, "Error logging in.")
		time.Sleep(t.Delay)
		return LOGIN
	}
	return quicktask.ErrorTaskState

	// return handleLoginResponse(c, b)
}
