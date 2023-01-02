package logs

import (
	"fmt"
	"log"
	"net/url"
	"strconv"

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

func logCheckoutDiscord(checkout *CheckoutLogRequest, discordWebhook string) {
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
	go logCheckoutBackend(checkout)
	go logCheckoutDiscord(checkout, discordWebhook)
}

// func Webhook(new_id data.ID, idx int) {
// 	var webhookURL = os.Getenv("SPAM")
// 	n_size := len(new_id.Results[idx].SellNows)
// 	var fields []data.Fields
// 	for i := 0; i < n_size; i++ {
// 		fields = append(fields, data.Fields{
// 			Name:   "Payout",
// 			Value:  "[" + strconv.Itoa(new_id.Results[idx].SellNows[i].Price) + " €" + " | " + new_id.Results[idx].SellNows[i].Size + "]" + "(" + "https://sell.wethenew.com/sell-now/" + strconv.Itoa(new_id.Results[idx].SellNows[i].ID) + "?holding-Lab" + ")",
// 			Inline: true,
// 		})
// 	}
// 	payload := &data.Top{
// 		Username:  "Wethenew Monitor",
// 		AvatarURL: Image_URL,
// 		Content:   "",
// 		Embeds: []data.Embeds{
// 			{
// 				Title: new_id.Results[idx].Name,
// 				// Description: "Sell Now",
// 				Color:  1999236,
// 				Fields: fields,
// 				Thumbnail: data.Thumbnail{
// 					URL: new_id.Results[idx].Image,
// 				},
// 				Footer: data.Footer{
// 					IconURL: Image_URL,
// 					Text:    "Wethenew | Holding-Lab " + Time(),
// 				},
// 			},
// 		},
// 	}
// 	payloadBuf := new(bytes.Buffer)
// 	_ = json.NewEncoder(payloadBuf).Encode(payload)

// 	if webhookURL == "" {
// 		fmt.Println("SET DISCORD_WEBHOOK_URL ENV VAR")
// 	}
// 	SendWebhook, err := http.NewRequest("POST", webhookURL, payloadBuf)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	SendWebhook.Header.Set("content-type", "application/json")

// 	sendWebhookRes, err := client.Do(SendWebhook)
// 	if err != nil {
// 		fmt.Print(err)
// 	}
// 	if sendWebhookRes.StatusCode != 204 {
// 		fmt.Printf("Webhook failed with status %d\n", sendWebhookRes.StatusCode)
// 	}
// 	defer sendWebhookRes.Body.Close()
// }
