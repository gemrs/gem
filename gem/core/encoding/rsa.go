package encoding

import (
	"bytes"
	"io"

	"github.com/gemrs/gem/gem/core/crypto"
)

type RSABlock struct {
	Codable
}

type RSADecodeArgs struct {
	Key       *crypto.Keypair
	BlockSize int
}

func (rsa *RSABlock) Encode(buf io.Writer, flags interface{}) {
	panic("not implemented")
}

func (rsa *RSABlock) Decode(buf io.Reader, flags interface{}) {
	args := flags.(RSADecodeArgs)
	key := args.Key
	size := args.BlockSize

	// Get the block into a big.Int
	data := make([]byte, size)
	n, err := buf.Read(data)
	if err != nil {
		panic(err)
	}

	if n != size {
		panic(io.EOF)
	}

	// Decrypt into buffer
	msgBuf := bytes.NewBuffer(key.Decrypt(data))
	rsa.Codable.Decode(msgBuf, nil)
}
