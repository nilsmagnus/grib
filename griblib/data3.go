package griblib

import (
	"fmt"
	"io"

	"github.com/nilsmagnus/grib/internal/reader"
)

// Data3 is a Grid point data - complex packing and spatial differencing
// http://www.nco.ncep.noaa.gov/pmb/docs/grib2/grib2_temp5-3.shtml
type Data3 struct {
	Data2
	SpatialOrderDifference uint8 `json:"spatialOrderDifference"`
	OctetsNumber           uint8 `json:"octetsNumber"`
}

func (template *Data3) applySpacialDifferencing(section7Data []int64, minsd int64, ival1 int64, ival2 int64) {
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
}

func (template *Data3) extractSpacingDifferentialValues(bitReader *reader.BitReader) (int64, int64, int64, error) {
	var ival1 int64
	var ival2 int64
	var minsd int64

	rc := int(template.OctetsNumber) * 8
	if rc != 0 {
		var err error
		ival1, err = bitReader.ReadInt(rc)
		if err != nil {
			return minsd, ival1, ival2, fmt.Errorf("Spacial differencing Value 1: %s", err.Error())
		}

		if template.SpatialOrderDifference == 2 {
			ival2, err = bitReader.ReadInt(rc)
			if err != nil {
				return minsd, ival1, ival2, fmt.Errorf("Spacial differencing Value 2: %s", err.Error())
			}
		}

		minsd, err = bitReader.ReadInt(rc)
		if err != nil {
			return minsd, ival1, ival2, fmt.Errorf("Spacial differencing Reference: %s", err.Error())
		}
	}

	return minsd, ival1, ival2, nil
}

// ParseData3 parses data3 struct from the reader into the an array of floating-point values
func ParseData3(dataReader io.Reader, dataLength int, template *Data3) ([]float64, error) {

	//
	// Init reader
	//
	bitReader, err := reader.New(dataReader, dataLength)
	if err != nil {
		return nil, err
	}

	//
	//  Extract Spatial differencing values, if using DRS Template 5.3
	//
	minsd, ival1, ival2, err := template.extractSpacingDifferentialValues(bitReader)
	if err != nil {
		return nil, fmt.Errorf("Spacial differencing Value 1: %s", err.Error())
	}

	//
	// Extract Bit Group Parameters
	//
	bitGroups, err := template.extractBitGroupParameters(bitReader)
	if err != nil {
		return nil, fmt.Errorf("Groups: %s", err.Error())
	}

	//
	//  Test to see if the group widths and lengths are consistent with number of
	//  values, and length of section 7.
	//
	if err := checkLengths(bitGroups, dataLength); err != nil {
		return nil, fmt.Errorf("Check length: %s", err.Error())
	}

	//
	//  For each group, unpack data values
	//
	section7Data, ifldmiss, err := template.extractData(bitReader, bitGroups)
	if err != nil {
		return nil, fmt.Errorf("Data extract: %s", err.Error())
	}

	//
	// Apply spacing differencing
	//
	template.applySpacialDifferencing(section7Data, minsd, ival1, ival2)

	return template.scaleValues(section7Data, ifldmiss), nil
}
