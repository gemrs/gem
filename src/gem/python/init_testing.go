// For some reason, go test doesn't like InitAndLock(),
// I guess it's doing something funky.
// For this reason, an alternate init() is provided which calls just Initialize()

// +build test_python

package python

import (
	"fmt"
	"os"

	"github.com/qur/gopy/lib"
)

func pythonInit() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("initializing python: %s", r)
			os.Exit(1)
		}
	}()

	py.Initialize()

	SetTypeConvFuncs()

	/* Create our globals */
	if globals, err := py.NewDict(); err != nil {
		panic(err)
	} else if builtins, err := py.GetBuiltins(); err != nil {
		panic(err)
	} else if err = globals.SetItemString("__builtins__", builtins); err != nil {
		panic(err)
	}

}

func pythonExit() {
	lock := py.NewLock()
	py.Finalize()
	lock.Unlock()
}

func init() {
	pythonInit()
}
