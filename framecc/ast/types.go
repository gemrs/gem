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

type Type interface {
	// If the byte length is variable, error
	Lengthable
	Identifiable
}

/* Standard Types */
type LengthSpec interface {
	Lengthable
	ConstantExpr() (bool, error)
}

type FixedLength struct {
	Length int
}

func (f *FixedLength) ConstantExpr() (bool, error) {
	return true, nil
}

func (f *FixedLength) ByteLength() (int, error) {
	return f.Length, nil
}

/* A length spec which is determined at runtime by evaluating an integer field */
type ReferenceLength struct {
	Field *Field
}

func (r *ReferenceLength) ConstantExpr() (bool, error) {
	return false, nil
}

func (r *ReferenceLength) ByteLength() (int, error) {
	return 0, ErrVariableType
}

// A Fixed length string (eg. string[256])
type StringType struct {
	Length LengthSpec
}

func (s *StringType) Identifier() string {
	return "string"
}

func (s *StringType) ByteLength() (int, error) {
	return s.Length.ByteLength()
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
	return 0, ErrVariableType
}

type ArrayType struct {
	Object Type
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
