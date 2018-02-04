package rt3

import (
	"bytes"
	"compress/bzip2"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"

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
	dataBuffer       *bytes.Buffer
}

func NewContainer(compression CompressionType, data []byte) *Container {
	switch compression {
	case CompressionNone:
		return &Container{
			CompressionType: compression,
			Size:            len(data),
			Data:            data,
			Version:         -1,
		}
	default:
		panic(fmt.Errorf("compressed container encode not implemented"))
	}
}

func (struc *Container) Read(p []byte) (n int, err error) {
	if struc.dataBuffer == nil {
		struc.dataBuffer = new(bytes.Buffer)
		struc.dataBuffer.Write(struc.Data)
	}

	return struc.dataBuffer.Read(p)
}

func (struc *Container) Encode(w io.Writer, flags interface{}) {
	buf := encoding.WrapWriter(w)
	buf.PutU8(int(struc.CompressionType))
	buf.PutU32(struc.Size)

	if struc.CompressionType != CompressionNone {
		panic(fmt.Errorf("compressed container encode not implemented"))
	}

	buf.PutBytes(struc.Data)
	if struc.Version != -1 {
		buf.Put16(struc.Version)
	}
}

func (struc *Container) Decode(r io.Reader, flags interface{}) {
	buf := encoding.WrapReader(r)
	struc.CompressionType = CompressionType(buf.GetU8())

	struc.Size = buf.GetU32()

	if struc.CompressionType == CompressionNone {
		struc.Data = buf.GetBytes(struc.Size)
	} else {
		struc.UncompressedSize = buf.GetU32()
		var err error

		data := buf.GetBytes(struc.Size)

		switch struc.CompressionType {
		case CompressionBzip2:
			struc.Data, err = headerlessBzip2Decompress(data)

		case CompressionGzip:
			var inBuf, outBuf bytes.Buffer
			inBuf.Write(data)

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
			panic(fmt.Errorf("Container decompression failed: %v", err))
		}

		if len(struc.Data) != struc.UncompressedSize {
			panic(fmt.Errorf("Container decompression length mismatch"))
		}

	}

	var tmp16 encoding.Uint16
	err := encoding.TryDecode(&tmp16, buf, encoding.IntNilFlag)
	if err != nil {
		struc.Version = int(tmp16)
	} else {
		struc.Version = -1
	}
}

func headerlessBzip2Decompress(compressed []byte) ([]byte, error) {
	header := []byte{'B', 'Z', 'h', '1'}
	bzipData := append(header, compressed...)
	reader := bytes.NewReader(bzipData)
	bzip2Reader := bzip2.NewReader(reader)
	uncompressed, err := ioutil.ReadAll(bzip2Reader)
	if err != nil {
		return nil, err
	}

	return uncompressed, nil
}
