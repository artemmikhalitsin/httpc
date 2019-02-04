package main

import (
	"flag"
	"os"

	"github.com/artemmikhalitsin/httpc/headers"
)

// CommonOptions contains the options which are
// the same for each command
type CommonOptions struct {
	Verbose    bool
	Headers    headers.List
	InlineData string
	InputFile  string
	OutputFile string
}

func main() {
	// Common flags
	var common CommonOptions
	flag.BoolVar(&common.Verbose, "v", false, "Verbose output")
	flag.Var(&common.Headers, "h", "Request header in key:value format")
	flag.StringVar(&common.InlineData, "d", "", "Inline data to be added as "+
		" body of a POST request")
	flag.StringVar(&common.InputFile, "f", "", "Text file, the context of which"+
		" should be added as body of a POST request")
	flag.StringVar(&common.OutputFile, "o", "", "Filename to which the output is written to")

	// Method command
	helpCommand := flag.NewFlagSet("help", flag.ExitOnError)
	getCommand := flag.NewFlagSet("get", flag.ExitOnError)
	postCommand := flag.NewFlagSet("post", flag.ExitOnError)

	if len(os.Args) < 2 {
		PrintUsage()
		os.Exit(1)
	}
	// Parse common flags
	flag.CommandLine.Parse(os.Args[2:])

	// Handle commands
	switch os.Args[1] {
	case "get":
		getCommand.Parse(flag.Args())
	case "post":
		postCommand.Parse(flag.Args())
	case "help":
		helpCommand.Parse(flag.Args())
	default:
		flag.PrintDefaults()
	}

	if getCommand.Parsed() {
		GetCommand(common, getCommand.Args())
	}

	if postCommand.Parsed() {
		PostCommand(common, postCommand.Args())
	}

	if helpCommand.Parsed() {
		HelpCommand(helpCommand.Args())
	}
}
