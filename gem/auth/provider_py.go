package auth

import (
	"fmt"
	"github.com/qur/gopy/lib"
)

func createAuthConstants(module *py.Module) error {
	for c := AuthPending; c < AuthEnd; c++ {
		pyInt := py.NewInt(int(c))
		if err := module.AddObject(fmt.Sprintf("%v", c), pyInt); err != nil {
			panic(err)
		}
	}
	return nil
}
