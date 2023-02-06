package fiver

import (
	"strings"
	"time"

	"github.com/eagle/handler/logs"
	"github.com/eagle/handler/task"
)

func fillUser(t *task.Task) task.TaskState {
	logs.LogBlue(t, "filling user...")

	var data = strings.NewReader("------WebKitFormBoundary5vIfyM2mVYAlDZO0\r\nContent-Disposition: form-data; name=\"user[username]\"\r\n\r\nemanuTTle2321s\r\n------WebKitFormBoundary5vIfyM2mVYAlDZO0\r\nContent-Disposition: form-data; name=\"user[password]\"\r\n\r\nCaccolafritta3!!\r\n------WebKitFormBoundary5vIfyM2mVYAlDZO0\r\nContent-Disposition: form-data; name=\"user[email]\"\r\n\r\nnuiqwnd2@gmail.com\r\n------WebKitFormBoundary5vIfyM2mVYAlDZO0\r\nContent-Disposition: form-data; name=\"funnel\"\r\n\r\nstandard\r\n------WebKitFormBoundary5vIfyM2mVYAlDZO0\r\nContent-Disposition: form-data; name=\"guest_checkout_token\"\r\n\r\nundefined\r\n------WebKitFormBoundary5vIfyM2mVYAlDZO0\r\nContent-Disposition: form-data; name=\"redirect_url\"\r\n\r\n\r\n------WebKitFormBoundary5vIfyM2mVYAlDZO0--\r\n")

	req, err := t.Client.NewRequest().
		SetURL("https://www.fiverr.com/users").
		SetMethod("POST").
		SetHeadersFiverUser().
		SetBodyReader(data).
		Do()

	if err != nil {
		logs.LogErr(t, "failed to fill user, retrying...")
		return FILL_USER
	}

	if req.StatusCode() != 200 {
		logs.LogErr(t, "failed to fill user, retrying...", req.StatusCode())
		return FILL_USER
	}

	logs.LogCyan(t, "filled user...")

	time.Sleep(3 * time.Second)
	return INITIALIZE

}
