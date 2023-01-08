package thebrokenarm_monitor

import (
	"time"

	"github.com/eagle/handler/logs"
	"github.com/eagle/handler/quicktask"
)

func guest(t *quicktask.Quicktask) quicktask.TaskState {
	logs.LogQuick(t, "checking stock...")
	data := "token=c21124404c5def43c52a677dc3c1b525&id_product=" + t.Pid + "&id_product=" + t.Pid + "&add=1&action=update"

	_, err := t.Client.NewRequest().
		SetURL("https://www.the-broken-arm.com/en/panier").
		SetMethod("POST").
		SetCartHeadersTBA().
		SetHeader("cookie", TBAInternalQuick.Cookies).
		SetBody(data).
		Do()

	if err != nil {
		// handle error and retry
		return GUEST
	}

	return handleAddToCart(t)
}

func handleAddToCart(t *quicktask.Quicktask) quicktask.TaskState {
	err := t.Client.LatestResponse.BodyAsJSON(&dataResponse)
	if err != nil {
		logs.LogQuickErr(t, "failed to add to cart, retrying...")
		time.Sleep(t.Delay)
		return GUEST
	}
	if !dataResponse.Success {
		logs.LogQuickErr(t, "error adding to cart, retrying...")
		time.Sleep(t.Delay)
		return GUEST
	}
	if dataResponse.Quantity == nil {
		logs.LogQuickErr(t, "product out of stock, retrying...")
		time.Sleep(t.Delay)
		return GUEST
	}

	logs.LogQuickCart(t, "added to cart")
	// console.AddCart()

	return CHECKOUT
}
