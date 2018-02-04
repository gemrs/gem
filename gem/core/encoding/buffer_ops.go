package encoding

func (b *Buffer) Get8(flags interface{}) int {
	var tmp8 Int8
	tmp8.Decode(b, flags)
	return int(tmp8)
}

func (b *Buffer) Get16(flags interface{}) int {
	var tmp16 Int16
	tmp16.Decode(b, flags)
	return int(tmp16)
}

func (b *Buffer) Get24(flags interface{}) int {
	var tmp24 Int24
	tmp24.Decode(b, flags)
	return int(tmp24)
}

func (b *Buffer) Get32(flags interface{}) int {
	var tmp32 Int32
	tmp32.Decode(b, flags)
	return int(tmp32)
}

func (b *Buffer) Get64(flags interface{}) int64 {
	var tmp64 Int64
	tmp64.Decode(b, flags)
	return int64(tmp64)
}

func (b *Buffer) GetU8(flags interface{}) int {
	var tmp8 Uint8
	tmp8.Decode(b, flags)
	return int(tmp8)
}

func (b *Buffer) GetU16(flags interface{}) int {
	var tmp16 Uint16
	tmp16.Decode(b, flags)
	return int(tmp16)
}

func (b *Buffer) GetU24(flags interface{}) int {
	var tmp24 Uint24
	tmp24.Decode(b, flags)
	return int(tmp24)
}

func (b *Buffer) GetU32(flags interface{}) int {
	var tmp32 Uint32
	tmp32.Decode(b, flags)
	return int(tmp32)
}

func (b *Buffer) GetU64(flags interface{}) uint64 {
	var tmp64 Uint64
	tmp64.Decode(b, flags)
	return uint64(tmp64)
}

func (b *Buffer) GetBytes(length int) []byte {
	var tmpBytes Bytes
	tmpBytes.Decode(b, length)
	return []byte(tmpBytes)
}

func (b *Buffer) GetStringZ() string {
	var tmpString String
	tmpString.Decode(b, 0)
	return string(tmpString)
}

func (b *Buffer) Put8(i int, flags interface{}) {
	Int8(i).Encode(b, flags)
}

func (b *Buffer) Put16(i int, flags interface{}) {
	Int16(i).Encode(b, flags)
}

func (b *Buffer) Put24(i int, flags interface{}) {
	Int24(i).Encode(b, flags)
}

func (b *Buffer) Put32(i int, flags interface{}) {
	Int32(i).Encode(b, flags)
}

func (b *Buffer) Put64(i int64, flags interface{}) {
	Int64(i).Encode(b, flags)
}

func (b *Buffer) PutU8(i int, flags interface{}) {
	Uint8(i).Encode(b, flags)
}

func (b *Buffer) PutU16(i int, flags interface{}) {
	Uint16(i).Encode(b, flags)
}

func (b *Buffer) PutU24(i int, flags interface{}) {
	Uint24(i).Encode(b, flags)
}

func (b *Buffer) PutU32(i int, flags interface{}) {
	Uint32(i).Encode(b, flags)
}

func (b *Buffer) PutU64(i uint64, flags interface{}) {
	Uint64(i).Encode(b, flags)
}

func (b *Buffer) PutBytes(data []byte) {
	Bytes(data).Encode(b, len(data))
}

func (b *Buffer) PutStringZ(s string) {
	String(s).Encode(b, 0)
}
