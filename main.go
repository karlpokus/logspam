package main

import (
	"os"
	"flag"
)

var verbose = flag.Bool("v", false, "Toggle verbose output")

func main() {
	flag.Parse()
	sampleRate := 5
	start(os.Stdin, os.Stderr, sampleRate, *verbose)
}
