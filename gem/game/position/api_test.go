package position_test

import (
	"testing"

	"github.com/qur/gopy/lib"

	"github.com/sinusoids/gem/gem/python/test"
)

func TestLogAPI(t *testing.T) {
	_ = py.NewLock()

	test.PythonTest("test/api/gem/game/position/", t)
}
