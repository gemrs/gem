package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"regexp"
	"strings"

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

var (
	types         arrayFlags
	includeFuncs  arrayFlags
	excludeFuncs  arrayFlags
	includeFields arrayFlags
	excludeFields arrayFlags
)

func filterFunc(inc, exc arrayFlags) gopygen.FilterFunc {
	return func(name string) bool {
		if len(inc) > 0 {
			for _, rule := range inc {
				if match, err := regexp.MatchString(rule, name); err != nil {
					log.Fatalf("regexp rule invalid: %v", rule)
				} else if match {
					return true
				}
			}
		} else if len(exc) > 0 {
			for _, rule := range exc {
				if match, err := regexp.MatchString(rule, name); err != nil {
					log.Fatalf("regexp rule invalid: %v", rule)
				} else if match {
					return false
				}
			}
		}
		return true
	}
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("gopygen: ")
	flag.Var(&types, "type", "specify a type to generate bindings for")
	flag.Var(&includeFuncs, "incfunc", "specify regular expression matching function names to include")
	flag.Var(&excludeFuncs, "excfunc", "specify regular expression matching function names to exclude")
	flag.Var(&includeFields, "incfield", "specify regular expression matching field names to include")
	flag.Var(&excludeFields, "excfield", "specify regular expression matching field names to exclude")
	flag.Parse()

	// BUG(tom): For now, just accept a single file as an argument
	args := flag.Args()
	filename := args[0]

	if len(includeFuncs) > 0 && len(excludeFuncs) > 0 {
		log.Fatalf("func include and exclude flags are mutually exclusive")
	}

	if len(includeFields) > 0 && len(excludeFields) > 0 {
		log.Fatalf("field include and exclude flags are mutually exclusive")
	}

	src, err := gopygen.Process(filename, types,
		filterFunc(includeFuncs, excludeFuncs),
		filterFunc(includeFields, excludeFields))
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
