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

	return r
}

func (r *Request) SetHeadersMonitor() *Request {
	r.SetHeader("authority", "www.the-broken-arm.com")
	r.SetHeader("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8")
	r.SetHeader("accept-language", "en-GB,en;q=0.7")
	r.SetHeader("cache-control", "max-age=0")
	// r.SetHeader("cookie", "PHPSESSID=s2vjaar1s1hm5nvim1so9ro949; PrestaShop-b255acdcaf89d3f7cc8c1687088165cb=def5020075b6c74e228f544648ec9a547d7519a420b574216105675770664ce3123b5c3086a35cd89e24af1a06f2aba2e71a6c67cc85f257322b6d0809d6e4d3289d6f219ffe692ca852e8c17ad9e8fb15019894e4ba19f2d65dae4060c1faeda28317b1d178817b8ad964fc3625b0c293990ef70ef5b700c3c00fcd6cf428bbe69a4b5f703cb84585a2d4554448c39b3fb01c432bf4935ad4b0cc83d01426d7e8837d9c0ce62ae534b55037453aee89cb0c30507f6dc38bc1f0cb0e827f021a84867666950ec2646259c31f2dce3f2a1cccd0bb544ab17e7fe07327a1ffdb8da8334d1c419295384d411aa53ea97f7a21ab95d1902d2a980dd80f9f966f63fabdbb0fa6f8ddf01ee917609d4873ffa3c622578d4a3dbff265da2f48a0abae5e39053cf692687c75ba9cc01ed3; tarteaucitron=!analytics=true!gtag=true; cf_clearance=3TiED9b5DcA1rfD6mMJf.5ktsvE9ChwhGcLzKmWJ2Hw-1673965737-0-150")
	r.SetHeader("if-modified-since", "Mon, 23 Jan 2023 03:00:02 GMT")
	r.SetHeader("if-none-match", `"18337d-5f2e598775dd7-gzip"`)
	r.SetHeader("sec-ch-ua", `"Not_A Brand";v="99", "Brave";v="109", "Chromium";v="109"`)
	r.SetHeader("sec-ch-ua-mobile", "?0")
	r.SetHeader("sec-ch-ua-platform", `"macOS"`)
	r.SetHeader("sec-fetch-dest", "document")
	r.SetHeader("sec-fetch-mode", "navigate")
	r.SetHeader("sec-fetch-site", "none")
	r.SetHeader("sec-fetch-user", "?1")
	r.SetHeader("sec-gpc", "1")
	r.SetHeader("upgrade-insecure-requests", "1")
	r.SetHeader("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36")

	return r
}

//Fiver Request

func (r *Request) SetHeadersFiver() *Request {
	r.SetHeader("authority", "www.fiverr.com")
	r.SetHeader("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8")
	r.SetHeader("accept-language", "en-GB,en;q=0.8")
	r.SetHeader("cache-control", "no-cache")
	r.SetHeader("cookie", "_pxhd=p1uXuQl8k1w1TYtkNjM5dIZKG/r5LeSiGO8wDYqS8wPH90S4OHWLgWmz3Rk/KVFr499i9ZOFny1yap/q7aTilQ==:9YD7MVSyEivVEP7hcecKs3WkN0XEFSIhHR/7j-ldIziTRNdYaTQE6aQP3BxKOdmYyzW7I1HuN6xQXO7RlJcyNuDb398EatIMYzyx5B15aVU=; __cf_bm=O_FxsdZWqAtpsz_fKqhJq0ylCeQB2Dwj5X4.ljvyqpE-1675205624-0-AeWZNmzvJ4V+LqN0YvXOJrW/q+OKJoVdAkcM/U105kbh8C2T+/pFIX5zA1PxOUeCRCWeGL2rhL1N7TXhKyScrPM=; __cfruid=29ef2dfd6b17fa950aaa7a088ec1142fdb5512da-1675205624; pxcts=1d4b4ea0-a1ba-11ed-a6a3-5a4265727665; _pxvid=1d4b414e-a1ba-11ed-a6a3-5a4265727665; u_guid=1675205627000-18cce759f5ff4bb4efd8b985f04b23ea91d90398; logged_out_currency=EUR; page_views=5; _pxff_cc=U2FtZVNpdGU9TGF4Ow==; _pxff_rf=1; _px3=704a58f76cff8fa94580aaba13f22df0c2b4ad35da51bd3036ee82a01f32ecfe:Ot+KqjtDREiZ1fqYN+xcZNUoET1+CQtcTwBBU/zNErkrLGQ4yVoGHI01dqm3c02yYREHCNKxnRiA08mm9XTd8A==:1000:uxaUw0PHg1IbjY/PoHQ3lwPdo71ZlizgGDIjU4TzeTXWHKgLERUyDIbqJTIIHbm8ODxwdZjSEMvoqUAl6hyqZ3QMiQfZrF5TZYB1g1EZrFrTR1NmTqp3QKjAq8Pq73Z2JTT0YK/3OY8+S1Q/hWCdYSJAdiRx8yNXY2FXzeiOlNN/XHIQWwRIVOVYyWVLUE5NjFrBAiOXQHq2Qm/wasqmLQ==; _pxde=9bda4b2f93182fba33b9c65582a94b18481715d2279fc2a242eac59c6ae59ff0:eyJ0aW1lc3RhbXAiOjE2NzUyMDU4NzEzNTcsImZfa2IiOjAsImlwY19pZCI6W119; forterToken=348a48c33038413c9b0486ad3fbf2b4b_1675205870084_314_dUAL43b-mnts_15ck")
	r.SetHeader("pragma", "no-cache")
	r.SetHeader("sec-ch-ua", `"Not_A Brand";v="99", "Brave";v="109", "Chromium";v="109"`)
	r.SetHeader("sec-ch-ua-mobile", "?0")
	r.SetHeader("sec-ch-ua-platform", `"macOS"`)
	r.SetHeader("sec-fetch-dest", "document")
	r.SetHeader("sec-fetch-mode", "navigate")
	r.SetHeader("sec-fetch-site", "none")
	r.SetHeader("sec-fetch-user", "?1")
	r.SetHeader("sec-gpc", "1")
	r.SetHeader("upgrade-insecure-requests", "1")
	r.SetHeader("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36")

	return r
}

