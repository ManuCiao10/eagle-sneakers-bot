package client

import (
	"fmt"

	http "github.com/saucesteals/fhttp"
)

var (
	UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36"
)

func (r *Request) SetDefaultHeadersTBA() *Request {
	r.SetHeader("User-Agent", UserAgent)
	r.SetHeader("authority", "www.the-broken-arm.com")
	r.SetHeader("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8")
	r.SetHeader("accept-language", "en-GB,en;q=0.9")
	r.SetHeader("cache-control", "max-age=0")
	r.SetHeader("sec-ch-ua", `"Not?A_Brand";v="8", "Chromium";v="108", "Brave";v="108"`)
	r.SetHeader("sec-ch-ua-mobile", "?0")
	r.SetHeader("sec-ch-ua-platform", `"macOS"`)
	r.SetHeader("sec-fetch-dest", "document")
	r.SetHeader("sec-fetch-mode", "navigate")
	r.SetHeader("sec-fetch-site", "none")
	r.SetHeader("sec-fetch-user", "?1")
	r.SetHeader("sec-gpc", "1")
	r.SetHeader("upgrade-insecure-requests", "1")

	return r
}

func (r *Request) SetCookie(cookies []*http.Cookie) *Request {
	fmt.Println("set cookie...")
	r.cookies = append(r.cookies, cookies...)
	return r
}
