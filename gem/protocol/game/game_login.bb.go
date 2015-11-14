// Generated by bbc; DO NOT EDIT
package game

import (
	"io"

	"github.com/gemrs/gem/gem/encoding"
)

type InboundLoginBlock struct {
	LoginType       encoding.Uint8
	LoginLen        encoding.Uint8
	Magic           encoding.Uint8
	Revision        encoding.Uint16
	MemType         encoding.Uint8
	ArchiveCRCs     [9]encoding.Uint32
	SecureBlockSize encoding.Uint8
}

func (struc *InboundLoginBlock) Encode(buf io.Writer, flags interface{}) (err error) {
	err = struc.LoginType.Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	err = struc.LoginLen.Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	err = struc.Magic.Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	err = struc.Revision.Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	err = struc.MemType.Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	for i := 0; i < 9; i++ {
		err = struc.ArchiveCRCs[i].Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
		if err != nil {
			return err
		}
	}

	err = struc.SecureBlockSize.Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	return err
}

func (struc *InboundLoginBlock) Decode(buf io.Reader, flags interface{}) (err error) {
	err = struc.LoginType.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	err = struc.LoginLen.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	err = struc.Magic.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	err = struc.Revision.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	err = struc.MemType.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	for i := 0; i < 9; i++ {
		err = struc.ArchiveCRCs[i].Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
		if err != nil {
			return err
		}
	}

	err = struc.SecureBlockSize.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	return err
}

type InboundSecureLoginBlock struct {
	Magic     encoding.Uint8
	ISAACSeed [4]encoding.Uint32
	ClientUID encoding.Uint32
	Username  encoding.JString
	Password  encoding.JString
}

func (struc *InboundSecureLoginBlock) Encode(buf io.Writer, flags interface{}) (err error) {
	err = struc.Magic.Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	for i := 0; i < 4; i++ {
		err = struc.ISAACSeed[i].Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
		if err != nil {
			return err
		}
	}

	err = struc.ClientUID.Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	err = struc.Username.Encode(buf, 0)
	if err != nil {
		return err
	}

	err = struc.Password.Encode(buf, 0)
	if err != nil {
		return err
	}

	return err
}

func (struc *InboundSecureLoginBlock) Decode(buf io.Reader, flags interface{}) (err error) {
	err = struc.Magic.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	for i := 0; i < 4; i++ {
		err = struc.ISAACSeed[i].Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
		if err != nil {
			return err
		}
	}

	err = struc.ClientUID.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	err = struc.Username.Decode(buf, 0)
	if err != nil {
		return err
	}

	err = struc.Password.Decode(buf, 0)
	if err != nil {
		return err
	}

	return err
}

type OutboundLoginResponse struct {
	Response encoding.Uint8
	Rights   encoding.Uint8
	Flagged  encoding.Uint8
}

func (struc *OutboundLoginResponse) Encode(buf io.Writer, flags interface{}) (err error) {
	err = struc.Response.Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	err = struc.Rights.Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	err = struc.Flagged.Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	return err
}

func (struc *OutboundLoginResponse) Decode(buf io.Reader, flags interface{}) (err error) {
	err = struc.Response.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	err = struc.Rights.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	err = struc.Flagged.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	return err
}

type OutboundLoginResponseUnsuccessful struct {
	Response encoding.Uint8
}

func (struc *OutboundLoginResponseUnsuccessful) Encode(buf io.Writer, flags interface{}) (err error) {
	err = struc.Response.Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	return err
}

func (struc *OutboundLoginResponseUnsuccessful) Decode(buf io.Reader, flags interface{}) (err error) {
	err = struc.Response.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return err
	}

	return err
}
