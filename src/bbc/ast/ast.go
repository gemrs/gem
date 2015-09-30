package ast

type File struct {
	Name  string
	Scope *Scope
}

func NewFile(filename string) *File {
	return &File{
		Name:  filename,
		Scope: NewScope(),
	}
}

type Scope struct {
	S []Node
}

func NewScope() *Scope {
	return &Scope{make([]Node, 0)}
}

func (s *Scope) Add(decl Node) {
	s.S = append(s.S, decl)
}

func (s *Scope) Identifier() string {
	return "scope"
}

func (s *Scope) ByteLength() (int, error) {
	accum := 0
	for _, f := range s.S {
		len, err := f.ByteLength()
		if err != nil {
			return 0, err
		}

		accum = accum + len
	}

	return accum, nil
}

type Struct struct {
	Name  string
	Scope *Scope
}

func NewStruct(name string) *Struct {
	return &Struct{
		Name:  name,
		Scope: NewScope(),
	}
}

func (s Struct) Identifier() string {
	return s.Name
}

func (s Struct) ByteLength() (int, error) {
	return s.Scope.ByteLength()
}

type Field struct {
	Name string
	Type Node
}

type FrameSize int

func (f Field) Identifier() string {
	return f.Name
}

func (f Field) ByteLength() (int, error) {
	return f.Type.ByteLength()
}

//go:generate stringer -type=FrameSize
const (
	SzFixed FrameSize = iota
	SzVar8
	SzVar16
)

type Frame struct {
	Name   string
	Number int
	Size   FrameSize
	Object Node
}

func (s Frame) Identifier() string {
	return s.Name
}

func (s Frame) ByteLength() (int, error) {
	return s.Object.ByteLength()
}
