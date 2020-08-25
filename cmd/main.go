package main

import (
	"os"
	"log"

	"logspam"
)

func main() {
	sampleRate := 5
	err := logspam.Start(sampleRate, os.Stdin, os.Stderr)
	if err != nil {
		log.Printf("Input reading err: %s\n", err)
		return
	}
	log.Println("logspam exited")
}
