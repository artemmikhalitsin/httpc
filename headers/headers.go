package headers

import (
	"errors"
	"flag"
	"fmt"
	"strings"
)

// Header represents a single http header
type Header struct {
	Key   string
	Value string
}

// ToString returns a string representation of a Header
func (h *Header) ToString() string {
	return fmt.Sprintf("%s:%s", h.Key, h.Value)
}

// NewHeader creates a Header by parsing a string
// that is in the format "key:value"
func NewHeader(str string) (Header, error) {
	parsed := strings.Split(str, ":")
	if len(parsed) != 2 {
		return Header{}, errors.New("Headers must be in format key:value")
	}
	return Header{parsed[0], parsed[1]}, nil
}

// List represents a list of Headers
// Implements the flag.Value interface
type List []Header

// String returns a string representation of a List.
// Implemented to satisfy flag.Value interface
func (list *List) String() string {
	var headerStrings []string
	for _, element := range *list {
		header := element.ToString()
		headerStrings = append(headerStrings, header)
	}
	return strings.Join(headerStrings, ", ")
}

// Set creates a Header from a string and attaches it to the list
// Implemented to satisfy flag.Value interface
func (list *List) Set(value string) error {
	newHeader, err := NewHeader(value)
	if err != nil {
		fmt.Println(err.Error())
		flag.PrintDefaults()
	}
	*list = append(*list, newHeader)
	return nil
}
