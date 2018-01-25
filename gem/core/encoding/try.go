package encoding

import (
	"fmt"
	"io"
)

func TryEncode(e Encodable, buf io.Writer, flags_ interface{}) (err error) {
	/*defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()*/
	if e == nil {
		return fmt.Errorf("cannot encode nil message")
	}
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
