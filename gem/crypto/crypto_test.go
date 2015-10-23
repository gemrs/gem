package crypto

import (
	"testing"
)

func TestRoundTrip(t *testing.T) {
	key, err := GeneratePrivateKey(512)
	if err != nil {
		t.Fatal(err)
	}

	text := "THIS IS SOME TEXT!"

	encrypted := key.Encrypt([]byte(text))
	decrypted := string(key.Decrypt(encrypted))

	if decrypted != text {
		t.Errorf("roundtrip text mismatched")
	}
}
