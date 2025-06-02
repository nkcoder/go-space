package main

import (
	"sync"
	"testing"
)

func TestUpdateMessage(t *testing.T) {

	// Setup waitgroup
	var wg sync.WaitGroup
	wg.Add(1)

	// Call the function
	testString := "test message"
	go updateMessage(testString, &wg)

	wg.Wait()

	// Verify the output contains our test string
	if string(msg) != testString {
		t.Errorf("Expected output to be %q, got %q", testString+"\n", string(msg))
	}
}
