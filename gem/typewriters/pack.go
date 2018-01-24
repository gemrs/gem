package typewriters

import (
	"fmt"
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

func (p Pack) Write(w io.Writer, t typewriter.Type) error {
	_, ok := t.FindTag(p)
	if !ok {
		return nil
	}

	fmt.Fprintf(w, "func (o %v) Encode(w io.Writer, flags interface{}) {\n", t.Name)
	fmt.Fprintf(w, "server.Proto.Encode(o).Encode(w, flags)\n")
	fmt.Fprintf(w, "}\n\n")
	return nil
}
