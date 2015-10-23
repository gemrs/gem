package main

import (
	"os"

	"github.com/qur/gopy/lib"

	_ "gem"
	_ "gem/archive"
	_ "gem/auth"
	_ "gem/engine"
	_ "gem/engine/event"
	_ "gem/event"
	_ "gem/game"
	_ "gem/game/event"
	_ "gem/game/player"
	_ "gem/game/position"
	_ "gem/game/server"
	"gem/python"
	_ "gem/runite"
	_ "gem/task"
	"gem/util/safe"
)

func main() {
	run(os.Args[1:], false)
}

// This is split into its own function to allow test to invoke the python interpreter
func run(args []string, testing bool) {
	// Disable panic recovery?
	safe.Unsafe = testing

	python.LinkModules()
	args = append([]string{"gem"}, args...)
	py.Main(args)
	py.Finalize()
}
