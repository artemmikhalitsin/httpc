package client

import (
	"io"

	"github.com/artemmikhalitsin/httpc/headers"
)

// Request represents an HTTP request
type Request struct {
	Method  string
	Host    string
	Path    string
	Headers headers.List
	Body    io.Reader
}

// NewRequest creates a new request from parameters
func NewRequest(method, host, path string, headers headers.List, body io.Reader) *Request {
	return &Request{
		Method:  method,
		Host:    host,
		Path:    path,
		Headers: headers,
		Body:    body,
	}
}
