// libracore
//
// Copyright 2019 by KeyFuse Labs
//
// GPLv3 License

package xbase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuffer(t *testing.T) {
	writer := NewBuffer()
	writer.WriteU8(0)
	writer.WriteBool(true)
	writer.WriteBytes([]byte{0x01, 0x02})
	writer.WriteU16(65534)
	writer.WriteU32(653400000)
	writer.WriteU64(uint64(43000000000))

	datas := writer.Bytes()
	reader := NewBufferReader(datas)

	// U8.
	{
		v, err := reader.ReadU8()
		assert.Nil(t, err)
		want := uint8(0)
		assert.Equal(t, want, v)
	}

	// bool.
	{
		v, err := reader.ReadBool()
		assert.Nil(t, err)
		want := true
		assert.Equal(t, want, v)
	}

	// bytes.
	{
		v, err := reader.ReadBytes(2)
		assert.Nil(t, err)
		want := []byte{0x01, 0x02}
		assert.Equal(t, want, v)
	}

	// U16.
	{
		end := reader.End()
		assert.False(t, end)

		v, err := reader.ReadU16()
		assert.Nil(t, err)
		want := uint32(65534)
		assert.Equal(t, want, v)
	}

	// Seek, Len.
	{
		seek := reader.Seek()
		assert.Equal(t, 6, seek)

		pos := reader.Len()
		assert.Equal(t, 18, pos)
	}

	// U32.
	{
		end := reader.End()
		assert.False(t, end)

		v, err := reader.ReadU32()
		assert.Nil(t, err)
		want := uint32(653400000)
		assert.Equal(t, want, v)
	}

	// U64.
	{

		v, err := reader.ReadU64()
		assert.Nil(t, err)
		want := uint64(43000000000)
		assert.Equal(t, want, v)
	}

	// Error.
	{
		end := reader.End()
		assert.True(t, end)

		_, err := reader.ReadU8()
		assert.NotNil(t, err)

		_, err = reader.ReadBytes(1)
		assert.NotNil(t, err)

		_, err = reader.ReadU16()
		assert.NotNil(t, err)

		_, err = reader.ReadU32()
		assert.NotNil(t, err)
	}

	// Reset.
	{
		writer.Reset()
	}
}
