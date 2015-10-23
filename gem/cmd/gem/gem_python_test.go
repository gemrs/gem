// +build test_python

package main

import (
	"testing"
)

// Launch py.test tests
func TestPython(t *testing.T) {
	run([]string{
		"-c", "import pytest; pytest.main(\"-s ../../../../content/\")",
	}, true)
}
