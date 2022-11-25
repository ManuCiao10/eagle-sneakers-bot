package rich_presence

import (
	// "log"
	// "fmt"
	"fmt"
	"time"

	"github.com/hugolgst/rich-go/client"
)

// var (
// 	startingTime    = time.Now()
// 	CurrentActivity = client.Activity{
// 		State:      "Browsing Menu",
// 		Details:    "Eating Pussy",
// 		LargeImage: "../dio.png",
// 		LargeText:  "Eagle",
// 		SmallImage: "../10.png",
// 		SmallText:  "Eagle",
// 		Party: &client.Party{
// 			ID:         "-1",
// 			Players:    1,
// 			MaxPlayers: 1,
// 		},
// 		Timestamps: &client.Timestamps{
// 			Start: &startingTime,
// 		},
// 		Buttons: []*client.Button{
// 			{
// 				Label: "Twitter",
// 				Url:   "https://twitter.com/eagle_aio",
// 			},
// 			{
// 				Label: "Website",
// 				Url:   "https://eagleaio.com/",
// 			},
// 		},
// 	}
// )

func Initialize() {
	err := client.Login("797931316783611935")
	if err != nil {
		panic(err)
	}

	now := time.Now()
	err = client.SetActivity(client.Activity{
		Details:    "I'm running on rich-go :)",
		State:      "Heyy!!!",
		LargeImage: "largeimageid",
		LargeText:  "This is the large image :D",
		SmallImage: "smallimageid",
		SmallText:  "And this is the small image",
		Party:      &client.Party{ID: "-1", Players: 15, MaxPlayers: 24},
		Timestamps: &client.Timestamps{Start: &now},
		Secrets:    &client.Secrets{},
		Buttons:    []*client.Button{{Label: "GitHub", Url: "https://github.com/hugolgst/rich-go"}},
	})

	if err != nil {
		panic(err)
	}

	// Discord will only show the presence if the app is running
	// Sleep for a few seconds to see the update
	fmt.Println("Sleeping...")
	time.Sleep(time.Second * 4)
}
