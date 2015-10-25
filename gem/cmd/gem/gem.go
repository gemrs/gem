package main

import (
	"os"

	"github.com/qur/gopy/lib"

	_ "github.com/sinusoids/gem/gem/python/api"
)

func main() {
	py.NewLock()
	py.Main(os.Args)
	py.Finalize()
}
