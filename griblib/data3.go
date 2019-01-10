package griblib

import (
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
func ParseData3(dataReader io.Reader, dataLength int, template *Data3) ([]float64, error) {

	var ival1 int64
	var ival2 int64
	var minsd int64

	//
	// Init reader
	//
	bitReader := makeBitReader(dataReader, dataLength)

	//
	//  Extract Spatial differencing values, if using DRS Template 5.3
	//
	rc := int(template.OctetsNumber) * 8
	if rc != 0 {
		var err error
		ival1, err = bitReader.readInt(rc)
		if err != nil {
			return []float64{}, err
		}

		if template.SpatialOrderDifference == 2 {
			ival2, err = bitReader.readInt(rc)
			if err != nil {
				return []float64{}, err
			}
		}

		minsd, err = bitReader.readInt(rc)
		if err != nil {
			return []float64{}, err
		}
	}

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

	// Spatial differencing is a pre-processing before group splitting at encoding time.
	// It is intended to reduce the size of sufficiently smooth fields, when combined with
	// a splitting scheme as described in Data Representation Template 5.2. At order 1,
	// an initial field of values f is replaced by a new field of values g,
	// where g1 = f1, g2 = f2 - f1, ..., gn = fn - fn-1.
	// At order 2, the field of values g is itself replaced by a new field of values h,
	// where h1 = f1, h2 = f2, h3 = g3 - g2, ..., hn = gn - gn-1.
	// To keep values positive, the overall minimum of the resulting field (either gmin or
	// hmin) is removed. At decoding time, after bit string unpacking, the original scaled
	// values are recovered by adding the overall minimum and summing up recursively.
	switch template.SpatialOrderDifference {
	case 1:
		// first order
		section7Data[0] = ival1

		for n := int(1); n < len(section7Data); n++ {
			section7Data[n] = section7Data[n] + section7Data[n-1] + minsd
		}
	case 2:
		// second order

		section7Data[0] = ival1
		section7Data[1] = ival2

		for n := int(2); n < len(section7Data); n++ {
			section7Data[n] = section7Data[n] + (2 * section7Data[n-1]) - section7Data[n-2] + minsd
		}
	}

	return template.scaleValues(section7Data, ifldmiss), nil
}
