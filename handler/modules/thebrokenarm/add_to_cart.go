package thebrokenarm

import (
	"fmt"

	"github.com/eagle/handler/logs"
	"github.com/eagle/handler/task"
)

func addToCart(t *task.Task) task.TaskState {
	// if len(t.Pid) > 10 {
	// 	t.Pid = splitPid(t.Pid)
	// }

	data := "token=7f2711d779a862633d2f07dcadfe0f08&id_product=" + t.Pid + "&id_customization=0&author=&friend_name=&friend_email=&id_product=" + t.Pid + " &add=1&action=update"

	_, err := t.Client.NewRequest().
		SetURL("https://www.the-broken-arm.com/en/panier").
		SetMethod("POST").
		SetDefaultHeadersTBA().
		SetBody(data).
		SetCookie(t.Client.LatestResponse.Cookies()).
		Do()

	if err != nil {
		// handle error and retry
		return ADD_TO_CART
	}
	fmt.Println(t.Client.LatestResponse.Cookies())
	return handleAddToCart(t)

}

func handleAddToCart(t *task.Task) task.TaskState {
	fmt.Println(t.Client.LatestResponse.StatusCode())
	if t.Client.LatestResponse.StatusCode() != 200 {
		logs.LogErr(t, "failed to add to cart, retrying...")
		return ADD_TO_CART
	}

	logs.LogSuccess(t, "added to cart")
	// console.AddCart()

	return CHECKOUT
}
