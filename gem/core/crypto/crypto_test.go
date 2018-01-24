package crypto

import (
	"io/ioutil"
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
		t.Error("roundtrip text mismatched")
	}
}

func TestLoadStore(t *testing.T) {
	key, err := GeneratePrivateKey(512)
	if err != nil {
		t.Fatal(err)
	}

	file, err := ioutil.TempFile("", "gem")
	if err != nil {
		t.Error(err)
	}
	file.Close()

	path := file.Name()

	text := "THIS IS SOME TEXT!"

	encrypted := key.Encrypt([]byte(text))

	err = key.Store(path)
	if err != nil {
		t.Error(err)
	}

	key, err = LoadPrivateKey(path)
	if err != nil {
		t.Error(err)
	}

	decrypted := string(key.Decrypt(encrypted))

	if decrypted != text {
		t.Error("roundtrip text mismatched with stored/retrieved key")
	}
}
