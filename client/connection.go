package client

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
)

type Connection struct {
	Host       string
	Connection net.TCPConn
}

func NewConnection(host string) *Connection {
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

	return &Connection{
		Host:       host,
		Connection: *conn,
	}
}

// Close closes the open TCP connection
func (c *Connection) Close() error {
	err := c.Connection.Close()
	if err != nil {
		return err
	}
	return nil
}

// WriteRequest writes a request to an open TCP connection
func (c *Connection) WriteRequest(r *Request) (string, error) {
	// Write the method, path, protocol, host
	request := fmt.Sprintf("%s %s HTTP/1.0\r\n"+
		"Host: %s\r\n", r.Method, r.Path, r.Host)
	// Write headers
	for _, header := range r.Headers {
		request += header.ToString() + "\r\n"
	}
	// Write body if exists
	if r.Body != nil {
		var bodyBuffer bytes.Buffer
		io.Copy(&bodyBuffer, r.Body)

		request += fmt.Sprintf("Content-Length: %d\r\n\r\n", bodyBuffer.Len())
		request += fmt.Sprintf(bodyBuffer.String())
	}
	// Finish the request with carriage return
	request += "\r\n"

	_, err := fmt.Fprint(&c.Connection, request)

	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	// Copy response into the buffer
	_, err = io.Copy(&buf, &c.Connection)

	if err != nil {
		return "", err
	}

	// Return the response
	return buf.String(), nil
}
