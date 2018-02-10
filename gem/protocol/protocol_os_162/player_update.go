package protocol_os_162

import (
	"io"

	"github.com/gemrs/gem/gem/core/encoding"
	"github.com/gemrs/gem/gem/protocol/protocol_os"
)

// +gen define_outbound:"Pkt40,SzVar16"
type PlayerUpdate protocol_os.PlayerUpdate

func (struc PlayerUpdate) Encode(w io.Writer, flags interface{}) {
	config := protocol_os.PlayerUpdateConfig{
		TranslatePlayerFlags: translatePlayerFlags,
		EncodeAppearanceBlock: func(w encoding.Writer, srcBlock []byte) {
			block := make([]byte, len(srcBlock))
			for i := range srcBlock {
				block[i] = srcBlock[len(block)-i-1]
			}
			w.PutU8(len(block), encoding.IntNegate)
			w.Write(block)
		},
	}
	protocol_os.PlayerUpdate(struc).Encode(w, config)
}
