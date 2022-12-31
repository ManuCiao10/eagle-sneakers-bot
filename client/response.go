package client

// BodyAsString returns the response body as a string
func (r *Response) BodyAsString() string {
	return string(r.body)
}

// StatusCode returns the response status code
func (r *Response) StatusCode() int {
	return r.statusCode
}
