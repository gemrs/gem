package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/qur/gopy/lib"

	"github.com/sinusoids/gem/gem"
)

func main() {
	pythonInit()

	py.Main(os.Args)
}

type registerFunc func(*py.Module) error

var moduleRegisterFuncs = []registerFunc{
	gem.RegisterSysLog,
	gem.RegisterEngine,
	gem.RegisterLogModule,
}

func pythonInit() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("initializing python: %s", r)
			os.Exit(1)
		}
	}()

	py.Initialize()

	var err error
	var module *py.Module
	if module, err = py.InitModule("gem", []py.Method{}); err != nil {
		panic(err)
	}

	/* Register modules */
	for _, registerFunc := range moduleRegisterFuncs {
		if err = registerFunc(module); err != nil {
			panic(err)
		}
	}

	/* Create our globals */
	if globals, err := py.NewDict(); err != nil {
		panic(err)
	} else if builtins, err := py.GetBuiltins(); err != nil {
		panic(err)
	} else if err = globals.SetItemString("__builtins__", builtins); err != nil {
		panic(err)
	}

	/* Create our logger object */
	if syslog, err := gem.InitSysLog(); err != nil {
		panic(err)
	} else if err = module.AddObject("syslog", syslog); err != nil {
		panic(err)
	}

	/* Make sure we catch SIGTERM and clean up python gracefully */
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		<-c
		py.SetInterrupt()
		os.Exit(1)
	}()
}
