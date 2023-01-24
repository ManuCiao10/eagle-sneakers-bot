package thebrokenarmonitor

import (
	"encoding/xml"
	"fmt"
	"regexp"
	"time"

	"github.com/eagle/client"
)

var links = []string{}

func FrontendLink() {
	fmt.Println("[Frontend] starting...")

	for {
		client, err := client.NewClient()

		if err != nil {
			fmt.Println("[Frontend] error getting client")
		}

		response, err := client.NewRequest().
			SetURL(frontendLink).
			SetMethod("GET").
			SetHeadersMonitor().
			Do()

		if err != nil {
			fmt.Println("[Frontend] error getting products")
			time.Sleep(3 * time.Second)
		}

		if response.StatusCode() == 200 {
			var data Urlset

			xml := xml.Unmarshal(response.Body(), &data)

			if xml != nil {
				fmt.Println("[Frontend] error parsing xml")
				time.Sleep(2 * time.Second)
				continue
			}

			fmt.Println("[Frontend] Successfully fetched data...")
			if firstRun {
				firstRun = false
				for url := range data.URL {
					url := data.URL[url].Loc
					// links = append(links, url)
					fmt.Println("[Frontend] Added link: ", url)
				}

			}

			if len(links) == len(data.URL) {
				fmt.Println("[Frontend] No new links found...")
				time.Sleep(2 * time.Second)
			}

			for url := range data.URL {
				if !contains(links, data.URL[url].Loc) {
					links = append(links, data.URL[url].Loc)
					fmt.Println("[Frontend] Added link: ", data.URL[url].Loc)

					for {
						fmt.Println("[Frontend] Getting product info...")
						url_test := "https://www.the-broken-arm.com/en/men/9427-43569-nike-air-force-1-low-slam-jam.html"
						response, err := client.NewRequest().
							// SetURL(data.URL[url].Loc).
							SetURL(url_test).
							SetMethod("GET").
							SetDefaultHeadersTBA().
							Do()

						if err != nil {
							fmt.Println("[Frontend] error getting product info")
							continue
						}

						//from response get the class "form-control form-control-select"

						//use regex to get the options
						getData(response.BodyAsString())

						// send to monitor
						if len(data.URL[url].Image) == 0 {
							continue
						}
						image := data.URL[url].Image[0].Loc

						if len(data.URL[url].Image[0].Title) == 0 {
							continue
						}

						title := data.URL[url].Image[0].Title

						// send to monitor
						fmt.Println("[Frontend] Sending to monitor...")

						// _, err := client.NewRequest().
						// 	SetURL("http://localhost:8080/").
						// 	SetMethod("POST").
						// 	SetHeadersMonitor().
						// 	Do()

						if err != nil {
							fmt.Println("[Frontend] error sending to monitor")
							time.Sleep(3 * time.Second)
							continue
						}
						fmt.Println(image, title)

						break

					}

				}

			}

		} else {
			fmt.Println("[Frontend] error getting products")
			time.Sleep(3 * time.Second)
		}

	}

}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func getData(body string) string {
	// fmt.Println(body)
	regex := regexp.MustCompile(`product-actions`)
	match := regex.FindStringSubmatch(body)
	
	// fmt.Println(match[0])
	fmt.Println(match[1])
	return match[1]

}
