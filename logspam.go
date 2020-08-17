// TODO: check out go pipelines (returning channels)
package main

import (
	"fmt"
	"time"
	"math"
)

func main() {
	in := make(chan struct{})
	stop := make(chan struct{})
	out := make(chan int)
	go timer(stop)
	go tally(in, stop, out)
	go calc(out)
	listen(in)
}

func listen(in chan struct{}) {
	// TODO: read from stdin or interface
	for {
		in <-struct{}{}
		time.Sleep(1 * time.Second / 10)
	}
}

func timer(stop chan struct{}) {
	for {
		<-time.After(5 * time.Second)
		stop <-struct{}{}
	}
}

func tally(in, stop chan struct{}, out chan int) {
	var tally int
	for {
		select {
		case <-in:
			tally++
			//fmt.Println("log") // debug
		case <-stop:
			out <-tally
			tally = 0
		}
	}
}

// output speed every 5s
func calc(out chan int) {
	for tally := range out {
		if tally == 0 {
			fmt.Println("no logs in the last sample")
			continue
		}
		speed := math.Round(float64(tally) / float64(5) * 100) / 100
		fmt.Printf("last sample: %.1f lines/s\n", speed)
	}
}
