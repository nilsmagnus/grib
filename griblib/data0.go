package griblib

import (
	"io"
	"math"

	"github.com/nilsmagnus/grib/internal/reader"
)

const INT_MAX = 9223372036854775807

// Data0 is a Grid point data - simple packing
// http://www.nco.ncep.noaa.gov/pmb/docs/grib2/grib2_doc/grib2_temp5-0.shtml
//
//	| Octet Number | Content
//	-----------------------------------------------------------------------------------------
//	| 12-15	     | Reference value (R) (IEEE 32-bit floating-point value)
//	| 16-17	     | Binary scale factor (E)
//	| 18-19	     | Decimal scale factor (D)
//	| 20	         | Number of bits used for each packed value for simple packing, or for each
//	|              | group reference value for complex packing or spatial differencing
//	| 21           | Type of original field values
//	|              |    - 0 : Floating point
//	|              |    - 1 : Integer
//	|              |    - 2-191 : reserved
//	|              |    - 192-254 : reserved for Local Use
//	|              |    - 255 : missing
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

func (template Data0) scaleFunc() func(uintValue int64) float64 {
	ref, scale := template.getRefScale()
	return func(value int64) float64 {
		signed := int64(value)
		return ref + float64(signed)*scale
	}
}

// ParseData0 parses data0 struct from the reader into the an array of floating-point values
func ParseData0(dataReader io.Reader, dataLength int, template *Data0) ([]float64, error) {

	fld := []float64{}

	if dataLength == 0 {
		return fld, nil
	}

	scaleStrategy := template.scaleFunc()

	bitReader, err := reader.New(dataReader, dataLength)
	if err != nil {
		return fld, err
	}

	dataSize := int64(math.Floor(
		float64(8*dataLength) / float64(template.Bits),
	))

	uintDataSlice, errRead := bitReader.ReadUintsBlock(int(template.Bits), dataSize, false)
	if errRead != nil {
		return []float64{}, errRead
	}

	for _, uintValue := range uintDataSlice {
		fld = append(fld, scaleStrategy(int64(uintValue)))
	}

	return fld, nil
}
