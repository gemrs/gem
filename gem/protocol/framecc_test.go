package protocol

import (
	"bytes"
	"encoding/hex"
	"reflect"
	"testing"

	"github.com/sinusoids/gem/gem/encoding"
)

//go:generate framecc test_frame.frame test_frame.gen.go

var testFrame = TestFrame{
	Message:  "test string",
	Values8:  [...]encoding.Int8{20, 88, 100, 255},
	Values16: [...]encoding.Int16{2000, 30000},
	Struc1: EmbeddedStruct{
		A: 60000,
		B: 60000,
		C: 60000,
	},
	Struc2: [...]EmbeddedStruct{
		{
			A: 1234,
			B: 2345,
			C: 3456,
		},
		{
			A: 4567,
			B: 5678,
			C: 6789,
		},
	},
}

func bufferToHex(buf *bytes.Buffer) string {
	return hex.EncodeToString(buf.Bytes())
}

func hexToBuffer(buf string) *bytes.Buffer {
	data, err := hex.DecodeString(buf)
	if err != nil {
		panic(err)
	}

	return bytes.NewBuffer(data)
}
l
func TestRoundTrip(t *testing.T) {
	buffer := bytes.NewBuffer([]byte{})

	if err := testFrame.Encode(buffer, 0); err != nil {
		t.Errorf("%v", err)
	}

	buffer = bytes.NewBuffer(buffer.Bytes())

	var frame TestFrame
	if err := frame.Decode(buffer, 0); err != nil {
		t.Errorf("%v", err)
	}

	if !reflect.DeepEqual(frame, testFrame) {
		t.Errorf("Decoded data mismatch: \n%#v\n%#v", testFrame, frame)
	}
}
