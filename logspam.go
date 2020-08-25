package logspam

import (
	"bufio"
	"io"
	"log"
	"math"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Start configures logging, the input-, and output source,
// combines channels and goroutines and starts listening for input
func Start(r io.Reader, w io.Writer, sampleRate int, verbose bool) {
	log.SetOutput(w)
	log.SetFlags(0)
	if verbose {
		log.Printf("logspam started. SampleRate %ds\n", sampleRate)
	}

	in := make(chan []byte)
	stop := make(chan struct{})
	out := make(chan int)
	errc := make(chan error)

	go tally(in, stop, out)
	go calc(out, sampleRate)
	go timer(stop, sampleRate)
	go listen(in, r, errc)

	select {
	case err := <-errc:
		if err != nil {
			log.Printf("Input reading err: %s\n", err)
		}
	case <-interrupt():
		// noop
	}
	if verbose {
		log.Println("logspam exited")
	}
}

// interrupt returns a chan that recieves interrupt signals
func interrupt() <-chan os.Signal {
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	return sigc
}

// listen sends input to the in chan and when done, any error encountered
// reading, to the err chan
func listen(in chan []byte, r io.Reader, errc chan error) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		in <- scanner.Bytes()
	}
	errc <- scanner.Err()
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
