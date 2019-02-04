package client

import (
	"fmt"
	"io"
	"net/url"
	"strings"

	"github.com/artemmikhalitsin/httpc/headers"
)

// DoRequest performs a general request, where you can specify the method
func DoRequest(method, uri string, headers headers.List, body io.Reader) (*Response, error) {
	// Parse URL
	parsed, err := url.ParseRequestURI(uri)

	if err != nil {
		fmt.Println("Unable to parse request URL. " +
			"Did you forget the protocol? (http://)")
	}

	if parsed.Scheme != "http" {
		fmt.Println("httpc only supports HTTP requests :(")
	}

	host := parsed.Host
	if !strings.Contains(host, ":") {
		// No port defined: use 80
		host += ":80"
	}
	path := parsed.Path
	if path == "" {
		// No path defined: use "/"
		path = "/"
	}
	if parsed.RawQuery != "" {
		// Attach query if exists
		path += "?" + parsed.RawQuery
	}

	// Establish connection
	conn := NewConnection(host)
	// Make sure the connection is closed in the end
	defer conn.Close()
	// Build a request from the parsed params
	req := NewRequest(method, host, path, headers, body)
	// Write the request to the connection
	res, err := conn.WriteRequest(req)
	var response *Response

	if err != nil {
		return response, err
	}

	response = NewResponse(res)

	return response, nil
}

// Get performs a GET request and returns the response
func Get(uri string, headers headers.List) (*Response, error) {
	return DoRequest("GET", uri, headers, nil)
}

// Post performs a POST request and returns the response
func Post(uri string, headers headers.List, body io.Reader) (*Response, error) {
	return DoRequest("POST", uri, headers, body)
}
