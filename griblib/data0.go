package griblib

import (
	"bytes"
	"fmt"
	"io"
	"math"
)

//Data0 is a Grid point data - simple packing
// http://www.nco.ncep.noaa.gov/pmb/docs/grib2/grib2_doc/grib2_temp5-0.shtml
//    | Octet Number | Content
//    -----------------------------------------------------------------------------------------
//    | 12-15	     | Reference value (R) (IEEE 32-bit floating-point value)
//    | 16-17	     | Binary scale factor (E)
//    | 18-19	     | Decimal scale factor (D)
//    | 20	         | Number of bits used for each packed value for simple packing, or for each
//    |              | group reference value for complex packing or spatial differencing
//    | 21           | Type of original field values
//    |              |    - 0 : Floating point
//    |              |    - 1 : Integer
//    |              |    - 2-191 : reserved
//    |              |    - 192-254 : reserved for Local Use
//    |              |    - 255 : missing
type Data0 struct {
	Reference    float32 `json:"reference"`
	BinaryScale  uint16  `json:"binaryScale"`
	DecimalScale uint16  `json:"decimalScale"`
	Bits         uint8   `json:"bits"`
	Type         uint8   `json:"type"`
}

func (template Data0) getRefScale() (float64, float64) {
	bscale := math.Pow(2.0, float64(template.BinaryScale))
	dscale := math.Pow(10.0, -float64(template.DecimalScale))

	scale := bscale * dscale
	ref := dscale * float64(template.Reference)

	return ref, scale
}

func (template Data0) scaleFunc() func(uintValue uint64) float64 {
	ref, scale := template.getRefScale()
	return func(value uint64) float64 {
		signed := int64(value)
		return ref + float64(signed)*scale
	}
}

// ParseData0 parses data0 struct from the reader into the an array of floating-point values
func ParseData0(dataReader io.Reader, dataLength int, template *Data0) []float64 {

	fld := []float64{}

	if dataLength == 0 {
		return fld
	}

	rawData := make([]byte, dataLength)
	bytesRead, errRead := dataReader.Read(rawData)
	if errRead != nil {
		panic(errRead)
	}
	fmt.Printf("read: %d\n", bytesRead)

	ref, scale := template.getRefScale()

	buffer := bytes.NewBuffer(rawData)
	bitReader := newReader(buffer)

	dataSize := int(math.Floor(
		float64(8*dataLength) / float64(template.Bits),
	))

	uintDataSlice, errRead := bitReader.readUintsBlock(int(template.Bits), dataSize)
	if errRead != nil {
		panic(errRead)
	}

	for _, uintValue := range uintDataSlice {
		signed := int64(uintValue)
		fld = append(fld, ref+float64(signed)*scale)
	}

	return fld
}
