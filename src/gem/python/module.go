package python

import (
	"fmt"
	"strings"

	"github.com/qur/gopy/lib"
)

var gemModules = map[string]*py.Module{}

func InitModule(name string, methods []py.Method) (*py.Module, error) {
	module, err := py.InitModule(name, methods)
	if err != nil {
		return module, err
	}
	fmt.Printf("registered %v\n", name)
	gemModules[name] = module
	return module, err
}

// LinkModules ensures that all modules are inserted as objects into their parent packages
// Since we're initing modules in an indeterminite order, we do the linking step as a second pass,
// performed after all modules have been initialized
func LinkModules() {
	for module, object := range gemModules {
		idx := strings.LastIndex(module, ".")
		if idx == -1 {
			continue
		}
		packagePath := module[:idx]
		parent, ok := gemModules[packagePath]
		if !ok {
			panic(fmt.Sprintf("parent package of %v doesn't exist", module))
		}

		module = module[idx+1:]
		parent.AddObject(module, object)
	}
}
