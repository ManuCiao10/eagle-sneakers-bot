package thebrokenarm

import (
	"encoding/json"
	"time"

	"github.com/eagle/handler/logs"
	"github.com/eagle/handler/task"
)

var success addToCartResponse

type addToCartResponse struct {
	Success bool `json:"success"`
}

// func getTokenCart(t *task.Task) string {

func addToCart(t *task.Task) task.TaskState {
	logs.LogPurple(t, "adding to cart...")

	// token = getTokenCart(t)

	data := "token=7f2711d779a862633d2f07dcadfe0f08&id_product=" + t.Pid + "&id_product=" + t.Pid + "&add=1&action=update"

	_, err := t.Client.NewRequest().
		SetURL("https://www.the-broken-arm.com/en/panier").
		SetMethod("POST").
		SetCartHeadersTBA().
		SetHeader("cookie", TBAInternal.Cookies).
		SetBody(data).
		Do()

	if err != nil {
		// handle error and retry
		return ADD_TO_CART
	}

	return handleAddToCart(t)
}

func handleAddToCart(t *task.Task) task.TaskState {

	if err := json.Unmarshal(t.Client.LatestResponse.Body(), &success); err != nil {
		logs.LogErr(t, "failed to add to cart, retrying...")
		time.Sleep(t.Delay)
		return ADD_TO_CART
	}

	if !success.Success {
		logs.LogErr(t, "product out of stock, retrying...")
		time.Sleep(t.Delay)
		return ADD_TO_CART
	}
	logs.LogCyan(t, "added to cart")
	// console.AddCart()

	return CHECKOUT
}
