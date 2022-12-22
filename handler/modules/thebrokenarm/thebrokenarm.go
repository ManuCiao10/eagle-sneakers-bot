package thebrokenarm

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/eagle/handler/client"
	"github.com/eagle/handler/task"
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

func HandleSessionResponse(t *task.Task) task.TaskState {
	//handle response index body to take the ID for inizialize the challenge

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
		SetURL("https://www.the-broken-arm.com/en/").
		SetMethod("GET").
		// SetDefaultHeadersTBA().
		// Do()

	if err != nil {
		err_("REQUEST ERROR")
		return task.ErrorTaskState
	}

	fmt.Print("GETTING SESSION ...")

	// return Checkout(t)
	return HandleSessionResponse(t)

}

func Initialize(t *task.Task) task.TaskState {
	rand.Seed(time.Now().UnixNano())
	//handle more modes when they are added
	if !Contains([]string{"login", "normal"}, t.Mode) {
		err_("MODE IS NOT SUPPORTED FOR THIS SITE")
		return task.ErrorTaskState
	}

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

	jar := client.NewCookieJar(nil)
	options := []client.HttpClientOption{
		client.WithTimeout(30),
		// client.WithClientProfile(client.Chrome_105),
		client.WithNotFollowRedirects(),
		client.WithCookieJar(jar), // create cookieJar instance and pass it as argument
		client.WithProxyUrl(proxyURL),
		//client.WithInsecureSkipVerify(),
	}

	client, err := client.NewHttpClient(client.NewNoopLogger(), options...)
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
