package reader_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/nilsmagnus/grib/internal/mocks"
	"github.com/nilsmagnus/grib/internal/reader"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	t.Run("no error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		rd := mocks.NewMockReader(ctrl)
		rd.EXPECT().Read(gomock.Any()).AnyTimes().Return(42, nil)

		bitReader, err := reader.New(rd, 42)
		assert.NoError(t, err)
		assert.NotNil(t, bitReader)
	})

	t.Run("not enough data to read", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		rd := mocks.NewMockReader(ctrl)
		rd.EXPECT().Read(gomock.Any()).AnyTimes().Return(0, io.EOF)

		bitReader, err := reader.New(rd, 42)
		assert.Error(t, err)
		assert.Nil(t, bitReader)
	})

	t.Run("unexpected error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		rd := mocks.NewMockReader(ctrl)
		rd.EXPECT().Read(gomock.Any()).AnyTimes().Return(0, io.ErrClosedPipe)

		bitReader, err := reader.New(rd, 42)
		assert.Error(t, err)
		assert.Nil(t, bitReader)
	})
}

func TestReadInt(t *testing.T) {
	t.Run("ok - positive number", func(t *testing.T) {
		bitReader, err := reader.New(bytes.NewBuffer([]byte{0x28, 0xE5, 0x2B}), 3)
		require.NoError(t, err)
		require.NotNil(t, bitReader)

		integer, err := bitReader.ReadInt(10)
		assert.NoError(t, err)
		assert.Equal(t, int64(163), integer)
	})

	t.Run("ok - negative number", func(t *testing.T) {
		bitReader, err := reader.New(bytes.NewBuffer([]byte{0xA8, 0xE5, 0x2B}), 3)
		require.NoError(t, err)
		require.NotNil(t, bitReader)

		integer, err := bitReader.ReadInt(10)
		assert.NoError(t, err)
		assert.Equal(t, int64(-163), integer)
	})

	t.Run("not enough bits", func(t *testing.T) {
		bitReader, err := reader.New(bytes.NewBuffer([]byte{0xA8, 0xE5, 0x2B}), 3)
		require.NoError(t, err)
		require.NotNil(t, bitReader)

		_, err = bitReader.ReadInt(64)
		assert.Error(t, err)
	})
}

func TestReadUintsBlock(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		bitReader, err := reader.New(bytes.NewBuffer([]byte{0xA8, 0xE5, 0x2B, 0xf4}), 4)
		require.NoError(t, err)
		require.NotNil(t, bitReader)

		data, err := bitReader.ReadUintsBlock(4, 8, false)
		assert.NoError(t, err)
		assert.Len(t, data, 8)
		assert.Equal(t, []uint64([]uint64{10, 8, 14, 5, 2, 11, 15, 4}), data)
	})

	t.Run("not enough data", func(t *testing.T) {
		bitReader, err := reader.New(bytes.NewBuffer([]byte{0xA8, 0xE5, 0x2B, 0xf4}), 4)
		require.NoError(t, err)
		require.NotNil(t, bitReader)

		_, err = bitReader.ReadUintsBlock(4, 10, false)
		assert.Error(t, err)
	})
}
