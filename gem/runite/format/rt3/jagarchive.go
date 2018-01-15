package rt3

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"strings"

	"github.com/dsnet/compress/bzip2"

	"github.com/gemrs/gem/gem/encoding"
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
		hash = (hash*61 + int32(c)) - 32
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
	err = fileCount.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	dataOffset := 2 + int(fileCount*10)
	struc.index = make([]ArchiveFileIndex, fileCount)
	struc.files = make(map[int32][]byte)
	for i := range struc.index {
		err2 := struc.index[i].Decode(buf, nil)
		if err2 != nil {
			return err2
		}

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
	err = struc.UncompressedSize.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	err = struc.CompressedSize.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	err = struc.Body.Decode(buf, nil)
	if err != nil {
		return err
	}

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
	bzipData := append([]byte{'B', 'Z', 'h', '1'}, compressed...)
	reader := bytes.NewReader(bzipData)
	config := bzip2.ReaderConfig{}
	bzip2Reader, err := bzip2.NewReader(reader, &config)
	if err != nil {
		return nil, err
	}

	uncompressed, err := ioutil.ReadAll(bzip2Reader)
	if err != nil {
		return nil, err
	}

	return uncompressed, nil
}
