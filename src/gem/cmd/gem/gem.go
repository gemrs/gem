package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/qur/gopy/lib"

	"gem"
)

func main() {
	_ = pythonInit()
	py.Main(os.Args)
	py.Finalize()
}

func pythonInit() *py.Lock {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("initializing python: %s", r)
			os.Exit(1)
		}
	}()

	lock := py.InitAndLock()

	if err := gem.InitPyModule(); err != nil {
		panic(err)
	}

	/* Create our globals */
	if globals, err := py.NewDict(); err != nil {
		panic(err)
	} else if builtins, err := py.GetBuiltins(); err != nil {
		panic(err)
	} else if err = globals.SetItemString("__builtins__", builtins); err != nil {
		panic(err)
	}

	/* Make sure we catch SIGTERM and clean up python gracefully */
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		<-c

		lock := py.NewLock()
		py.Finalize()
		lock.Unlock()

		os.Exit(0)
	}()

	return lock
}
