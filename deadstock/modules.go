package deadstock

import (
	// "fmt"
	"fmt"
	"os"
	"strconv"
	"time"

	// "io"
	"io/ioutil"
	"math/rand"
	"strings"

	// "golang.org/x/net/html"
	// goHttp "net/http"

	http "github.com/bogdanfinn/fhttp"
	tls_client "github.com/bogdanfinn/tls-client"
	"github.com/fatih/color"
)

func Print_err(msg string) {
	color.Red("[Eagle 0.0.2]"+"["+time.Now().Format("15:04:05.000000")+"]"+" %s", msg)
	os.Exit(0)
}

func Print(msg string) {
	color.Magenta("[Eagle 0.0.2]"+"["+time.Now().Format("15:04:05.000000")+"]"+" %s", msg)
}

func Print_cart(msg string) {
	color.Cyan("[Eagle 0.0.2]"+"["+time.Now().Format("15:04:05.000000")+"]"+" %s", msg)
}

func Print_err_cart(msg string) {
	color.Yellow("[Eagle 0.0.2]"+"["+time.Now().Format("15:04:05.000000")+"]"+" %s", msg)
}
func preload_cart(client tls_client.HttpClient) string{
	req, err := http.NewRequest(http.MethodGet, "https://www.sugar.it/catalog/product/view/id/212183#eagle", nil)
	if err != nil {
		Print_err("REQUEST ERROR")
	}
	// req.Header = http.Header{
	// 	"autority":                  {"www.sugar.it"},
	// 	"accept":                    {"text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"},
	// 	"accept-language":           {"it-IT,it;q=0.9,en-US;q=0.8,en;q=0.7,de;q=0.6,fr;q=0.5"},
	// 	"cache-control":             {"no-cache"},
	// 	"pragma":                    {"no-cache"},
	// 	"referer":                   {"https://www.sugar.it/"},
	// 	"sec-ch-ua":                 {`"Chromium";v="106", "Google Chrome";v="106", "Not;A=Brand";v="99"`},
	// 	"sec-ch-ua-mobile":          {"?0"},
	// 	// "set-cookie":                {"PHPSESSID=759a5b96a9dc6838c5cc7cefe02be550; path=/; domain=.sugar.it; HttpOnly"},
	// 	"sec-ch-ua-platform":        {`"macOS"`},
	// 	"sec-fetch-dest":            {"document"},
	// 	"sec-fetch-mode":            {"navigate"},
	// 	"sec-fetch-site":            {"none"},
	// 	"sec-fetch-user":            {"?1"},
	// 	"upgrade-insecure-requests": {"1"},
	// 	"user-agent":                {"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36"},

	// 	http.HeaderOrderKey: {
	// 		"autority",
	// 		"accept",
	// 		"accept-language",
	// 		"cache-control",
	// 		"sec-ch-ua",
	// 		"sec-ch-ua-mobile",
	// 		"sec-ch-ua-platform",
	// 		"sec-fetch-dest",
	// 		"sec-fetch-mode",
	// 		"sec-fetch-site",
	// 		"sec-fetch-user",
	// 		"upgrade-insecure-requests",
	// 		"user-agent",
	// 	},
	// }
	resp, err := client.Do(req)
	if err != nil {
		Print_err("RESPONSE ERROR CART")
	}
	fmt.Println(req.Cookies())
	if resp.StatusCode != 200 {
		Print_err("STATUS CODE " + strconv.Itoa(resp.StatusCode))
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		Print_err("BODY ERROR")
	}
	defer resp.Body.Close()
	uenc := get_cart_url(string(bodyText))
	if len(uenc) == 0 {
		Print_err("CONNECTION ERROR")
	}
	return uenc
}

