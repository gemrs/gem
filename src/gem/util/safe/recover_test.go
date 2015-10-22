package safe

import (
	"fmt"
	"testing"

	"gem/log"
)

func TestRecover(t *testing.T) {
	logger := log.NewMock("recover_test")
	panicMsg := "testing safe.Recover"
	func() {
		defer Recover(logger)
		panic(panicMsg)
	}()

	searchStr := fmt.Sprintf("Recovered from panic in game client handler: %v", panicMsg)
	if !logger.HasLogged(searchStr) {
		t.Error("Couldn't find recovery message")
	}

	if !logger.HasLogged("TestRecover") {
		t.Error("Couldn't find stack trace")
	}
}