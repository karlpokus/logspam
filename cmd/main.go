package main

import (
	"os"
	"flag"

	"logspam"
)

var verbose = flag.Bool("v", false, "Toggle verbose output")

func main() {
	flag.Parse()
	sampleRate := 5
	logspam.Start(os.Stdin, os.Stderr, sampleRate, *verbose)
}
