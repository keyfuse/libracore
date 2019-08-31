// libracore
//
// Copyright 2019 by KeyFuse Labs
//
// GPLv3 License

package xbase

import (
	"fmt"
)

// Buffer -- a buffer writer/reader for bitcoin.
type Buffer struct {
	seek   int
	length int
	data   []byte
}

// NewBuffer -- creates a new buffer.
func NewBuffer() *Buffer {
	return &Buffer{}
}

// NewBufferReader -- creates a new buffer with data initialize.
func NewBufferReader(data []byte) *Buffer {
	return &Buffer{
		data:   data,
		length: len(data),
	}
}

func (b *Buffer) check(l int) error {
	if (b.seek + l) > b.length {
		return fmt.Errorf("error.seek[%v].required[%v].length[%v]", b.seek, l, b.length)
	}
	return nil
}

// WriteU8 -- write uint8(byte) to buffer.
func (b *Buffer) WriteU8(v uint8) {
	b.data = append(b.data, v)
	b.length++
}

// ReadU8 -- read a byte from the buffer.
func (b *Buffer) ReadU8() (uint8, error) {
	if err := b.check(1); err != nil {
		return 0, err
	}
	v := b.data[b.seek]
	b.seek++
	return v, nil
}

// WriteBool -- write bool.
func (b *Buffer) WriteBool(v bool) {
	var bv byte = 0x00
	if v {
		bv = 0x01
	}
	b.data = append(b.data, bv)
	b.length++
}

// ReadBool -- read bool.
func (b *Buffer) ReadBool() (bool, error) {
	var v bool

	if err := b.check(1); err != nil {
		return v, err
	}
	data := b.data[b.seek]
	if data == 0x01 {
		v = true
	}
	b.seek++
	return v, nil
}

// WriteU16 -- write uint16 with little-endian to the buffer.
func (b *Buffer) WriteU16(v uint32) {
	b.data = append(b.data, byte(v))
	b.data = append(b.data, byte(v>>8))
	b.length += 2
}

// ReadU16 -- read uint16 from the buffer which with the little-endian byteorder.
func (b *Buffer) ReadU16() (uint32, error) {
	if err := b.check(2); err != nil {
		return 0, err
	}

	v := uint32(b.data[b.seek]) |
		uint32(b.data[b.seek+1])<<8
	b.seek += 2
	return v, nil
}

// WriteU32 -- write uint32 with little-endian to the buffer.
func (b *Buffer) WriteU32(v uint32) {
	b.data = append(b.data, byte(v))
	b.data = append(b.data, byte(v>>8))
	b.data = append(b.data, byte(v>>16))
	b.data = append(b.data, byte(v>>24))
	b.length += 4
}

// ReadU32 -- read uint32 from the buffer which with the little-endian byteorder.
func (b *Buffer) ReadU32() (uint32, error) {
	if err := b.check(4); err != nil {
		return 0, err
	}

	v := uint32(b.data[b.seek]) |
		uint32(b.data[b.seek+1])<<8 |
		uint32(b.data[b.seek+2])<<16 |
		uint32(b.data[b.seek+3])<<24
	b.seek += 4
	return v, nil
}

// WriteU64 -- write uint64 to the buffer with little-endian byteorder.
func (b *Buffer) WriteU64(v uint64) {
	b.data = append(b.data, byte(v))
	b.data = append(b.data, byte(v>>8))
	b.data = append(b.data, byte(v>>16))
	b.data = append(b.data, byte(v>>24))
	b.data = append(b.data, byte(v>>32))
	b.data = append(b.data, byte(v>>40))
	b.data = append(b.data, byte(v>>48))
	b.data = append(b.data, byte(v>>56))
	b.length += 8
}

// ReadU64 -- read uint64 from the buffer which with the little-endian byteorder.
func (b *Buffer) ReadU64() (uint64, error) {
	if err := b.check(8); err != nil {
		return 0, err
	}

	v := uint64(b.data[b.seek]) |
		uint64(b.data[b.seek+1])<<8 |
		uint64(b.data[b.seek+2])<<16 |
		uint64(b.data[b.seek+3])<<24 |
		uint64(b.data[b.seek+4])<<32 |
		uint64(b.data[b.seek+5])<<40 |
		uint64(b.data[b.seek+6])<<48 |
		uint64(b.data[b.seek+7])<<56
	b.seek += 8
	return v, nil
}

// WriteBytes -- write bytes to the buffer.
func (b *Buffer) WriteBytes(v []byte) {
	b.data = append(b.data, v...)
	b.length += len(v)
}

// ReadBytes -- read l-bytes from the buffer.
func (b *Buffer) ReadBytes(l int) ([]byte, error) {
	if l == 0 {
		return nil, nil
	}

	if err := b.check(l); err != nil {
		return nil, err
	}
	v := make([]byte, l)
	copy(v, b.data[b.seek:b.seek+l])
	b.seek += l
	return v, nil
}

// Bytes -- returns all the datas in the buffer.
func (b *Buffer) Bytes() []byte {
	return b.data
}

// Seek -- returns the seek postion.
func (b *Buffer) Seek() int {
	return b.seek
}

// Len -- returns the buffer length.
func (b *Buffer) Len() int {
	return b.length
}

// End -- returns whether the seek is end or not.
func (b *Buffer) End() bool {
	return b.seek >= b.length
}

// Reset -- reset the seek, length and the data.
func (b *Buffer) Reset() {
	b.seek = 0
	b.length = 0
	b.data = b.data[0:0]
}
