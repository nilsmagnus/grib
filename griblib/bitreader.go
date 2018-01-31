package griblib

import (
	//"fmt"
	"bufio"
	"io"
)

// BitReader is undocomented
type BitReader struct {
	reader io.ByteReader
	byte   byte
	offset byte
}

func newReader(r io.ByteReader) *BitReader {
	return &BitReader{r, 0, 0}
}

func (r *BitReader) readBit() (bool, error) {
	if r.offset == 8 {
		r.offset = 0
	}
	if r.offset == 0 {
		var err error
		if r.byte, err = r.reader.ReadByte(); err != nil {
			return false, err
		}
	}
	bit := (r.byte & (0x80 >> r.offset)) != 0
	r.offset++
	return bit, nil
}

func (r *BitReader) readUint(nbits int) (uint64, error) {
	var result uint64
	for i := nbits - 1; i >= 0; i-- {
		bit, err := r.readBit()

		if err != nil {
			return 0, err
		}
		if bit {
			result |= 1 << uint(i)
		}
	}

	return result, nil
}

func (r *BitReader) readInt(nbits int) (int64, error) {
	var result int64
	var negative int64 = 1
	for i := nbits - 1; i >= 0; i-- {
		bit, err := r.readBit()

		if err != nil {
			return 0, err
		}
		if i == (nbits-1) && bit {
			negative = -1
		} else if bit {
			result |= 1 << uint(i)
		}
	}
	return negative * result, nil
}

func (r *BitReader) readUintsBlock(bits int, count int) ([]uint64, error) {
	//fmt.Println("Reading", bits, "bits", count, "x")
	data := make([]uint64, count)
	var err error

	if bits != 0 {
		for i := 0; i != count; i++ {
			data[i], err = r.readUint(bits)
			if err != nil {
				return data, err
			}

			//fmt.Println(data[i])
		}

		// if we are not fitting last byte seek to byte end
		//rest := (bits * count) % 8
		//if rest != 0 {
		//	r.offset += byte(8 - int64(rest))
		//}
		r.offset = 0

	}
	return data, nil
}

func (r *BitReader) readIntsBlock(bits int, count int) ([]int64, error) {
	//fmt.Println("Reading", bits, "bits", count, "x")
	data := make([]int64, count)
	var err error

	if bits != 0 {
		for i := 0; i != count; i++ {
			data[i], err = r.readInt(bits)
			if err != nil {
				return data, err
			}
			//fmt.Println(data[i])
		}

	}
	return data, nil
}

////////////////////////////////////////////////////////////////////////////////////

// bitReader wraps an io.Reader and provides the ability to read values,
// bit-by-bit, from it. Its Read* methods don't return the usual error
// because the error handling was verbose. Instead, any error is kept and can
// be checked afterwards.
type bitReader struct {
	r    io.ByteReader
	n    uint64
	bits uint
	err  error
}

// newBitReader returns a new bitReader reading from r. If r is not
// already an io.ByteReader, it will be converted via a bufio.Reader.
func newBitReader(r io.Reader) bitReader {
	byter, ok := r.(io.ByteReader)
	if !ok {
		byter = bufio.NewReader(r)
	}
	return bitReader{r: byter}
}

func (br *bitReader) incrByte() {
	br.bits = 0
}

// readBits64 reads the given number of bits and returns them in the
// least-significant part of a uint64. In the event of an error, it returns 0
// and the error can be obtained by calling Err().
func (br *bitReader) readBits64(bits uint) (n uint64) {
	for bits > br.bits {
		b, err := br.r.ReadByte()
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
		if err != nil {
			br.err = err
			return 0
		}
		br.n <<= 8
		br.n |= uint64(b)
		br.bits += 8
	}

	// br.n looks like this (assuming that br.bits = 14 and bits = 6):
	// Bit: 111111
	//      5432109876543210
	//
	//         (6 bits, the desired output)
	//        |-----|
	//        V     V
	//      0101101101001110
	//        ^            ^
	//        |------------|
	//           br.bits (num valid bits)
	//
	// This the next line right shifts the desired bits into the
	// least-significant places and masks off anything above.
	n = (br.n >> (br.bits - bits)) & ((1 << bits) - 1)
	br.bits -= bits
	return
}

func (br *bitReader) readIntsBlock(bits int, count int, compensateByte bool) ([]int64, error) {
	//fmt.Println("Reading", bits, "bits", count, "x")
	data := make([]int64, count)

	if bits != 0 {
		for i := 0; i != count; i++ {
			data[i] = int64(br.readBits64(uint(bits)))
			//fmt.Println(data[i])
		}

		if compensateByte {
			// if we are not fitting last byte seek to byte end
			//rest := (bits * count) % 8
			//if rest != 0 {
			//	r.offset += byte(8 - int64(rest))
			//}
			br.n = 0
		}
	}
	return data, nil
}

func (br *bitReader) readBits(bits uint) (n int) {
	n64 := br.readBits64(bits)
	return int(n64)
}

func (br *bitReader) readBit() bool {
	n := br.readBits(1)
	return n != 0
}

func (br *bitReader) tryReadBit() (bit byte, ok bool) {
	if br.bits > 0 {
		br.bits--
		return byte(br.n>>br.bits) & 1, true
	}
	return 0, false
}

func (br *bitReader) Err() error {
	return br.err
}
