package encoding

import (
	"io"
	"testing"
	"time"
)

var data = []byte{5, 10, 15, 20, 25, 30, 35, 40, 45, 50}

func TestTry(t *testing.T) {
	var err error
	buffer := NewBufferBytes(data)

	// Test that Try works when happy
	pos := buffer.i
	err = buffer.Try(func(b *Buffer) error {
		for i := 0; i < 4; i++ {
			x, err := b.ReadByte()
			if err != nil {
				return err
			}
			if x != data[i] {
				t.Error("data read mismatch: got %v expected %v", x, data[i])
			}
		}
		return nil
	})

	if err != nil {
		t.Error("try returned error: %v", err)
	}

	if buffer.i != (pos + 4) {
		t.Error("position mismatch after reading 4 bytes successfully: got %v expected %v", buffer.i, (pos + 4))
	}

	// Test that Try resets our position on error
	pos = buffer.i
	err = buffer.Try(func(b *Buffer) error {
		for i := 0; i < 10; i++ {
			x, err := b.ReadByte()
			if err != nil {
				return err
			}
			if (i+pos) < len(data) && x != data[i+pos] {
				t.Error("data read mismatch: got %v expected %v", x, data[i+pos])
			}
		}
		return nil
	})

	if err != io.EOF {
		t.Error("try didn't return EOF: %v", err)
	}

	if buffer.i != pos {
		t.Error("position mismatch after reading past EOF: got %v expected %v", buffer.i, pos)
	}

}

func TestCopySemantics(t *testing.T) {
	dataCopy := make([]byte, len(data))
	copy(dataCopy, data)
	buffer := NewBufferBytes(dataCopy)
	for i := range dataCopy {
		dataCopy[i]++
	}

	for _, d := range data {
		b, err := buffer.ReadByte()
		if err != nil {
			t.Error("ReadByte returned error: %v", err)
		}
		if b != d {
			t.Error("original slice was modified")
		}
	}
}

func TestCopySemanticsOnPeek(t *testing.T) {
	buffer := NewBufferBytes(data)

	copied, err := buffer.Peek(2)
	if err != nil {
		t.Error("Peek returned error: %v", err)
	}

	for i := range copied {
		if copied[i] != data[i] {
			t.Error("Peeked data incorrect")
		}
		copied[i]++
	}

	actual := make([]byte, 2)
	_, err = buffer.Read(actual)
	if err != nil {
		t.Error("Read returned error: %v", err)
	}

	for i := range copied {
		if actual[i] == copied[i] {
			t.Error("original slice was modified")
		}
		if actual[i] != data[i] {
			t.Error("original slice was modified")
		}
	}
}

func TestCopySemanticsOnWrite(t *testing.T) {
	dataCopy := make([]byte, len(data))
	copy(dataCopy, data)
	buffer := NewBuffer()
	n, err := buffer.Write(dataCopy)
	if err != nil {
		t.Error("Write returned error: %v", err)
	}
	if n != len(dataCopy) {
		t.Error("Write was partial: %v", err)
	}

	for i := range dataCopy {
		dataCopy[i]++
	}

	for _, d := range data {
		b, err := buffer.ReadByte()
		if err != nil {
			t.Error("ReadByte returned error: %v", err)
		}
		if b != d {
			t.Error("original slice was modified")
		}
	}
}

func TestTrim(t *testing.T) {
	buffer := NewBufferBytes(data)

	l := len(buffer.s)
	trimBytes := len(data) / 2
	err := buffer.Try(func(b *Buffer) error {
		for i := 0; i < trimBytes; i++ {
			x, err := b.ReadByte()
			if err != nil {
				return err
			}
			if x != data[i] {
				t.Error("data read mismatch: got %v expected %v", x, data[i])
			}
		}
		return nil
	})

	if err != nil {
		t.Error("try returned error: %v", err)
	}

	// We don't expect to have trimmed yet
	if len(buffer.s) != l {
		t.Error("data was discarded before trim!")
	}

	buffer.Trim()

	if len(buffer.s) != l-trimBytes {
		t.Error("data wasn't discarded by trim!")
	}

	err = buffer.Try(func(b *Buffer) error {
		for i := 0; i < trimBytes; i++ {
			x, err := b.ReadByte()
			if err != nil {
				return err
			}
			if x != data[i+trimBytes] {
				t.Error("data read mismatch: got %v expected %v", x, data[i+trimBytes])
			}
		}
		return nil
	})

	if err != nil {
		t.Error("try returned error: %v", err)
	}
}

func TestTrimLock(t *testing.T) {
	buffer := NewBufferBytes(data)

	signal := make(chan int, 1)
	go func() {
		err := buffer.Try(func(b *Buffer) error {
			select {
			case <-time.After(5 * time.Second):

			case <-signal:
				t.Error("Trimmed before lock was released")
			}
			return nil
		})

		if err != nil {
			t.Error("try returned error: %v", err)
		}
	}()

	time.Sleep(1 * time.Second) // sleep, to ensure Try has a chance to lock

	// This should block for 4 seconds, then trim
	buffer.Trim()
	signal <- 1

	time.Sleep(1 * time.Second) // sleep, to ensure Try has a chance to pick the signal up
	select {
	case <-signal:
	default:
		t.Error("signal was consumed elsewhere")
	}
}
