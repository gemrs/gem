package rt3

import (
	"bytes"
	"compress/bzip2"
	"errors"
	"io"
	"io/ioutil"
	"strings"

	"github.com/gemrs/gem/gem/core/encoding"
)

var ErrFileMissing = errors.New("missing file")

type JagArchive struct {
	header arcHeader
	index  []ArchiveFileIndex
	files  map[int32][]byte
}

func NewJagArchive(data []byte) (*JagArchive, error) {
	archive := &JagArchive{}
	reader := bytes.NewReader(data)
	err := archive.Decode(reader, nil)
	if err != nil {
		return nil, err
	}
	return archive, nil
}

func jagHash(name string) int32 {
	name = strings.ToUpper(name)
	hash := int32(0)
	for _, c := range name {
		hash = (hash*61 + c) - 32
	}
	return hash
}

func (struc *JagArchive) File(name string) ([]byte, error) {
	if data, ok := struc.files[jagHash(name)]; ok {
		return data, nil
	}
	return nil, ErrFileMissing
}

func (struc *JagArchive) Decode(buf io.Reader, flags interface{}) (err error) {
	err = struc.header.Decode(buf, nil)
	if err != nil {
		return err
	}

	buf = bytes.NewReader(struc.header.Body)

	var fileCount encoding.Uint16
	fileCount.Decode(buf, encoding.IntNilFlag)

	dataOffset := 2 + int(fileCount*10)
	struc.index = make([]ArchiveFileIndex, fileCount)
	struc.files = make(map[int32][]byte)
	for i := range struc.index {
		struc.index[i].Decode(buf, nil)

		ident := int32(struc.index[i].Identifier)
		dataSize := int(struc.index[i].CompressedSize)

		struc.files[ident] =
			struc.header.Body[dataOffset : dataOffset+dataSize]

		if !struc.header.Decompressed {
			struc.files[ident], err = headerlessBzip2Decompress(struc.files[ident])
			if err != nil {
				return err
			}
		}

		dataOffset += int(struc.index[i].CompressedSize)
	}

	return nil
}

type arcHeader struct {
	UncompressedSize encoding.Uint24
	CompressedSize   encoding.Uint24
	Body             encoding.Bytes
	Decompressed     bool
}

func (struc *arcHeader) Decode(buf io.Reader, flags interface{}) (err error) {
	struc.UncompressedSize.Decode(buf, encoding.IntNilFlag)

	struc.CompressedSize.Decode(buf, encoding.IntNilFlag)

	struc.Body.Decode(buf, nil)

	if int(struc.UncompressedSize) != int(struc.CompressedSize) {
		uncompressed, err := headerlessBzip2Decompress([]byte(struc.Body))
		if err != nil {
			return err
		}

		struc.Decompressed = true
		struc.Body = encoding.Bytes(uncompressed)
	}

	return nil
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
