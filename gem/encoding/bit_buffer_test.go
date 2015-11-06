package encoding

import (
	"fmt"
	"testing"
)

func binaryDump(s []byte) string {
	str := ""
	for _, b := range s {
		str = fmt.Sprintf("%v%b ", str, b)
	}
	return str
}

func TestBitWrite1(t *testing.T) {
	buffer := NewBuffer()

	bits := NewBitBuffer(buffer)
	bits.WriteBit(true)
	bits.Close()

	b := buffer.Bytes()
	if len(b) != 1 {
		t.Error("byte size mismatch: %v", len(b))
	}

	if b[0] != (1 << 7) {
		t.Error("bit write mismatch: %v", binaryDump(b))
	}
}

func TestBitWrite2(t *testing.T) {
	buffer := NewBuffer()

	bits := NewBitBuffer(buffer)
	bits.Write(3, 7) // write three 1s
	bits.Write(3, 0) // write three 0s
	bits.Write(2, 2) // write 10
	bits.Close()

	b := buffer.Bytes()
	if len(b) != 1 {
		t.Error("byte size mismatch: %v", len(b))
	}

	if b[0] != 0xE2 {
		t.Error("bit write mismatch: %v", binaryDump(b))
	}
}

func TestBitWrite3(t *testing.T) {
	buffer := NewBuffer()

	bits := NewBitBuffer(buffer)
	bits.Write(11, 0x7FF) // write 11 1s
	bits.Close()

	b := buffer.Bytes()
	if len(b) != 2 {
		t.Error("byte size mismatch: %v", len(b))
	}

	if b[0] != 0xFF ||
		b[1] != 0xE0 {
		t.Error("bit write mismatch: %v", binaryDump(b))
	}
}

func TestBitWrite4(t *testing.T) {
	buffer := NewBuffer()

	buffer.WriteByte(0x0F)
	bits := NewBitBuffer(buffer)
	bits.Write(11, 0x7FF) // write 11 1s
	bits.Close()
	buffer.WriteByte(0xF0)

	b := buffer.Bytes()
	if len(b) != 4 {
		t.Error("byte size mismatch: %v", len(b))
	}

	if b[0] != 0x0F ||
		b[1] != 0xFF ||
		b[2] != 0xE0 ||
		b[3] != 0xF0 {
		t.Error("bit write mismatch: %v", binaryDump(b))
	}
}
