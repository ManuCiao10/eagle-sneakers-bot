package client

import (
	"C"
	"errors"
	"fmt"
	"io"
	"net/url"

	http "github.com/saucesteals/fhttp"
	"github.com/saucesteals/mimic"
)

var (
	ErrNoCertificates = errors.New("no certificates in client")
	latestVersion     = mimic.MustGetLatestVersion(mimic.PlatformWindows)
	m, _              = mimic.Chromium(mimic.BrandChrome, latestVersion)
	proxy             = ""
)

// NewClient Takes in the optional arguments: proxy, servername
func NewClient(parameters ...string) (*Client, error) {
	if len(parameters) > 0 && len(parameters[0]) > 0 {
		proxy = parameters[0]
	}

	newClient, err := createClient(proxy)
	if err != nil {
		createCResponse(&Response{Error: err.Error()})
	}

	return &Client{
		client:         newClient,
		LatestResponse: &Response{},
	}, nil
}

// NewRequest creates a new request under a specified http client
func (c *Client) NewRequest() *Request {
	return &Request{
		client: c,
		header: make(http.Header),
	}
}

// Do will send the specified request
func (c *Client) Do(r *http.Request) (*Response, error) {
	resp, err := c.client.Do(r)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	_ = resp.Body.Close()
	if err != nil {
		return nil, err
	}

	// https://help.socketlabs.com/docs/how-to-fix-error-only-one-usage-of-each-socket-address-protocolnetwork-addressport-is-normally-permitted
	// https://www.geeksforgeeks.org/http-headers-connection/#:~:text=close%20This%20close%20connection%20directive,want%20your%20connection%20to%20close.
	r.Close = true // perhaps set this to false?

	response := &Response{
		headers:    resp.Header,
		body:       body,
		status:     resp.Status,
		statusCode: resp.StatusCode,
		cookies:    resp.Cookies(),
	}

	c.LatestResponse = response
	if Debug {
		fmt.Printf("%s %s\n", r.Method, r.URL.String())
		fmt.Printf("Response Body: %s\n", response.BodyAsString())
	}

	return response, nil
}

// AddCookie adds a new cookie to the request client cookie jar
func (c *Client) AddCookie(u *url.URL, cookie []*http.Cookie) *Request {
	c.client.Jar.SetCookies(u, cookie)
	return &Request{
		client: c,
		header: make(http.Header),
	}
}

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

func createClient(proxy string) (*http.Client, error) {
	transport, err := createTransport(proxy)
	if err != nil {
		return nil, err
	}

	return &http.Client{
		Transport: m.ConfigureTransport(transport),
	}, nil
}
