package encoding

import (
	"bytes"
	"io"

	"github.com/sinusoids/gem/gem/crypto"
)

type RSABlock struct {
	Codable
}

type RSADecodeArgs struct {
	Key       *crypto.Keypair
	BlockSize int
}

func (rsa *RSABlock) Encode(buf io.Writer, flags interface{}) error {
	panic("not implemented")
}

func (rsa *RSABlock) Decode(buf io.Reader, flags interface{}) error {
	args := flags.(RSADecodeArgs)
	key := args.Key
	size := args.BlockSize

	// Get the block into a big.Int
	data := make([]byte, size)
	n, err := buf.Read(data)
	if err != nil {
		return err
	}

	if n != size {
		return io.EOF
	}

	// Decrypt into buffer
	msgBuf := bytes.NewBuffer(key.Decrypt(data))
	return rsa.Codable.Decode(msgBuf, nil)
}
