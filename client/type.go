package client

import (
	"C"
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
