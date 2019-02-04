package main

import (
	"fmt"
	"os"
)

// PrintUsage prints the intended format
// in which the program should be ran
func PrintUsage() {
	fmt.Println("Usage: httpc (get|post) [-v] (-h \"key:value\")* " +
		"[-d inline-data] [-f file] URL")
}

// HelpCommand handles carrying out
// a help command of the cli
func HelpCommand(args []string) {
	if len(args) < 1 {
		printProgramHelp()
		os.Exit(0)
	}
	command := args[0]

	switch command {
	case "get":
		printGetHelp()
	case "post":
		printPostHelp()
	default:
		fmt.Printf("Unknown command: %s", command)
		os.Exit(1)
	}
	os.Exit(0)
}

// Prints the general use of the program
// and a list of commands
func printProgramHelp() {
	fmt.Println("httpc is a curl-like application but supports HTTP protocol only")
	fmt.Println("Usage:\n" +
		"httpc command [arguments]")
	fmt.Println("The commands are:")
	fmt.Printf("   %-8s %s\n", "get", "executes a HTTP GET request and prints the response.")
	fmt.Printf("   %-8s %s\n", "post", "executes a HTTP POST request and prints the response.")
	fmt.Printf("   %-8s %s\n", "help", "prints this screen.")
	fmt.Println()
	fmt.Println("Use \"httpc help [command]\" for more information about a command.")
}

// Prints the help text for a get command
func printGetHelp() {
	fmt.Println("usage: httpc get [-v] [-h key:value] URL")
	fmt.Println()
	fmt.Println("Get executes a HTTP GET request for a given URL.")
	fmt.Println()
	fmt.Printf("   %-14s %s\n", "-v", "Prints the detail of the response "+
		"such as protocol, status, and headers.")
	fmt.Printf("   %-14s %s\n", "-h key:value", "Associates headers to HTTP Request "+
		"with the format 'key:value'.")
}

func printPostHelp() {
	fmt.Println("usage: httpc post [-v] [-h key:value] [-d inline-data] [-f file] URL")
	fmt.Println()
	fmt.Println("Post executes a HTTP POST request for a given URL with inline " +
		"data or from file.")
	fmt.Println()
	fmt.Printf("   %-14s %s\n", "-v", "Prints the detail of the response such as "+
		"protocol, status, and headers.")
	fmt.Printf("   %-14s %s\n", "-h key:value", "Associates headers to HTTP Request "+
		"with the format 'key:value'.")
	fmt.Printf("   %-14s %s\n", "-d string", "Associates an inline data to the "+
		"body HTTP POST request.")
	fmt.Printf("   %-14s %s\n", "-f string", "Associates the content of a file "+
		" to the body HTTP POST request.")
	fmt.Println()
	fmt.Println("Either [-d] or [-f] can be used but not both.")
}
