package id

import (
	"testing"
)

func TestGenerator(t *testing.T) {
	generator := Generator(0)
	for i := 0; i < 10; i++ {
		x := <-generator
		if x != i {
			t.Error("Generated ID mismatch")
		}
	}
}
