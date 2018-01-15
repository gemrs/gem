// Code generated by bbc; DO NOT EDIT.
package rt3

import (
	"io"

	"github.com/gemrs/gem/gem/encoding"
)

type ArchiveFileIndex struct {
	Identifier       encoding.Uint32
	UncompressedSize encoding.Uint24
	CompressedSize   encoding.Uint24
}

func (struc *ArchiveFileIndex) Encode(buf io.Writer, flags interface{}) (err error) {
	err = struc.Identifier.Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	err = struc.UncompressedSize.Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	err = struc.CompressedSize.Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	return err
}

func (struc *ArchiveFileIndex) Decode(buf io.Reader, flags interface{}) (err error) {
	err = struc.Identifier.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	err = struc.UncompressedSize.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	err = struc.CompressedSize.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	return err
}