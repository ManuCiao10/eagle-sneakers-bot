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

func (r *Request) SetHeadersNike() *Request {
	r.SetHeader("authority", "accounts.nike.com")
	r.SetHeader("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8")
	r.SetHeader("accept-language", "en-GB,en-US;q=0.9,en;q=0.8")
	r.SetHeader("cache-control", "no-cache")
	// r.SetHeader("cookie", "anonymousId=A2EA02C36E5F4DFC208E78E85B0933D0; emperor_id=43fa760f-a768-4e5f-b904-adee3d10f3c3; did=0dac9f6e-2f63-4db3-9eeb-bb2a8c8088e9; AnalysisUserId=2.23.155.29.28597167575946623; audience_segmentation_performed=true; AKA_A2=A; geoloc=cc=IT,rc=,tp=vhigh,tz=GMT+1,la=45.47,lo=9.20; s_ecid=MCMID%7C18534471974391590551536879385143941762; AMCVS_F0935E09512D2C270A490D4D%40AdobeOrg=1; AMCV_F0935E09512D2C270A490D4D%40AdobeOrg=1994364360%7CMCMID%7C18534471974391590551536879385143941762%7CMCAID%7CNONE%7CMCOPTOUT-1675766666s%7CNONE%7CvVersion%7C3.4.0; bm_mi=CA85C431B79462E9A9CAC482FB9A0A39~YAAQvjRoaD8zy9OFAQAAkT8MKxKRPAF6ha9l+nKnnVvVCHhhqO2TNwRDN9YELBF/nl4McXT8an6HZmAPXkjIUN2o/4Xskc5IoGmC9ReWP2wvvuYpoSJcRZkPRGXgIKRB87VFV0MvI+cl4F2Uog2JP4X9pBQDSJRfby2Y8UOhkvOrkfuzv/5DCXCJKO504Ai5ZqVkiJ5M6mWwHHjZCWbjq2yy1mb51zLhpRV/j3EWmn7xN/3ZGI0dtmywEpdbQGBjkTtd6a0hacICKnY09uVKwvghl1Cdn6EhV8e+9q+OLcX/RKFYRChVsp6N+SNYoP1wzKWkCnH0Q5w=~1; ak_bmsc=47C9C25132F3A444B50E863A8300937D~000000000000000000000000000000~YAAQvjRoaI4zy9OFAQAAp0MMKxLYLdX38Hqoo1IJ7OAlT9DfkNjCZGuKH3xUFfFOsOxA+gskprRsPHpgi+VLaBRcoj7HiaUwZVC2xxlYsozom8nMM1i4T3ImvyamqTDtvebqj35033uXYUDI+HYsTOXR+7/qittx9JLujjol3x8qXy3F6iabXrzGegvb1OgRzvDke1JyiIRmvdhstfODKcmFHoDgKoNRokk7/MYsQcNjn2SaZREjgXbjEKu9v6F6xs6cMK/seKV/NpRKw/Ti/kr/FeOctmhC9KtQsMPewJxgO7BA0O38kQ/UQSi4uk6bw5Wtdy4e+ILkUqF7mUvMBto3+e6OTSEzs9KLrkH5gBN+xKXdG9mpzMLIlPXPsk8A7EWYKQh0mWbQA+A1mIjfb0dg1t5CFZMioCg4cbVtBC8xuS1d47FKA0hQQDxVCW6y4bfmsnH7PmXd3qiJhKXLT9gZ26l0/5T3hYi4xDrWp2EFOKY9tLQSAjBQa30c/iYsQvlbiEDkeUEMgrk=; bm_sz=66E01324A6F589FDB30E4BFB411B376F~YAAQPDRoaO86lM6FAQAAcpEOKxIQlCI9c6qEqoEaGO+HD6YJ30NXz3JL8K0In4EedsUCtUnn+0/uyGzxyX+x3dYt38XnIWLzZzBoAIiuyH+5ivCYb//eGU5/MsO2r7dOxNjF9ySed4Nz49RqHKO79rFkp4RfrGB+NgY3FVq+rkCQX6lU+cBKxeEGyzjFiF2jAB2pRq77rF0ySUiN+psxAtVBEnJu+HYvx6xb+ReriP/eEOaXaC3icCKKuzAqRh2MJZYbmcCO4miUMaQpIxf6cwjRQEqjRtJe9bpzIar3ZgEsLeZiRZnk8/hHuUVfNTBmTiT9YUzvEPW1ZPiEAONAlbOONypm+Kw97awkg8gWNHawjIayHcLqIxX+kXxMaSHBPvkW4FgfPNxrauo1dx2NCmHfRU27yroKccII9L/Ok2Xn/SsF6mh3cV+Jog==~4473665~4601924; ak_bmsc_nke-2.2-ssn=09yQ7DPqf2Wakuvq4waRtnjUPbaf9CT3bgZh9uhLQlyesij6IIZnZiJRZzWU3qwuAAFaB6qBIXRZbXqfZueduS9R2Uxli8h7SRmrO6EHc77ul5UxcLe0kQPhNdDN1vPpT3voArY9jF8xOrj6H3FMN1Yjayvkqc8fORIzoV2qMlOHlkCV5FeakA2; ak_bmsc_nke-2.2=09yQ7DPqf2Wakuvq4waRtnjUPbaf9CT3bgZh9uhLQlyesij6IIZnZiJRZzWU3qwuAAFaB6qBIXRZbXqfZueduS9R2Uxli8h7SRmrO6EHc77ul5UxcLe0kQPhNdDN1vPpT3voArY9jF8xOrj6H3FMN1Yjayvkqc8fORIzoV2qMlOHlkCV5FeakA2; bm_sv=81FB7F4D20DAAB87D3CD87AFB8707A9C~YAAQtjRoaOR6MiaGAQAAS2kVKxIqJxEsacn/NZxfogbXp/KQIXJcfDtniEz4+Uk43EaC8UymfOjEZrElJo8xJMv5+NhmkCAKJW0WnmjNxl26GzyYDswjATCwTLHkByOThS97qDz3xb1DlmlfwrLIFEVqZphG+WHn4sMp2y0vEaTgUv4b5CGK5r3hgKq/ctEkrTsrLOWx5HrzQgUYU5VnVHXN5oVNkdj1GG4VsCYyZKeoPPf8GevTiKkfL1f2Ux4=~1; _abck=9D2E407A6A904D33B2AE30DB6EDEC3C2~-1~YAAQtjRoaBB7MiaGAQAA6GsVKwlFsT6zK714KE0+U9cihj0MuxHkUD9T/QLTB6XBQGVNFO2w1Tf4FfDPZPPFrmCPUjG8B/0KaCHFW8CHl3DBquoZbtoCfqw5RWzcnejSBvP1Qk80brZM+6YccgZAXyktpmNP9gI2OEE2KR5Q1dOzk3KWYqT9D2YFmiX+nzW+btzVhxggk0N5MumjT1fe6Fda+PkyIrAXVs7f8FRWEE/4IQnu83UdXmcE2Su0/aNNzUq/NmTk1yBSJ/rBhcTSjjFje/PUW6gyAWhQEvQbKp1pJ5oDwDufVnuNgGVeE4aOx5W5w/WdSTvDebxmONZV8aFAIhdu42Vyb3Ojo6mgUKjvG3disdiSppn02sAlt25fTqmpMf6k9FzEKsfSx82gprJhL158qTipctUsBa8I/f/+wKbLQRQUoFYEAJnE+RLRbcHJ8OGJiSUb0KF/h8bQ4E9rW6VO4RRzN8BPzW9FHS5R8eOLf3oqh4HT66Is/RIVcLK4pm2yQ6hk4T4vopGO17jziME=~-1~-1~-1")
	r.SetHeader("pragma", "no-cache")
	r.SetHeader("referer", "https://www.nike.com/")
	r.SetHeader("sec-ch-ua", `"Not_A Brand";v="99", "Brave";v="109", "Chromium";v="109"`)
	r.SetHeader("sec-ch-ua-mobile", "?0")
	r.SetHeader("sec-ch-ua-platform", `"macOS"`)
	r.SetHeader("sec-fetch-dest", "document")
	r.SetHeader("sec-fetch-mode", "navigate")
	r.SetHeader("sec-fetch-site", "same-site")
	r.SetHeader("sec-fetch-user", "?1")
	r.SetHeader("sec-gpc", "1")
	r.SetHeader("upgrade-insecure-requests", "1")
	r.SetHeader("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36")

	return r
}
