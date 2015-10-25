package log_test

import (
	"testing"

	"github.com/qur/gopy/lib"

	"github.com/sinusoids/gem/gem/python/test"
)

func TestEventAPI(t *testing.T) {
	_ = py.NewLock()

	test.PythonTest("test/api/gem/log/", t)
}
