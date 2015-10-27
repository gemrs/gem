package pybind_test

import (
	"fmt"
	"testing"

	"github.com/qur/gopy/lib"

	"github.com/sinusoids/gem/gem/python/modules"
	"github.com/sinusoids/gem/gem/util/safe"
	"github.com/sinusoids/gem/pybind"
)

var testCat = `
import test_module
the_cat = test_module.cat("Paws")
the_cat.give_cheeseburger()
the_cat.rename("Garfield")
`

func TestPyBind(t *testing.T) {
	_ = py.NewLock()
	safe.Unsafe = true

	if main, err := py.NewDict(); err != nil {
		t.Fatal(err)
	} else if g, err := py.GetBuiltins(); err != nil {
		t.Fatal(err)
	} else if err := main.SetItemString("__builtins__", g); err != nil {
		t.Fatal(err)
	} else if module, err := modules.Init("test_module", []py.Method{}); err != nil {
		t.Fatal(err)
	} else if err := RegisterCat(module); err != nil {
		t.Fatal(err)
	} else if _, err := py.RunString(testCat, py.FileInput, main, nil); err != nil {
		t.Fatal(err)
	} else if a, err := main.GetItemString("the_cat"); err != nil {
		t.Fatal(err)
	} else if the_cat, ok := a.(*Cat); !ok {
		t.Errorf("Unable to extract test object")
	} else {
		if the_cat.Name != "Garfield" || the_cat.Cheeseburgers != 1 {
			t.Errorf("Object's properties weren't updated")
		}
		t.Logf("%V", the_cat)
	}
}

type Cat struct {
	py.BaseObject

	Name          string
	Cheeseburgers int
}

func NewCat(name string) (*Cat, error) {
	args := pybind.ReflectValues(name)
	argsObjs, err := pybind.ConvertOut(args)
	if err != nil {
		return nil, err
	}

	argsTuple, err := py.PackTuple(argsObjs...)
	if err != nil {
		return nil, err
	}

	cat, err := CatDef.New(CatDef.Type, argsTuple, nil)
	if err != nil {
		return nil, err
	}

	return cat.(*Cat), nil
}

func InitCat(c *Cat, name string) error {
	c.Name = name
	c.Cheeseburgers = 0
	return nil
}

func (c *Cat) Rename(name string) {
	c.Name = name
}

func (c *Cat) GiveCheeseburger() {
	c.Cheeseburgers++
}

var CatDef = py.Class{
	Name:    "cat",
	Flags:   py.TPFLAGS_BASETYPE,
	Pointer: (*Cat)(nil),
	New:     pybind.WrapConstructor(InitCat),
}

func RegisterCat(module *py.Module) error {
	var err error
	var class *py.Type
	if class, err = CatDef.Create(); err != nil {
		return err
	}

	if err = module.AddObject(CatDef.Name, class); err != nil {
		return err
	}

	return nil
}

func (c *Cat) Py_rename(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fmt.Println("calling rename")

	fn := pybind.Wrap(c.Rename)
	return fn(args, kwds)
}

func (c *Cat) Py_give_cheeseburger(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fmt.Println("calling burgers")

	fn := pybind.Wrap(c.GiveCheeseburger)
	return fn(args, kwds)
}
