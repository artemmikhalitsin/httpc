package client

import (
	"regexp"
	"strconv"
)

type Response struct {
	Status  int
	Headers string
	Body    string
	Raw     string
}

func NewResponse(raw string) *Response {
	re, _ := regexp.Compile(`(?ms)HTTP/\d\.\d (\d+) \w+\r\n(.*)\r\n\r\n(.*)`)
	details := re.FindStringSubmatch(raw)
	status, _ := strconv.Atoi(details[1])
	return &Response{
		Status:  status,
		Headers: details[2],
		Body:    details[3],
		Raw:     raw,
	}
}
