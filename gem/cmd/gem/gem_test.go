package main

import (
	"testing"

	"github.com/qur/gopy/lib"
)

// Launch py.test tests
func TestPython(t *testing.T) {
	_ = py.NewLock()

	run([]string{
		"-c", "import pytest; pytest.main(\"-s ../../../content/\")",
	}, true)
}
