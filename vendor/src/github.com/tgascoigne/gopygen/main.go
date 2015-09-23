package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
	"regexp"

	"github.com/tgascoigne/gopygen/gopygen"
)

type arrayFlags []string

func (i *arrayFlags) String() string {
	return fmt.Sprintf("[%v]", strings.Join(*i, " "))
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

var types, includeFuncs, excludeFuncs arrayFlags

func filterFunction(name string) bool {
	if len(includeFuncs) > 0 {
		for _, rule := range includeFuncs {
			if match, err := regexp.MatchString(rule, name); err != nil {
				log.Fatalf("regexp rule invalid: %v", rule)
			} else if match {
				return true
			}
		}
	} else if len(excludeFuncs) > 0 {
		for _, rule := range excludeFuncs {
			if match, err := regexp.MatchString(rule, name); err != nil {
				log.Fatalf("regexp rule invalid: %v", rule)
			} else if match {
				return false
			}
		}
	}
	return true
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("gopygen: ")
	flag.Var(&types, "type", "specify a type to generate bindings for")
	flag.Var(&includeFuncs, "include", "specify regular expression matching function names to include")
	flag.Var(&excludeFuncs, "exclude", "specify regular expression matching function names to exclude")
	flag.Parse()

	// BUG(tom): For now, just accept a single file as an argument
	args := flag.Args()
	filename := args[0]

	if len(includeFuncs) > 0 && len(excludeFuncs) > 0 {
		log.Fatalf("include and exclude flags are mutually exclusive")
	}

	src, err := gopygen.Process(filename, types, filterFunction)
	if err != nil {
		log.Fatalf("processing input: %s", err)
	}

	basename := strings.TrimSuffix(filename, filepath.Ext(filename))
	outputName := fmt.Sprintf("%v.py.go", basename)

	err = ioutil.WriteFile(outputName, []byte(src), 0644)
	if err != nil {
		log.Fatalf("writing output: %s", err)
	}

	log.Printf("wrote output: %v", outputName)
}
