package wg

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_gr(t *testing.T) {
	stdOut := os.Stdout

	// Get reader and writer
	r, w, _ := os.Pipe()
	os.Stdout = w

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go gr("A", wg)

	wg.Wait()

	_ = w.Close()

	result, err := io.ReadAll(r)
	if err != nil {
		panic("Panic!")
	}
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "A") {
		t.Error("Output does not contain A")
	}
}
