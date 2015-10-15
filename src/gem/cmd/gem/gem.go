package main

import (
	"os"

	"github.com/qur/gopy/lib"

	_ "gem"
	_ "gem/archive"
	_ "gem/auth"
	_ "gem/event"
	_ "gem/game"
	_ "gem/game/player"
	_ "gem/game/server"
	"gem/python"
	_ "gem/runite"
	_ "gem/task"
)

func main() {
	python.LinkModules()
	py.Main(os.Args)
	py.Finalize()
}
