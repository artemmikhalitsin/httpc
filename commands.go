package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/artemmikhalitsin/httpc/client"
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

	fullRes, _ := client.Get(requestURL, opt.Headers)
	// Split into response details and content
	res := strings.Split(fullRes, "\r\n\r\n")
	content := res[1]

	if opt.Verbose {
		// Verbose: Print details and content
		fmt.Println(fullRes)
	} else {
		// Non-verbose : Print only content
		fmt.Println(content)
	}
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

	fullRes, _ := client.Post(requestURL, opt.Headers, body)
	// Split into response details and content
	res := strings.Split(fullRes, "\r\n\r\n")
	content := res[1]

	if opt.Verbose {
		// Verbose: Print details and content
		fmt.Println(fullRes)
	} else {
		// Non-verbose: Print just content
		fmt.Println(content)
	}
}
