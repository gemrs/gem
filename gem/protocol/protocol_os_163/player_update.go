package protocol_os_163

import (
	"io"

	"github.com/gemrs/gem/gem/protocol/protocol_os"
)

// +gen define_outbound:"Pkt3,SzVar16"
type PlayerUpdate protocol_os.PlayerUpdate

func (struc PlayerUpdate) Encode(w io.Writer, flags interface{}) {
	protocol_os.PlayerUpdate(struc).Encode(w, translatePlayerFlags)
}
