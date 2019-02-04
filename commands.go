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
	out := os.Stdout
	// If output file is specified, write there instead
	if opt.OutputFile != "" {
		file, err := os.Create(opt.OutputFile)
		if err != nil {
			fmt.Println("Error opening file for output")
			os.Exit(1)
		}
		defer file.Close()
		out = file
	}

	requestURL := args[0]

	response, _ := client.Get(requestURL, opt.Headers)
	result := response.Raw
	if !opt.Verbose {
		result = response.Body
	}

	fmt.Fprintf(out, result)
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

	out := os.Stdout
	// If output file is specified, write there instead
	if opt.OutputFile != "" {
		file, err := os.Create(opt.OutputFile)
		if err != nil {
			fmt.Println("Error opening file for output")
			os.Exit(1)
		}
		defer file.Close()
		out = file
	}

	var body io.Reader
	if opt.InlineData != "" {
		body = strings.NewReader(opt.InlineData)
	}
	if opt.InputFile != "" {
		f, err := os.Open(opt.InputFile)
		if err != nil {
			fmt.Println("Error opening file" + opt.InputFile)
			os.Exit(1)
		}
		defer f.Close()
		body = f
	}

	requestURL := args[0]

	response, _ := client.Post(requestURL, opt.Headers, body)
	result := response.Raw

	if !opt.Verbose {
		// Non-verbose: Print only the  content
		result = response.Body
	}

	fmt.Fprintf(out, result)
}
