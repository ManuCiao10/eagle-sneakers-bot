package client

import (
	"errors"
	"fmt"
	"io"

	http "github.com/saucesteals/fhttp"
	"github.com/saucesteals/mimic"
)

var (
	ErrNoCertificates = errors.New("no certificates in client")
	latestVersion     = mimic.MustGetLatestVersion(mimic.PlatformWindows)
	m, _              = mimic.Chromium(mimic.BrandChrome, latestVersion)
	proxy             = ""
)

func createClient(proxy string) (*http.Client, error) {
	transport, err := createTransport(proxy)
	if err != nil {
		return nil, err
	}

	return &http.Client{
		Transport: m.ConfigureTransport(transport),
	}, nil
}

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

// SetURL sets the url of the request
func (r *Request) SetURL(url string) *Request {
	r.url = url
	return r
}

// SetMethod sets the method of the request
func (r *Request) SetMethod(method string) *Request {
	r.method = method
	return r
}

// AddHeader adds a specified header to the request
// If the header already exists, the value will be appended by the new specified value
// If the header does not exist, the header will be set to the specified value
func (r *Request) AddHeader(key, value string) *Request {
	if header, ok := r.header[key]; ok {
		header = append(header, value)
		r.header[key] = header
	} else {
		r.header[key] = []string{value}
	}
	return r
}

// SetHeader sets a specified header to the request
// This overrides any previously set values of the specified header
func (r *Request) SetHeader(key, value string) *Request {
	r.header[key] = []string{value}
	return r
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
