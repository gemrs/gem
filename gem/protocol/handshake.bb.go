// Generated by bbc; DO NOT EDIT
package protocol

import (
	"io"

	"github.com/gemrs/gem/gem/encoding"
)

type InboundServiceSelect struct {
	Service encoding.Uint8
}

func (struc *InboundServiceSelect) Encode(buf io.Writer, flags interface{}) (err error) {
	err = struc.Service.Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	return err
}

func (struc *InboundServiceSelect) Decode(buf io.Reader, flags interface{}) (err error) {
	err = struc.Service.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	return err
}

type InboundGameHandshake struct {
	NameHash encoding.Uint8
}

func (struc *InboundGameHandshake) Encode(buf io.Writer, flags interface{}) (err error) {
	err = struc.NameHash.Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	return err
}

func (struc *InboundGameHandshake) Decode(buf io.Reader, flags interface{}) (err error) {
	err = struc.NameHash.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	return err
}

type OutboundGameHandshake struct {
	ignored         [8]encoding.Uint8
	loginRequest    encoding.Uint8
	ServerISAACSeed [2]encoding.Uint32
}

func (struc *OutboundGameHandshake) Encode(buf io.Writer, flags interface{}) (err error) {
	for i := 0; i < 8; i++ {
		err = struc.ignored[i].Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
		if err != nil {
			return err
		}
	}

	err = struc.loginRequest.Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	for i := 0; i < 2; i++ {
		err = struc.ServerISAACSeed[i].Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
		if err != nil {
			return err
		}
	}

	return err
}

func (struc *OutboundGameHandshake) Decode(buf io.Reader, flags interface{}) (err error) {
	for i := 0; i < 8; i++ {
		err = struc.ignored[i].Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
		if err != nil {
			return err
		}
	}

	err = struc.loginRequest.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	for i := 0; i < 2; i++ {
		err = struc.ServerISAACSeed[i].Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
		if err != nil {
			return err
		}
	}

	return err
}

type OutboundUpdateHandshake struct {
	ignored [8]encoding.Uint8
}

func (struc *OutboundUpdateHandshake) Encode(buf io.Writer, flags interface{}) (err error) {
	for i := 0; i < 8; i++ {
		err = struc.ignored[i].Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
		if err != nil {
			return err
		}
	}

	return err
}

func (struc *OutboundUpdateHandshake) Decode(buf io.Reader, flags interface{}) (err error) {
	for i := 0; i < 8; i++ {
		err = struc.ignored[i].Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
		if err != nil {
			return err
		}
	}

	return err
}
