package ast

import (
	"bytes"
	"fmt"
)

var ErrVariableType = fmt.Errorf("Can't calculate length of a variable-length type")
var ErrNoSuchType = fmt.Errorf("No such type")
var ErrNoSuchField = fmt.Errorf("No such field")
var ErrFieldNotInteger = fmt.Errorf("Field in array dimension expr not an integer")

type Identifiable interface {
	Identifier() string
}

type Lengthable interface {
	ByteLength() (int, error)
}

type Node interface {
	// If the byte length is variable, error
	Lengthable
	Identifiable
}

/* Standard Types */
type LengthSpec interface {
	Lengthable
	ConstantExpr() (bool, error)
}

type StaticLength struct {
	Length int
}

func (f *StaticLength) ConstantExpr() (bool, error) {
	return true, nil
}

func (f *StaticLength) ByteLength() (int, error) {
	return f.Length, nil
}

/* A reference to a type declaration to be resolved post-parse */
type DeclReference struct {
	DeclName string
}

func (r *DeclReference) Identifier() string {
	return fmt.Sprintf("<unresolved reference %v>", r.DeclName)
}

func (r *DeclReference) ConstantExpr() (bool, error) {
	return false, nil
}

func (r *DeclReference) ByteLength() (int, error) {
	return 0, ErrVariableType
}

/* A length spec which is determined at runtime by evaluating an integer field */
type DynamicLength struct {
	Field Node
}

func (r *DynamicLength) Identifier() string {
	return "dynamic length"
}

func (r *DynamicLength) ConstantExpr() (bool, error) {
	return false, nil
}

func (r *DynamicLength) ByteLength() (int, error) {
	return 0, ErrVariableType
}

// A Fixed length string (eg. string[256])
type StringBaseType struct {}

func (s *StringBaseType) Identifier() string {
	return "string"
}

func (s *StringBaseType) ByteLength() (int, error) {
	return 1, nil
}

type IntegerFlag int

const (
	IntNegate IntegerFlag = (1 << iota)
	IntInv128
	IntOfs128
	IntLittleEndian
	IntPDPEndian
	IntRPDPEndian
)

type IntegerType struct {
	Signed    bool
	Bitsize   int
	Modifiers IntegerFlag
}

func (i *IntegerType) Set(flag IntegerFlag) {
	i.Modifiers = i.Modifiers | flag
}

func (s *IntegerType) Identifier() string {
	var buf bytes.Buffer
	if s.Signed {
		buf.WriteString("u")
	}
	buf.WriteString("int")
	buf.WriteString(fmt.Sprintf("%v", s.Bitsize))
	return buf.String()
}

func (s *IntegerType) ByteLength() (int, error) {
	return (s.Bitsize / 8), ErrVariableType
}

type ArrayType struct {
	Object Node
	Length LengthSpec
}

func (a *ArrayType) Identifier() string {
	return "string"
}

func (a *ArrayType) ByteLength() (int, error) {
	baseLength, err := a.Object.ByteLength()
	if err != nil {
		return 0, err
	}
	mulLength, err := a.Length.ByteLength()
	if err != nil {
		return 0, err
	}

	return baseLength * mulLength, nil
}
