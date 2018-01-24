package encoding

import "io"

func TryEncode(e Encodable, buf io.Writer, flags_ interface{}) (err error) {
	/*defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()*/
	e.Encode(buf, flags_)
	return nil
}

func TryDecode(e Decodable, buf io.Reader, flags_ interface{}) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	e.Decode(buf, flags_)
	return nil
}