// cookies
// porxies
// keep the session
// do to not pirnt the golanf error
func payload_cart(uenc string, client tls_client.HttpClient) bool {
	var data = strings.NewReader(`------WebKitFormBoundaryeujuuUMOAExGTwND
	Content-Disposition: form-data; name="product"

	258156
	------WebKitFormBoundaryeujuuUMOAExGTwND
	Content-Disposition: form-data; name="selected_configurable_option"

	258150
	------WebKitFormBoundaryeujuuUMOAExGTwND
	Content-Disposition: form-data; name="related_product"


	------WebKitFormBoundaryeujuuUMOAExGTwND
	Content-Disposition: form-data; name="item"

	258156
	------WebKitFormBoundaryeujuuUMOAExGTwND
	Content-Disposition: form-data; name="form_key"

	BXNMKB1Z4E8hx7E1
	------WebKitFormBoundaryeujuuUMOAExGTwND
	Content-Disposition: form-data; name="super_attribute[99]"

	120
	------WebKitFormBoundaryeujuuUMOAExGTwND--`)

	req, err := http.NewRequest(http.MethodPost, "https://www.sugar.it/checkout/cart/add/uenc"+uenc, data)
	if err != nil {
		Print_err("REQUEST ERROR")
	}
	// req.Header.Set("authority", "www.sugar.it")
	// req.Header.Set("accept", "application/json, text/javascript, */*; q=0.01")
	// req.Header.Set("accept-language", "it-IT,it;q=0.9,en-US;q=0.8,en;q=0.7,de;q=0.6,fr;q=0.5")
	// req.Header.Set("cache-control", "no-cache")
	// // req.Header.Set("content-type", "multipart/form-data; boundary=----WebKitFormBoundaryeujuuUMOAExGTweE")
	// // req.Header.Set("set-cookie", "PHPSESSID=b0b0b0b0b0b0b0b0b0b0b0b0b0b0b0b0; path=/; domain=.sugar.it; HttpOnly")
	// // req.Header.Set("cookie", "rmStore=ald:20220924_1801|atrv:nmrHekKy67Q-4dg2BmmR8wQ5hCXCLzqi6Q; _gcl_au=1.1.756736159.1665812079; CookieConsent={stamp:%27-1%27%2Cnecessary:true%2Cpreferences:true%2Cstatistics:true%2Cmarketing:true%2Cver:1%2Cutc:1665812079043%2Ciab2:%27%27%2Cregion:%27CA%27}; _ga=GA1.1.1822933619.1665812079; mage-translation-storage=%7B%7D; mage-translation-file-version=%7B%7D; sugar_newsletter=1; _clck=1ljfpq6|1|f5q|0; mage-cache-storage=%7B%7D; mage-cache-storage-section-invalidation=%7B%7D; private_content_version=need_version; mage-messages=; recently_viewed_product=%7B%7D; recently_viewed_product_previous=%7B%7D; recently_compared_product=%7B%7D; recently_compared_product_previous=%7B%7D; product_data_storage=%7B%7D; PHPSESSID=9f01505fe6b6619f39476e69c2d62d51; X-Magento-Vary=c58cc7336841735bf5ef13185766282824a9d073; _hjSessionUser_2226440=eyJpZCI6IjJmMGI3NTcxLWIzMzEtNWUyMC04ZTIzLTc4YzQ5ZTNlNjNkZiIsImNyZWF0ZWQiOjE2NjU4MTIwNzkyMjAsImV4aXN0aW5nIjp0cnVlfQ==; _hjIncludedInSessionSample=0; _hjSession_2226440=eyJpZCI6IjVmNDI2ZTgwLTdlMGUtNGM3Ni04MjQ1LTI3Nzc1ZjJjMDQ0YiIsImNyZWF0ZWQiOjE2NjU4Njc1OTMyNzQsImluU2FtcGxlIjpmYWxzZX0=; _hjAbsoluteSessionInProgress=0; form_key=4JisAObqEaqk98cn; mage-cache-sessid=true; _clsk=1ob1izd|1665867627527|5|1|h.clarity.ms/collect; section_data_ids=%7B%22customer%22%3A1665867594%2C%22compare-products%22%3A1665867594%2C%22last-ordered-items%22%3A1665867594%2C%22cart%22%3A1665867595%2C%22directory-data%22%3A1665867594%2C%22review%22%3A1665867594%2C%22instant-purchase%22%3A1665867594%2C%22persistent%22%3A1665867594%2C%22captcha%22%3A1665867594%2C%22wishlist%22%3A1665867628%2C%22recently_viewed_product%22%3A1665867594%2C%22recently_compared_product%22%3A1665867594%2C%22product_data_storage%22%3A1665867594%2C%22paypal-billing-agreement%22%3A1665867594%2C%22checkout-fields%22%3A1665867594%2C%22collection-point-result%22%3A1665867594%2C%22pickup-location-result%22%3A1665867594%7D; _ga_1TT1ERKS8Z=GS1.1.1665867592.11.1.1665867921.60.0.0")
	// req.Header.Set("origin", "https://www.sugar.it")
	// req.Header.Set("pragma", "no-cache")
	// req.Header.Set("referer", "https://www.sugar.it/catalog/product/view/id/212183")
	// req.Header.Set("sec-ch-ua", `"Chromium";v="106", "Google Chrome";v="106", "Not;A=Brand";v="99"`)
	// req.Header.Set("sec-ch-ua-mobile", "?0")
	// req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	// req.Header.Set("sec-fetch-dest", "empty")
	// req.Header.Set("sec-fetch-mode", "cors")
	// req.Header.Set("sec-fetch-site", "same-origin")
	// req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36")
	// req.Header.Set("x-requested-with", "XMLHttpRequest")
	response, err := client.Do(req)
	if err != nil {
		Print_err("RESPONSE ERROR")
	}
	// defer response.Body.Close()
	// f, _ := os.Create("output.txt")
	// defer f.Close()
	// response.Write(f)
	fmt.Println(req.Cookies())
	fmt.Println(response.Cookies())
	if response.StatusCode == 200 {
		Print_cart("ADDED TO CART " + "<|" + response.Status + "|>")
		return true
	} else {
		Print_err("STATUS CODE " + strconv.Itoa(response.StatusCode))
		return false
	}

}

