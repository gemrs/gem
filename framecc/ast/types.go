package ast

import (
	"fmt"
	"bytes"
)

var ErrVariableType = fmt.Errorf("Can't calculate length of a variable-length type")
var ErrNoSuchType = fmt.Errorf("No such type")
var ErrNoSuchField = fmt.Errorf("No such field")
var ErrFieldNotInteger = fmt.Errorf("Field in array dimension expr not an integer")

type Identifiable interface{
	Identifier() string
}

type Type interface{
	// If the byte length is variable, error
	ByteLength() (int, error)
	Identifiable
}

/* Standard Primitives */

// A Fixed length string (eg. string[256])
type StringType struct {
	Length int
}

func (s *StringType) Identifier() string {
	return "string"
}

func (s *StringType) ByteLength() (int, error) {
	return s.Length, nil
}


// A variable length string (eg. string[LocalField])
type VariableStringType struct {
	Field *Field
}

func (s *VariableStringType) Identifier() string {
	return "varstring"
}

func (s *VariableStringType) ByteLength() (int, error) {
	return 0, ErrVariableType
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
	Signed bool
	Bitsize int
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
