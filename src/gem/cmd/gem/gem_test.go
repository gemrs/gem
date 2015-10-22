package main

import (
	"os/exec"
	"testing"
)

// Launch py.test tests
func TestPython(t *testing.T) {
	// Find the py.test binary
	pyTest, err := exec.LookPath("py.test")
	if err != nil {
		t.Fatal("Couldn't locate py.test")
	}
	run([]string{
		pyTest, "../../../../content/",
	})
}
