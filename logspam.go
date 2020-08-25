package logspam

import (
	"bufio"
	"io"
	"log"
	"math"
	"time"
)

// Start grabs the reader and writer, combines channels and goroutines
// and starts listening for input
func Start(sampleRate int, r io.Reader, w io.Writer) error {
	log.SetOutput(w)
	log.SetFlags(0)

	in := make(chan []byte)
	stop := make(chan struct{})
	out := make(chan int)

	go tally(in, stop, out)
	go calc(out, sampleRate)
	go timer(stop, sampleRate)

	return listen(in, r)
}

// listen sends input to the in chan
func listen(in chan []byte, r io.Reader) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		in <- scanner.Bytes()
	}
	return scanner.Err()
}

// timer notifyes on the stop chan every sampleRate seconds
func timer(stop chan struct{}, sampleRate int) {
	for {
		<-time.After(time.Duration(sampleRate) * time.Second)
		stop <- struct{}{}
	}
}

// tally selects between incrementing a tally and sending that tally on
// the out chan
func tally(in chan []byte, stop chan struct{}, out chan int) {
	var tally int
	for {
		select {
		case <-in:
			tally++
		case <-stop:
			out <- tally
			tally = 0
		}
	}
}

// calc outputs speed every sampleRate seconds
func calc(out chan int, sampleRate int) {
	for tally := range out {
		if tally == 0 {
			log.Println("no logs in the last sample")
			continue
		}
		speed := math.Round(float64(tally)/float64(sampleRate)*100) / 100
		log.Printf("last sample: %.1f lines/s\n", speed)
	}
}
