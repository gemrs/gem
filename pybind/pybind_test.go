package pybind_test

import (
	"fmt"
	"testing"

	"github.com/qur/gopy/lib"

	"github.com/gemrs/gem/gem/python/modules"
	"github.com/gemrs/gem/gem/util/safe"
	"github.com/gemrs/gem/pybind"
)

var testCatOut = `
import test_module
the_cat = test_module.cat("Paws")
the_cat.give_cheeseburger()
the_cat.rename("Garfield")
`

var testCatIn = `
from test_module import the_cat
the_cat.give_cheeseburger()
the_cat.rename("Garfield")
`

// TestPyBindOut tests creating objects in python and extracting them
func TestPyBindOut(t *testing.T) {
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
	} else if _, err := py.RunString(testCatOut, py.FileInput, main, nil); err != nil {
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

// TestPyBindIn tests creating objects in go and inserting them into python
func TestPyBindIn(t *testing.T) {
	_ = py.NewLock()
	safe.Unsafe = true

	var main *py.Dict
	var module *py.Module
	var err error
	if main, err = py.NewDict(); err != nil {
		t.Fatal(err)
	} else if g, err := py.GetBuiltins(); err != nil {
		t.Fatal(err)
	} else if err := main.SetItemString("__builtins__", g); err != nil {
		t.Fatal(err)
	} else if module, err = modules.Init("test_module", []py.Method{}); err != nil {
		t.Fatal(err)
	} else if err := RegisterCat(module); err != nil {
		t.Fatal(err)
	}

	the_cat := NewCat("Paws")
	the_cat.GiveCheeseburger()

	if err := module.AddObject("the_cat", the_cat); err != nil {
		t.Fatal(err)
	} else if _, err := py.RunString(testCatIn, py.FileInput, main, nil); err != nil {
		t.Fatal(err)
	} else if a, err := main.GetItemString("the_cat"); err != nil {
		t.Fatal(err)
	} else if the_cat, ok := a.(*Cat); !ok {
		t.Errorf("Unable to extract test object")
	} else {
		if the_cat.Name != "Garfield" || the_cat.Cheeseburgers != 2 {
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

func (c *Cat) Init(name string) {
	fmt.Println("calling init")
	c.Name = name
	c.Cheeseburgers = 0
}

func (c *Cat) Rename(name string) {
	fmt.Println("calling rename")
	c.Name = name
}

func (c *Cat) GiveCheeseburger() {
	fmt.Println("calling burgers")
	c.Cheeseburgers++
}

func (c *Cat) Py_rename(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(c.Rename)
	return fn(args, kwds)
}

func (c *Cat) Py_give_cheeseburger(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(c.GiveCheeseburger)
	return fn(args, kwds)
}

var CatDef = pybind.Define("cat", (*Cat)(nil))
var RegisterCat = pybind.GenerateRegisterFunc(CatDef)
var NewCat = pybind.GenerateConstructor(CatDef).(func(string) *Cat)
