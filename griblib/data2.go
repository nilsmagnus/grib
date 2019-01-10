package griblib

import (
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
	GroupMethod            uint8  `json:"groupMethod"`            // 22
	MissingValue           uint8  `json:"missingValue"`           // 23
	MissingSubstitute1     uint32 `json:"missingSubstitute1"`     // 24-27
	MissingSubstitute2     uint32 `json:"missingSubstitute2"`     // 28-31
	NG                     uint32 `json:"ng"`                     // 32-35
	GroupWidths            uint8  `json:"groupWidths"`            // 36
	GroupWidthsBits        uint8  `json:"groupWidthsBits"`        // 37
	GroupLengthsReference  uint32 `json:"groupLengthsReference"`  // 38-41
	GroupLengthIncrement   uint8  `json:"groupLengthIncrement"`   // 42
	GroupLastLength        uint32 `json:"groupLastLength"`        // 43-46
	GroupScaledLengthsBits uint8  `json:"groupScaledLengthsBits"` // 47
}

func (template *Data2) missingValueSubstitute() (float64, float64) {
	var missingValueSubstitute1 float64
	var missingValueSubstitute2 float64
	if template.MissingValue == 1 {
		missingValueSubstitute1 = float64(template.MissingSubstitute1)
	} else if template.MissingValue == 2 {
		missingValueSubstitute1 = float64(template.MissingSubstitute1)
		missingValueSubstitute2 = float64(template.MissingSubstitute1)
	}
	return missingValueSubstitute1, missingValueSubstitute2
}

func (template *Data2) scaleValues(section7Data []int64, ifldmiss []int64) []float64 {
	fld := []float64{}

	scaleStrategy := template.scaleFunc()
	missingValueSubstitute1, missingValueSubstitute2 := template.missingValueSubstitute()

	if template.MissingValue == 0 {
		// no missing values
		for _, dataValue := range section7Data {
			fld = append(fld, scaleStrategy(dataValue))
		}
	}
	if template.MissingValue == 1 || template.MissingValue == 2 {
		// missing values included
		non := 0
		for n, dataValue := range section7Data {
			switch ifldmiss[n] {
			case 0:
				non++
				fld = append(fld, scaleStrategy(dataValue))
			case 1:
				fld = append(fld, missingValueSubstitute1)
			case 2:
				fld = append(fld, missingValueSubstitute2)
			}
		}
	}

	return fld
}

// ParseData2 parses data2 struct from the reader into the an array of floating-point values
func ParseData2(dataReader io.Reader, dataLength int, template *Data2) ([]float64, error) {

	//
	// Init reader
	//
	bitReader := makeBitReader(dataReader, dataLength)

	bitGroups, err := template.extractBitGroupParameters(bitReader)
	if err != nil {
		return []float64{}, err
	}

	//
	//  Test to see if the group widths and lengths are consistent with number of
	//  values, and length of section 7.
	//
	if err := checkLengths(bitGroups, dataLength); err != nil {
		return []float64{}, err
	}

	//
	//  For each group, unpack data values
	//
	non := int64(0)
	section7Data := []int64{}
	ifldmiss := []int64{}

	fmt.Printf("offset %d\n", bitReader.offset)

	switch template.MissingValue {
	case 0:
		for _, bitGroup := range bitGroups {
			var tmp []int64
			if bitGroup.Width != 0 {
				tmp, err = bitReader.readIntsBlock(int(bitGroup.Width), int64(bitGroup.Length), false)
				if err != nil {
					fmt.Printf("ERROR %s\n", err.Error())
				}
			} else {
				tmp = bitGroup.zeroGroup()
			}

			for _, elt := range tmp {
				section7Data = append(section7Data, elt+int64(bitGroup.Reference))
			}
		}

	case 1, 2:
		// missing values included
		n := 0
		for _, bitGroup := range bitGroups {
			if bitGroup.Width != 0 {
				msng1 := math.Pow(2.0, float64(bitGroup.Width)) - 1
				msng2 := msng1 - 1

				var err error
				ifldmiss, err = bitReader.readIntsBlock(int(bitGroup.Width), int64(bitGroup.Length), false)
				if err != nil {
					return []float64{}, err
				}

				for k := 0; k < int(bitGroup.Length); k++ {
					if section7Data[n] == int64(msng1) {
						ifldmiss[n] = 1
						//section7Data[n]=0
						n++
						continue
					}
					if template.MissingValue == 2 && section7Data[n] == int64(msng2) {
						ifldmiss[n] = 2
						//section7Data[n]=0
						n++
						continue
					}
					ifldmiss[n] = 0
					//section7Data[non++]=section7Data[n]+references[j]
					n++
				}
			} else {
				msng1 := math.Pow(2.0, float64(template.Bits)) - 1
				msng2 := msng1 - 1
				if bitGroup.Reference == uint64(msng1) {
					for l := n; l < n+int(bitGroup.Length); l++ {
						ifldmiss[l] = 1
					}
				} else if template.MissingValue == 2 && bitGroup.Reference == uint64(msng2) {
					for l := n; l < n+int(bitGroup.Length); l++ {
						ifldmiss[l] = 2
					}
				} else {
					for l := n; l < n+int(bitGroup.Length); l++ {
						ifldmiss[l] = 0
					}
					for l := non; l < non+int64(bitGroup.Length); l++ {
						section7Data[l] = int64(bitGroup.Reference)
					}
					non += int64(bitGroup.Length)
				}
				n = n + int(bitGroup.Length)
			}
		}
	}

	return template.scaleValues(section7Data, ifldmiss), nil
}
