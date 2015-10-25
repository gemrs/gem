package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/qur/gopy/lib"

	_ "github.com/sinusoids/gem/gem/python/api"
	"github.com/sinusoids/gem/gem/util/safe"
)

func PythonTest(testDir string, t *testing.T) {
	// Disable panic recovery
	safe.Unsafe = true

	gopath := os.Getenv("GOPATH")
	pythonDir := fmt.Sprintf("%s/src/github.com/sinusoids/gem/content", gopath) // There's probably a better way to locate this

	pyLaunchTest := fmt.Sprintf(`import pytest; test_result = pytest.main("-s %v/%v")`, pythonDir, testDir)

	if main, err := py.NewDict(); err != nil {
		t.Fatal(err)
	} else if g, err := py.GetBuiltins(); err != nil {
		t.Fatal(err)
	} else if err := main.SetItemString("__builtins__", g); err != nil {
		t.Fatal(err)
	} else if _, err := py.RunString(pyLaunchTest, py.FileInput, main, nil); err != nil {
		t.Fatal(err)
	} else if a, err := main.GetItemString("test_result"); err != nil {
		t.Fatal(err)
	} else if b, ok := a.(*py.Int); !ok || b.Int() != 0 {
		t.Errorf("pytest returned non-zero exit code: %v", b.Int())
	}
}
