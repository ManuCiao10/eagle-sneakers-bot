package thebrokenarm

// func HandleSessionResponse(t *task.Task) task.TaskState {
// 	//handle response index body to take the ID for inizialize the challenge

// 	fmt.Print(t.Client.LatestResponse.StatusCode())
// 	if t.Client.LatestResponse.StatusCode() != 200 {
// 		// retry
// 		time.Sleep(t.Delay)
// 		return GET_SESSION
// 	}

// 	//parse the body to get the challenge ID
// 	if utils.Debug {
// 		fmt.Println(t.Delay)
// 		fmt.Println(t.Client.LatestResponse.Body())
// 	}
// 	return task.ContinueTaskState

// }

// func getSession(t *task.Task) task.TaskState {
// 	// if t.Mode == "login" {
// 	// 	return Login(t)
// 	// }

// 	//find the cookies for the session
// 	_, err := t.Client.NewRequest().
// 		SetURL("https://www.shoezgallery.com/en/").
// 		SetMethod("GET").
// 		SetDefaultHeadersTBA().
// 		Do()

// 	if err != nil {
// 		err_("REQUEST ERROR")
// 		return task.ErrorTaskState
// 	}
// 	log.Print("GETTING SESSION ...")
// 	return HandleSessionResponse(t)

// }
