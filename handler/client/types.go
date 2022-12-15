package client

import (
	"net/http"

	sessionjar "github.com/juju/persistent-cookiejar"
)

type Client struct {
	client         *http.Client
	jar            *sessionjar.Jar
	LatestResponse *Response
}

type Response struct {
	headers http.Header

	body []byte

	status     string
	statusCode int
	cookies    []*http.Cookie
}
