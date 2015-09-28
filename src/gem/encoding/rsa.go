package encoding

import (
	"bytes"
	"io"
	"math/big"

	"gem/crypto"
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

	ciphertext := new(big.Int)
	ciphertext.SetBytes(data)

	// Decrypt into msg
	msg := new(big.Int).Exp(ciphertext, key.D, key.N)

	// Decode into
	msgBuf := bytes.NewBuffer(msg.Bytes())
	return rsa.Codable.Decode(msgBuf, nil)
}
