package thebrokenarm

import (
	"bytes"

	"github.com/eagle/handler/logs"
	"github.com/eagle/handler/task"
)

func addToCart(t *task.Task) task.TaskState {
	logs.LogPurple(t, "checking stock...")
	// if len(t.Pid) > 10 {
	// 	t.Pid = splitPid(t.Pid)
	// }

	data := "token=7f2711d779a862633d2f07dcadfe0f08&id_product=" + t.Pid + "&id_customization=0&author=&friend_name=&friend_email=&id_product=" + t.Pid + " &add=1&action=update"

	_, err := t.Client.NewRequest().
		SetURL("https://www.the-broken-arm.com/en/panier").
		SetMethod("POST").
		SetDefaultHeadersTBA().
		SetBody(data).
		Do()

	if err != nil {
		// handle error and retry
		return ADD_TO_CART
	}

	return handleAddToCart(t)

}

func handleAddToCart(t *task.Task) task.TaskState {
	if bytes.Contains(t.Client.LatestResponse.Body(), []byte("cloudflare")) {
		logs.LogErr(t, "failed to add to cart, retrying...")
		return ADD_TO_CART
	}
	logs.LogSuccess(t, "added to cart")
	// console.AddCart()

	return CHECKOUT
}
