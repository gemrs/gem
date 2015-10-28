package auth

import (
	"fmt"

	"github.com/qur/gopy/lib"

	"github.com/sinusoids/gem/pybind"
)

var ProviderImplDef = pybind.Define("ProviderImpl", (*ProviderImpl)(nil))
var RegisterProviderImpl = pybind.GenerateRegisterFunc(ProviderImplDef)
var NewProviderImpl = pybind.GenerateConstructor(ProviderImplDef).(func() *ProviderImpl)

func createAuthConstants(module *py.Module) error {
	for c := AuthPending; c < AuthEnd; c++ {
		pyInt := py.NewInt(int(c))
		if err := module.AddObject(fmt.Sprintf("%v", c), pyInt); err != nil {
			panic(err)
		}
	}
	return nil
}
