// Code generated by bbc; DO NOT EDIT.
package update_protocol

import (
	"io"

	"github.com/gemrs/gem/gem/encoding"
)

type OutboundUpdateResponse struct {
	Index encoding.Int8
	File  encoding.Int16
	Size  encoding.Int16
	Chunk encoding.Int8
	Data  encoding.Bytes
}

func (struc *OutboundUpdateResponse) Encode(buf io.Writer, flags interface{}) (err error) {
	err = struc.Index.Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	err = struc.File.Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	err = struc.Size.Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	err = struc.Chunk.Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	err = struc.Data.Encode(buf, 500)
	if err != nil {
		return err
	}

	return err
}

func (struc *OutboundUpdateResponse) Decode(buf io.Reader, flags interface{}) (err error) {
	err = struc.Index.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	err = struc.File.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	err = struc.Size.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	err = struc.Chunk.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	err = struc.Data.Decode(buf, 500)
	if err != nil {
		return err
	}

	return err
}

type InboundUpdateRequest struct {
	Index    encoding.Int8
	File     encoding.Int16
	Priority encoding.Int8
}

func (struc *InboundUpdateRequest) Encode(buf io.Writer, flags interface{}) (err error) {
	err = struc.Index.Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	err = struc.File.Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	err = struc.Priority.Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	return err
}

func (struc *InboundUpdateRequest) Decode(buf io.Reader, flags interface{}) (err error) {
	err = struc.Index.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	err = struc.File.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	err = struc.Priority.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	return err
}
