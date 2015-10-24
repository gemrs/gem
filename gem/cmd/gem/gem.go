package main

import (
	"os"

	"github.com/qur/gopy/lib"

	_ "github.com/sinusoids/gem/gem"
	_ "github.com/sinusoids/gem/gem/archive"
	_ "github.com/sinusoids/gem/gem/auth"
	_ "github.com/sinusoids/gem/gem/engine"
	_ "github.com/sinusoids/gem/gem/engine/event"
	_ "github.com/sinusoids/gem/gem/event"
	_ "github.com/sinusoids/gem/gem/game"
	_ "github.com/sinusoids/gem/gem/game/event"
	_ "github.com/sinusoids/gem/gem/game/player"
	_ "github.com/sinusoids/gem/gem/game/position"
	_ "github.com/sinusoids/gem/gem/game/server"
	"github.com/sinusoids/gem/gem/python"
	_ "github.com/sinusoids/gem/gem/runite"
	_ "github.com/sinusoids/gem/gem/task"
	"github.com/sinusoids/gem/gem/util/safe"
)

func main() {
	py.NewLock()
	run(os.Args[1:], false)
	py.Finalize()
}

// This is split into its own function to allow test to invoke the python interpreter
func run(args []string, testing bool) {
	// Disable panic recovery?
	safe.Unsafe = testing

	python.LinkModules()
	args = append([]string{"gem"}, args...)
	py.Main(args)
}
