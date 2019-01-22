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
		fmt.Println("httpc is a curl-like application but supports HTTP protocol only")
		PrintUsage()
		fmt.Println("The commands are:")
		fmt.Printf("   %-8s %s\n", "get", "executes a HTTP GET request and prints the response.")
		fmt.Printf("   %-8s %s\n", "post", "executes a HTTP POST request and prints the response.")
		fmt.Printf("   %-8s %s\n", "help", "prints this screen.")
		fmt.Println("Use \"httpc help [command]\" for more information about a command.")
		os.Exit(1)
	}
	command := args[0]

	switch command {
	case "get":
		printGetHelp()
	default:
		fmt.Printf("Unknown command: %s", command)
	}
}

// Prints the help text for a get command
func printGetHelp() {
	fmt.Println("usage: httpc get [-v] [-h key:value] URL\n")
	fmt.Println("Get executes a HTTP GET request for a given URL.\n")

	fmt.Printf("   %-14s %s\n", "-v", "Prints the detail of the response "+
		"such as protocol, status, and headers.")
	fmt.Printf("   %-14s %s\n", "-h key:value", "Associates headers to HTTP Request "+
		"with the format 'key:value'.")
}
