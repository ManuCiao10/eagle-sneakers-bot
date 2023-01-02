package thebrokenarm

import (
	// "strings"

	"strings"
	"time"

	"github.com/eagle/handler/logs"
	"github.com/eagle/handler/task"
)

func checkout(t *task.Task) task.TaskState {
	logs.LogPurple(t, "preparing to checkout")

	data := strings.NewReader(`payment_option=card&amount=2600&currency=eur&stripe_auto_save_card=false&card_form_payment=true&save_card_form=false&payment_request=false`)

	_, err := t.Client.NewRequest().
		SetURL("https://www.the-broken-arm.com/en/module/stripe_official/createIntent").
		SetMethod("POST").
		SetDefaultHeadersTBA().
		SetBodyReader(data).
		Do()

	if err != nil {
		// handle error and retry
		return CHECKOUT
	}

	return handleCheckout(t)
}

func handleCheckout(t *task.Task) task.TaskState {
	if t.Client.LatestResponse.StatusCode() != 200 {
		logs.LogErr(t, "failed to checkout, retrying...")
		time.Sleep(t.Delay)
		return CHECKOUT
	}

	logs.LogSuccess(t, "checkout")
	// console.AddCheckout()
	time.Sleep(t.Delay)

	return CHECKOUT
}
