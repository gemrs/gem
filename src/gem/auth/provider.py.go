// Generated by gopygen; DO NOT EDIT
package auth

import (
	"fmt"

	"github.com/qur/gopy/lib"

	"github.com/tgascoigne/gopygen/gopygen"
)

// Sometimes we might generate code which doesn't use some of the above imports
// Use them here just in case
var _ = fmt.Sprintf("")
var _ = gopygen.Dummy

var ProviderImplDef = py.Class{
	Name:    "ProviderImpl",
	Flags:   py.TPFLAGS_BASETYPE,
	Pointer: (*ProviderImpl)(nil),
}

// Registers this type with a python module
func RegisterProviderImpl(module *py.Module) error {
	var err error
	var class *py.Type
	if class, err = ProviderImplDef.Create(); err != nil {
		return err
	}

	if err = module.AddObject("ProviderImpl", class); err != nil {
		return err
	}

	return nil
}

func (p *ProviderImpl) Py_LookupProfile(_args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	lock := py.NewLock()
	defer lock.Unlock()

	var err error
	_ = err
	args := _args.Slice()
	if len(args) != 2 {
		return nil, fmt.Errorf("Py_LookupProfile: parameter length mismatch")
	}

	in_0, err := gopygen.TypeConvIn(args[0], "string")
	if err != nil {
		return nil, err
	}

	in_1, err := gopygen.TypeConvIn(args[1], "string")
	if err != nil {
		return nil, err
	}

	res0, res1 := p.LookupProfile(in_0.(string), in_1.(string))

	out_0, err := gopygen.TypeConvOut(res0, "*player.Profile")
	if err != nil {
		return nil, err
	}

	out_1, err := gopygen.TypeConvOut(res1, "AuthResponse")
	if err != nil {
		return nil, err
	}

	return py.PackTuple(out_0, out_1)

}
