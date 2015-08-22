package ast

type File struct {
	Types  map[string]Type
	Frames map[string]*Frame
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
	Type Type
}

type FrameSize int

const (
	SzFixed FrameSize = iota
	SzVar8
	SzVar16
)

type Frame struct {
	Name   string
	Number int
	Size   FrameSize
	Object Type
}
