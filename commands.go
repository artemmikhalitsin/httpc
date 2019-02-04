package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"net/url"
	"os"
	"strings"
)

// GetCommand handles carrying out
// a get command of the cli
func GetCommand(opt CommonOptions, args []string) {
	// Guards
	// Must include a URL as first argument
	if len(args) < 1 {
		fmt.Println("Please specify the URL of the GET request")
		os.Exit(1)
	}
	// Must not have -d or -f specified
	if opt.InlineData != "" || opt.InputFile != "" {
		fmt.Println("GET requests cannot include a body")
		os.Exit(1)
	}
	requestURL := args[0]

	parsed, err := url.ParseRequestURI(requestURL)

	if err != nil {
		fmt.Println("Unable to parse request URL. " +
			"Did you forget the protocol? (http://)")
		os.Exit(1)
	}

	if parsed.Scheme != "http" {
		fmt.Print("httpc only supports HTTP requests")
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
	// Resolve the TCP address
	addr, err := net.ResolveTCPAddr("tcp", host)
	if err != nil {
		fmt.Println("Unable to resolve the given address")
	}

	// Open TCP connect
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		os.Exit(1)
	}

	if err != nil {
		fmt.Println("Error establishing connection to " + host)
	}

	request := fmt.Sprintf("GET %s HTTP/1.0\n"+
		"Host: %s\n\n", path, host)

	fmt.Fprint(conn, request)

	var buf bytes.Buffer
	// Copy response into the buffer
	io.Copy(&buf, conn)

	fmt.Println(buf.String())
}

// PostCommand handles carrying out
// a get command of the cli
func PostCommand(opt CommonOptions, args []string) {
	// Must include a URL
	if len(args) < 1 {
		fmt.Println("Please specify the URL of the POST request")
		os.Exit(1)
	}
	// Must provide either -d or -f but not both
	// == performs the xor operation
	if (opt.InlineData != "") == (opt.InputFile != "") {
		fmt.Println("Please specify either inline data or a file, but not both")
		os.Exit(1)
	}

	var body io.Reader
	if opt.InlineData != "" {
		body = strings.NewReader(opt.InlineData)
	}
	if opt.InputFile != "" {
		f, err := os.Open(opt.InputFile)
		if err != nil {
			fmt.Println("Error opening file" + opt.InputFile)
		}
		body = f
	}

	requestURL := args[0]

	parsed, err := url.ParseRequestURI(requestURL)

	if err != nil {
		fmt.Println("Unable to parse request URL. " +
			"Did you forget the protocol? (http://)")
		os.Exit(1)
	}

	if parsed.Scheme != "http" {
		fmt.Print("httpc only supports HTTP requests")
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
	// Resolve the TCP address
	addr, err := net.ResolveTCPAddr("tcp", host)
	if err != nil {
		fmt.Println("Unable to resolve the given address")
	}

	// Open TCP connect
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		os.Exit(1)
	}

	if err != nil {
		fmt.Println("Error establishing connection to " + host)
	}

	request := fmt.Sprintf("POST %s HTTP/1.0\r\n"+
		"Host: %s\r\n", path, host)

	var bodyBuffer bytes.Buffer
	io.Copy(&bodyBuffer, body)

	request += fmt.Sprintf("Content-Length: %d\r\n\r\n", bodyBuffer.Len())
	request += fmt.Sprintf(bodyBuffer.String() + "\r\n")

	fmt.Fprint(conn, request)

	var buf bytes.Buffer
	// Copy response into the buffer
	io.Copy(&buf, conn)

	fmt.Println(buf.String())
}