func (r *Request) SetHeadersFiverLogin() *Request {
	r.SetHeader("authority", "it.fiverr.com")
	r.SetHeader("accept", "application/json")
	r.SetHeader("accept-language", "en-GB,en;q=0.8")
	r.SetHeader("cache-control", "no-cache")
	r.SetHeader("content-type", "multipart/form-data; boundary=----WebKitFormBoundaryQvBXcpmMiTG40it8")
	r.SetHeader("cookie", "_pxhd=yXjrBHNfAQ1tbwXAcQ4PSxMkE1hQB5Xqn7GWOyEP5dstq79-ONIfimNRjSXy5fu-tGUYmCpy-qcj3-qIa345Yw==:hz5-LK4K/4Aah7I-L8kCnKCim-u5/F-0sI0vFibQRVFp/0WCaeMpkHWG-7xGiy8mYPtQHyQYFJG2/OYPhDwJAcNT/MAw16eXIvMae14UtWM=; u_guid=1675203866000-7f96f574368b419d970c929452e7dfcdbcd590a0; logged_out_currency=EUR; __cf_bm=9puBkQmbAdu5CRrJd_8TCOyl3otxhNOis7RtMa.cQfA-1675203866-0-AZCK/JW1PHv7X0VIcO9zpeUtigcq/YrvELN4TY3hCXduIM69CHgebraknDpLZRZTQc9UFdI/5tF7qv3U++54QzU=; __cfruid=26d82a1197e077797da9f442d7808761a5c8aa00-1675203866; pxcts=0650674b-a1b6-11ed-be51-784b74434f77; _pxvid=050a79da-a1b6-11ed-abd8-42796a656170; _pxff_cc=U2FtZVNpdGU9TGF4Ow==; _pxff_rf=1; forterToken=e722f6a8e64f4ed0ac2864935bde5fc8_1675204649368_965_dUAL43b-mnts-ants_15ck; _pxde=e6aeb4180c97e5c1e58053e121e14e873da81611d6279c6d956e2c9ba8da5d09:eyJ0aW1lc3RhbXAiOjE2NzUyMDQ2NjAzMzIsImZfa2IiOjAsImlwY19pZCI6W119")
	r.SetHeader("fvrr-page-ctx-id", "21d2e2e3ca508d88d9d1561d2ec535e3")
	r.SetHeader("origin", "https://it.fiverr.com")
	r.SetHeader("pragma", "no-cache")
	r.SetHeader("referer", "https://it.fiverr.com/")
	r.SetHeader("sec-ch-ua", `"Not_A Brand";v="99", "Brave";v="109", "Chromium";v="109"`)
	r.SetHeader("sec-ch-ua-mobile", "?0")
	r.SetHeader("sec-ch-ua-platform", `"macOS"`)
	r.SetHeader("sec-fetch-dest", "empty")
	r.SetHeader("sec-fetch-mode", "cors")
	r.SetHeader("sec-fetch-site", "same-origin")
	r.SetHeader("sec-gpc", "1")
	r.SetHeader("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36")
	r.SetHeader("x-csrf-token", "1676414248.BxlfmBYYfLIyfvccnxLJZ591ytRrtPTHYykbTnQUhRg=")
	r.SetHeader("x-requested-with", "XMLHttpRequest")

	return r
}

