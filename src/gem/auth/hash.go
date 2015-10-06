package auth

import (
	"gopkg.in/hlandau/passlib.v1"
)

func HashPassword(plain string) string {
	hash, err := passlib.Hash(plain)
	if err != nil {
		panic(err)
	}
	return hash
}
