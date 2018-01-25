// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
	Implementation adapted from Needham and Wheeler's paper:
	http://www.cix.co.uk/~klockstone/xtea.pdf

	A precalculated look up table is used during encryption/decryption for values that are based purely on the key.
*/

package xtea

// XTEA is based on 64 rounds.
// .. but we want 32
const numRounds = 32

// blockToUint32 reads an 8 byte slice into two uint32s.
// The block is treated as big endian.
func blockToUint32(src []byte) (uint32, uint32) {
	r0 := uint32(src[0])<<24 | uint32(src[1])<<16 | uint32(src[2])<<8 | uint32(src[3])
	r1 := uint32(src[4])<<24 | uint32(src[5])<<16 | uint32(src[6])<<8 | uint32(src[7])
	return r0, r1
}

// uint32ToBlock writes two uint32s into an 8 byte data block.
// Values are written as big endian.
func uint32ToBlock(v0, v1 uint32, dst []byte) {
	dst[0] = byte(v0 >> 24)
	dst[1] = byte(v0 >> 16)
	dst[2] = byte(v0 >> 8)
	dst[3] = byte(v0)
	dst[4] = byte(v1 >> 24)
	dst[5] = byte(v1 >> 16)
	dst[6] = byte(v1 >> 8)
	dst[7] = byte(v1 >> 0)
}

// decryptBlock decrypt a single 8 byte block using XTEA.
func decryptBlock(c *Cipher, dst, src []byte) {
	v0, v1 := blockToUint32(src)

	delta := uint32(0x9E3779B9)
	sum := delta * numRounds

	// Two rounds of XTEA applied per loop
	for i := 0; i < numRounds; i++ {
		v1 -= ((v0<<4 ^ v0>>5) + v0) ^ (sum + c.key[(sum>>11)&3])
		sum -= delta
		v0 -= ((v1<<4 ^ v1>>5) + v1) ^ (sum + c.key[sum&3])
	}

	uint32ToBlock(v0, v1, dst)
}
