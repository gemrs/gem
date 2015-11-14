package safe

import (
	"runtime"

	"github.com/gemrs/gem/gem/log"
)

// Set Unsafe to disable recovery (for testing)
var Unsafe = false

// recover captures panics in the game client handler and prints a stack trace
func Recover(log log.Log) {
	if err := recover(); err != nil {
		if Unsafe {
			log.Notice("github.com/gemrs/gem/gem/safe: Unsafe mode enabled; not recovering")
			panic(err)
		}
		stack := make([]byte, 1024*10)
		runtime.Stack(stack, true)
		log.Error("Recovered from panic in game client handler: %v", err)
		log.Debug(string(stack))
	}
}
