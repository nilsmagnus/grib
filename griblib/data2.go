package griblib

import (
	"fmt"
	"io"

	"github.com/nilsmagnus/grib/internal/reader"
)

// Data2 is a Grid point data - complex packing
// http://www.nco.ncep.noaa.gov/pmb/docs/grib2/grib2_doc/grib2_temp5-2.shtml
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
//	| 22           | Group splitting method used (see Code Table 5.4)
//	|              |    - 0 : Row by Row Splitting
//	|              |    - 1 : General Group Splitting
//	|              |    - 2-191 : Reserved
//	|              |    - 192-254 : Reserved for Local Use
//	|              |    - 255 : Missing
//	| 23           | Missing value management used (see Code Table 5.5)
//	|              |    - 0 : No explicit missing values included within the data values
//	|              |    - 1 : Primary missing values included within the data values
//	|              |    - 2 : Primary and secondary missing values included within the data values
//	|              |    - 3-191 : Reserved
//	|              |    - 192-254 : Reserved for Local Use
//	|              |    - 255 : Missing
//	| 24-27        | Primary missing value substitute
//	| 28-31        | Secondary missing value substitute
//	| 32-35        | NG ― number of groups of data values into which field is split
//	| 36           | Reference for group widths (see Note 12)
//	| 37           | Number of bits used for the group widths (after the reference value in octet 36
//	|              | has been removed)
//	| 38-41        | Reference for group lengths (see Note 13)
//	| 42           | Length increment for the group lengths (see Note 14)
//	| 43-46        | True length of last group
//	| 47           | Number of bits used for the scaled group lengths (after subtraction of the
//	|              | reference value given in octets 38-41 and division by the length increment
//	|              | given in octet 42)
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
	fld := make([]float64, len(section7Data))

	scaleStrategy := template.scaleFunc()
	missingValueSubstitute1, missingValueSubstitute2 := template.missingValueSubstitute()

	if template.MissingValue == 0 {
		// no missing values
		for n, dataValue := range section7Data {
			fld[n] = scaleStrategy(dataValue)
		}
	}
	if template.MissingValue == 1 || template.MissingValue == 2 {
		// missing values included
		for n, dataValue := range section7Data {
			switch ifldmiss[n] {
			case 0:
				fld[n] = scaleStrategy(dataValue)
			case 1:
				fld[n] = missingValueSubstitute1
			case 2:
				fld[n] = missingValueSubstitute2
			}
		}
	}

	return fld
}

func (template *Data2) extractData(bitReader *reader.BitReader, bitGroups []bitGroupParameter) ([]int64, []int64, error) {
	var totalLength uint64
	for _, group := range bitGroups {
		totalLength += group.Length
	}
	section7Data := make([]int64, totalLength)
	ifldmiss := make([]int64, 0, totalLength)
	s7i := 0

	for _, bitGroup := range bitGroups {
		tmp, err := bitGroup.readData(bitReader)
		if err != nil {
			return section7Data, ifldmiss, fmt.Errorf("bitGroup read: %s", err.Error())
		}

		missingValueBits := bitGroup.Width
		if missingValueBits == 0 {
			missingValueBits = uint64(template.Bits)
		}

		missingValues := []uint64{1<<missingValueBits - 1, 1<<missingValueBits - 2}

		switch template.MissingValue {
		case 0:
			ifldmiss = append(ifldmiss, make([]int64, len(tmp))...)
			for _, elt := range tmp {
				section7Data[s7i] = elt + int64(bitGroup.Reference)
				s7i++
			}

		case 1:
			for _, elt := range tmp {
				if bitGroup.Reference == missingValues[0] {
					section7Data[s7i] = -1
					s7i++
					ifldmiss = append(ifldmiss, 1)
				} else {
					section7Data[s7i] = elt + int64(bitGroup.Reference)
					s7i++
					ifldmiss = append(ifldmiss, 0)
				}
			}

		case 2:
			for _, elt := range tmp {
				if bitGroup.Reference == missingValues[0] || bitGroup.Reference == missingValues[1] {
					section7Data[s7i] = -1
					s7i++

					if bitGroup.Reference == missingValues[0] {
						ifldmiss = append(ifldmiss, 1)
					} else {
						ifldmiss = append(ifldmiss, 2)
					}
				} else {
					section7Data[s7i] = elt + int64(bitGroup.Reference)
					s7i++

					ifldmiss = append(ifldmiss, 0)
				}
			}
		}
	}

	return section7Data, ifldmiss, nil
}

// ParseData2 parses data2 struct from the reader into the an array of floating-point values
func ParseData2(dataReader io.Reader, dataLength int, template *Data2) ([]float64, error) {

	//
	// Init reader
	//
	bitReader, err := reader.New(dataReader, dataLength)
	if err != nil {
		return nil, err
	}

	//
	// Extract Bit Group Parameters
	//
	bitGroups, err := template.extractBitGroupParameters(bitReader)
	if err != nil {
		return nil, err
	}

	//
	//  Test to see if the group widths and lengths are consistent with number of
	//  values, and length of section 7.
	//
	if err := checkLengths(bitGroups, dataLength); err != nil {
		return nil, err
	}

	//
	//  For each group, unpack data values
	//
	section7Data, ifldmiss, err := template.extractData(bitReader, bitGroups)
	if err != nil {
		return nil, fmt.Errorf("Data extract: %s", err.Error())
	}

	return template.scaleValues(section7Data, ifldmiss), nil
}
