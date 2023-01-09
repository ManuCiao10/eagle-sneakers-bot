package thebrokenarm_monitor

import (
	"time"

	"github.com/eagle/handler/logs"
	"github.com/eagle/handler/quicktask"
)

func checkout(t *quicktask.Quicktask) quicktask.TaskState {
	logs.LogQuickSess(t, "checking out...")

	_, err := t.Client.NewRequest().
		SetURL("https://www.the-broken-arm.com/en/module/paypal/ScInit?credit_card=0&getToken=1&getToken=1&source_page=false").
		SetMethod("GET").
		CheckoutHeders().
		SetHeader("cookie", TBAInternalQuick.Cookies).
		Do()

	if err != nil {
		logs.LogQuickErr(t, "failed to checkout, retrying...")
		time.Sleep(t.Delay)
		return CHECKOUT
	}

	return handleResponseCheckout(t)
}

func handleResponseCheckout(t *quicktask.Quicktask) quicktask.TaskState {
	if t.Client.LatestResponse.StatusCode() != 200 {
		time.Sleep(t.Delay)
		return CHECKOUT
	}

	err := t.Client.LatestResponse.BodyAsJSON(&payPal)
	if err != nil {
		logs.LogQuickErr(t, "failed to checkout, retrying...")
		time.Sleep(t.Delay)
		return CHECKOUT
	}

	if !payPal.Success {
		logs.LogQuickErr(t, "failed to checkout, retrying...")
		time.Sleep(t.Delay)
		return CHECKOUT
	}

	t.CheckoutData.Link = "https://www.paypal.com/checkoutnow?token=" + payPal.Token
	logs.LogQuickSuccess(t, "checked out successfully")
	t.CheckoutData.Status = "paypalsuccess"

	size := dataResponse.Cart.Products[0].Attributes.Taille
	if size == "" {
		size = "N/A"
	}
	t.CheckoutData.Size = size
	t.CheckoutData.Image_url = dataResponse.Cart.Products[0].Images[0].Medium.URL
	t.CheckoutData.ProductMSKU = t.Pid
	t.CheckoutData.Price = dataResponse.Cart.Products[0].RegularPriceAmount
	t.CheckoutData.Mode = "paypal"
	t.CheckoutData.ProductName = "Dunk Low to add"

	// console.AddCheckout()
	return quicktask.DoneTaskState
}
