package thebrokenarm

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	hclient "github.com/eagle/handler/client"
	"github.com/eagle/handler/task"
	"github.com/eagle/handler/utils"
)

// func Loading() {
// 	fmt.Print("\033[H\033[2J")
// 	utils.Banner()
// 	utils.Directory("thebrokenarm")

// 	csv_index := utils.SelectMode(utils.Version() + utils.Time() + "PLEASE SELECT CSV:")
// 	task_name := CvsIndex(csv_index, "thebrokenarm")
// 	if task_name == "UNEXPECTED" {
// 		err_("UNEXPECTED ERROR")
// 	}

// 	CvsInfo(task_name, "thebrokenarm")

// 	for _, t := range tasks {
// 		Initialize(t)
// 	}

// }

func HandleSessionResponse(t *task.Task) task.TaskState {
	//handle response index body to take the ID for inizialize the challenge

	fmt.Print(t.Client.LatestResponse.StatusCode())
	if t.Client.LatestResponse.StatusCode() != 200 {
		// retry
		time.Sleep(t.Delay)
		return GET_SESSION
	}

	//parse the body to get the challenge ID
	if utils.Debug {
		fmt.Println(t.Delay)
		fmt.Println(t.Client.LatestResponse.Body())
	}
	return task.ContinueTaskState

}

func getSession(t *task.Task) task.TaskState {
	// if t.Mode == "login" {
	// 	return Login(t)
	// }

	//find the cookies for the session
	_, err := t.Client.NewRequest().
		SetURL("https://www.shoezgallery.com/en/").
		SetMethod("GET").
		SetDefaultHeadersTBA().
		Do()

	if err != nil {
		err_("REQUEST ERROR")
		return task.ErrorTaskState
	}
	log.Print("GETTING SESSION ...")
	return HandleSessionResponse(t)

}

func Initialize(t *task.Task) task.TaskState {
	rand.Seed(time.Now().UnixNano())
	//handle more modes when they are added
	if !Contains([]string{"login", "normal"}, t.Mode) {
		err_("MODE IS NOT SUPPORTED FOR THIS SITE")
		return task.ErrorTaskState
	}

	//check if all the variable in the task are set (size, mode, profile existing etc..)
	//---> TODO
	taskProfile := GetProfile(t)
	if taskProfile.ID == "not_found" {
		err_("PROFILE NOT FOUND")
		return task.ErrorTaskState
	}

	p := GetProxyList(t)
	if p.ID == "not_found" {
		err_("PROXY LIST NOT FOUND")
		return task.ErrorTaskState
	}

	proxyURL := ProxyToUrl(p.ProxyList[rand.Intn(len(p.ProxyList))])

	t.CheckoutData.Proxy = proxyURL

	//fix using the proxy
	client, err := hclient.NewClient(proxyURL)

	if err != nil {
		log.Println(err)
	}

	if err != nil {
		err_("CLIENT ERROR")
		return task.ErrorTaskState
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

	return getSession(t)

}
