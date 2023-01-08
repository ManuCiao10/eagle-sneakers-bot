package client

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

func (r *Request) SetCartHeadersTBA() *Request {
	r.SetHeader("authority", "www.the-broken-arm.com")
	r.SetHeader("accept", "application/json, text/javascript, */*; q=0.01")
	r.SetHeader("accept-language", "en-GB,en;q=0.9")
	r.SetHeader("content-type", "application/x-www-form-urlencoded; charset=UTF-8")
	r.SetHeader("origin", "https://www.the-broken-arm.com")
	r.SetHeader("sec-ch-ua", `"Not?A_Brand";v="8", "Chromium";v="108", "Brave";v="108"`)
	r.SetHeader("sec-ch-ua-mobile", "?0")
	r.SetHeader("sec-ch-ua-platform", `"macOS"`)
	r.SetHeader("sec-fetch-dest", "empty")
	r.SetHeader("sec-fetch-mode", "cors")
	r.SetHeader("sec-fetch-site", "same-origin")
	r.SetHeader("sec-gpc", "1")
	r.SetHeader("User-Agent", UserAgent)
	r.SetHeader("x-requested-with", "XMLHttpRequest")

	return r
}

//SetLoginHeadersTBA sets the headers for the login request

func (r *Request) SetLoginHeadersTBA() *Request {
	r.SetHeader("authority", "www.the-broken-arm.com")
	r.SetHeader("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8")
	r.SetHeader("accept-language", "en-GB,en;q=0.9")
	r.SetHeader("cache-control", "max-age=0")
	r.SetHeader("content-type", "application/x-www-form-urlencoded")
	r.SetHeader("origin", "https://www.the-broken-arm.com")
	r.SetHeader("referer", "https://www.the-broken-arm.com/en/connexion")
	r.SetHeader("sec-ch-ua", `"Not?A_Brand";v="8", "Chromium";v="108", "Brave";v="108"`)
	r.SetHeader("sec-ch-ua-mobile", "?0")
	r.SetHeader("sec-ch-ua-platform", `"macOS"`)
	r.SetHeader("sec-fetch-dest", "document")
	r.SetHeader("sec-fetch-mode", "navigate")
	r.SetHeader("sec-fetch-site", "same-origin")
	r.SetHeader("sec-fetch-user", "?1")
	r.SetHeader("sec-gpc", "1")
	r.SetHeader("upgrade-insecure-requests", "1")
	r.SetHeader("user-agent", UserAgent)

	return r
}

func (r *Request) CheckoutHeders() *Request {
	r.SetHeader("authority", "www.the-broken-arm.com")
	r.SetHeader("accept", "*/*")
	r.SetHeader("accept-language", "en-GB,en;q=0.9")
	// r.SetHeader("referer", "https://www.the-broken-arm.com/en/panier?action=show")
	r.SetHeader("sec-ch-ua", `"Not?A_Brand";v="8", "Chromium";v="108", "Brave";v="108"`)
	r.SetHeader("sec-ch-ua-mobile", "?0")
	r.SetHeader("sec-ch-ua-platform", `"macOS"`)
	r.SetHeader("sec-fetch-dest", "empty")
	r.SetHeader("sec-fetch-mode", "cors")
	r.SetHeader("sec-fetch-site", "same-origin")
	r.SetHeader("sec-gpc", "1")
	r.SetHeader("user-agent", UserAgent)
	r.SetHeader("x-requested-with", "XMLHttpRequest")
	r.SetHeader("cookie", "PHPSESSID=s2vjaar1s1hm5nvim1so9ro949; tarteaucitron=!analytics=true!gtag=true; cf_clearance=zKFT6ybXKv_0XMJrje.yAEVeZTgxmXBA5k0pd9PlusM-1673194788-0-150; PrestaShop-b255acdcaf89d3f7cc8c1687088165cb=def5020075b6c74e228f544648ec9a547d7519a420b574216105675770664ce3123b5c3086a35cd89e24af1a06f2aba2e71a6c67cc85f257322b6d0809d6e4d3289d6f219ffe692ca852e8c17ad9e8fb15019894e4ba19f2d65dae4060c1faeda28317b1d178817b8ad964fc3625b0c293990ef70ef5b700c3c00fcd6cf428bbe69a4b5f703cb84585a2d4554448c39b3fb01c432bf4935ad4b0cc83d01426d7e8837d9c0ce62ae534b55037453aee89cb0c30507f6dc38bc1f0cb0e827f021a84867666950ec2646259c31f2dce3f2a1cccd0bb544ab17e7fe07327a1ffdb8da8334d1c419295384d411aa53ea97f7a21ab95d1902d2a980dd80f9f966f63fabdbb0fa6f8ddf01ee917609d4873ffa3c622578d4a3dbff265da2f48a0abae5e39053cf692687c75ba9cc01ed3; __cf_bm=a9OQAppYrihjeJbpC21MfkHNpWlv0yuAJ1KznhOp8lM-1673198487-0-Afh4vkCqgkda67Ksrd+SzLrql4BjI4PUtvplJisoEcrkOKYo2pzxOFfDykEml9eCSvS+prAfMgI3UwPwaUIYEWDFQ+iYP2WV5VKvXLU4Ykb4vwg0byHlO0olL7U4hE1jPQS26fLWL9JMKghM4CzROvs=")

	return r
}
