package typewriters

import (
	"io"

	"github.com/clipperhouse/typewriter"
)

type TypeCollector interface {
	Visit(t typewriter.Type) error
	Collect(w io.Writer) error
}
