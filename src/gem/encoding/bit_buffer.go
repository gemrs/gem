package encoding

import (
	"io"
)

// BitBuffer provides bit-level access to a ByteWriter
// Currently only write access is supported
type BitBuffer struct {
	writer io.Writer

	/* i is the index into c of the next bit
	   i = 7 is the most significant bit */
	i int
	c byte /* c is the current byte we're writing to */
}

// NewBitBuffer creates a new BitBuffer which
func NewBitBuffer(w io.Writer) *BitBuffer {
	return &BitBuffer{
		writer: w,
		i:      7,
		c:      0,
	}
}

// Write puts nbits bits from value into the buffer
func (b *BitBuffer) Write(nbits int, value uint32) error {
	for i := uint(nbits); i > 0; i-- {
		err := b.WriteBit((value & (1 << (i - 1))) != 0)
		if err != nil {
			return err
		}
	}
	return nil
}

// WriteBit puts a single bit into the buffer
func (b *BitBuffer) WriteBit(v bool) error {
	if v {
		b.c |= (1 << uint(b.i))
	}

	b.i--

	if b.i < 0 {
		b.Flush()
	}
	return nil
}

// Close flushes the current byte and invalidates the BitBuffer
func (b *BitBuffer) Close() {
	if b.i != 7 {
		b.Flush()
	}
	b.writer = nil
}

// Flush flushes the current byte and prepares a new one for writing
func (b *BitBuffer) Flush() error {
	_, err := b.writer.Write([]byte{b.c})
	b.c = 0
	b.i = 7
	return err
}
