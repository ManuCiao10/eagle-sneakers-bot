package fiver

import (
	"math/rand"
	"strings"
	"time"

	"github.com/eagle/handler/logs"
	"github.com/eagle/handler/task"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func fillData(t *task.Task) task.TaskState {
	logs.LogPurple(t, "filling data...")

	email := RandStringRunes(6) + "@gmail.com"

	var data = strings.NewReader("------WebKitFormBoundaryQvBXcpmMiTG40it8\r\nContent-Disposition: form-data; name=\"email\"\r\n\r\n%s\r\n------WebKitFormBoundaryQvBXcpmMiTG40it8\r\nContent-Disposition: form-data; name=\"funnel\"\r\n\r\nstandard\r\n------WebKitFormBoundaryQvBXcpmMiTG40it8\r\nContent-Disposition: form-data; name=\"guest_checkout_token\"\r\n\r\nundefined\r\n------WebKitFormBoundaryQvBXcpmMiTG40it8\r\nContent-Disposition: form-data; name=\"redirect_url\"\r\n\r\nhttps://it.fiverr.com/manuedesign/do-everything-in-golang?utm_source=654037&utm_medium=cx_affiliate&utm_campaign=&afp=&cxd_token=654037_23371827&show_join=true\r\n------WebKitFormBoundaryQvBXcpmMiTG40it8--\r\n" + email)

	req, err := t.Client.NewRequest().
		SetURL("https://www.fiverr.com/validate_email").
		SetMethod("POST").
		SetHeadersFiverLogin().
		SetBodyReader(data).
		Do()

	if err != nil {
		logs.LogErr(t, "failed to fill data, retrying...")
		return FILL_DATA
	}

	if req.StatusCode() != 201 {
		logs.LogErr(t, "failed to fill data, retrying...")
		return FILL_DATA
	}

	logs.LogBlue(t, "filled email...")

	time.Sleep(3 * time.Second)
	return FILL_USER
}
