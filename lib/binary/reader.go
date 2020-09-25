package binary

import "encoding/binary"

type ByteReader struct {
	Data []byte

	i uint
}

func (b *ByteReader) ReadUint32() (value uint32) {
	value = binary.LittleEndian.Uint32(b.Data[b.i : b.i+4])

	b.i += 4

	return
}

func (b *ByteReader) ReadUint8() (value uint8) {
	value = b.Data[b.i]

	b.i++

	return
}

func (b *ByteReader) Remaining() []uint8 {
	return b.Data[b.i:]
}

