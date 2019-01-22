package main

import (
	"fmt"
	"os"
)

// PrintUsage prints the intended format
// in which the program should be ran
func PrintUsage() {
	fmt.Println("Usage:")
	fmt.Println("httpc (get|post) [-v] (-h \"key:value\")* [-d inline-data] [-f file] URL")
}

// HelpCommand handles carrying out
// a help command of the cli
func HelpCommand(args []string) {
	if len(args) < 1 {
		fmt.Println("Please specify a command to see help")
		os.Exit(1)
	}

	command := args[0]

	switch command {
	default:
		fmt.Printf("Unknown command: %s", command)
	}
}

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
	fmt.Println("Command: get")
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
	// == equivalent to xor operation
	if (opt.InlineData != "") == (opt.InputFile != "") {
		fmt.Println("Please specify either inline data or a file, but not both")
		os.Exit(1)
	}
	fmt.Println("Command: post")
}
