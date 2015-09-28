package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"os"
)

var ErrInvalidKey = errors.New("invalid private key")

func GeneratePrivateKey(bits int) (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, bits)
}

func LoadPrivateKey(path string) (*rsa.PrivateKey, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(bytes)
	if block == nil {
		return nil, ErrInvalidKey
	}

	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

func StorePrivateKey(path string, key *rsa.PrivateKey) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	pkcs1 := x509.MarshalPKCS1PrivateKey(key)
	block := pem.Block{
		Type:    "GEM PRIVATE KEY",
		Headers: nil,
		Bytes:   pkcs1,
	}

	return pem.Encode(file, &block)
}
