package encoding

import (
	"errors"
	"io"
	"sync"
)

type ReadWriter interface {
	Reader
	Writer
	io.Seeker
}

// Buffer is similar to bytes.Buffer, but allows us to save and recall the read pointer,
// through the use of Try. Suitable for a long-living socket read buffer, as Trim allows
// us to discard and garbage collect data we've already dealt with.
type Buffer struct {
	s []byte

	i int
	m sync.Mutex

	readFromBuffer []byte

	srcReader  io.Reader
	destWriter io.Writer
}

type ReadFunc func(b *Buffer) error

func NewBuffer() *Buffer {
	return &Buffer{
		s:              make([]byte, 0),
		i:              0,
		readFromBuffer: make([]byte, 512),
	}
}

func WrapReader(r io.Reader) Reader {
	buf := NewBuffer()
	buf.srcReader = r
	return buf
}

func WrapWriter(w io.Writer) Writer {
	buf := NewBuffer()
	buf.destWriter = w
	return buf
}

func NewBufferBytes(s []byte) *Buffer {
	buffer := NewBuffer()
	buffer.s = append(buffer.s, s...)
	return buffer
}

func (b *Buffer) Len() int {
	b.m.Lock()
	defer b.m.Unlock()

	return len(b.s[b.i:])
}

// Trim discards all data before the current read pointer.
// blocks until we get a lock
func (b *Buffer) Trim() {
	b.m.Lock()
	defer b.m.Unlock()

	// perform a copy, so that the old array (and the discarded data) can be garbage collected
	oldSlice := b.s
	b.s = make([]byte, len(oldSlice)-b.i)
	copy(b.s, oldSlice[b.i:])
	b.i = 0
}

func (b *Buffer) Pos() int {
	b.m.Lock()
	defer b.m.Unlock()

	return b.i
}

// Try saves the current position, calls cb, and if cb returns an error, restores the previous position
// locks the buffer to trimming, to ensure we can always pop back to the original position
// since the trim mutex is locked until cb returns, deadlock can occur with incorrect usage
func (b *Buffer) Try(cb ReadFunc) error {
	b.m.Lock()
	defer b.m.Unlock()

	oldPtr := b.i
	err := cb(b)
	if err != nil {
		b.i = oldPtr
	}
	return err
}

func (b *Buffer) Read(p []byte) (n int, err error) {
	if b.i >= len(b.s) {
		// Buffer is empty
		if len(p) == 0 {
			return
		}

		// If we're wrapping a reader, buffer from it
		if b.srcReader != nil {
			buf := make([]byte, len(p))
			_, err := b.srcReader.Read(buf)
			if err != nil {
				return 0, err
			}
			b.s = append(b.s, buf...)
		} else {
			return 0, io.EOF
		}
	}
	n = copy(p, b.s[b.i:])
	b.i += n
	return n, nil
}

func (b *Buffer) ReadByte() (c byte, err error) {
	s := make([]byte, 1)
	n, err := b.Read(s)
	if err != nil {
		return 0, err
	}
	if n != 1 {
		return 0, io.EOF
	}
	return s[0], nil
}

func (b *Buffer) ReadFrom(r io.Reader) (int64, error) {
	n, err := r.Read(b.readFromBuffer)
	if n > 0 {
		m, err2 := b.Write(b.readFromBuffer[:n])
		if m != n {
			panic("Buffer.ReadFrom: couldn't write entire buffer")
		}
		if err2 != nil {
			panic("Buffer.ReadFrom: error appending to buffer")
		}
	}
	return int64(n), err
}

func (b *Buffer) Write(p []byte) (n int, err error) {
	b.m.Lock()
	defer b.m.Unlock()

	if b.destWriter != nil {
		// If we're piping to a Writer...
		i, err := b.destWriter.Write(p)
		b.s = append(b.s, p[:i]...)
		return i, err
	}

	b.s = append(b.s, p...)
	return len(p), nil
}

func (b *Buffer) WriteByte(c byte) error {
	_, err := b.Write([]byte{c})
	return err
}

func (b *Buffer) WriteTo(w io.Writer) (int64, error) {
	n, err := w.Write(b.s[b.i:])
	b.i += n
	return int64(n), err
}

func (b *Buffer) Peek(n int) ([]byte, error) {
	if b.i+n > len(b.s) {
		return nil, io.EOF
	}

	data := b.s[b.i : b.i+n]
	copied := make([]byte, n)
	i := copy(copied, data)
	if i != n {
		return nil, io.EOF
	}

	return copied, nil
}

func (b *Buffer) Seek(offset int64, whence int) (int64, error) {
	var abs int64
	switch whence {
	case 0:
		abs = offset
	case 1:
		abs = int64(b.i) + offset
	case 2:
		abs = int64(len(b.s)) + offset
	default:
		return 0, errors.New("encoding.Buffer.Seek: invalid whence")
	}
	if abs < 0 {
		return 0, errors.New("encoding.Buffer.Seek: negative position")
	}
	if int(abs) > len(b.s) {
		return 0, errors.New("encoding.Buffer.Seek: out of bounds")
	}
	b.i = int(abs)
	return abs, nil
}

func (b *Buffer) Bytes() []byte {
	return b.s
}
