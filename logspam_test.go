package logspam

import (
	"strings"
	"testing"

	"github.com/karlpokus/bufw"
)

func TestLogspam(t *testing.T) {
	r := strings.NewReader(`line 1
		line 2
		line 3`)
	w := bufw.New(true)
	errc := make(chan error)
	go func() {
		sampleRate := 1
		errc <- Start(sampleRate, r, w)
	}()
	err := w.Wait()
	if err != nil {
		t.Fatal(err)
	}
	err = <-errc
	if err != nil {
		t.Fatal(err)
	}
	output := w.String()
	expected := "last sample: 3.0 lines/s"
	if output != expected {
		t.Fatalf("%s does not match %s", output, expected)
	}
}
