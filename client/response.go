package client

import (
	"encoding/json"
	"log"

	http "github.com/saucesteals/fhttp"
)

func createCResponse(resp *Response) {
	errorJson, _ := json.Marshal(resp)
	log.Println(string(errorJson))

}

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

// BodyAsJSON unmarshalls the current response body to the specified data structure
func (r *Response) BodyAsJSON(data interface{}) error {
	return json.Unmarshal(r.body, data)
}

// get response cookies
func (r *Response) Cookies() []*http.Cookie {
	return r.cookies
}

// get response cookies as string
func (r *Response) CookiesAsString() string {
	var cookies string
	for _, cookie := range r.cookies {
		cookies += cookie.String() + ";"
	}
	return cookies
}
