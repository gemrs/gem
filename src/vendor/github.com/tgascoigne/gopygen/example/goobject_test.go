package example_test

import (
	"fmt"
	"testing"

	"./"
	"github.com/qur/gopy/lib"
)

func initPythonLib() *py.Dict {
	py.Initialize()

	m, err := py.InitModule("gen", []py.Method{})
	if err != nil {
		panic(err)
	}

	err = example.RegisterGoObject(m)
	if err != nil {
		panic(err)
	}

	glob, err := py.NewDict()
	if err != nil {
		panic(err)
	}

	g, err := py.GetBuiltins()
	if err != nil {
		panic(err)
	}

	err = glob.SetItemString("__builtins__", g)
	if err != nil {
		panic(err)
	}

	return glob
}

func TestAttributes(t *testing.T) {
	py.Initialize()
	globals := initPythonLib()

	/* Create an object and set some attributes */
	var pyString = `import gen
obj = gen.GoObject()
obj.X = 5
obj.Y = 2
`
	_, err := py.RunString(pyString, py.FileInput, globals, nil)
	if err != nil {
		t.Fatal(err)
	}

	goObj_, err := globals.GetItemString("obj")
	if err != nil {
		t.Fatal(err)
	}
	goObj := goObj_.(*example.GoObject)

	t.Logf("X: %v", goObj.X)
	t.Logf("Y: %v", goObj.Y)
	if goObj.X != 5 || goObj.Y != 2 {
		t.Error("attribute value mismatch")
	}

	/* Modify attributes in go-land */
	goObj.X *= 2
	goObj.Y *= 2

	/* Access attributes and check we got the right values back */
	pyString = `
retrievedX = obj.X
retrievedY = obj.Y
`
	_, err = py.RunString(pyString, py.FileInput, globals, nil)
	if err != nil {
		t.Fatal(err)
	}

	retrievedX_, err := globals.GetItemString("retrievedX")
	if err != nil {
		t.Fatal(err)
	}
	retrievedX := retrievedX_.(*py.Int).Int()

	retrievedY_, err := globals.GetItemString("retrievedY")
	if err != nil {
		t.Fatal(err)
	}
	retrievedY := retrievedY_.(*py.Int).Int()

	t.Logf("X: %v", retrievedX)
	t.Logf("Y: %v", retrievedY)
	if retrievedX != 10 || retrievedY != 4 {
		t.Error("attribute value mismatch")
	}

	py.Finalize()
}

func TestIgnoredAttributes(t *testing.T) {
	py.Initialize()
	globals := initPythonLib()

	var pyString = `import gen
obj = gen.GoObject()
obj.Z = 1
`
	_, err := py.RunString(pyString, py.FileInput, globals, nil)
	if err == nil {
		t.Error("ignored attribute didn't throw error")
	}

	py.Finalize()
}

func TestObjectAllocation(t *testing.T) {
	py.Initialize()

	obj, err := example.GoObject{X: 2, Y: 3}.Alloc()
	if err != nil {
		t.Fatal(err)
	}

	if obj.X != 2 || obj.Y != 3 {
		t.Error("member field value mismatch")
	}

	py.Finalize()
}

func TestObjectSharing(t *testing.T) {
	py.Initialize()
	globals := initPythonLib()

	obj, err := example.GoObject{X: 2, Y: 3}.Alloc()
	if err != nil {
		t.Fatal(err)
	}

	var pyString = `
def testObjectSharing(obj):
	retrievedX = obj.X
	retrievedY = obj.Y
	return (retrievedX, retrievedY)
`
	_, err = py.RunString(pyString, py.FileInput, globals, nil)
	if err != nil {
		t.Fatal(err)
	}

	pyFunc_, err := globals.GetItemString("testObjectSharing")
	if err != nil {
		t.Fatal(err)
	}
	pyFunc := pyFunc_.Base()

	argsTuple, err := py.PackTuple(obj)
	if err != nil {
		t.Fatal(err)
	}

	result, err := pyFunc.CallObject(argsTuple)
	if err != nil {
		t.Fatal(err)
	}

	var retrievedX, retrievedY int
	py.ParseTuple(result.(*py.Tuple), "ii", &retrievedX, &retrievedY)

	t.Logf("X: %v", retrievedX)
	t.Logf("Y: %v", retrievedY)
	if retrievedX != 2 || retrievedY != 3 {
		t.Error("attribute value mismatch")
	}

	py.Finalize()
}

func TestMethodWrapping_1(t *testing.T) {
	py.Initialize()
	globals := initPythonLib()

	for _, f := range []string{"FooBar_1", "FooBar_2"} {
		var pyString = `import gen
obj = gen.GoObject()
obj.X = 1
obj.Y = 2
result = obj.%s(5)
`
		_, err := py.RunString(fmt.Sprintf(pyString, f), py.FileInput, globals, nil)
		if err != nil {
			t.Fatal(err)
		}

		result_, err := globals.GetItemString("result")
		if err != nil {
			t.Fatal(err)
		}
		result := result_.(*py.Int).Int()

		t.Logf("result: %v", result)
		if result != 10 {
			t.Error("result value mismatch")
		}

	}
	py.Finalize()
}

func TestMethodWrapping_2(t *testing.T) {
	py.Initialize()
	globals := initPythonLib()

	for _, f := range []string{"FooBar_3", "FooBar_4"} {
		var pyString = `import gen
obj = gen.GoObject()
szA, szB = obj.%s("cat", "mouse")
`
		_, err := py.RunString(fmt.Sprintf(pyString, f), py.FileInput, globals, nil)
		if err != nil {
			t.Fatal(err)
		}

		szA_, err := globals.GetItemString("szA")
		if err != nil {
			t.Fatal(err)
		}
		szA := szA_.(*py.Int).Int()

		szB_, err := globals.GetItemString("szB")
		if err != nil {
			t.Fatal(err)
		}
		szB := szB_.(*py.Int).Int()

		t.Logf("szA: %v", szA)
		t.Logf("szB: %v", szB)
		if szA != 3 || szB != 5 {
			t.Error("return value mismatch")
		}

	}

	py.Finalize()
}

func TestMethodWrapping_3(t *testing.T) {
	py.Initialize()
	globals := initPythonLib()

	var pyString = `import gen
obj = gen.GoObject()
obj.FooBar_5(200)
result = obj.FooBar_6()
`
	_, err := py.RunString(pyString, py.FileInput, globals, nil)
	if err != nil {
		t.Fatal(err)
	}

	result_, err := globals.GetItemString("result")
	if err != nil {
		t.Fatal(err)
	}
	result := result_.(*py.Int).Int()

	t.Logf("result: %v", result)
	if result != 200 {
		t.Error("return value mismatch")
	}

	py.Finalize()
}
