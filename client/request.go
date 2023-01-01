package client

import "strings"

func (r *Request) SetBody(body string) *Request {
	r.body = strings.NewReader(body)
	return r
}
