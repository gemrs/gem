package ast

type File struct {
	Name string
	Decls  map[string]Node
}

func NewFile(filename string) *File {
	return &File{
		Name: filename,
		Decls: make(map[string]Node),
	}
}

type Struct struct {
	Name   string
	Fields []*Field
}

func (s Struct) Identifier() string {
	return s.Name
}

func (s Struct) ByteLength() (int, error) {
	accum := 0
	for _, f := range s.Fields {
		len, err := f.Type.ByteLength()
		if err != nil {
			return 0, err
		}

		accum = accum + len
	}

	return accum, nil
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
