package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_calculateIncome(t *testing.T) {
	stdOut := os.Stdout
	r, w, _ := os.Pipe()

	os.Stdout = w

	calculateIncome()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "Final bank balance: 83720") {
		t.Errorf("Expected 'Final bank balance: 83720' in output, but got %s", output)
	}
}
