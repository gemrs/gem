package encoding

// This huge lump of ugly was transpiled from the original java and fixed up by hand
type Huffman struct {
	masks []int
	bits  []byte
	keys  []int
}

const INT_MIN_VALUE = 0x80000000

func NewHuffman(huffmanTable []byte) (rcvr *Huffman) {
	rcvr = &Huffman{}
	length := len(huffmanTable)
	rcvr.masks = make([]int, length)
	rcvr.bits = huffmanTable
	var3 := make([]int, 33)
	rcvr.keys = make([]int, 8)
	var4 := 0

	for i := 0; i < length; i++ {
		var6 := huffmanTable[i]
		if var6 != 0 {
			var7 := 1 << (32 - var6)
			var8 := var3[var6]
			rcvr.masks[i] = var8
			var var9 int
			var var10 int
			var var11 int
			var var12 int
			if var8&var7 != 0 {
				var9 = var3[var6-1]
			} else {
				var9 = var8 | var7

				for var10 = int(var6 - 1); var10 >= 1; var10-- {
					var11 = var3[var10]
					if var8 != var11 {
						break
					}

					var12 = 1 << uint(32-var10)
					if var11&var12 != 0 {
						var3[var10] = var3[var10-1]
						break
					}

					var3[var10] = var11 | var12
				}
			}

			var3[var6] = var9

			for var10 = int(var6 + 1); var10 <= 32; var10++ {
				if var8 == var3[var10] {
					var3[var10] = var9
				}
			}

			var10 = 0

			for var11 = 0; var11 < int(var6); var11++ {
				var12 = int(uint32(INT_MIN_VALUE) >> uint(var11))
				if var8&var12 != 0 {
					if rcvr.keys[var10] == 0 {
						rcvr.keys[var10] = var4
					}
					var10 = rcvr.keys[var10]
				} else {
					var10++
				}

				if var10 >= len(rcvr.keys) {
					var13 := make([]int, len(rcvr.keys)*2)

					for var14 := 0; var14 < len(rcvr.keys); var14++ {
						var13[var14] = rcvr.keys[var14]
					}

					rcvr.keys = var13
				}

				var12 = int(uint32(var12) >> 1)
			}

			rcvr.keys[var10] = ^i
			if var10 >= var4 {
				var4 = var10 + 1
			}
		}
	}
	return
}

func (rcvr *Huffman) Compress(data []byte) []byte {
	output := make([]byte, len(data)*2)
	size := rcvr.compress(data, 0, len(data), output, 0)
	return output[:size]
}

func (rcvr *Huffman) compress(uncompressed []byte, fromOffset int, toOffset int, outBuffer []byte, var5 int) int {
	key := int64(0)
	bitpos := int64(var5 << 3)
	toOffset += fromOffset
	for ; fromOffset < toOffset; fromOffset++ {
		data := uncompressed[fromOffset] & 255
		mask := int64(rcvr.masks[data])
		bits := int64(rcvr.bits[data])
		if bits == 0 {
			panic("")
		}

		var11 := bitpos >> 3
		var12 := bitpos & 7
		key &= -var12 >> 31
		var13 := (var12+bits-1)>>3 + var11
		var12 += 24
		key |= mask >> uint64(var12)
		outBuffer[var11] = byte(key)
		if var11 < var13 {
			var11++
			var12 -= 8
			key = mask >> uint64(var12)
			outBuffer[var11] = byte(key)
			if var11 < var13 {
				var11++
				var12 -= 8
				key = mask >> uint64(var12)
				outBuffer[var11] = byte(key)
				if var11 < var13 {
					var11++
					var12 -= 8
					key = mask >> uint64(var12)
					outBuffer[var11] = byte(key)
					if var11 < var13 {
						var11++
						var12 -= 8
						key = mask << uint32((32-var12)%32)
						outBuffer[var11] = byte(key)
					}
				}
			}
		}
		bitpos += bits
	}
	return int(uint32(bitpos+7)>>3 - uint32(var5))
}

func (rcvr *Huffman) Decompress(compressed []byte, length int) []byte {
	output := make([]byte, length*2)
	rcvr.decompress(compressed, 0, output, 0, length)
	return output[:length]
}

func (rcvr *Huffman) decompress(compressed []byte, var2 int, decompressed []byte, var4 int, decompressedLength int) int {
	if decompressedLength == 0 {
		return 0
	} else {
		var6 := 0
		decompressedLength += var4
		pos := var2
		for true {
			var8 := int8(compressed[pos])

			if var8 < 0 {
				var6 = rcvr.keys[var6]
			} else {
				var6++
			}

			var var9 int
			var9 = rcvr.keys[var6]
			if var9 < 0 {
				decompressed[var4] = byte(^var9)
				var4++
				if var4 >= decompressedLength {
					break
				}
				var6 = 0
			}

			if var8&64 != 0 {
				var6 = rcvr.keys[var6]
			} else {
				var6++
			}

			var9 = rcvr.keys[var6]
			if (var9) < 0 {
				decompressed[var4] = byte(^var9)
				var4++
				if var4 >= decompressedLength {
					break
				}
				var6 = 0
			}

			if var8&32 != 0 {
				var6 = rcvr.keys[var6]
			} else {
				var6++
			}

			var9 = rcvr.keys[var6]
			if (var9) < 0 {
				decompressed[var4] = byte(^var9)
				var4++
				if var4 >= decompressedLength {
					break
				}
				var6 = 0
			}

			if var8&16 != 0 {
				var6 = rcvr.keys[var6]
			} else {
				var6++
			}

			var9 = rcvr.keys[var6]
			if (var9) < 0 {
				decompressed[var4] = byte(^var9)
				var4++
				if var4 >= decompressedLength {
					break
				}
				var6 = 0
			}

			if var8&8 != 0 {
				var6 = rcvr.keys[var6]
			} else {
				var6++
			}

			var9 = rcvr.keys[var6]
			if (var9) < 0 {
				decompressed[var4] = byte(^var9)
				var4++
				if var4 >= decompressedLength {
					break
				}
				var6 = 0
			}

			if var8&4 != 0 {
				var6 = rcvr.keys[var6]
			} else {
				var6++
			}

			var9 = rcvr.keys[var6]
			if (var9) < 0 {
				decompressed[var4] = byte(^var9)
				var4++
				if var4 >= decompressedLength {
					break
				}
				var6 = 0
			}

			if var8&2 != 0 {
				var6 = rcvr.keys[var6]
			} else {
				var6++
			}

			var9 = rcvr.keys[var6]
			if (var9) < 0 {
				decompressed[var4] = byte(^var9)
				var4++
				if var4 >= decompressedLength {
					break
				}
				var6 = 0
			}

			if var8&1 != 0 {
				var6 = rcvr.keys[var6]
			} else {
				var6++
			}

			var9 = rcvr.keys[var6]
			if (var9) < 0 {
				decompressed[var4] = byte(^var9)
				var4++
				if var4 >= decompressedLength {
					break
				}
				var6 = 0
			}
			pos++
		}
		return pos + 1 - var2
	}
}
