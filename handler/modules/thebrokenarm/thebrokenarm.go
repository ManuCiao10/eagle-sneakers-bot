package thebrokenarm

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/eagle/handler/client"
	"github.com/eagle/handler/utils"
)

func Loading() {
	fmt.Print("\033[H\033[2J")
	utils.Banner()
	utils.Directory("thebrokenarm")

	csv_index := utils.SelectMode(utils.Version() + utils.Time() + "PLEASE SELECT CSV:")
	task_name := CvsIndex(csv_index, "thebrokenarm")
	if task_name == "UNEXPECTED" {
		err_("UNEXPECTED ERROR")
	}

	CvsInfo(task_name, "thebrokenarm")

	for _, t := range tasks {
		Initialize(t)
	}

}

func Request(t *Task) TaskState {
	_, err := t.Client.NewRequest().
		SetURL("https://www.thebrokenarm.com/products/" + t.Pid + ".json").
		SetMethod("GET").
		Do()

	if err != nil {
		err_("REQUEST ERROR")
		return ErrorTaskState
	}

	// return Checkout(t)
	return ErrorTaskState

}

func Initialize(t *Task) TaskState {
	rand.Seed(time.Now().UnixNano())
	//handle more modes when they are added
	if !Contains([]string{"login", "normal"}, t.Mode) {
		err_("MODE IS NOT SUPPORTED FOR THIS SITE")
		return ErrorTaskState
	}

	taskProfile := GetProfile(t)
	if taskProfile.ID == "not_found" {
		err_("PROFILE NOT FOUND")
		return ErrorTaskState
	}

	p := GetProxyList(t)
	if p.ID == "not_found" {
		err_("PROXY LIST NOT FOUND")
		return ErrorTaskState
	}

	proxyURL := ProxyToUrl(p.ProxyList[rand.Intn(len(p.ProxyList))])

	t.CheckoutData.Proxy = proxyURL
	client, err := client.NewClient(proxyURL)

	if err != nil {
		err_("CLIENT ERROR")
		return ErrorTaskState
	}

	if t.Size == "random" {
		t.Size = RandomSize()
	} else {
		t.Size = SplitSize(t.Size)
	}

	t.CheckoutData.Website = "thebrokenarm"
	t.CheckoutData.Mode = t.Mode
	t.CheckoutData.ProductMSKU = t.Pid
	t.CheckoutData.Size = t.Size

	t.Client = client

	return Request(t)

}
