package rt3

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"

	"github.com/gemrs/gem/gem/core/encoding"
)

type CompressionType int

const (
	CompressionNone CompressionType = iota
	CompressionBzip2
	CompressionGzip
)

type Container struct {
	CompressionType  CompressionType
	Size             int
	UncompressedSize int
	Version          int
	Data             []byte
}

func (struc *Container) Encode(buf io.Writer, flags interface{}) {
	encoding.Uint8(struc.CompressionType).Encode(buf, encoding.IntNilFlag)
	encoding.Uint16(struc.Size).Encode(buf, encoding.IntNilFlag)

	if struc.CompressionType != CompressionNone {
		panic(fmt.Errorf("compressed container encode not implemented"))
	}

	encoding.Bytes(struc.Data).Encode(buf, nil)
	if struc.Version != -1 {
		encoding.Uint16(struc.Version).Encode(buf, encoding.IntNilFlag)
	}
}

func (struc *Container) Decode(buf io.Reader, flags interface{}) {
	var tmp8 encoding.Uint8
	var tmp16 encoding.Uint16
	var tmp32 encoding.Uint32
	var tmpBytes encoding.Bytes

	tmp8.Decode(buf, encoding.IntNilFlag)
	struc.CompressionType = CompressionType(tmp8)

	tmp32.Decode(buf, encoding.IntNilFlag)
	struc.Size = int(tmp32)

	tmpBytes.Decode(buf, struc.Size)

	if struc.CompressionType == CompressionNone {
		struc.Data = []byte(tmpBytes)
	} else {
		tmp32.Decode(buf, encoding.IntNilFlag)
		struc.UncompressedSize = int(tmp32)
		var err error

		switch struc.CompressionType {
		case CompressionBzip2:
			struc.Data, err = headerlessBzip2Decompress([]byte(tmpBytes))

		case CompressionGzip:
			var inBuf, outBuf bytes.Buffer
			inBuf.Write([]byte(tmpBytes))

			reader, err := gzip.NewReader(&inBuf)
			if err != nil {
				panic(fmt.Errorf("Cannot create gzip reader"))
			}

			io.Copy(&outBuf, reader)
			reader.Close()

			struc.Data = outBuf.Bytes()

		default:
			panic(fmt.Errorf("Unknown container compression type"))
		}

		if err != nil {
			panic(fmt.Errorf("Container decompression failed"))
		}

		if len(struc.Data) != struc.UncompressedSize {
			panic(fmt.Errorf("Container decompression length mismatch"))
		}

	}

	err := encoding.TryDecode(&tmp16, buf, encoding.IntNilFlag)
	if err != nil {
		struc.Version = int(tmp16)
	} else {
		struc.Version = -1
	}
}
