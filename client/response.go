package client

import (
	"log"

	http "github.com/saucesteals/fhttp"
)

// Header returns the response headers
func (r *Response) Header() http.Header {
	return r.headers
}

// BodyAsString returns the response body as a string
func (r *Response) BodyAsString() string {
	return string(r.body)
}

// StatusCode returns the response status code
func (r *Response) StatusCode() int {
	return r.statusCode
}

// Body returns the response body
func (r *Response) Body() []byte {
	return r.body
}

func (c *Client) SaveCookies() {
	if c.client.Jar != nil {
		err := c.jar.Save()
		if err != nil {
			log.Println(err)
			return
		}
	}
}

// get response cookies
func (r *Response) Cookies() []*http.Cookie {
	
	return r.cookies
}

