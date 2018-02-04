package encoding

import (
	"bytes"
	"io"
	"io/ioutil"

	"github.com/gemrs/gem/fork/golang.org/x/crypto/xtea"
)

type XTEABlock struct {
	Codable
}

type XTEADecodeArgs struct {
	Key [4]uint32
}

func (block *XTEABlock) Encode(buf io.Writer, flags interface{}) {
	panic("not implemented")
}

func (block *XTEABlock) Decode(buf io.Reader, flags interface{}) {
	args := flags.(XTEADecodeArgs)

	encryptedBlock, err := ioutil.ReadAll(buf)
	if err != nil {
		panic(err)
	}
	decryptedBlock := DecryptXteaBlock(encryptedBlock, args.Key[:])

	msgBuf := bytes.NewBuffer(decryptedBlock)
	block.Codable.Decode(msgBuf, nil)
}

func DecryptXteaBlock(block []byte, keys []uint32) []byte {
	cipher, err := xtea.NewCipher(keys[0:4])
	if err != nil {
		panic(err)
	}

	buf := NewBufferBytes(block)
	var deciphered bytes.Buffer
	blockIn := make([]byte, xtea.BlockSize)
	blockOut := make([]byte, xtea.BlockSize)
	for {
		n, err := buf.Read(blockIn)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		if n < xtea.BlockSize {
			deciphered.Write(blockIn)
			break
		}

		cipher.Decrypt(blockOut, blockIn)
		deciphered.Write(blockOut)
	}

	return deciphered.Bytes()
}
