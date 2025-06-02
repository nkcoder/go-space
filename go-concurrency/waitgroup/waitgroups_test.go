package main

import (
	"io"
	"os"
	"sync"
	"testing"
	"time"
)

func TestPrintSomething(t *testing.T) {
	// Capture stdout to verify output
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Setup waitgroup
	var wg sync.WaitGroup
	wg.Add(1)

	// Call the function
	testString := "test message"
	go printSomething(testString, &wg)

	wg.Wait()

	// Restore stdout
	w.Close()
	os.Stdout = oldStdout

	// Read captured output
	output, _ := io.ReadAll(r)

	// Verify the output contains our test string
	if string(output) != testString+"\n" {
		t.Errorf("Expected output to be %q, got %q", testString+"\n", string(output))
	}

	// The WaitGroup counter should be 0 now because printSomething calls Done()
	// We can verify this by calling Add(1) and then Wait() - it should not block
	wg.Add(1)
	waitChannel := make(chan struct{})

	go func() {
		wg.Wait()
		close(waitChannel)
	}()

	// Call Done to decrement the counter to 0
	wg.Done()

	// Wait for a short time to see if the goroutine completes
	select {
	case <-waitChannel:
		// Success - waitgroup reached 0
	case <-time.After(10 * time.Millisecond):
		t.Error("WaitGroup did not complete in time, Done() may not have been called")
	}
}
