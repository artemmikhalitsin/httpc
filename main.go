package main

import (
	"flag"
	"fmt"
	"os"
)

type headerFlags []string

// Implement the flag.Value interface
func (i *headerFlags) String() string {
	return ""
}

func (i *headerFlags) Set(value string) error {
	return nil
}

func main() {
	method := os.Args[1]
	verbosePtr := flag.Bool("v", false, "Verbose output")

	flag.CommandLine.Parse(os.Args[2:])
	fmt.Println("method:", method)
	fmt.Println("verbose:", *verbosePtr)
	fmt.Println("tail:", flag.Args())
}
