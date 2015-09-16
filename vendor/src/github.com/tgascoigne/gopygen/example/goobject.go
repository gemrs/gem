package example

import (
	"github.com/qur/gopy/lib"
)

//go:generate gopygen $GOFILE GoObject

type GoObject struct {
	py.BaseObject
	X          int `Py:"X"`
	Y          int `Py:"Y"`
	Z          int
	unexported int `Py:"unexp"`
	Map        map[int]interface{}
	Slice      []string
}

func (o *GoObject) FooBar_1(z int) int {
	return o.X * o.Y * z
}

func (o *GoObject) FooBar_2(z int) (a int) {
	a = o.X * o.Y * z
	return
}

func (o *GoObject) FooBar_3(a, b string) (x, y int) {
	return len(a), len(b)
}

func (o *GoObject) FooBar_4(a string, b string) (int, int) {
	return len(a), len(b)
}

func (o *GoObject) FooBar_5(z int) {
	o.Z = z
}

func (o *GoObject) FooBar_6() int {
	return o.Z
}

// Non-pointer reciever - ignored
func (o GoObject) BarFoo(z int) int {
	return o.X * o.Y * z
}
