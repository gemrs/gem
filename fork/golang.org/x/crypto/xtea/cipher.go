// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package xtea implements XTEA encryption, as defined in Needham and Wheeler's
// 1997 technical report, "Tea extensions."
package xtea

// For details, see http://www.cix.co.uk/~klockstone/xtea.pdf

import "strconv"

// The XTEA block size in bytes.
const BlockSize = 8

// A Cipher is an instance of an XTEA cipher using a particular key.
// table contains a series of precalculated values that are used each round.
type Cipher struct {
	key [4]uint32
}

type KeySizeError int

func (k KeySizeError) Error() string {
	return "crypto/xtea: invalid key size " + strconv.Itoa(int(k))
}

// NewCipher creates and returns a new Cipher.
// The key argument should be the XTEA key.
// XTEA only supports 128 bit (16 byte) keys.
func NewCipher(key []uint32) (*Cipher, error) {
	k := len(key)
	switch k {
	default:
		return nil, KeySizeError(k)
	case 4:
		break
	}

	c := new(Cipher)
	copy(c.key[0:4], key[0:4])

	return c, nil
}

// BlockSize returns the XTEA block size, 8 bytes.
// It is necessary to satisfy the Block interface in the
// package "crypto/cipher".
func (c *Cipher) BlockSize() int { return BlockSize }

// Decrypt decrypts the 8 byte buffer src using the key k and stores the result in dst.
func (c *Cipher) Decrypt(dst, src []byte) { decryptBlock(c, dst, src) }