func (r *Request) SetHeadersFiverUser() *Request {
	r.SetHeader("authority", "www.fiverr.com")
	r.SetHeader("accept", "application/json")
	r.SetHeader("accept-language", "en-GB,en;q=0.8")
	r.SetHeader("cache-control", "no-cache")
	r.SetHeader("content-type", "multipart/form-data; boundary=----WebKitFormBoundary5vIfyM2mVYAlDZO0")
	r.SetHeader("cookie", "_pxhd=p1uXuQl8k1w1TYtkNjM5dIZKG/r5LeSiGO8wDYqS8wPH90S4OHWLgWmz3Rk/KVFr499i9ZOFny1yap/q7aTilQ==:9YD7MVSyEivVEP7hcecKs3WkN0XEFSIhHR/7j-ldIziTRNdYaTQE6aQP3BxKOdmYyzW7I1HuN6xQXO7RlJcyNuDb398EatIMYzyx5B15aVU=; __cf_bm=O_FxsdZWqAtpsz_fKqhJq0ylCeQB2Dwj5X4.ljvyqpE-1675205624-0-AeWZNmzvJ4V+LqN0YvXOJrW/q+OKJoVdAkcM/U105kbh8C2T+/pFIX5zA1PxOUeCRCWeGL2rhL1N7TXhKyScrPM=; __cfruid=29ef2dfd6b17fa950aaa7a088ec1142fdb5512da-1675205624; pxcts=1d4b4ea0-a1ba-11ed-a6a3-5a4265727665; _pxvid=1d4b414e-a1ba-11ed-a6a3-5a4265727665; u_guid=1675205627000-18cce759f5ff4bb4efd8b985f04b23ea91d90398; logged_out_currency=EUR; page_views=5; _px3=20fe2f680ace363b42b4930bc372dee68b20f768a081b7eb21393a7b152c8193:7bdzR+DTFqztzqrDU/jNZ8tQ/Kau6OP6JyNPThpuRffBOksIlgmdzybx9U524s4hNVDh3gyJn8D/JgQgZvP0Aw==:1000:nWhbAQkZh93vohtQIIJBU5fJ+B2rkbsRLjrk1GwvFD/0dX3cUYaBNG/qNu0UTrHMiykGKVPdN6P8loLuuMSQyDCH3WX8srIRaXeYKEa1FKZ1TZwQO3E9fUwwfbymfZHEDKQ9Vre7/ODMQM7Sf+Jfp4ixe0O6On1Y3o2Gy4p22w1pBz/JspxFQ0IuZILQXDheS5mMLBC9H/lfxG+GE0rXbg==; forterToken=348a48c33038413c9b0486ad3fbf2b4b_1675205874396_519_dUAL43b-mnts-ants_15ck; _pxde=49491222d4db04b46d0b8312f66fcf5f15b99ae7b614d34ad0d6f309af1af42f:eyJ0aW1lc3RhbXAiOjE2NzUyMDU5OTQzNjQsImZfa2IiOjAsImlwY19pZCI6W119")
	r.SetHeader("fvrr-page-ctx-id", "f715178a2cc2140a53e69f463d6501e3")
	r.SetHeader("origin", "https://www.fiverr.com")
	r.SetHeader("pragma", "no-cache")
	r.SetHeader("referer", "https://www.fiverr.com/")
	r.SetHeader("sec-ch-ua", `"Not_A Brand";v="99", "Brave";v="109", "Chromium";v="109"`)
	r.SetHeader("sec-ch-ua-mobile", "?0")
	r.SetHeader("sec-ch-ua-platform", `"macOS"`)
	r.SetHeader("sec-fetch-dest", "empty")
	r.SetHeader("sec-fetch-mode", "cors")
	r.SetHeader("sec-fetch-site", "same-origin")
	r.SetHeader("sec-gpc", "1")
	r.SetHeader("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36")
	r.SetHeader("x-csrf-token", "1676415473.gNVtT0VhytQ6oN1ggnYfZrw4FcxuE2nfonyrJyftQyM=")
	r.SetHeader("x-requested-with", "XMLHttpRequest")

	return r

}
