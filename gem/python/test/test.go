package test

// #include <python2.7/Python.h>
// #cgo LDFLAGS: -lpython2.7
import "C"

import (
	"fmt"
	"os"
	"testing"
	"unsafe"

	"github.com/qur/gopy/lib"

	_ "github.com/gemrs/gem/gem/python/api"
	"github.com/gemrs/gem/gem/util/safe"
)

func PythonTest(testDir string, t *testing.T) {
	// Disable panic recovery
	safe.Unsafe = true

	gopath := os.Getenv("GOPATH")
	pythonDir := fmt.Sprintf("%s/src/github.com/gemrs/gem/content", gopath) // There's probably a better way to locate this

	pyLaunchTest := fmt.Sprintf(`import pytest; test_result = pytest.main("-s %v/%v")`, pythonDir, testDir)

	argv := make([]*C.char, len(os.Args))

	for i, arg := range os.Args {
		argv[i] = C.CString(arg)
		defer C.free(unsafe.Pointer(argv[i]))
	}

	C.PySys_SetArgv(C.int(len(argv)), &argv[0])
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
		t.Error("pytest returned non-zero exit code: %v", b.Int())
	}
}
