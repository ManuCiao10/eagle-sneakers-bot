package client

import (
	"C"
	"encoding/json"
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

// BodyAsJSON unmarshalls the current response body to the specified data structure
func (r *Response) BodyAsJSON(data interface{}) error {
	return json.Unmarshal(r.body, data)
}

// get response cookies
func (r *Response) Cookies() []*http.Cookie {
	return r.cookies
}

func createCResponse(resp *Response) *C.char {
	errorJson, _ := json.Marshal(resp)
	return C.CString(string(errorJson))
}
