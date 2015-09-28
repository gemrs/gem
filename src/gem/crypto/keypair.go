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

type Keypair struct {
	*rsa.PrivateKey
}

func GeneratePrivateKey(bits int) (*Keypair, error) {
	key, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, err
	}

	return &Keypair{key}, nil
}

func LoadPrivateKey(path string) (*Keypair, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(bytes)
	if block == nil {
		return nil, ErrInvalidKey
	}

	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return &Keypair{key}, nil
}

func (key *Keypair) Store(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	pkcs1 := x509.MarshalPKCS1PrivateKey(key.PrivateKey)
	block := pem.Block{
		Type:    "GEM PRIVATE KEY",
		Headers: nil,
		Bytes:   pkcs1,
	}

	return pem.Encode(file, &block)
}
