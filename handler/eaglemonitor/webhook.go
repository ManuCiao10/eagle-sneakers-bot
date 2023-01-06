package eaglemonitor

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"time"
)

var (
	cookieJar, _ = cookiejar.New(nil)
	img          = "https://media.discordapp.net/attachments/1013517214906859540/1039155134556536894/01IHswd8_400x400.jpeg"
)
var client = &http.Client{
	Jar: cookieJar,
}

// add hidden webhook field
func monitorWebhook(checkout *MonitorDetected, discordWebhook string) {
	Fields := []Fields{
		{
			Name:   "PID",
			Value:  checkout.pid,
			Inline: false,
		},
		{
			Name:   "Proxy List",
			Value:  checkout.proxy,
			Inline: true,
		},
		{
			Name: "Task File",
			//add hidden value
			Value:  checkout.taskFile,
			Inline: true,
		},
		{
			Name:   "Delay",
			Value:  fmt.Sprintf("%d", checkout.delay),
			Inline: true,
		},
		{
			Name:   "Store",
			Value:  checkout.store,
			Inline: true,
		},
		{
			Name:   "Size",
			Value:  checkout.size,
			Inline: true,
		},
		{
			Name:   "Task Quantity",
			Value:  fmt.Sprintf("%d", checkout.taskQuantity),
			Inline: true,
		},
	}

	payload := &Top{
		Username:  "EagleBot",
		AvatarURL: img,
		Embeds: []Embeds{
			{
				Title:  "**Monitor Detected!**",
				Color:  1999236,
				Fields: Fields,
				Thumbnail: Thumbnail{
					URL: img,
				},
				Footer: Footer{
					IconURL: img,
					Text:    "EagleBot | " + time.Now().Format("15:04:05"),
				},
			},
		},
	}
	payloadBuf := new(bytes.Buffer)
	_ = json.NewEncoder(payloadBuf).Encode(payload)

	SendWebhook, err := http.NewRequest("POST", discordWebhook, payloadBuf)
	if err != nil {
		fmt.Println(err)
	}
	SendWebhook.Header.Set("content-type", "application/json")

	sendWebhookRes, err := client.Do(SendWebhook)
	if err != nil {
		fmt.Print(err)
	}
	if sendWebhookRes.StatusCode != 204 {
		fmt.Printf("Webhook failed with status %d\n", sendWebhookRes.StatusCode)
	}
	defer sendWebhookRes.Body.Close()
}
