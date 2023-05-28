# Discord Presence

Discord Rich Presence is a feature that allows you to display your activity inside your **Discord** profile when you are using software or playing a connected game.

### Inizialization

To implement Discord Rich Presence, we can use the `"github.com/hugolgst/rich-go/client"` library, which supports Linux, macOS, and Windows.

To get your DISCORD\_APP\_ID, you should create a new application on the [Discord portal](https://discord.com/developers/applications).

```go
var (
	startingTime    = time.Now()
	CurrentActivity = client.Activity{
		State:      "",
		Details:    "<Version> 0.0.23",
		LargeImage: "preview",
		LargeText:  "",
		SmallImage: "<image_name>",
		SmallText:  "",
		Timestamps: &client.Timestamps{
			Start: &startingTime,
		},
	}
)

func Initialize() {
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
```