func Module_deadstock(profile []Info) {
	defer timer("main")()
	rand.Seed(time.Now().UnixNano())
	Print("PREPARING SESSION") // + strings.ToUpper(profile[0].Profile_name)
	options := []tls_client.HttpClientOption{
		tls_client.WithTimeout(7),
		tls_client.WithClientProfile(tls_client.Chrome_105),
		// tls_client.WithNotFollowRedirects(),
		//tls_client.WithProxyUrl("http://user:pass@host:ip"),
		tls_client.WithInsecureSkipVerify(),
		// tls_client.WithCookieJar(cJar), // create cookieJar instance and pass it as argument
	}
	client, err := tls_client.NewHttpClient(tls_client.NewNoopLogger(), options...)
	if err != nil {
		Print_err("CLIENT ERROR")
	}

	// webkit := randomString(16)
	var uenc = preload_cart(client)
	if len(uenc) == 0 {
		Print_err("CONNECTION ERROR")
	}

	//----------------------------------------------------------------//

	if payload_cart(string(uenc), client) {
		onestepcheckout(string(uenc), client)
	} else {
		Print_err("PAYLOAD ERROR")
	}

}

func onestepcheckout(uenc string, client tls_client.HttpClient) {
	Print("PREPARING CHECKOUT")
	r, err := http.NewRequest("GET", "https://www.sugar.it/onestepcheckout/", nil)
	if err != nil {
		Print_err("REQUEST ERROR")
	}
	// r.Header.Set("authority", "www.sugar.it")
	// r.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	// r.Header.Set("accept-language", "it-IT,it;q=0.9,en-US;q=0.8,en;q=0.7,de;q=0.6,fr;q=0.5")
	// r.Header.Set("cache-control", "no-cache")
	// // r.Header.Set("set-cookie", "PHPSESSID=759a5b96a9dc6838c5cc7cefe02be550; path=/; domain=.sugar.it; HttpOnly")
	// // r.Header.Set("cookie", "X-Magento-Vary=c58cc7336841735bf5ef13185766282824a9d073; rmStore=ald:20220924_1801|atrv:nmrHekKy67Q-4dg2BmmR8wQ5hCXCLzqi6Q; _gcl_au=1.1.1605308678.1665258211; CookieConsent={stamp:%27-1%27%2Cnecessary:true%2Cpreferences:true%2Cstatistics:true%2Cmarketing:true%2Cver:1%2Cutc:1665258211606%2Ciab2:%27%27%2Cregion:%27CA%27}; _ga=GA1.1.1518477362.1665258212; mage-translation-storage=%7B%7D; mage-translation-file-version=%7B%7D; sugar_newsletter=1; _clck=70oxf|1|f5j|0; _hjFirstSeen=1; _hjSession_2226440=eyJpZCI6IjEwZmJlYmJmLTY0ZWUtNDA3Zi1iMTc2LTJmMjFkMzJlZGFkYiIsImNyZWF0ZWQiOjE2NjUyNTgyMTIyNTcsImluU2FtcGxlIjpmYWxzZX0=; _hjAbsoluteSessionInProgress=0; form_key=BXNMKB1Z4E8hx7E1; mage-cache-storage=%7B%7D; mage-cache-storage-section-invalidation=%7B%7D; mage-cache-sessid=true; recently_viewed_product=%7B%7D; recently_viewed_product_previous=%7B%7D; recently_compared_product=%7B%7D; recently_compared_product_previous=%7B%7D; product_data_storage=%7B%7D; _hjSessionUser_2226440=eyJpZCI6ImRjZmQ5OTdmLTZlYzAtNWFlYS1iMDRkLTdmMTY5OWU2MzAwNSIsImNyZWF0ZWQiOjE2NjUyNTgyMTE5MDUsImV4aXN0aW5nIjp0cnVlfQ==; private_content_version=e62144218afb7fe41a3a64ee5e84b006; _hjIncludedInPageviewSample=1; _hjIncludedInSessionSample=0; PHPSESSID=48d5f2625c05351736e52b3fa8cb8018; mage-messages=; __stripe_mid=4723292b-3a63-450d-9830-726dfa3412116411af; __stripe_sid=088b8c90-3593-46c4-ae2f-449f08254d44233a21; _ga_1TT1ERKS8Z=GS1.1.1665258211.1.1.1665259758.39.0.0; _clsk=1dcrruj|1665259758761|14|1|h.clarity.ms/collect; section_data_ids=%7B%22customer%22%3A1665258813%2C%22compare-products%22%3A1665258813%2C%22last-ordered-items%22%3A1665258813%2C%22cart%22%3A1665259750%2C%22directory-data%22%3A1665258813%2C%22review%22%3A1665258813%2C%22instant-purchase%22%3A1665258813%2C%22persistent%22%3A1665258813%2C%22captcha%22%3A1665258813%2C%22wishlist%22%3A1665259759%2C%22recently_viewed_product%22%3A1665258813%2C%22recently_compared_product%22%3A1665258813%2C%22product_data_storage%22%3A1665258813%2C%22paypal-billing-agreement%22%3A1665258813%2C%22checkout-fields%22%3A1665258813%2C%22collection-point-result%22%3A1665258813%2C%22pickup-location-result%22%3A1665258813%7D")
	// r.Header.Set("pragma", "no-cache")
	// r.Header.Set("referer", "https://www.sugar.it/checkout/cart/add/uenc"+uenc)
	// r.Header.Set("sec-ch-ua", `"Google Chrome";v="105", "Not)A;Brand";v="8", "Chromium";v="105"`)
	// r.Header.Set("sec-ch-ua-mobile", "?0")
	// r.Header.Set("sec-ch-ua-platform", `"macOS"`)
	// r.Header.Set("sec-fetch-dest", "document")
	// r.Header.Set("sec-fetch-mode", "navigate")
	// r.Header.Set("sec-fetch-site", "same-origin")
	// r.Header.Set("sec-fetch-user", "?1")
	// r.Header.Set("upgrade-insecure-requests", "1")
	// r.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36")
	resp, _ := client.Do(r)
	bodyText1, _ := ioutil.ReadAll(resp.Body)
	if err != nil {
		Print_err("RESPONSE ERROR")
	}
	fmt.Println(r.Cookies())
	fmt.Println(resp.Cookies())
	// c, _ := os.Create("cart.txt")
	// c.Write(bodyText1)
	// defer c.Close()

	// resp.Write(c)
	// fmt.Printf("%s\n", bodyText1)
	// fmt.Println(resp.StatusCode)
	entity_id := get_identity_id(string(bodyText1))
	fmt.Println(entity_id)

}

