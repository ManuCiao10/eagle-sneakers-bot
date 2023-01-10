package logs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"
	"time"

	"github.com/eagle/client"
)

func logCheckoutBackend(checkout *CheckoutLogRequest) {
	checkoutClient, _ := client.NewClient()
	// checkout.AllowPublic = loading.Data.Settings.Settings.AllowPublicWebhook

	_, err := checkoutClient.NewRequest().
		SetURL("https://api.eagle.com/api/checkout").
		SetMethod("POST").
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "*/*").
		// SetHeader("Authorization", "Bearer "+auth.AuthToken).
		// SetJSONBody(checkout).
		Do()

	if err != nil {
		log.Println(err.Error())
	}
}

func LogCheckoutDiscord(checkout *CheckoutLogRequest, discordWebhook string) {
	var title string
	var color int

	if checkout.Status == "success" {
		title = "**Successful Checkout!**"
		color = 2524623
	} else if checkout.Status == "denied" {
		title = "**Checkout Failed!**"
		color = 16711680
	} else {
		return // invalid status
	}

	checkoutClient, _ := client.NewClient()

	requestData := fmt.Sprintf(`{"content":null,"embeds":[{"title":"%s","description":"%s","color":%d,"fields":[{"name":"MSKU","value":"%s","inline":true},{"name":"Mode","value":"%s","inline":true},{"name":"Size","value":"[%s](https://quicktask.hellasaio.com/quicktask?product_id=%s&siteId=%s&size=%s)","inline":true},{"name":"Checkout Time","value":"%sms","inline":true},{"name":"Price","value":"€%.2f","inline":true},{"name":"Store","value":"%s","inline":true},{"name":"Quicktask Link","value":"[Link](https://quicktask.hellasaio.com/quicktask?product_id=%s&siteId=%s&size=random)","inline":true}],"thumbnail":{"url":"%s"}}],"attachments":[]}`,
		title, checkout.ProductName, color, checkout.ProductMSKU, checkout.Mode, checkout.Size, checkout.ProductMSKU, checkout.TaskEnd, url.QueryEscape(checkout.Size), strconv.Itoa(checkout.CheckoutMs), checkout.Price, checkout.Website, checkout.ProductMSKU, checkout.Status, checkout.ImageUrl)

	_, err := checkoutClient.NewRequest().
		SetURL(discordWebhook).
		SetMethod("POST").
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "*/*").
		SetBody(requestData).
		Do()

	if err != nil {
		log.Fatalln(err.Error())
	}
}

func LogCheckout(checkout *CheckoutLogRequest, discordWebhook string) {
	if checkout.Status == "paypalsuccess" {
		go logPaypalDiscord(checkout, discordWebhook)
	}

	// go logCheckoutBackend(checkout)
	// go LogCheckoutDiscord(checkout, discordWebhook)
}

func LogTimeout(discordWebhook string) {
	checkoutClient, _ := client.NewClient()

	requestData := fmt.Sprintf(`{"content":null,"embeds":[{"title":"**Checkout Timeout!**","description":"Checkout timed out...","color":16711680,"fields":[{"name":"Store","value":"%s","inline":true}],"thumbnail":{"url":"https://media.discordapp.net/attachments/1013517214906859540/1039155134556536894/01IHswd8_400x400.jpeg"}}],"attachments":[]}`, "N/A")

	_, err := checkoutClient.NewRequest().
		SetURL(discordWebhook).
		SetMethod("POST").
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "*/*").
		SetBody(requestData).
		Do()

	if err != nil {
		log.Fatalln(err.Error())
	}
}

var (
	cookieJar, _ = cookiejar.New(nil)
)
var clientChechout = &http.Client{
	Jar: cookieJar,
}

// add hidden webhook field
func MonitorWebhook(checkout *MonitorDetected, discordWebhook string) {
	Fields := []Fields{
		{
			Name:   "PID",
			Value:  checkout.Pid,
			Inline: false,
		},
		{
			Name:   "Proxy List",
			Value:  checkout.Proxy,
			Inline: true,
		},
		{
			Name:   "Task File",
			Value:  checkout.TaskFile,
			Inline: true,
		},
		{
			Name:   "Delay",
			Value:  fmt.Sprintf("%d", checkout.Delay),
			Inline: true,
		},
		{
			Name:   "Store",
			Value:  checkout.Store,
			Inline: true,
		},
		{
			Name:   "Size",
			Value:  checkout.Size,
			Inline: true,
		},
		{
			Name:   "Task Quantity",
			Value:  fmt.Sprintf("%d", checkout.TaskQuantity),
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

	sendWebhookRes, err := clientChechout.Do(SendWebhook)
	if err != nil {
		fmt.Print(err)
	}
	if sendWebhookRes.StatusCode != 204 {
		fmt.Printf("Webhook failed with status %d\n", sendWebhookRes.StatusCode)
	}
	defer sendWebhookRes.Body.Close()
}

func logPaypalDiscord(checkout *CheckoutLogRequest, discordWebhook string) {
	Fields := []Fields{
		{
			Name:   "Pid",
			Value:  checkout.ProductMSKU,
			Inline: true,
		},
		{
			Name:   "Mode",
			Value:  checkout.Mode,
			Inline: true,
		},
		{
			Name:   "Size",
			Value:  checkout.Size,
			Inline: true,
		},
		{
			Name:   "Checkout Time",
			Value:  fmt.Sprintf("%dms", checkout.CheckoutMs),
			Inline: true,
		},
		{
			Name:   "Price",
			Value:  checkout.Price,
			Inline: true,
		},
		{
			Name:   "Website",
			Value:  checkout.Website,
			Inline: true,
		},

		{
			Name:   "Link",
			Value:  "[PayPal Link](" + checkout.PayPal + ")",
			Inline: true,
		},
	}
	payload := &Top{
		Username:  "EagleBot",
		AvatarURL: img,
		Embeds: []Embeds{
			{
				Title:       "**Paypal Success!**",
				Color:       1999236,
				Fields:      Fields,
				Description: checkout.ProductName,
				Thumbnail: Thumbnail{
					URL: checkout.ImageUrl,
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

	sendWebhookRes, err := clientChechout.Do(SendWebhook)
	if err != nil {
		fmt.Print(err)
	}
	if sendWebhookRes.StatusCode != 204 {
		fmt.Printf("Webhook failed with status %d\n", sendWebhookRes.StatusCode)
	}
	defer sendWebhookRes.Body.Close()

}
