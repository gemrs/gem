package safe

import (
	"runtime"

	"gem/log"
)

// recover captures panics in the game client handler and prints a stack trace
func Recover(log *log.Module) {
	if err := recover(); err != nil {
		stack := make([]byte, 1024*10)
		runtime.Stack(stack, true)
		log.Criticalf("Recovered from panic in game client handler: %v", err)
		log.Debug(string(stack))
	}
}
