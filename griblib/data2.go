package griblib

import (
	"bytes"
	"fmt"
	"io"
	"math"
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

	rawData := make([]byte, dataLength)
	dataReader.Read(rawData)

	var err error

	var missingValueSubstitute1 float64
	var missingValueSubstitute2 float64
	numberOfGroups := int(template.NG)
	if template.MissingValue == 1 {
		missingValueSubstitute1 = float64(template.MissingSubstitute1)
	} else if template.MissingValue == 2 {
		missingValueSubstitute1 = float64(template.MissingSubstitute1)
		missingValueSubstitute2 = float64(template.MissingSubstitute1)
	}

	//
	// Init reader
	//

	buffer := bytes.NewBuffer(rawData)

	bitReader := newReader(buffer)

	//
	//  Extract Each Group's reference value
	//
	references, err := bitReader.readUintsBlock(int(template.Bits), numberOfGroups)
	if err != nil {
		panic(err)
	}

	//
	//  Extract Each Group's bit width
	//
	widths, err := bitReader.readUintsBlock(int(template.GroupWidthsBits), numberOfGroups)
	if err != nil {
		panic(err)
	}

	for j := 0; j < numberOfGroups; j++ {
		widths[j] += uint64(template.GroupWidths)
	}

	//
	//  Extract Each Group's length (number of values in each group)
	//
	lengths, err := bitReader.readUintsBlock(int(template.GroupScaledLengthsBits), numberOfGroups)
	if err != nil {
		panic(err)
	}

	for j := 0; j < numberOfGroups; j++ {
		lengths[j] = (lengths[j] * uint64(template.GroupLengthIncrement)) + uint64(template.GroupLengthsReference)
	}
	lengths[template.NG-1] = uint64(template.GroupLastLength)

	//
	//  Test to see if the group widths and lengths are consistent with number of
	//  values, and length of section 7.
	//

	totBit := 0
	totLen := 0

	for j := 0; j < numberOfGroups; j++ {
		totBit += int(widths[j]) * int(lengths[j])
		totLen += int(lengths[j])
	}

	if totBit/8 > int(dataLength) {
		fmt.Println(totLen)
		panic("Checksum err")
	}

	//
	//  For each group, unpack data values
	//
	non := 0
	section7Data := []int64{}
	ifldmiss := []int64{}

	if template.MissingValue == 0 {
		n := 0
		for j := 0; j < numberOfGroups; j++ {
			if widths[j] != 0 {
				tmp, _ := bitReader.readUintsBlock(int(widths[j]), int(lengths[j]))
				for _, elt := range tmp {
					section7Data = append(section7Data, int64(elt))
				}
				//section7Data = append(section7Data, tmp...)

				for k := 0; k < int(lengths[j]); k++ {
					section7Data[n] = section7Data[n] + int64(references[j])
					n++
				}
			} else {
				for l := n; l < n+int(lengths[j]); l++ {
					section7Data = append(section7Data, int64(references[j]))
				}
				n = n + int(lengths[j])
			}
		}

	} else if template.MissingValue == 1 || template.MissingValue == 2 {
		// missing values included
		n := 0
		for j := 0; j < numberOfGroups; j++ {
			if widths[j] != 0 {
				msng1 := math.Pow(2.0, float64(widths[j])) - 1
				msng2 := msng1 - 1

				ifldmiss, err = bitReader.readIntsBlock(int(widths[j]), int(lengths[j]))
				if err != nil {
					panic(err)
				}

				for k := 0; k < int(lengths[j]); k++ {
					if section7Data[n] == int64(msng1) {
						ifldmiss[n] = 1
						//section7Data[n]=0
					} else if template.MissingValue == 2 && section7Data[n] == int64(msng2) {
						ifldmiss[n] = 2
						//section7Data[n]=0
					} else {
						ifldmiss[n] = 0
						//section7Data[non++]=section7Data[n]+references[j]
					}
					n++
				}
			} else {
				msng1 := math.Pow(2.0, float64(template.Bits)) - 1
				msng2 := msng1 - 1
				if references[j] == uint64(msng1) {
					for l := n; l < n+int(lengths[j]); l++ {
						ifldmiss[l] = 1
					}
				} else if template.MissingValue == 2 && references[j] == uint64(msng2) {
					for l := n; l < n+int(lengths[j]); l++ {
						ifldmiss[l] = 2
					}
				} else {
					for l := n; l < n+int(lengths[j]); l++ {
						ifldmiss[l] = 0
					}
					for l := non; l < non+int(lengths[j]); l++ {
						section7Data[l] = int64(references[j])
					}
					non += int(lengths[j])
				}
				n = n + int(lengths[j])
			}
		}
	}

	fld := make([]float64, len(section7Data))

	scaleStrategy := template.scaleFunc()

	if template.MissingValue == 0 {
		// no missing values
		for i, dataValue := range section7Data {
			fld[i] = scaleStrategy(dataValue)
		}
	} else if template.MissingValue == 1 || template.MissingValue == 2 {
		// missing values included
		non := 0
		for n, dataValue := range section7Data {
			if ifldmiss[n] == 0 {
				non++
				fld[n] = scaleStrategy(dataValue)
			} else if ifldmiss[n] == 1 {
				fld[n] = missingValueSubstitute1
			} else if ifldmiss[n] == 2 {
				fld[n] = missingValueSubstitute2
			}

		}
	}

	return fld
}
