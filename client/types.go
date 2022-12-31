package client

import (
	"io"

	sessionjar "github.com/juju/persistent-cookiejar"
	http "github.com/saucesteals/fhttp"
)

type Client struct {
	client         *http.Client
	jar            *sessionjar.Jar
	LatestResponse *Response
}

type Request struct {
	client *Client

	method, url, host string

	header http.Header

	body io.Reader

	cookies []*http.Cookie
}

type Session struct {
	Client    *http.Client
	Headers   map[string]string
	Cookies   map[string]string
	Randomize bool
}

type Response struct {
	headers http.Header

	body []byte

	status     string
	statusCode int
	cookies    []*http.Cookie

	Error string
}