// //-------------------------------------------------//
// var data1 = strings.NewReader(`{"cartId":"Lm9PxqYEQdwZ8PG1oTjKl51XiJJxjTGl","paymentMethod":{"method":"paypal_express","po_number":null,"additional_data":null},"billingAddress":{"countryId":"IT","region":"Italia","company":"","telephone":"3662299421","postcode":"50121","city":"Firenze","firstname":"emanuele","lastname":"ardinghi","customAttributes":[{"attribute_code":"gender","value":"1"}],"saveInAddressBook":null}}`)
// req3, err := http.NewRequest("POST", "https://www.sugar.it/rest/default/V1/guest-carts/Lm9PxqYEQdwZ8PG1oTjKl51XiJJxjTGl/discount-payment-method", data1)
// if err != nil {
// 	log.Fatal(err)
// }
// req3.Header.Set("authority", "www.sugar.it")
// req3.Header.Set("accept", "*/*")
// req3.Header.Set("accept-language", "it-IT,it;q=0.9,en-US;q=0.8,en;q=0.7,de;q=0.6,fr;q=0.5")
// req3.Header.Set("cache-control", "no-cache")
// req3.Header.Set("content-type", "application/json")
// req3.Header.Set("cookie", "rmStore=ald:20220924_1801|atrv:nmrHekKy67Q-4dg2BmmR8wQ5hCXCLzqi6Q; _gcl_au=1.1.1605308678.1665258211; CookieConsent={stamp:%27-1%27%2Cnecessary:true%2Cpreferences:true%2Cstatistics:true%2Cmarketing:true%2Cver:1%2Cutc:1665258211606%2Ciab2:%27%27%2Cregion:%27CA%27}; _ga=GA1.1.1518477362.1665258212; mage-translation-storage=%7B%7D; mage-translation-file-version=%7B%7D; sugar_newsletter=1; _clck=70oxf|1|f5j|0; mage-cache-storage=%7B%7D; mage-cache-storage-section-invalidation=%7B%7D; recently_viewed_product=%7B%7D; recently_viewed_product_previous=%7B%7D; recently_compared_product=%7B%7D; recently_compared_product_previous=%7B%7D; product_data_storage=%7B%7D; _hjSessionUser_2226440=eyJpZCI6ImRjZmQ5OTdmLTZlYzAtNWFlYS1iMDRkLTdmMTY5OWU2MzAwNSIsImNyZWF0ZWQiOjE2NjUyNTgyMTE5MDUsImV4aXN0aW5nIjp0cnVlfQ==; mage-messages=; __stripe_mid=4723292b-3a63-450d-9830-726dfa3412116411af; X-Magento-Vary=c58cc7336841735bf5ef13185766282824a9d073; _hjIncludedInSessionSample=0; _hjSession_2226440=eyJpZCI6IjYwODkzNDE5LWI4ZmYtNGVjOC1iMDAwLTFiYjE3ZGM0NGJkZCIsImNyZWF0ZWQiOjE2NjUyNzMyNjUxNjAsImluU2FtcGxlIjpmYWxzZX0=; _hjIncludedInPageviewSample=1; _hjAbsoluteSessionInProgress=0; form_key=MIqbsArjefNmX8y6; mage-cache-sessid=true; private_content_version=6db6bb45bf63d621d330d916f72e9efd; PHPSESSID=dccfdfb6c2c11be3454ac1c74235b674; _ga_1TT1ERKS8Z=GS1.1.1665273262.2.1.1665273300.22.0.0; _clsk=2j80xi|1665273300857|6|1|h.clarity.ms/collect; __stripe_sid=d8df2385-f922-4d07-8d58-8ff1641f3a5d296072; section_data_ids=%7B%22customer%22%3A1665273302%2C%22compare-products%22%3A1665273302%2C%22last-ordered-items%22%3A1665273302%2C%22cart%22%3A1665273303%2C%22directory-data%22%3A1665273302%2C%22review%22%3A1665273302%2C%22instant-purchase%22%3A1665273302%2C%22persistent%22%3A1665273302%2C%22captcha%22%3A1665273302%2C%22wishlist%22%3A1665273302%2C%22recently_viewed_product%22%3A1665273302%2C%22recently_compared_product%22%3A1665273302%2C%22product_data_storage%22%3A1665273302%2C%22paypal-billing-agreement%22%3A1665273302%2C%22checkout-fields%22%3A1665273302%2C%22collection-point-result%22%3A1665273302%2C%22pickup-location-result%22%3A1665273302%7D")
// req3.Header.Set("origin", "https://www.sugar.it")
// req3.Header.Set("pragma", "no-cache")
// req3.Header.Set("referer", "https://www.sugar.it/onestepcheckout/")
// req3.Header.Set("sec-ch-ua", `"Google Chrome";v="105", "Not)A;Brand";v="8", "Chromium";v="105"`)
// req3.Header.Set("sec-ch-ua-mobile", "?0")
// req3.Header.Set("sec-ch-ua-platform", `"macOS"`)
// req3.Header.Set("sec-fetch-dest", "empty")
// req3.Header.Set("sec-fetch-mode", "cors")
// req3.Header.Set("sec-fetch-site", "same-origin")
// req3.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36")
// req3.Header.Set("x-requested-with", "XMLHttpRequest")
// resp3, err := client.Do(req3)
// if err != nil {
// 	log.Fatal(err)
// }
// fmt.Println(resp3.Status)
// var data4 = strings.NewReader(`{"addressInformation":{"shipping_address":{"countryId":"IT","region":"Italia","street":["via orcagna"],"company":"","telephone":"3662299421","postcode":"50121","city":"Firenze","firstname":"emanuele","lastname":"ardinghi","customAttributes":[{"attribute_code":"gender","value":"1"}],"extension_attributes":{}},"billing_address":{"countryId":"IT","region":"Italia","street":["via orcagna"],"company":"","telephone":"3662299421","postcode":"50121","city":"Firenze","firstname":"emanuele","lastname":"ardinghi","customAttributes":[{"attribute_code":"gender","value":"1"}],"saveInAddressBook":null},"shipping_method_code":"bestway","shipping_carrier_code":"tablerate"},"customerAttributes":{"gender":"1"},"additionInformation":{"register":false,"same_as_shipping":true}}`)
// req4, err := http.NewRequest("POST", "https://www.sugar.it/rest/default/V1/guest-carts/QUQVOfYWtseeiUjDoaZQ6hcLngPIx7WS/checkout-information", data4)
// if err != nil {
// 	log.Fatal(err)
// }
// req4.Header.Set("authority", "www.sugar.it")
// req4.Header.Set("accept", "*/*")
// req4.Header.Set("accept-language", "it-IT,it;q=0.9,en-US;q=0.8,en;q=0.7,de;q=0.6,fr;q=0.5")
// req4.Header.Set("cache-control", "no-cache")
// req4.Header.Set("content-type", "application/json")
// req4.Header.Set("cookie", "rmStore=ald:20220924_1801|atrv:nmrHekKy67Q-4dg2BmmR8wQ5hCXCLzqi6Q; _gcl_au=1.1.1605308678.1665258211; CookieConsent={stamp:%27-1%27%2Cnecessary:true%2Cpreferences:true%2Cstatistics:true%2Cmarketing:true%2Cver:1%2Cutc:1665258211606%2Ciab2:%27%27%2Cregion:%27CA%27}; _ga=GA1.1.1518477362.1665258212; mage-translation-storage=%7B%7D; mage-translation-file-version=%7B%7D; sugar_newsletter=1; _clck=70oxf|1|f5j|0; mage-cache-storage=%7B%7D; mage-cache-storage-section-invalidation=%7B%7D; recently_viewed_product=%7B%7D; recently_viewed_product_previous=%7B%7D; recently_compared_product=%7B%7D; recently_compared_product_previous=%7B%7D; product_data_storage=%7B%7D; _hjSessionUser_2226440=eyJpZCI6ImRjZmQ5OTdmLTZlYzAtNWFlYS1iMDRkLTdmMTY5OWU2MzAwNSIsImNyZWF0ZWQiOjE2NjUyNTgyMTE5MDUsImV4aXN0aW5nIjp0cnVlfQ==; mage-messages=; __stripe_mid=4723292b-3a63-450d-9830-726dfa3412116411af; X-Magento-Vary=c58cc7336841735bf5ef13185766282824a9d073; _hjIncludedInSessionSample=0; _hjSession_2226440=eyJpZCI6IjYwODkzNDE5LWI4ZmYtNGVjOC1iMDAwLTFiYjE3ZGM0NGJkZCIsImNyZWF0ZWQiOjE2NjUyNzMyNjUxNjAsImluU2FtcGxlIjpmYWxzZX0=; _hjIncludedInPageviewSample=1; _hjAbsoluteSessionInProgress=0; form_key=MIqbsArjefNmX8y6; mage-cache-sessid=true; private_content_version=6db6bb45bf63d621d330d916f72e9efd; PHPSESSID=dccfdfb6c2c11be3454ac1c74235b674; _ga_1TT1ERKS8Z=GS1.1.1665273262.2.1.1665273300.22.0.0; _clsk=2j80xi|1665273300857|6|1|h.clarity.ms/collect; __stripe_sid=d8df2385-f922-4d07-8d58-8ff1641f3a5d296072; section_data_ids=%7B%22customer%22%3A1665273302%2C%22compare-products%22%3A1665273302%2C%22last-ordered-items%22%3A1665273302%2C%22cart%22%3A1665273303%2C%22directory-data%22%3A1665273302%2C%22review%22%3A1665273302%2C%22instant-purchase%22%3A1665273302%2C%22persistent%22%3A1665273302%2C%22captcha%22%3A1665273302%2C%22wishlist%22%3A1665273302%2C%22recently_viewed_product%22%3A1665273302%2C%22recently_compared_product%22%3A1665273302%2C%22product_data_storage%22%3A1665273302%2C%22paypal-billing-agreement%22%3A1665273302%2C%22checkout-fields%22%3A1665273302%2C%22collection-point-result%22%3A1665273302%2C%22pickup-location-result%22%3A1665273302%7D")
// req4.Header.Set("origin", "https://www.sugar.it")
// req4.Header.Set("pragma", "no-cache")
// req4.Header.Set("referer", "https://www.sugar.it/onestepcheckout/")
// req4.Header.Set("sec-ch-ua", `"Google Chrome";v="105", "Not)A;Brand";v="8", "Chromium";v="105"`)
// req4.Header.Set("sec-ch-ua-mobile", "?0")
// req4.Header.Set("sec-ch-ua-platform", `"macOS"`)
// req4.Header.Set("sec-fetch-dest", "empty")
// req4.Header.Set("sec-fetch-mode", "cors")
// req4.Header.Set("sec-fetch-site", "same-origin")
// req4.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36")
// req4.Header.Set("x-requested-with", "XMLHttpRequest")
// resp4, err := client.Do(req4)
// if err != nil {
// 	log.Fatal(err)
// }
// fmt.Println(resp4.Status)
// // fmt.Println(resp4)
// req5, err := http.NewRequest("GET", "https://www.sugar.it/paypal/express/start/", nil)
// if err != nil {
// 	log.Fatal(err)
// }
// req5.Header.Set("authority", "www.sugar.it")
// req5.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
// req5.Header.Set("accept-language", "it-IT,it;q=0.9,en-US;q=0.8,en;q=0.7,de;q=0.6,fr;q=0.5")
// req5.Header.Set("cache-control", "no-cache")
// req5.Header.Set("cookie", "rmStore=ald:20220924_1801|atrv:nmrHekKy67Q-4dg2BmmR8wQ5hCXCLzqi6Q; _gcl_au=1.1.1605308678.1665258211; CookieConsent={stamp:%27-1%27%2Cnecessary:true%2Cpreferences:true%2Cstatistics:true%2Cmarketing:true%2Cver:1%2Cutc:1665258211606%2Ciab2:%27%27%2Cregion:%27CA%27}; _ga=GA1.1.1518477362.1665258212; mage-translation-storage=%7B%7D; mage-translation-file-version=%7B%7D; sugar_newsletter=1; _clck=70oxf|1|f5j|0; mage-cache-storage=%7B%7D; mage-cache-storage-section-invalidation=%7B%7D; recently_viewed_product=%7B%7D; recently_viewed_product_previous=%7B%7D; recently_compared_product=%7B%7D; recently_compared_product_previous=%7B%7D; product_data_storage=%7B%7D; _hjSessionUser_2226440=eyJpZCI6ImRjZmQ5OTdmLTZlYzAtNWFlYS1iMDRkLTdmMTY5OWU2MzAwNSIsImNyZWF0ZWQiOjE2NjUyNTgyMTE5MDUsImV4aXN0aW5nIjp0cnVlfQ==; mage-messages=; __stripe_mid=4723292b-3a63-450d-9830-726dfa3412116411af; X-Magento-Vary=c58cc7336841735bf5ef13185766282824a9d073; _hjIncludedInSessionSample=0; _hjSession_2226440=eyJpZCI6IjYwODkzNDE5LWI4ZmYtNGVjOC1iMDAwLTFiYjE3ZGM0NGJkZCIsImNyZWF0ZWQiOjE2NjUyNzMyNjUxNjAsImluU2FtcGxlIjpmYWxzZX0=; _hjIncludedInPageviewSample=1; _hjAbsoluteSessionInProgress=0; form_key=MIqbsArjefNmX8y6; mage-cache-sessid=true; private_content_version=6db6bb45bf63d621d330d916f72e9efd; PHPSESSID=dccfdfb6c2c11be3454ac1c74235b674; _ga_1TT1ERKS8Z=GS1.1.1665273262.2.1.1665273300.22.0.0; _clsk=2j80xi|1665273300857|6|1|h.clarity.ms/collect; __stripe_sid=d8df2385-f922-4d07-8d58-8ff1641f3a5d296072; section_data_ids=%7B%22customer%22%3A1665274302%2C%22compare-products%22%3A1665273302%2C%22last-ordered-items%22%3A1665273302%2C%22cart%22%3A1665274303%2C%22directory-data%22%3A1665273302%2C%22review%22%3A1665273302%2C%22instant-purchase%22%3A1665273302%2C%22persistent%22%3A1665273302%2C%22captcha%22%3A1665273302%2C%22wishlist%22%3A1665273302%2C%22recently_viewed_product%22%3A1665273302%2C%22recently_compared_product%22%3A1665273302%2C%22product_data_storage%22%3A1665273302%2C%22paypal-billing-agreement%22%3A1665273302%2C%22checkout-fields%22%3A1665273302%2C%22collection-point-result%22%3A1665273302%2C%22pickup-location-result%22%3A1665273302%2C%22messages%22%3A2000%7D")
// req5.Header.Set("pragma", "no-cache")
// req5.Header.Set("referer", "https://www.sugar.it/onestepcheckout/")
// req5.Header.Set("sec-ch-ua", `"Google Chrome";v="105", "Not)A;Brand";v="8", "Chromium";v="105"`)
// req5.Header.Set("sec-ch-ua-mobile", "?0")
// req5.Header.Set("sec-ch-ua-platform", `"macOS"`)
// req5.Header.Set("sec-fetch-dest", "document")
// req5.Header.Set("sec-fetch-mode", "navigate")
// req5.Header.Set("sec-fetch-site", "same-origin")
// req5.Header.Set("sec-fetch-user", "?1")
// req5.Header.Set("upgrade-insecure-requests", "1")
// req5.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36")
// resp5, err := client.Do(req5)
// if err != nil {
// 	log.Fatal(err)
// }
// fmt.Println(resp5.Status)
// fmt.Println(req5)
// var data6 = strings.NewReader(`{"addressInformation":{"shipping_address":{"countryId":"IT","region":"Italia","street":["via orcagna"],"company":"","telephone":"3662299421","postcode":"50121","city":"Firenze","firstname":"emanuele","lastname":"ardinghi","customAttributes":[{"attribute_code":"gender","value":"1"}],"extension_attributes":{}},"billing_address":{"countryId":"IT","region":"Italia","street":["via orcagna"],"company":"","telephone":"3662299421","postcode":"50121","city":"Firenze","firstname":"emanuele","lastname":"ardinghi","customAttributes":[{"attribute_code":"gender","value":"1"}],"saveInAddressBook":null},"shipping_method_code":"bestway","shipping_carrier_code":"tablerate"},"customerAttributes":{"gender":"1"},"additionInformation":{"register":false,"same_as_shipping":true}}`)
// req6, err := http.NewRequest("POST", "https://www.sugar.it/rest/default/V1/guest-carts/QUQVOfYWtseeiUjDoaZQ6hcLngPIx7WS/checkout-information", data6)
// if err != nil {
// 	log.Fatal(err)
// }
// req6.Header.Set("authority", "www.sugar.it")
// req6.Header.Set("accept", "*/*")
// req6.Header.Set("accept-language", "it-IT,it;q=0.9,en-US;q=0.8,en;q=0.7,de;q=0.6,fr;q=0.5")
// req6.Header.Set("cache-control", "no-cache")
// req6.Header.Set("content-type", "application/json")
// req6.Header.Set("cookie", "rmStore=ald:20220924_1801|atrv:nmrHekKy67Q-4dg2BmmR8wQ5hCXCLzqi6Q; _gcl_au=1.1.1605308678.1665258211; CookieConsent={stamp:%27-1%27%2Cnecessary:true%2Cpreferences:true%2Cstatistics:true%2Cmarketing:true%2Cver:1%2Cutc:1665258211606%2Ciab2:%27%27%2Cregion:%27CA%27}; _ga=GA1.1.1518477362.1665258212; mage-translation-storage=%7B%7D; mage-translation-file-version=%7B%7D; sugar_newsletter=1; mage-cache-storage=%7B%7D; mage-cache-storage-section-invalidation=%7B%7D; recently_viewed_product=%7B%7D; recently_viewed_product_previous=%7B%7D; recently_compared_product=%7B%7D; recently_compared_product_previous=%7B%7D; product_data_storage=%7B%7D; _hjSessionUser_2226440=eyJpZCI6ImRjZmQ5OTdmLTZlYzAtNWFlYS1iMDRkLTdmMTY5OWU2MzAwNSIsImNyZWF0ZWQiOjE2NjUyNTgyMTE5MDUsImV4aXN0aW5nIjp0cnVlfQ==; mage-messages=; __stripe_mid=4723292b-3a63-450d-9830-726dfa3412116411af; X-Magento-Vary=c58cc7336841735bf5ef13185766282824a9d073; _hjSession_2226440=eyJpZCI6IjYwODkzNDE5LWI4ZmYtNGVjOC1iMDAwLTFiYjE3ZGM0NGJkZCIsImNyZWF0ZWQiOjE2NjUyNzMyNjUxNjAsImluU2FtcGxlIjpmYWxzZX0=; _hjAbsoluteSessionInProgress=0; form_key=MIqbsArjefNmX8y6; mage-cache-sessid=true; private_content_version=6db6bb45bf63d621d330d916f72e9efd; __stripe_sid=d8df2385-f922-4d07-8d58-8ff1641f3a5d296072; PHPSESSID=2e9c32d3fec0a97bc0314dd1e30658e7; _clck=70oxf|1|f5k|0; _clsk=2j80xi|1665274331170|7|1|h.clarity.ms/collect; _hjIncludedInPageviewSample=1; _hjIncludedInSessionSample=0; _ga_1TT1ERKS8Z=GS1.1.1665273262.2.1.1665274331.59.0.0; section_data_ids=%7B%22customer%22%3A1665274332%2C%22compare-products%22%3A1665273302%2C%22last-ordered-items%22%3A1665273302%2C%22cart%22%3A1665274333%2C%22directory-data%22%3A1665273302%2C%22review%22%3A1665273302%2C%22instant-purchase%22%3A1665273302%2C%22persistent%22%3A1665273302%2C%22captcha%22%3A1665273302%2C%22wishlist%22%3A1665274334%2C%22recently_viewed_product%22%3A1665273302%2C%22recently_compared_product%22%3A1665273302%2C%22product_data_storage%22%3A1665273302%2C%22paypal-billing-agreement%22%3A1665273302%2C%22checkout-fields%22%3A1665273302%2C%22collection-point-result%22%3A1665273302%2C%22pickup-location-result%22%3A1665273302%7D")
// req6.Header.Set("origin", "https://www.sugar.it")
// req6.Header.Set("pragma", "no-cache")
// req6.Header.Set("referer", "https://www.sugar.it/onestepcheckout/")
// req6.Header.Set("sec-ch-ua", `"Google Chrome";v="105", "Not)A;Brand";v="8", "Chromium";v="105"`)
// req6.Header.Set("sec-ch-ua-mobile", "?0")
// req6.Header.Set("sec-ch-ua-platform", `"macOS"`)
// req6.Header.Set("sec-fetch-dest", "empty")
// req6.Header.Set("sec-fetch-mode", "cors")
// req6.Header.Set("sec-fetch-site", "same-origin")
// req6.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36")
// req6.Header.Set("x-requested-with", "XMLHttpRequest")
// resp6, err := client.Do(req6)
// if err != nil {
// 	log.Fatal(err)
// }
// fmt.Println(resp6)
// time.Sleep(2 * time.Second)

