package main

import (
	"strings"
	"testing"

	"github.com/karlpokus/bufw"
)

func TestLogspam(t *testing.T) {
	r := strings.NewReader(`line 1
		line 2
		line 3`)
	w := bufw.New()
	go func() {
		sampleRate := 1
		verbose := false
		start(r, w, sampleRate, verbose)
	}()
	err := w.Wait()
	if err != nil {
		t.Fatal(err)
	}
	output := w.String()
	expected := "last sample: 3.0 lines/s"
	if output != expected {
		t.Fatalf("%s does not match %s", output, expected)
	}
}
