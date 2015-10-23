package id

import (
	"testing"
)

func TestGenerator(t *testing.T) {
	generator := Generator()
	for i := 0; i < 10; i++ {
		x := <-generator
		if x != i {
			t.Errorf("Generated ID mismatch")
		}
	}
}
