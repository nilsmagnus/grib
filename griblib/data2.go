package griblib

import (
	"io"
)

//Data2 is a Grid point data - complex packing
// http://www.nco.ncep.noaa.gov/pmb/docs/grib2/grib2_doc/grib2_temp5-2.shtml
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
//    | 22           | Group splitting method used (see Code Table 5.4)
//    |              |    - 0 : Row by Row Splitting
//    |              |    - 1 : General Group Splitting
//    |              |    - 2-191 : Reserved
//    |              |    - 192-254 : Reserved for Local Use
//    |              |    - 255 : Missing
//    | 23           | Missing value management used (see Code Table 5.5)
//    |              |    - 0 : No explicit missing values included within the data values
//    |              |    - 1 : Primary missing values included within the data values
//    |              |    - 2 : Primary and secondary missing values included within the data values
//    |              |    - 3-191 : Reserved
//    |              |    - 192-254 : Reserved for Local Use
//    |              |    - 255 : Missing
//    | 24-27        | Primary missing value substitute
//    | 28-31        | Secondary missing value substitute
//    | 32-35        | NG ― number of groups of data values into which field is split
//    | 36           | Reference for group widths (see Note 12)
//    | 37           | Number of bits used for the group widths (after the reference value in octet 36
//    |              | has been removed)
//    | 38-41        | Reference for group lengths (see Note 13)
//    | 42           | Length increment for the group lengths (see Note 14)
//    | 43-46        | True length of last group
//    | 47           | Number of bits used for the scaled group lengths (after subtraction of the
//    |              | reference value given in octets 38-41 and division by the length increment
//    |              | given in octet 42)
type Data2 struct {
	Data0
	GroupMethod            uint8  `json:"groupMethod"`
	MissingValue           uint8  `json:"missingValue"`
	MissingSubstitute1     uint32 `json:"missingSubstitute1"`
	MissingSubstitute2     uint32 `json:"missingSubstitute2"`
	NG                     uint32 `json:"ng"`
	GroupWidths            uint8  `json:"groupWidths"`
	GroupWidthsBits        uint8  `json:"groupWidthsBits"`
	GroupLengthsReference  uint32 `json:"groupLengthsReference"` // 13
	GroupLengthIncrement   uint8  `json:"groupLengthIncrement"`  // 14
	GroupLastLength        uint32 `json:"groupLastLength"`       // 15
	GroupScaledLengthsBits uint8  `json:"groupScaledLengthsBits"`
}

// ParseData2 parses data2 struct from the reader into the an array of floating-point values
func ParseData2(dataReader io.Reader, dataLength int, template *Data2) []float64 {

	fld := []float64{}

	// TODO: replace following by final code
	for i := 0; i < 20; i++ {
		fld = append(fld, 0)
	}

	return fld
}
