package ast

import (
	"bytes"
	"fmt"
	"strings"
	"regexp"
	"strconv"
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
	Node
	ConstantExpr() (bool, error)
}

type StaticLength struct {
	Length int
}

func (f *StaticLength) Identifier() string {
	return "static length"
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
	Object   Node
	Meta     interface{}
}

func (r *DeclReference) Identifier() string {
	if r.Object == nil {
		return fmt.Sprintf("<unresolved reference %v>", r.DeclName)
	}
	return r.Object.Identifier()
}

func (r *DeclReference) ByteLength() (int, error) {
	if r.Object == nil {
		return 0, ErrVariableType
	}
	return r.Object.ByteLength()
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
type StringBaseType struct{}

func (s *StringBaseType) Identifier() string {
	return "string"
}

func (s *StringBaseType) ByteLength() (int, error) {
	return 1, nil
}

type IntegerFlag int

const (
	IntNilFlag IntegerFlag = (1 << iota)
	IntNegate
	IntInverse128
	IntOffset128
	IntLittleEndian
	IntPDPEndian
	IntRPDPEndian
	IntReverse
)

// String is used by the compiler to output human readable flags
// should be done a lot cleanre
func (f IntegerFlag) String() string {
	parts := make([]string, 0)

	if f&IntNegate != 0 {
		parts = append(parts, "encoding.IntNegate")
	}

	if f&IntInverse128 != 0 {
		parts = append(parts, "encoding.IntInverse128")
	}

	if f&IntOffset128 != 0 {
		parts = append(parts, "encoding.IntOffset128")
	}

	if f&IntLittleEndian != 0 {
		parts = append(parts, "encoding.IntLittleEndian")
	}

	if f&IntPDPEndian != 0 {
		parts = append(parts, "encoding.IntPDPEndian")
	}

	if f&IntRPDPEndian != 0 {
		parts = append(parts, "encoding.IntRPDPEndian")
	}

	if len(parts) == 0 {
		parts = append(parts, "0")
	}

	return strings.Join(parts, " | ")
}

type IntegerType struct {
	Signed    bool
	Bitsize   int
	Modifiers IntegerFlag
}

var integerRegexp = regexp.MustCompile("(u)?int(8|16|24|32|64)")
func ParseIntegerType(typ string) (*IntegerType, error) {
	if !integerRegexp.MatchString(typ) {
		return nil, fmt.Errorf("unrecognized integer type: %v", typ)
	}

	groups := integerRegexp.FindAllStringSubmatch(typ, -1)
	bitsize, err := strconv.Atoi(groups[0][2])
	if err != nil {
		return nil, err
	}

	return &IntegerType{
		Signed: groups[0][1] != "u",
		Bitsize: bitsize,
	}, nil
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
