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
	pythonInit()

	py.Main(os.Args)
}

func pythonInit() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("initializing python: %s", r)
			os.Exit(1)
		}
	}()

	py.Initialize()

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
		fmt.Printf("caught int")
		py.SetInterrupt()
		os.Exit(1)
	}()

}