// func randomString(n int) string {
// 	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
// 	sb := strings.Builder{}
// 	sb.Grow(n)
// 	for i := 0; i < n; i++ {
// 		sb.WriteByte(charset[rand.Intn(len(charset))])
// 	}
// 	return sb.String()
// }

func get_cart_url(bodyText string) string {
	if strings.Contains(bodyText, "503 Service Unavailable") {
		Print_err("503 SERVICE UNAVAILABLE")
	}
	return strings.Split(strings.Split(bodyText, "add/uenc")[1][:132], "\"")[0]

}

func get_identity_id(bodyText string) string {
	// content, _ := ioutil.ReadFile("cart.txt")
	if strings.Contains(bodyText, "entity_id") {
		return strings.Split(strings.Split(bodyText, "entity_id")[1], "\"")[2]
	} else {
		Print_err_cart("ADD TO CART FAILED [FAKE CART]")
	}
	return ""

}

//CHECK TO REVERSE SCRIPT TO GET THE TOKEN
//TLS 1.3

/*
STEP 1: Request to https://www.sugar.it/catalog/product/view/id/212183 TO get the entity ID for the CART
STEP 2: Request to https://www.sugar.it/checkout/cart/add/uenc/CART_ID/product/195475/
STEP 3: GET Request to https://www.sugar.it/onestepcheckout/ --> TAKE the entity_id in the html page
STEP 4: POST Request to https://www.sugar.it/rest/default/V1/guest-carts/entity_id/checkout-information with json data

check for size
keep the session
check for cookies
add the client
check concorrency request
add webkit genartor random 16 string
improve speed/ cpu handle
proxies
3DS handle checkout (if needed)
PP handle checkout (if needed)
ADD Mutex Monitor //cache control request & response https://lanre.wtf/blog/2017/07/24/roundtripper-go/
*/

/*
random string of lengh 16
chek if ttyjLrlqjsjaWhv2 is necessary
42 is the size fo the shoes
add client
check availability size and took a random one
add backgroud cli blue andchange color like the images
SSL Certificate Pinning
*/
