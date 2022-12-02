package rich_presence

import (
	"fmt"
	"time"

	"github.com/eagle/handler/loading"
	"github.com/hugolgst/rich-go/client"
)

var (
	startingTime    = time.Now()
	CurrentActivity = client.Activity{
		State:      "",
		Details:    "Version 0.0.23",
		LargeImage: "preview",
		LargeText:  "",
		SmallImage: "ghidra",
		SmallText:  "",
		Timestamps: &client.Timestamps{
			Start: &startingTime,
		},
	}
)

func Initialize() {
	// user_ = auth.Auth.Integrations.Discord.Username

	ID := loading.Data.Env.Env.DISCORD_APP_ID
	err := client.Login(ID)
	if err != nil {
		fmt.Println("Failed to start discord rich presence.")
		return
	}

	err = client.SetActivity(CurrentActivity)
	if err != nil {
		fmt.Println("Failed to set Eagle discord rich presence.")
		return
	}
}
