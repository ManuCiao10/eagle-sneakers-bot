package thebrokenarm

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/eagle/handler/logs"
	"github.com/eagle/handler/task"
)

var success addToCartResponse

type addToCartResponse struct {
	Success bool `json:"success"`
}

func addToCart(t *task.Task) task.TaskState {
	logs.LogPurple(t, "checking stock...")

	//token = getTokenCart(t)
	//id_product = t.Pid
	//fix cookies

	data := strings.NewReader(`token=c21124404c5def43c52a677dc3c1b525&id_product=8802&id_customization=0&group%5B6%5D=1647&friend_name=&friend_email=&id_product=8802&add=1&action=update`)

	_, err := t.Client.NewRequest().
		SetURL("https://www.the-broken-arm.com/en/panier").
		SetMethod("POST").
		SetCartHeadersTBA().
		SetHeader("cookie", TBAInternal.Cookies).
		SetBodyReader(data).
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
