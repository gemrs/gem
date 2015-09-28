package main

import (
	"flag"
	"fmt"
	"os"

	"gem/crypto"
)

var bits = flag.Int("bits", 512, "bitsize of the generated key")
var print = flag.Bool("print", false, "prints the given private key")
var key = flag.String("key", "", "specifies the key file to print/generate")

func main() {
	flag.Parse()
	if *key == "" {
		flag.Usage()
		os.Exit(1)
	}

	if *print {
		os.Exit(printKey(*key))
	} else {
		os.Exit(generateKey(*bits, *key))
	}
}

func printKey(path string) int {
	key, err := crypto.LoadPrivateKey(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to load key %v: %v\n", path, err)
		return 1
	}

	err = key.Validate()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to validate key %v: %v\n", path, err)
		return 1
	}

	fmt.Printf("Private Exponent: %v\n", key.D)
	fmt.Printf("Public Exponent: %v\n", key.PublicKey.E)
	fmt.Printf("Modulus: %v\n", key.PublicKey.N)
	return 0
}

func generateKey(bits int, path string) int {
	key, err := crypto.GeneratePrivateKey(bits)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to generate key %v: %v\n", path, err)
		return 1
	}

	err = key.Store(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to store key %v: %v\n", path, err)
		return 1
	}

	fmt.Printf("Stored private key: %v\n", path)

	return printKey(path)
}
