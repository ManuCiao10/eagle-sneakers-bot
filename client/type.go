package client

import (
	"C"
	"encoding/json"
	"fmt"
	"net/url"

	http "github.com/saucesteals/fhttp"
)

var (
	Debug = false
)

func createTransport(proxy string) (*http.Transport, error) {
	if len(proxy) != 0 {
		proxyUrl, err := url.Parse(proxy)
		if err != nil {
			return nil, err
		}
		return &http.Transport{Proxy: http.ProxyURL(proxyUrl)}, nil
	} else {
		return &http.Transport{}, nil
	}
}

func createCResponse(resp *Response) *C.char {
	errorJson, _ := json.Marshal(resp)
	return C.CString(string(errorJson))
}

// Do will send the request with all specified request values
func (r *Request) Do() (*Response, error) {
	req, err := http.NewRequest(r.method, r.url, r.body)
	if err != nil {
		return nil, err
	}

	for _, cookie := range r.cookies {
		if cookie != nil {
			req.AddCookie(cookie)
		}
	}

	req.Header = r.header

	if len(r.host) > 0 {
		req.Host = r.host
	}

	if Debug {
		fmt.Println("Request Body:", r.body)
	}

	return r.client.Do(req)
}
