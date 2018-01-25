package typewriters

import (
	"html/template"
	"io"

	"github.com/clipperhouse/typewriter"
)

func init() {
	err := typewriter.Register(&Pack{})
	if err != nil {
		panic(err)
	}
}

type Pack struct{}

func (p Pack) Name() string {
	return "pack_outgoing"
}

func (p Pack) Imports(t typewriter.Type) []typewriter.ImportSpec {
	return nil
}

var packTmpl = template.Must(template.New("").Parse(`
func (o {{.}}) Encode(w io.Writer, flags interface{}) {
	packet := server.Proto.Encode(o)
	if packet != nil {
		packet.Encode(w, flags)
	}
}
`))

func (p Pack) Write(w io.Writer, t typewriter.Type) error {
	_, ok := t.FindTag(p)
	if !ok {
		return nil
	}

	return packTmpl.Execute(w, t.Name)
}
