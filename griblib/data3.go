package griblib

import (
	"bytes"
	"fmt"
	"io"
	"math"
)

//Data3 is a Grid point data - complex packing and spatial differencing
// http://www.nco.ncep.noaa.gov/pmb/docs/grib2/grib2_temp5-3.shtml
type Data3 struct {
	Data2
	SpatialOrderDifference uint8 `json:"spatialOrderDifference"`
	OctetsNumber           uint8 `json:"octetsNumber"`
}

// ParseData3 parses data3 struct from the reader into the an array of floating-point values
func ParseData3(dataReader io.Reader, dataLength int, template *Data3) []float64 {

	rawData := make([]byte, dataLength)
	dataReader.Read(rawData)

	var ival1 int64
	var ival2 int64
	var minsd uint64
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
	//  Extract Spatial differencing values, if using DRS Template 5.3
	//
	rc := int(template.OctetsNumber) * 8
	if rc != 0 {
		ival1, err = bitReader.readInt(rc)

		if template.SpatialOrderDifference == 2 {
			ival2, err = bitReader.readInt(rc)
		}

		minsd, err = bitReader.readUint(rc)
	}

	//fmt.Println(" ival1", ival1, "ival2", ival2, "minsd", minsd, "octed bytes", rc/8)
	//fmt.Println("")
	//d.Seek(int64(int(dataLength) - oc * 3), 1)
	//return
	//os.Exit(0)

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

	// debug
	for j := 0; j < numberOfGroups; j++ {
		//fmt.Println(j, " - Reference", references[j], "Width", widths[j], "length", lengths[j])
	}

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
	section7Data := make([]uint64, 0)
	ifldmiss := make([]int64, 0)

	if template.MissingValue == 0 {
		n := 0
		for j := 0; j < numberOfGroups; j++ {
			if widths[j] != 0 {
				tmp, _ := bitReader.readUintsBlock(int(widths[j]), int(lengths[j]))
				section7Data = append(section7Data, tmp...)

				for k := 0; k < int(lengths[j]); k++ {
					section7Data[n] = section7Data[n] + references[j]
					n++
				}
			} else {
				for l := n; l < n+int(lengths[j]); l++ {
					section7Data = append(section7Data, references[j])
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
					if section7Data[n] == uint64(msng1) {
						ifldmiss[n] = 1
						//section7Data[n]=0
					} else if template.MissingValue == 2 && section7Data[n] == uint64(msng2) {
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
						section7Data[l] = references[j]
					}
					non += int(lengths[j])
				}
				n = n + int(lengths[j])
			}
		}
	}

	/*
	   if ( references != 0 ) free(references)
	   if ( widths != 0 ) free(widths)
	   if ( lengths != 0 ) free(lengths)
	*/
	//
	//  If using spatial differences, add overall min value, and
	//  sum up recursively
	//
	//printf("SAGod: %ld %ld\n",idrsnum,idrstmpl[16])

	itemp := non
	ndpts := numberOfGroups

	if template.SpatialOrderDifference == 1 {
		// first order
		section7Data[0] = uint64(ival1)

		if template.MissingValue == 0 {
			itemp = ndpts // no missing values
		}

		for n := 1; n < itemp; n++ {
			section7Data[n] = section7Data[n] + minsd
			section7Data[n] = section7Data[n] + section7Data[n-1]
		}
	} else if template.SpatialOrderDifference == 2 {
		// second order

		section7Data[0] = uint64(ival1)
		section7Data[1] = uint64(ival2)
		if template.MissingValue == 0 {
			itemp = ndpts
		}

		for n := 2; n < itemp; n++ {
			//section7Data[n] = section7Data[n] + int64(minsd)
			// WTF does this line do??? it seems to fuck up everything
			//section7Data[n] = section7Data[n] + (2 * section7Data[n-1]) - section7Data[n-2]
		}
	}

	fld := make([]float64, len(section7Data))
	binaryScale := math.Pow(2.0, float64(template.BinaryScale))
	decimalScale := math.Pow(10.0, -float64(template.DecimalScale))

	scaleStrategy := scaleFunc(binaryScale, decimalScale, template.Reference)

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

func scaleFunc(bScale float64, dScale float64, referenceValue float32) func(uintValue uint64) float64 {
	scale := bScale * dScale
	ref := scale * float64(referenceValue)
	return func(value uint64) float64 {
		signed := int64(value)
		return ref + float64(signed)*scale
	}
}
